<?php


require dirname(__FILE__) . "/vendor/autoload.php";



$address = "127.0.0.1:10005";
$client = new Proto\Client($address, ["credentials" => Grpc\ChannelCredentials::createInsecure(),]);
$request = new Proto\Request();
$request->setPath("/rpc/index");
$request->setData(
    [
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
    ]
);


list($response, $status) = $client->Call($request)->wait();
if ($status->code !== Grpc\STATUS_OK) {
    echo "ERROR: " . $status->code . ", " . $status->details . PHP_EOL;
    exit(1);
}


$res = $response->getResponse();
var_dump($res);
var_dump(json_decode($res, true));




