# tensorflow-client-with-Go

a tensorflow serving Go client
 
## 参考
* [github](https://github.com/SineYuan/tensorflow-demo)

## Run
* 下载代码，[github](https://github.com/daxiang6506/TF-Client-Go)
* VSCode中，`docker build images` 选择Dockerfile，生成镜像 `tf-client-go`
* 启动 `tensorflow serving` 服务端
  ```
  docker run -it --name model-server -v $(pwd):/bitnami/model-data daxiang6506/tensorflow-serving:1.8.0
  ```
* 启动 `tensorflow serving` 客户端
  ```
  docker run -p 1323:1323 -d --link model-server tf-client-go
  ```
## 修订
* 2017-7-16 增加 `pic` 文件夹