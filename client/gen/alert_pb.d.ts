import * as jspb from 'google-protobuf'

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';


export class Notification extends jspb.Message {
  getId(): string;
  setId(value: string): Notification;

  getSummary(): string;
  setSummary(value: string): Notification;

  getMessage(): string;
  setMessage(value: string): Notification;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Notification.AsObject;
  static toObject(includeInstance: boolean, msg: Notification): Notification.AsObject;
  static serializeBinaryToWriter(message: Notification, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Notification;
  static deserializeBinaryFromReader(message: Notification, reader: jspb.BinaryReader): Notification;
}

export namespace Notification {
  export type AsObject = {
    id: string,
    summary: string,
    message: string,
  }
}

export class SendAlertRequest extends jspb.Message {
  getNotification(): Notification | undefined;
  setNotification(value?: Notification): SendAlertRequest;
  hasNotification(): boolean;
  clearNotification(): SendAlertRequest;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendAlertRequest.AsObject;
  static toObject(includeInstance: boolean, msg: SendAlertRequest): SendAlertRequest.AsObject;
  static serializeBinaryToWriter(message: SendAlertRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendAlertRequest;
  static deserializeBinaryFromReader(message: SendAlertRequest, reader: jspb.BinaryReader): SendAlertRequest;
}

export namespace SendAlertRequest {
  export type AsObject = {
    notification?: Notification.AsObject,
  }
}

export class SendAlertResponse extends jspb.Message {
  getId(): string;
  setId(value: string): SendAlertResponse;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SendAlertResponse.AsObject;
  static toObject(includeInstance: boolean, msg: SendAlertResponse): SendAlertResponse.AsObject;
  static serializeBinaryToWriter(message: SendAlertResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SendAlertResponse;
  static deserializeBinaryFromReader(message: SendAlertResponse, reader: jspb.BinaryReader): SendAlertResponse;
}

export namespace SendAlertResponse {
  export type AsObject = {
    id: string,
  }
}

