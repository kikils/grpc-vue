/**
 * @fileoverview gRPC-Web generated client stub for proto
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.proto = require('./count_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.addNumServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.proto.addNumServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.addNumParams,
 *   !proto.proto.totalNum>}
 */
const methodDescriptor_addNumService_addNum = new grpc.web.MethodDescriptor(
  '/proto.addNumService/addNum',
  grpc.web.MethodType.UNARY,
  proto.proto.addNumParams,
  proto.proto.totalNum,
  /**
   * @param {!proto.proto.addNumParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.totalNum.deserializeBinary
);


/**
 * @param {!proto.proto.addNumParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.totalNum)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.totalNum>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServiceClient.prototype.addNum =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.addNumService/addNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_addNum,
      callback);
};


/**
 * @param {!proto.proto.addNumParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.totalNum>}
 *     Promise that resolves to the response
 */
proto.proto.addNumServicePromiseClient.prototype.addNum =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.addNumService/addNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_addNum);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.getTotalNumParams,
 *   !proto.proto.totalNum>}
 */
const methodDescriptor_addNumService_getTotalNum = new grpc.web.MethodDescriptor(
  '/proto.addNumService/getTotalNum',
  grpc.web.MethodType.UNARY,
  proto.proto.getTotalNumParams,
  proto.proto.totalNum,
  /**
   * @param {!proto.proto.getTotalNumParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.totalNum.deserializeBinary
);


/**
 * @param {!proto.proto.getTotalNumParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.totalNum)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.totalNum>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServiceClient.prototype.getTotalNum =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.addNumService/getTotalNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_getTotalNum,
      callback);
};


/**
 * @param {!proto.proto.getTotalNumParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.totalNum>}
 *     Promise that resolves to the response
 */
proto.proto.addNumServicePromiseClient.prototype.getTotalNum =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.addNumService/getTotalNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_getTotalNum);
};


module.exports = proto.proto;

