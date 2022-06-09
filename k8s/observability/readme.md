


```
minikube start --cache-images=true --memory 2048mb --cpus 8 --image-mirror-country=cn --registry-mirror="https://registry.docker-cn.com,https://docker.mirrors.ustc.edu.cn" 

minikube kubectl -- apply -f prom.yaml
```