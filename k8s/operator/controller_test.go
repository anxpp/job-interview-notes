package operator

import (
	"flag"
	"fmt"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/workqueue"
	"k8s.io/klog/v2"
	"testing"
	"time"
)

func TestController(t *testing.T) {
	var (
		kubeconfig string
		master     string
	)
	flag.StringVar(&kubeconfig, "kubeconfig", "./config.yaml", "kubeconfig file")
	flag.StringVar(&master, "master", "", "https://192.168.153.135:6443")
	// 读取构建config
	config, err := clientcmd.BuildConfigFromFlags(master, kubeconfig)
	if err != nil {
		klog.Fatal(err)
	}
	// 创建 k8s client
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatal(err)
	}
	// 从指定客户端、资源、命名空间和字段选择器创建一个新的 List-Watch
	podListWatcher := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())
	// 构建一个具有速率限制排队功能的 work queue
	queue := workqueue.NewRateLimitingQueue(workqueue.DefaultControllerRateLimiter())
	// 创建 indexer 和 infomer
	indexer, informer := cache.NewIndexerInformer(podListWatcher, &v1.Pod{}, 0, cache.ResourceEventHandlerFuncs{
		// 当有 Pod 创建时，根据 Delta Queue 弹出Object 生成对应的 Key，并加入 WorkQueue中，此处也可以根据 Object 的一些属性进行过滤
		AddFunc: func(obj interface{}) {
			key, err := cache.MetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
		UpdateFunc: nil,
		// Pod 删除操作
		DeleteFunc: func(obj interface{}) {
			// 再生成 key 之前检查对象。因为资源删除后有可能会有尽心重建操作，如果监听时错过了删除信息，会导致该条记录时陈旧的
			key, err := cache.DeletionHandlingMetaNamespaceKeyFunc(obj)
			if err == nil {
				queue.Add(key)
			}
		},
	}, cache.Indexers{})
	// 创建 controller
	controller := NewController(queue, indexer, informer)
	stop := make(chan struct{})
	defer close(stop)
	// 启动 controller
	go controller.Run(1, stop)
	select {}
}

type Controller struct {
	indexer  cache.Indexer                   // Indexer的引用
	queue    workqueue.RateLimitingInterface // WorkQueue的引用
	informer cache.Controller                // informer引用
}

func NewController(queue workqueue.RateLimitingInterface, indexer cache.Indexer, informer cache.Controller) *Controller {
	return &Controller{
		indexer:  indexer,
		queue:    queue,
		informer: informer,
	}
}

func (c *Controller) Run(threadiness int, stopCh chan struct{}) {
	defer runtime.HandleCrash()
	defer c.queue.ShuttingDown()
	klog.Info("Starting pod controller")
	// 启动 informer 线程
	// Run函数：
	// 1、运行一个 Reflector，并从 ListWatcher 中获取对象的通知放到队列中（Delta Queue）
	// 2、从队列去除对象并处理该对象的相关业务
	go c.informer.Run(stopCh)
	// 等待缓存同步队列
	if !cache.WaitForCacheSync(stopCh, c.informer.HasSynced) {
		runtime.HandleError(fmt.Errorf("Time out waiting for caches to sync\n"))
		return
	}
	// 启动多个 Worker 线程处理 WorkQueue 中的 Object
	for i := 0; i < threadiness; i++ {
		go wait.Until(c.runWorker, time.Second, stopCh)
	}
	<-stopCh
	klog.Info("Stopping Pod Controller")
}

func (c *Controller) runWorker() {
	// 无线循环接收并处理消息
	for c.processNextItem() {
	}
}

// 从 WorkQueue 中获取对象，并打印信息
func (c *Controller) processNextItem() bool {
	key, shutdown := c.queue.Get()
	// 退出
	if shutdown {
		return false
	}
	// 标记已处理
	defer c.queue.Done(key)
	// 打印 Key 对应的 Object 及相关信息
	err := c.syncToStdout(key.(string))
	c.handleError(err, key)
	return true
}

// 获取Key对应的 Object，并打印相关信息
func (c *Controller) syncToStdout(key string) error {
	obj, exists, err := c.indexer.GetByKey(key)
	if err != nil {
		klog.Errorf("Fetching object with key %s from store failed with %v", key, err)
		return err
	}
	if !exists {
		fmt.Printf("Pod %s does not exist anymore\n", key)
	} else {
		fmt.Printf("Sync/Add/Update for Pod %s\n", obj.(*v1.Pod).GetName())
	}
	return nil
}

func (c *Controller) handleError(err error, key interface{}) {
	if err != nil {
		klog.Errorf("handle error key=%v err=%v", key, err)
	}
}
