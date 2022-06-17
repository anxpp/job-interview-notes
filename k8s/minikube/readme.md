

```shell

export HTTP_PROXY=http://127.0.0.1:41091
export HTTPS_PROXY=https://127.0.0.1:41091
export NO_PROXY=localhost,127.0.0.1,172.16.29.155,10.96.0.0/12,192.168.59.0/24,192.168.39.0/24

minikube start \
  --extra-config=kubelet.CAdvisorPort=4194 \
  --image-repository=registry.cn-hangzhou.aliyuncs.com/google_containers 
```
