php环境需要grpc和protobuf扩展
可以在容器中运行，镜像：junyang7/php:7.2-fpm-alpine-2，注意网络需要设置成host模式，请求host需要时宿主机的真实IP
注意需要composer install
```text
docker run -d \
  --name test_php_container \
  --network host \
  -v $(pwd):/home/q/system/php \
  junyang7/php:7.2-fpm-alpine-2
  
```