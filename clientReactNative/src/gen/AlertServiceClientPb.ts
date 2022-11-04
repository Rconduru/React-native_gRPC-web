/**
 * @fileoverview gRPC-Web generated client stub for mobilePoc.alert
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';
import * as alert_pb from './alert_pb';


export class AlertServiceClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'text';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname;
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodInfoSendAlert = new grpcWeb.MethodDescriptor(
    '/mobilePoc.alert.AlertService/SendAlert',
    grpcWeb.MethodType.UNARY,
    alert_pb.SendAlertRequest,
    alert_pb.SendAlertResponse,
    (request: alert_pb.SendAlertRequest) => {
      return request.serializeBinary();
    },
    alert_pb.SendAlertResponse.deserializeBinary
  );

  sendAlert(
    request: alert_pb.SendAlertRequest,
    metadata: grpcWeb.Metadata | null): Promise<alert_pb.SendAlertResponse>;

  sendAlert(
    request: alert_pb.SendAlertRequest,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: alert_pb.SendAlertResponse) => void): grpcWeb.ClientReadableStream<alert_pb.SendAlertResponse>;

  sendAlert(
    request: alert_pb.SendAlertRequest,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: alert_pb.SendAlertResponse) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/mobilePoc.alert.AlertService/SendAlert',
        request,
        metadata || {},
        this.methodInfoSendAlert,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/mobilePoc.alert.AlertService/SendAlert',
    request,
    metadata || {},
    this.methodInfoSendAlert);
  }

  methodInfoListenAlert = new grpcWeb.MethodDescriptor(
    '/mobilePoc.alert.AlertService/ListenAlert',
    grpcWeb.MethodType.SERVER_STREAMING,
    google_protobuf_empty_pb.Empty,
    alert_pb.Notification,
    (request: google_protobuf_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    alert_pb.Notification.deserializeBinary
  );

  listenAlert(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata) {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/mobilePoc.alert.AlertService/ListenAlert',
      request,
      metadata || {},
      this.methodInfoListenAlert);
  }

}

