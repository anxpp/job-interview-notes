## 基础

### 数据类型

#### 基本数据类型

- bool
- byte
- char
- short
- int
- long
- float
- double

#### 包装类型

#### 缓存池

- bool
- byte
- short -128 ~ 127
- int -128 ~ 127
- char \u0000 ~ \u007F

### String

#### 不可变

#### String Pool

#### StringBuffer 和 StringBuilder

### 运算

### 关键字

### Object

### 继承

### 反射

### 异常

### 泛型

### 注解

### 引用类型

### 特性

## 容器

### 基础

### 设计模式

### 分析

- ArrayList
- Vector
- CopyOnWriteArrayList
- LinkedList
- HashMap
- ConcurrentHashMap
- LinkedHashMap
- WeakHashMap

## 并发

### 使用方式

- 继承 Thread: 不能继承其他类了，调用start方法

```
public class test0 {

    public static void main(String[] args) {
        Thread MyThread = new MyThread();
        MyThread.start();
    }
}

class MyThread extends Thread {
    @Override
    public void run() {
        System.out.println("hello myThread" + Thread.currentThread().getName());
    }
}
```

- 实现 Runnable：还可以继承其他类，需要附加到Thread实例上执行

```
public class test0 {

    public static void main(String[] args) {

        MyRunnable myRunnable = new MyRunnable();
        Thread thread = new Thread(myRunnable);
        thread.start();
    }
}

class MyRunnable implements Runnable{
    @Override
    public void run(){
        System.out.println("hello myRunnable" + Thread.currentThread().getName());
    }
}
```

- 实现 Callable：可以继承其他类，调用call方法，同能能获取到线程的返回结果

```
public class study2 {
    public static void main(String[] args) {
        MyCallable MyCallable = new MyCallable("张方兴");
        String call = null;
        try {
            call = MyCallable.call();
        } catch (Exception e) {
            e.printStackTrace();
        }
        System.out.println(call);
    }
}

class MyCallable implements Callable<String>{

    private String name;

    public MyCallable(String name) {
        this.name = name;
    }

    @Override
    public String call() throws Exception {
        return "call:" + name;
    }
} 
```

### 线程池

- 通过Runnable 或者 Callable创建任务对象
- 把创建完成的实现 Runnable/Callable接口的 对象直接交给 ExecutorService 执行: ExecutorService.execute（Runnable command））或者也可以把 Runnable 对象或Callable 对象提交给 ExecutorService 执行（ExecutorService.submit（Runnable task）或 ExecutorService.submit（Callable <T> task））
- 如果执行 ExecutorService.submit（…），ExecutorService 将返回一个实现Future接口的对象（我们刚刚也提到过了执行 execute()方法和 submit()方法的区别，submit()会返回一个 FutureTask 对象）。由于 FutureTask 实现了 Runnable，我们也可以创建 FutureTask，然后直接交给 ExecutorService 执行。
- 最后，主线程可以执行 FutureTask.get()方法来等待任务执行完成。主线程也可以执行 FutureTask.cancel（boolean mayInterruptIfRunning）来取消此任务的执行。

#### ThreadPoolExecutor

```
/**
     * 用给定的初始参数创建一个新的ThreadPoolExecutor。
     */
    public ThreadPoolExecutor(int corePoolSize,     //线程池的核心线程数量
                              int maximumPoolSize,  //线程池的最大线程数
                              long keepAliveTime,   //当线程数大于核心线程数时，多余的空闲线程存活的最长时间
                              TimeUnit unit,        //时间单位
                              BlockingQueue<Runnable> workQueue,    //任务队列，用来储存等待执行任务的队列
                              ThreadFactory threadFactory,          //线程工厂，用来创建线程，一般默认即可
                              RejectedExecutionHandler handler      //拒绝策略，当提交的任务过多而不能及时处理时，我们可以定制策略来处理任务
                               ) {
        if (corePoolSize < 0 ||
            maximumPoolSize <= 0 ||
            maximumPoolSize < corePoolSize ||
            keepAliveTime < 0)
            throw new IllegalArgumentException();
        if (workQueue == null || threadFactory == null || handler == null)
            throw new NullPointerException();
        this.corePoolSize = corePoolSize;
        this.maximumPoolSize = maximumPoolSize;
        this.workQueue = workQueue;
        this.keepAliveTime = unit.toNanos(keepAliveTime);
        this.threadFactory = threadFactory;
        this.handler = handler;
    }
```

- corePoolSize : 核心线程数线程数定义了最小可以同时运行的线程数量
- maximumPoolSize : 当队列中存放的任务达到队列容量的时候，当前可以同时运行的线程数量变为最大线程数
- workQueue: 当新任务来的时候会先判断当前运行的线程数量是否达到核心线程数，如果达到的话，新任务就会被存放在队列中

- keepAliveTime:当线程池中的线程数量大于 corePoolSize 的时候，如果这时没有新的任务提交，核心线程外的线程不会立即销毁，而是会等待，直到等待的时间超过了 keepAliveTime才会被回收销毁
- unit : keepAliveTime 参数的时间单位
- threadFactory :executor 创建新线程的时候会用到
- handler :饱和策略



### 基础机制

### 中断

### 互斥同步

### 协作

### 线程状态

### J.U.C

### 内存模型

### 线程安全

### 锁优化

### 最佳实践

## 虚拟机

### 运行时数据区域

#### 程序计数器

#### 虚拟机栈

#### 本地方法栈

#### 堆

#### 方法区

#### 运行时常量池

#### 直接内存

### GC

### 内存分配与回收

### 类加载机制

## IO

### 分类

### 

### 

### 

### 
