# prometheus 使用

## 概述

### 监控方法论

#### Google 的黄金指标

- 延迟：服务请求所需耗时
- 流量：服务容量需求
- 错误：请求失败的速率
- 饱和度：负载

#### Netflix 的 USE 方法

- Utilization使用率：系统资源的使用情况
- Saturation饱和度：资源饱和度
- Error错误：错误数

#### Weave Cloud 的 RED 方法

- Rate：每秒接收的请求数
- Errors：每秒失败的请求数
- Duration：每个请求所花费的时间

### 黑盒和白盒

- 黑盒监控：主要只用探针
- 黑盒监控：能够了解内部的实际运行状态

## 安装

### docker 安装 prometheus

```
docker run -d -p 9090:9090 -v /root/prom/prometheus.yml:/etc/prometheus/prometheus.yml -v /etc/localtime:/etc/localtime:ro --name=prom prom/prometheus
```

### 使用 go client 作为数据源

```
// ./goclient/main.go
```

### yml 配置数据源

```
  - job_name: myapp
    scrape_interval: 10s
    static_configs:
      - targets:
          - 192.168.50.206:8888
```

### promQL

```
myapp_processed_ops_total{}[5m]
``` 

### grafana

```
docker run -d -p 3000:3000 -v /etc/localtime:/etc/localtime:ro --name=grafana grafana/grafana
```

## promQL

### 介绍

#### Demo

- 获取主机当前可用内存

```
node_memory_free_bytes_total / (1024*1024)
```

- 基于两小时样本，预测未来24小时内磁盘是否会满

```
IF predict_linear(node_filesystem_free[2h], 24*3600)<0
```

增加告警：

```
ALERT DiskWillFullIn24Hours
    IF predict_linear(node_filesystem_free[2h], 24*3600)<0
```

- http 请求总数的常见 PromQL

```
# 查询请求总数
http_request_total

# 查询返回所有时间序列、指标http_request_total，以及给定job和handler的标签
http_request_total{job="apiserver", handler="/api/comments"}

# 查询状态码为200的请求总数
http_request_total{code="200"}

# 查询5分钟内的请求总量
http_request_total{}[5m]

# 所有http请求的总量
sum(http_request_total)

# 以server结尾的时间序列
http_request_total{job=~".*server"}

# 过滤4xx以外的状态码
http_request_total{code!~"4.."}

# 以1次每分钟的速率采集最近30分钟内的指标数据，然后返回这30分钟内距离当前时间最近的5分钟内的采集结果
rate(http_request_total[5m])[30m:1m]

# 以一次每秒的速率采集最近5分钟内的数据并将结果以时间序列的形式返回
rate(http_request_total[5m])
```

#### 数据类型

- 瞬时向量
- 区间向量
- 标量
- 字符串

#### 时间序列

- 指标：包括指标名称和一组标签
- 时间戳：精确到毫秒的时间戳
- 样本值：float64的浮点数

### 选择器

#### 匹配器

- 相等匹配器 = 
- 不相等匹配器 !=
- 正则表达式匹配器  =~
- 正则表达式相反匹配器 !~

#### 瞬时向量选择器

返回指定时间戳之前查新到的最新样本的瞬时向量

#### 区间向量选择器

```
http_request_total{}[5m]
```

单位可选：s,m,h,d,w,y(60*60*24*365秒)

#### 偏移量修改器

```
http_request_total{} offset 5m
http_request_total{}[5m] offset 5m
```

### 指标类型

#### 只增不减的计数器 Counter

可以结合topk、rate、increase、irate等计算

```
# 获取http请求量的增长速率
rate(http_request_total[5m])

# 统计每台机器每秒介绍的
sum without(device) (rate(http_request_total[5m]))
```

#### 可增可减的仪表盘 Gauge

经常结合 predict_linear 和 delta 使用

#### 直方图 Histogram

在一段时间范围内对数据进行采样，并将其计入可配置的存储桶中，后续可通过指定区间筛选样本，也可以统计样本总数。

#### 摘要