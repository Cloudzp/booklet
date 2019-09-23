# kubectl 常用命令简介

## kubectl弹性伸缩
```
$ kubectl scale --replicas=3 deployment/{name}
```

## 去掉污点
```
$ kubectl taint node {node_name} {label}-
```
