# gin-mesh
 基于网格的 api 服务
###安装
```
创建镜像  
make docker
镜像打包
make docker-image-tar
镜像上传到 k8s 导出镜像
docker load < ./gin-mesh.tar
安装
kubectl apply -f gin-mesh.yaml
kubectl apply -f gateway.yaml
```
###删除
```
kubectl delete gateway gin-mesh-gateway -n echo-grpc
kubectl delete virtualservice gin-mesh-vs -n echo-grpc
kubectl delete --ignore-not-found=true -f gin-mesh.yaml
```