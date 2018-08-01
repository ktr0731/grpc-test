import { grpc } from 'grpc-web-client';
import { SimpleRequest } from './api_pb.js';
import { Example } from './api_pb_service.js';

const req = new SimpleRequest();
// req.setName("tooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooo-looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong-teeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeext");
req.setName("ktr");
grpc.invoke(Example.ServerStreaming, {
    request: req,
    host: "http://localhost:50051",
    onMessage: message => {
        console.log("onMessage", message);
    },
    onEnd: res => {
        const { status, statusMessage, headers, message, trailers } = res;
        if (status === grpc.Code.OK && message) {
            console.log("all ok. : ", message.toObject());
            document.querySelector("#title").innerHTML = "It Works!";
        }
    }
});
