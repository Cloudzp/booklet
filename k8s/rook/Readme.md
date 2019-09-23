# rook-ceph安装

## 安装步骤
- 下载最新的版本压缩包**rook-release-1.0.zip**
```
$ ll
rook-release-1.0.zip

$ unzip rook-release-1.0.zip
rook-release-1.0

$ cd rook-release-1.0/cluster/examples/kubernetes/ceph
$ kubectl apply -f common.yaml
$ kubectl apply -f operator.yaml
$ kubectl apply -f cluster.yaml // cluster.yaml文件按照如下修改配置
$ kubectl apply -f storageclass.yaml
```
- [cluster.yaml](yaml/cluster.yaml)

---

## 关于rook-ceph的问题汇总
   关于部署使用中遇到的问题可以优先查看官[issue](https://rook.io/docs/rook/v0.8/common-issues.html)是否有相关的问题：
 
### 问题1. docker重启后rook-osd启动不起来
```
rook-ceph-osd-0-576fc688c6-rs6hb      0/1     CrashLoopBackOff   10         26m
rook-ceph-osd-1-75d69db689-n6xp9      0/1     Error              10         26m
```
- 关注[issue](https://github.com/rook/rook/issues/3157)
- 临时解决方案，重新安装

### 问题2. 安装指导创建rook-ceph完成后无法创建rook-ceph-mgr 的pod
- 检查节点发现有残留的容器没有删除，删除残留容器后重新部署ok；
在每个节点上执行：
```
docker ps -a |grep rook
```
