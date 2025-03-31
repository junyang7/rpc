docker run -d \
  --name test_php_container \
  --network host \
  -v $(pwd):/home/q/system/php \
  junyang7/php:7.2-fpm-alpine-2
