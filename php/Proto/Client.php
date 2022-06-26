<?php


namespace Proto;


use Grpc\BaseStub;



class Client extends BaseStub
{

    public function __construct($hostname, $opts, $channel = null)
    {
        parent::__construct($hostname, $opts, $channel);
    }

    public function Call(Request $argument, $metadata = [], $options = [])
    {
        return $this->_simpleRequest('/proto.Service/Call', $argument, ['\Proto\Response', 'decode'], $metadata, $options);
    }

}
