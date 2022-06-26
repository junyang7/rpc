<?php


namespace Proto;


use Grpc\MethodDescriptor;
use Grpc\ServerContext;
use Grpc\Status;


class Stub
{

    public function Request(Request $request, ServerContext $context)
    {
        $context->setStatus(Status::unimplemented());
        return null;
    }

    public final function getMethodDescriptors()
    {
        return [
            '/proto.Service/Call' => new MethodDescriptor(
                $this,
                'Call',
                '\Proto\Request',
                MethodDescriptor::UNARY_CALL
            ),
        ];
    }

}
