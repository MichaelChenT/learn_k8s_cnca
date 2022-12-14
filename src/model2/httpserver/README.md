## 作业三简介

镜像 
docker push liondocker2022/learn-k8s-cncamp:v1.0-lion

默认端口 8081

## 启动步骤
```bash
docker run --name httpserver --net host -d liondocker2022/learn-k8s-cncamp:v1.0-lion
```

### 如开启防火墙
```bash
firewall-cmd --zone=public --add-port 8081 --permanent
firewall-cmd --reload
```

### 查看访问信息
```
docker logs -f --tail=200 httpserver

curl http://ip:8081/healthz // 健康检查
curl http://ip:8081         // 访问信息
```
