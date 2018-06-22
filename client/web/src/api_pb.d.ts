// package: api
// file: api.proto

import * as jspb from "google-protobuf";

export class SimpleRequest extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SimpleRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SimpleRequest): SimpleRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SimpleRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SimpleRequest;
  static deserializeBinaryFromReader(message: SimpleRequest, reader: jspb.BinaryReader): SimpleRequest;
}

export namespace SimpleRequest {
  export type AsObject = {
    name: string,
  }
}

export class SimpleResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SimpleResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SimpleResponse): SimpleResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SimpleResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SimpleResponse;
  static deserializeBinaryFromReader(message: SimpleResponse, reader: jspb.BinaryReader): SimpleResponse;
}

export namespace SimpleResponse {
  export type AsObject = {
    message: string,
  }
}

export class Name extends jspb.Message {
  getFirstName(): string;
  setFirstName(value: string): void;

  getLastName(): string;
  setLastName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Name.AsObject;
  static toObject(includeInstance: boolean, msg: Name): Name.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Name, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Name;
  static deserializeBinaryFromReader(message: Name, reader: jspb.BinaryReader): Name;
}

export namespace Name {
  export type AsObject = {
    firstName: string,
    lastName: string,
  }
}

export class UnaryMessageRequest extends jspb.Message {
  hasName(): boolean;
  clearName(): void;
  getName(): Name | undefined;
  setName(value?: Name): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryMessageRequest): UnaryMessageRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryMessageRequest;
  static deserializeBinaryFromReader(message: UnaryMessageRequest, reader: jspb.BinaryReader): UnaryMessageRequest;
}

export namespace UnaryMessageRequest {
  export type AsObject = {
    name?: Name.AsObject,
  }
}

export class UnaryRepeatedRequest extends jspb.Message {
  clearNameList(): void;
  getNameList(): Array<string>;
  setNameList(value: Array<string>): void;
  addName(value: string, index?: number): string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryRepeatedRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryRepeatedRequest): UnaryRepeatedRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryRepeatedRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryRepeatedRequest;
  static deserializeBinaryFromReader(message: UnaryRepeatedRequest, reader: jspb.BinaryReader): UnaryRepeatedRequest;
}

export namespace UnaryRepeatedRequest {
  export type AsObject = {
    nameList: Array<string>,
  }
}

export class UnaryRepeatedMessageRequest extends jspb.Message {
  clearNameList(): void;
  getNameList(): Array<Name>;
  setNameList(value: Array<Name>): void;
  addName(value?: Name, index?: number): Name;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryRepeatedMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryRepeatedMessageRequest): UnaryRepeatedMessageRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryRepeatedMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryRepeatedMessageRequest;
  static deserializeBinaryFromReader(message: UnaryRepeatedMessageRequest, reader: jspb.BinaryReader): UnaryRepeatedMessageRequest;
}

export namespace UnaryRepeatedMessageRequest {
  export type AsObject = {
    nameList: Array<Name.AsObject>,
  }
}

export class UnaryRepeatedEnumRequest extends jspb.Message {
  clearGendersList(): void;
  getGendersList(): Array<Gender>;
  setGendersList(value: Array<Gender>): void;
  addGenders(value: Gender, index?: number): Gender;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryRepeatedEnumRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryRepeatedEnumRequest): UnaryRepeatedEnumRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryRepeatedEnumRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryRepeatedEnumRequest;
  static deserializeBinaryFromReader(message: UnaryRepeatedEnumRequest, reader: jspb.BinaryReader): UnaryRepeatedEnumRequest;
}

export namespace UnaryRepeatedEnumRequest {
  export type AsObject = {
    gendersList: Array<Gender>,
  }
}

export class UnaryMapRequest extends jspb.Message {
  getKvsMap(): jspb.Map<string, string>;
  clearKvsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryMapRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryMapRequest): UnaryMapRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryMapRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryMapRequest;
  static deserializeBinaryFromReader(message: UnaryMapRequest, reader: jspb.BinaryReader): UnaryMapRequest;
}

export namespace UnaryMapRequest {
  export type AsObject = {
    kvsMap: Array<[string, string]>,
  }
}

export class UnaryMapMessageRequest extends jspb.Message {
  getKvsMap(): jspb.Map<string, Name>;
  clearKvsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryMapMessageRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryMapMessageRequest): UnaryMapMessageRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryMapMessageRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryMapMessageRequest;
  static deserializeBinaryFromReader(message: UnaryMapMessageRequest, reader: jspb.BinaryReader): UnaryMapMessageRequest;
}

export namespace UnaryMapMessageRequest {
  export type AsObject = {
    kvsMap: Array<[string, Name.AsObject]>,
  }
}

export class UnaryOneofRequest extends jspb.Message {
  hasMsg(): boolean;
  clearMsg(): void;
  getMsg(): Name | undefined;
  setMsg(value?: Name): void;

  hasPlain(): boolean;
  clearPlain(): void;
  getPlain(): string;
  setPlain(value: string): void;

  getNameCase(): UnaryOneofRequest.NameCase;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryOneofRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryOneofRequest): UnaryOneofRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryOneofRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryOneofRequest;
  static deserializeBinaryFromReader(message: UnaryOneofRequest, reader: jspb.BinaryReader): UnaryOneofRequest;
}

export namespace UnaryOneofRequest {
  export type AsObject = {
    msg?: Name.AsObject,
    plain: string,
  }

  export enum NameCase {
    NAME_NOT_SET = 0,
    MSG = 1,
    PLAIN = 2,
  }
}

export class UnaryEnumRequest extends jspb.Message {
  getGender(): Gender;
  setGender(value: Gender): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): UnaryEnumRequest.AsObject;
  static toObject(includeInstance: boolean, msg: UnaryEnumRequest): UnaryEnumRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: UnaryEnumRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): UnaryEnumRequest;
  static deserializeBinaryFromReader(message: UnaryEnumRequest, reader: jspb.BinaryReader): UnaryEnumRequest;
}

export namespace UnaryEnumRequest {
  export type AsObject = {
    gender: Gender,
  }
}

export enum Gender {
  MALE = 0,
  FEMALE = 1,
}

