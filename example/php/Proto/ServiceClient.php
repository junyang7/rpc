<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Proto;

/**
 */
class ServiceClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Proto\Request $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     * @return \Grpc\UnaryCall
     */
    public function Call(\Proto\Request $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/proto.Service/Call',
        $argument,
        ['\Proto\Response', 'decode'],
        $metadata, $options);
    }

}
