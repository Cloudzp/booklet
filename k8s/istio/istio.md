# 姨死透

1. http://www.sohu.com/a/270131876_463994 (了解istio)
2. http://cizixs.com/2018/08/26/what-is-istio/

## 问题：
1. gateway 创建不成功，如下：
```
$ kubectl apply -f samples/bookinfo/networking/bookinfo-gateway.yaml
```
Error from server (Timeout): error when creating "samples/bookinfo/networking/bookinfo-gateway.yaml": Timeout: request did not complete within allowed duration
Error from server (Timeout): error when creating "samples/bookinfo/networking/bookinfo-gateway.yaml": Timeout: request did not complete within allowed durati
