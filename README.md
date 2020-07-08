Etcd 场景练习 (Go)
======================

DESCRIPTION
-------------
go-etcd场景练习  (Etcd V3 服务注册与发现)


PREREQUISITES
-------------

- This requires Go 1.14 or later
- Requires that [GOPATH is set](https://golang.org/doc/code.html#GOPATH)

```
$ go help gopath
$ # ensure the PATH contains $GOPATH/bin
$ export PATH=$PATH:$GOPATH/bin
```

INSTALL
-------

```
git clone https://github.com/Dingjianhui/go-etcd.git
cd go-etcd
go mod tidy

```

server/utils/Service.go 中配置Etcd  的节点更改为自己的Etcd节点信息

client/utils/Client.go 中配置Etcd  的节点更改为自己的Etcd节点信息

[使用docker模拟etcd集群的创建](http://dingjianhui.top/blog/2020/06/11/etcd-cluster-by-docker/)

TRY IT!
-------

- Run the server(服务注册)

  ```
  $ go run server/main.go
  ```
  

- Run the client(服务发现)

  ```
  $ go run client/main.go
  ```










