<?php

require_once "vendor/autoload.php";

$header = [
    "path" => "/rpc/test",
    "method" => "POST",
    "rid" => crc32(time()),
    "content-type" => "application/rpc",
    "token" => base64_encode(time()),
];
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

$req = new \Proto\Request();
$req->setHeader($header);
$req->setBody(json_encode($body));

$host = "10.222.96.105:60001";
$opt = [
    "credentials" => Grpc\ChannelCredentials::createInsecure(),
];
$channel = null;


$cli = new \Proto\ServiceClient($host, $opt, $channel);
list($response, $status) = $cli->Call($req)->wait();

if ($status->code !== Grpc\STATUS_OK) {
    throw new Exception("ERROR: " . $status->code . " | " . $status->details);
}

$res = $response->getResponse();
var_dump($res);
