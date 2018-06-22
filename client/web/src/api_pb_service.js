// package: api
// file: api.proto

var api_pb = require("./api_pb");
var grpc = require("grpc-web-client").grpc;

var Example = (function () {
  function Example() {}
  Example.serviceName = "api.Example";
  return Example;
}());

Example.Unary = {
  methodName: "Unary",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.SimpleRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryMessage = {
  methodName: "UnaryMessage",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryMessageRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryRepeated = {
  methodName: "UnaryRepeated",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryRepeatedRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryRepeatedMessage = {
  methodName: "UnaryRepeatedMessage",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryRepeatedMessageRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryRepeatedEnum = {
  methodName: "UnaryRepeatedEnum",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryRepeatedEnumRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryMap = {
  methodName: "UnaryMap",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryMapRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryMapMessage = {
  methodName: "UnaryMapMessage",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryMapMessageRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryOneof = {
  methodName: "UnaryOneof",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryOneofRequest,
  responseType: api_pb.SimpleResponse
};

Example.UnaryEnum = {
  methodName: "UnaryEnum",
  service: Example,
  requestStream: false,
  responseStream: false,
  requestType: api_pb.UnaryEnumRequest,
  responseType: api_pb.SimpleResponse
};

Example.ClientStreaming = {
  methodName: "ClientStreaming",
  service: Example,
  requestStream: true,
  responseStream: false,
  requestType: api_pb.SimpleRequest,
  responseType: api_pb.SimpleResponse
};

Example.ServerStreaming = {
  methodName: "ServerStreaming",
  service: Example,
  requestStream: false,
  responseStream: true,
  requestType: api_pb.SimpleRequest,
  responseType: api_pb.SimpleResponse
};

Example.BidiStreaming = {
  methodName: "BidiStreaming",
  service: Example,
  requestStream: true,
  responseStream: true,
  requestType: api_pb.SimpleRequest,
  responseType: api_pb.SimpleResponse
};

exports.Example = Example;

function ExampleClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

ExampleClient.prototype.unary = function unary(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.Unary, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryMessage = function unaryMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryMessage, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryRepeated = function unaryRepeated(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryRepeated, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryRepeatedMessage = function unaryRepeatedMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryRepeatedMessage, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryRepeatedEnum = function unaryRepeatedEnum(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryRepeatedEnum, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryMap = function unaryMap(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryMap, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryMapMessage = function unaryMapMessage(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryMapMessage, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryOneof = function unaryOneof(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryOneof, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

ExampleClient.prototype.unaryEnum = function unaryEnum(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  grpc.unary(Example.UnaryEnum, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          callback(Object.assign(new Error(response.statusMessage), { code: response.status, metadata: response.trailers }), null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
};

Example.prototype.clientStreaming = function clientStreaming() {
  throw new Error("Bi-directional streaming is not currently supported");
}

ExampleClient.prototype.serverStreaming = function serverStreaming(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Example.ServerStreaming, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.end.forEach(function (handler) {
        handler();
      });
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

Example.prototype.bidiStreaming = function bidiStreaming() {
  throw new Error("Client streaming is not currently supported");
}

exports.ExampleClient = ExampleClient;

