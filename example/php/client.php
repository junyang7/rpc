<?php

require_once "vendor/autoload.php";

// 设置请求头部信息
$header = [
    "path" => "/rpc/test",
    "method" => "POST",
    "rid" => crc32(time()),
    "content-type" => "application/rpc",
    "token" => base64_encode(time()),
];

// 设置请求体
$body = [
    "appId" => "8365174",
    "appKey" => "8djf923j23978dre34",
    "a" => "A",
    "b" => 0,
    "c" => json_encode(
        [
            "a" => "A",
            "b" => 0,
        ]
    )
];

// 创建 gRPC 请求对象
$req = new \Proto\Request();
$req->setHeader($header);
$req->setBody(json_encode($body));

// 设置 gRPC 服务地址
///
/// 注意，要改成宿主机IP
///
$host = "10.222.96.121:60001";

// 设置 gRPC 客户端选项（无安全性）
$opt = [
    "credentials" => Grpc\ChannelCredentials::createInsecure(),
];

// 创建 gRPC 客户端实例
$cli = new \Proto\ServiceClient($host, $opt);

// 调用 gRPC 服务，并等待响应
list($response, $status) = $cli->Call($req)->wait();

// 检查响应状态
if ($status->code !== Grpc\STATUS_OK) {
    throw new Exception("ERROR: " . $status->code . " | " . $status->details);
}

// 获取并输出响应
$res = $response->getResponse();
var_dump($res);
