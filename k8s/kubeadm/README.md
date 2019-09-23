# kubeadm
## 一、使用kubeadm安装集群
> 参考文档：https://feisky.gitbooks.io/kubernetes/components/kubeadm.html

### 1.安装docker
- 安装docker的yum源
````
yum install -y yum-utils device-mapper-persistent-data lvm2
yum-config-manager \
    --add-repo \
    https://download.docker.com/linux/centos/docker-ce.repo

````
- 查看历史最新的Docker版本
````
yum list docker-ce.x86_64  --showduplicates |sort -r
````

- 选择一个版本安装
````
yum install -y --setopt=obsoletes=0 \
  docker-ce-18.06.1.ce-3.el7

systemctl start docker
systemctl enable docker

````

### 2.安装kubeadm，kubelet
- 修改节点的yum源地址
````
cat <<EOF > /etc/yum.repos.d/kubernetes.repo
[kubernetes]
name=Kubernetes
baseurl=https://mirrors.aliyun.com/kubernetes/yum/repos/kubernetes-el7-x86_64/
enabled=1
gpgcheck=1
repo_gpgcheck=1
gpgkey=https://mirrors.aliyun.com/kubernetes/yum/doc/yum-key.gpg https://mirrors.aliyun.com/kubernetes/yum/doc/rpm-package-key.gpg
EOF
````

- 安装
````
yum install -y kubelet kubeadm kubectl
````
- 配置cgroup，设置节点的docker的cgroup与kubelet一致，否则启动会报如下错误
````
kubelet cgroup driver: "systemd" is different from docker cgroup driver: "cgroupfs"
````
Note: 修改kubelet的日志到文件中
```shell
$ cd /etc/systemd/system/kubelet.service.d
$ vi 10-kubeadm.conf
```
增加这两个参数 --logtostderr=false --log-dir=/home/data1/kubernetes/logs/kubelet

### 3. pull相关的image
国外镜像拉去需要翻墙，但万幸的是阿里提供了国内的镜像地址，将google镜像定时同步到自己的仓库，我们可以通过在国内阿里镜像仓库拉去镜像，然后`docker tag`
成google镜像的名称

```
docker pull registry.cn-hangzhou.aliyuncs.com/google_containers/{IMAGE}
docker tag registry.cn-hangzhou.aliyuncs.com/google_containers/{IMAGE} k8s.gcr.io/{IMAGE}
```

### 4. 安装集群组件
```
kubeadm init --pod-network-cidr 10.244.0.0/16 --kubernetes-version stable
```

## 二、使用kubeadm纳管节点

- 执行kubeadm join

在master节点上执行：
````
$ kubeadm token create --print-join-command

kubeadm join xxxx:6443 --token 78u9ke.4rbn5uwditomqlm1 --discovery-token-ca-cert-hash sha256:4d7f780a7a43e13bb03495fe28d44a54ddabde7fa4c0317ed01c8c2597a70fa9
````
- 将以上命令copy到要加入的主机上执行即可！

- 执行完成后在去master节点执行如下命令：
```
$ kubectl get csr
NAME                                                   AGE       REQUESTOR                 CONDITION
node-csr-c69HXe7aYcqkS1bKmH4faEnHAWxn6i2bHZ2mD04jZyQ   18s       system:bootstrap:878f07   Pending

$ kubectl certificate approve  node-csr-c69HXe7aYcqkS1bKmH4faEnHAWxn6i2bHZ2mD04jZyQ
```

## 三、网络插件安装
















## 问题汇总：
1. kubelet启动失败
````
[ERROR CRI]: unable to check if the container runtime at "/var/run/dockershim.sock" is running: exit status 1
````
[问题结局参考issue](https://github.com/kubernetes-sigs/cri-tools/issues/153)
在kubeadm启动时加上参数`--ignore-preflight-errors cri`

2. [preflight] Some fatal errors occurred:
    /proc/sys/net/bridge/bridge-nf-call-iptables contents are not set to 1
    
解决方案：
```
$ echo 1 > /proc/sys/net/bridge/bridge-nf-call-iptables
$ echo 1 > /proc/sys/net/bridge/bridge-nf-call-ip6tables
````
3. unable to fetch the kubeadm-config ConfigMap: failed to get config map: configmaps "kubeadm-config" is forbidden: 
User "system:bootstrap:y0gmey" cannot get configmaps in the namespace "kube-system"
解决：master和node的kubelet、kubeadm版本不一致，建议要和master的版本保持一致
