# timelinefeed

# 服务发现

```relationProjectServerAddr := "relationproject-service.timeline.svc.cluster.local:50051" // gRPC 服务器的地址
relationProjectConn, err := grpc.Dial(relationProjectServerAddr, grpc.WithInsecure())
relationProjectClient := helloworld.NewGreeterClient(relationProjectConn)

// ...

commentProjectServerAddr := "commentproject-service.timeline.svc.cluster.local:50052" // gRPC 服务器的地址
commentProjectConn, err := grpc.Dial(commentProjectServerAddr, grpc.WithInsecure())
commentProjectClient := helloworld.NewGreeterClient(commentProjectConn)

...

listenAddr := "0.0.0.0:50051" // 指定 gRPC 监听地址
lis, err := net.Listen("tcp", listenAddr)
if err != nil {
  log.Fatalf("failed to listen: %v", err)
}

grpcServer := grpc.NewServer()
greeterServer := helloworld_server.NewGreeterServerImpl() // 创建 GreeterServerImpl 实例
helloworld.RegisterGreeterServer(grpcServer, greeterServer)

// ...

listenAddr := "0.0.0.0:50052" // 指定 gRPC 监听地址
lis,err := net.Listen("tcp", listenAddr)
if err != nil {
  log.Fatalf("failed to listen: %v", err)
}

grpcServer := grpc.NewServer()
greeterServer := helloworld_server.NewGreeterServerImpl() // 创建 GreeterServerImpl 实例
helloworld.RegisterGreeterServer(grpcServer, greeterServer)
```

# 负载均衡
```
ClusterIP 服务是 Kubernetes 中的一种服务类型，它通过将流量均匀分发给后端的 Pod 来提供负载均衡功能：

创建服务：当你创建一个 ClusterIP 服务时，Kubernetes 控制平面会为该服务分配一个虚拟的 ClusterIP 地址。

服务关联：你可以通过选择器（selector）指定你希望服务关联的 Pod。服务会将流量转发给与选择器匹配的 Pod。

内部负载均衡：当流量到达 ClusterIP 地址时，Kubernetes 的内部负载均衡机制会将流量均匀地分发给关联的 Pod。这意味着每个请求将被发送到其中一个 Pod，以实现流量的均衡分配。

IPVS 或 iptables：Kubernetes 在不同的版本和配置中使用不同的技术来实现内部负载均衡。在一些 Kubernetes 部署中，使用 IPVS（IP Virtual Server）作为内部负载均衡技术，而在其他部署中，使用 iptables 规则进行负载均衡。

IPVS：IPVS 是一个基于内核的负载均衡器，通过使用虚拟服务器和调度算法来实现流量分发。

iptables：iptables 是 Linux 上常用的防火墙工具，它也可以用于负载均衡，通过配置 iptables 规则来将流量转发给后端的 Pod。

```
```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: bproject-deployment
  namespace: timeline
spec:
  replicas: 3
  selector:
    matchLabels:
      app: bproject
  template:
    metadata:
      labels:
        app: bproject
    spec:
      containers:
      - name: bproject-container
        image: yunyang1999/bproject
        ports:
        - containerPort: 50051
---
apiVersion: v1
kind: Service
metadata:
  name: bproject-service
  namespace: timeline
spec:
  selector:
    app: bproject
  ports:
  - protocol: TCP
    port: 50051
    targetPort: 50051
  type: ClusterIP
```

# 熔断
```
这个配置文件中使用了 Istio 的熔断机制来实现服务的熔断保护。下面是对配置文件中熔断相关部分的详细描述：

DestinationRule（目标规则）:

metadata.name: my-destinationrule：定义了目标规则资源的名称。

spec.host: aproject-service.timeline.svc.cluster.local：指定了要应用熔断策略的目标服务的主机名。

spec.trafficPolicy.outlierDetection：配置了异常检测的参数，用于触发熔断。

consecutive5xxErrors: 5：设置在5s内发生连续5个5xx错误时触发熔断。

interval: 5s：定义了异常检测的时间间隔为5秒。

baseEjectionTime: 30s：指定了熔断后的基础驱逐时间为30秒。

maxEjectionPercent: 50：设置了最大驱逐比例为50%，即当超过50%的请求失败时触发熔断。

```

```
apiVersion: networking.istio.io/v1alpha3
kind: Gateway
metadata:
  name: my-gateway
spec:
  selector:
    istio: ingressgateway
  servers:
  - port:
      number: 80
      name: http
      protocol: HTTP
    hosts:
    - "*"
---
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: my-virtualservice
spec:
  hosts:
  - "*"
  gateways:
  - my-gateway
  http:
  - match:
    - uri:
        prefix: /virtualservice_a/
    rewrite:
      uri: /
    route:
    - destination:
        host: aproject-service.timeline.svc.cluster.local
        port:
          number: 8080
---
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: my-destinationrule
spec:
  host: aproject-service.timeline.svc.cluster.local
  trafficPolicy:
    outlierDetection:
      consecutive5xxErrors: 5
      interval: 5s
      baseEjectionTime: 30s
      maxEjectionPercent: 50

```

# 镜像发布
```
FROM golang:1.19 AS builder：指定了构建阶段的基础镜像。在这里，使用了官方提供的Golang 1.19镜像作为构建环境。

WORKDIR /app：设置工作目录为/app，即在容器内部的/app路径下进行后续操作。

COPY . .：将本地当前目录中的所有文件和文件夹复制到容器的/app目录中。这里使用了两个点表示源路径和目标路径相同。

ENV GO111MODULE=on：设置环境变量GO111MODULE为on，启用Go模块支持。

ENV GOPROXY=https://goproxy.io,direct：设置环境变量GOPROXY为https://goproxy.io,direct，指定Go模块的代理地址。

RUN CGO_ENABLED=0 GOOS=linux go build -o ccommentproject：在容器中执行构建命令。这里使用go build命令编译Go应用程序，并将可执行文件命名为ccommentproject。CGO_ENABLED=0禁用了CGO，GOOS=linux指定了目标操作系统为Linux。

FROM ubuntu:20.04：指定了最终阶段的基础镜像。在这里，选择了官方提供的Ubuntu 20.04镜像作为最终的运行环境。

COPY --from=builder /app/ccommentproject /app/ccommentproject：从构建阶段的镜像中复制可执行文件ccommentproject到最终阶段的容器中的/app/ccommentproject路径。

ENV PORT=8080：设置环境变量PORT为8080，指定应用程序监听的端口号。

EXPOSE 8080：声明容器将监听的端口号为8080。这并不会自动打开端口，但是可以在运行容器时使用-p选项来映射宿主机的端口到容器的端口。

CMD ["/app/ccommentproject"]：在容器启动时执行的默认命令。这里指定了运行可执行文件ccommentproject作为容器的入口点。
```

```
# 构建阶段
FROM golang:1.19 AS builder

WORKDIR /app

COPY . .

# 设置环境变量
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.io,direct

# 构建项目
RUN CGO_ENABLED=0 GOOS=linux go build -o ccommentproject

# 最终阶段
FROM ubuntu:20.04

COPY --from=builder /app/ccommentproject /app/ccommentproject

ENV PORT=8080

EXPOSE 8080

CMD ["/app/ccommentproject"]
```

# 存储中间件配置
```
# Redis 配置文件

# 设置持久化方式为RDB和AOF的混合持久化
save 900 1
save 300 10
save 60 10000

# 设置RDB持久化文件名和目录
dbfilename dump.rdb
dir /var/lib/redis

# 启用AOF持久化
appendonly yes
appendfilename "appendonly.aof"

# 避免网络闪断导致的全量主从复制
repl-backlog-size 10mb
repl-backlog-ttl 3600

# 设置Redis实例的最大内存为8GB
maxmemory 8gb

# 当达到最大内存限制时，采用LRU算法删除键
maxmemory-policy volatile-lru

# 在主从同步时，将fork进程的耗时限制在50毫秒以内
repl-diskless-sync-delay 50

# 设置主从复制的超时时间
repl-timeout 60

# 启用TCP keepalive选项，保持长连接
tcp-keepalive 300

# 监听IP地址和端口号
bind 127.0.0.1
port 6379

# 启用日志记录
logfile "/var/log/redis/redis.log"
```

# JVM参数
```
-Xmx16g                       // 设置堆内存最大值为32GB
-Xms16g                       // 设置堆内存初始值为32GB
-XX:+UseG1GC                  // 使用G1垃圾收集器
-XX:MaxGCPauseMillis=200      // 设置最大停顿时间为200毫秒
-XX:ParallelGCThreads=16      // 设置并行GC线程数为16
-XX:+UseParallelRefProcEnabled   // 启用并行引用处理
-XX:-UseLargePages            // 禁用共享内存（大页）
```
