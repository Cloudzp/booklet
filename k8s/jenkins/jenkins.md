# helm 安装jenkins
### 安装
```
$ helm install --name irain-jenkins stable/jenkins --set Persistence.StorageClass=rook-ceph-block --set Master.AdminUser=irain --set Master.AdminPassword=irain2018  --namespace=jenkins \
--set Master.Cpu=2 --set Master.Memory=1Gi --set Master.JavaOpts="-Xms1g -Xmx1g"
```
- 参数说明：
  - 

### 卸载
```
$ helm delete irain-jenkins --purge=true
```
