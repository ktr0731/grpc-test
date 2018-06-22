import { grpc } from 'grpc-web-client';
import { Example, SimpleRequest } from './api_pb.js';

const req = new SimpleRequest();
req.setName("ktr");
grpc.unary(Example.Unary, {
    request: req,
    host: "localhost:50051",
    onEnd: res => {
        const { status, statusMessage, headers, message, trailers } = res;
        if (status === grpc.Code.OK && message) {
            console.log("all ok. got book: ", message.toObject());
        }
    }
});
