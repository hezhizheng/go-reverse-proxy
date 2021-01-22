# go-reverse-proxy

> 基于go实现简单http/https的反向代理服务

## 使用
- windows64 直接下载 release 
- 编译
```
git clone 
cd go-reverse-proxy

# 编译
go build -ldflags "-s -w"

# 执行启动
./go-reverse-proxy
```

## 参数说明
```
./go-reverse-proxy.exe -h

-p string
      本地监听的端口 (default "1874")
-r string
      需要代理的地址 (default "https://hzz.cool")
```