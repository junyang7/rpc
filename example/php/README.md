php环境需要grpc和protobuf扩展
可以在容器中运行，镜像：junyang7/php:7.2-fpm-alpine-2，注意网络需要设置成host模式，请求host需要时宿主机的真实IP
注意需要composer install
```text
docker run -d \
  --name test_php_container \ # 容器名称
  --network host \ # 网络类型设置成host，与主机共享
  -v /Users/junyang7/env/111/rpc/example/php:/home/q/system/php \ # 挂载项目路径到容器，方便宿主机上修改直接容器可见，容器内部执行
  junyang7/php:7.2-fpm-alpine-2 # 镜像
  
```