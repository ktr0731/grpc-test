// package: api
// file: api.proto

import * as api_pb from "./api_pb";
import {grpc} from "grpc-web-client";

type ExampleUnary = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.SimpleRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryMessage = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryMessageRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryRepeated = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryRepeatedRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryRepeatedMessage = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryRepeatedMessageRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryRepeatedEnum = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryRepeatedEnumRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryMap = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryMapRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryMapMessage = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryMapMessageRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryOneof = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryOneofRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleUnaryEnum = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.UnaryEnumRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleClientStreaming = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: true;
  readonly responseStream: false;
  readonly requestType: typeof api_pb.SimpleRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleServerStreaming = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof api_pb.SimpleRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

type ExampleBidiStreaming = {
  readonly methodName: string;
  readonly service: typeof Example;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof api_pb.SimpleRequest;
  readonly responseType: typeof api_pb.SimpleResponse;
};

export class Example {
  static readonly serviceName: string;
  static readonly Unary: ExampleUnary;
  static readonly UnaryMessage: ExampleUnaryMessage;
  static readonly UnaryRepeated: ExampleUnaryRepeated;
  static readonly UnaryRepeatedMessage: ExampleUnaryRepeatedMessage;
  static readonly UnaryRepeatedEnum: ExampleUnaryRepeatedEnum;
  static readonly UnaryMap: ExampleUnaryMap;
  static readonly UnaryMapMessage: ExampleUnaryMapMessage;
  static readonly UnaryOneof: ExampleUnaryOneof;
  static readonly UnaryEnum: ExampleUnaryEnum;
  static readonly ClientStreaming: ExampleClientStreaming;
  static readonly ServerStreaming: ExampleServerStreaming;
  static readonly BidiStreaming: ExampleBidiStreaming;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }
export type ServiceClientOptions = { transport: grpc.TransportConstructor }

interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: () => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}

export class ExampleClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: ServiceClientOptions);
  unary(
    requestMessage: api_pb.SimpleRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unary(
    requestMessage: api_pb.SimpleRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryMessage(
    requestMessage: api_pb.UnaryMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryMessage(
    requestMessage: api_pb.UnaryMessageRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryRepeated(
    requestMessage: api_pb.UnaryRepeatedRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryRepeated(
    requestMessage: api_pb.UnaryRepeatedRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryRepeatedMessage(
    requestMessage: api_pb.UnaryRepeatedMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryRepeatedMessage(
    requestMessage: api_pb.UnaryRepeatedMessageRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryRepeatedEnum(
    requestMessage: api_pb.UnaryRepeatedEnumRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryRepeatedEnum(
    requestMessage: api_pb.UnaryRepeatedEnumRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryMap(
    requestMessage: api_pb.UnaryMapRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryMap(
    requestMessage: api_pb.UnaryMapRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryMapMessage(
    requestMessage: api_pb.UnaryMapMessageRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryMapMessage(
    requestMessage: api_pb.UnaryMapMessageRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryOneof(
    requestMessage: api_pb.UnaryOneofRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryOneof(
    requestMessage: api_pb.UnaryOneofRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryEnum(
    requestMessage: api_pb.UnaryEnumRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  unaryEnum(
    requestMessage: api_pb.UnaryEnumRequest,
    callback: (error: ServiceError, responseMessage: api_pb.SimpleResponse|null) => void
  ): void;
  clientStreaming(): void;
  serverStreaming(requestMessage: api_pb.SimpleRequest, metadata?: grpc.Metadata): ResponseStream<api_pb.SimpleResponse>;
  bidiStreaming(): void;
}

