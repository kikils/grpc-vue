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
 *   !proto.proto.countNumparams,
 *   !proto.proto.totalNum>}
 */
const methodDescriptor_addNumService_countNum = new grpc.web.MethodDescriptor(
  '/proto.addNumService/countNum',
  grpc.web.MethodType.UNARY,
  proto.proto.countNumparams,
  proto.proto.totalNum,
  /**
   * @param {!proto.proto.countNumparams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.totalNum.deserializeBinary
);


/**
 * @param {!proto.proto.countNumparams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.totalNum)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.totalNum>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServiceClient.prototype.countNum =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.addNumService/countNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_countNum,
      callback);
};


/**
 * @param {!proto.proto.countNumparams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.totalNum>}
 *     Promise that resolves to the response
 */
proto.proto.addNumServicePromiseClient.prototype.countNum =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.addNumService/countNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_countNum);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.getRoomTotalNumParams,
 *   !proto.proto.roomTotalNums>}
 */
const methodDescriptor_addNumService_getRoomTotalNum = new grpc.web.MethodDescriptor(
  '/proto.addNumService/getRoomTotalNum',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.proto.getRoomTotalNumParams,
  proto.proto.roomTotalNums,
  /**
   * @param {!proto.proto.getRoomTotalNumParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.roomTotalNums.deserializeBinary
);


/**
 * @param {!proto.proto.getRoomTotalNumParams} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.proto.roomTotalNums>}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServiceClient.prototype.getRoomTotalNum =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/proto.addNumService/getRoomTotalNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_getRoomTotalNum);
};


/**
 * @param {!proto.proto.getRoomTotalNumParams} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.proto.roomTotalNums>}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServicePromiseClient.prototype.getRoomTotalNum =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/proto.addNumService/getRoomTotalNum',
      request,
      metadata || {},
      methodDescriptor_addNumService_getRoomTotalNum);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.addRoomParams,
 *   !proto.proto.roomInfo>}
 */
const methodDescriptor_addNumService_addRoom = new grpc.web.MethodDescriptor(
  '/proto.addNumService/addRoom',
  grpc.web.MethodType.UNARY,
  proto.proto.addRoomParams,
  proto.proto.roomInfo,
  /**
   * @param {!proto.proto.addRoomParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.roomInfo.deserializeBinary
);


/**
 * @param {!proto.proto.addRoomParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.roomInfo)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.roomInfo>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServiceClient.prototype.addRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.addNumService/addRoom',
      request,
      metadata || {},
      methodDescriptor_addNumService_addRoom,
      callback);
};


/**
 * @param {!proto.proto.addRoomParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.roomInfo>}
 *     Promise that resolves to the response
 */
proto.proto.addNumServicePromiseClient.prototype.addRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.addNumService/addRoom',
      request,
      metadata || {},
      methodDescriptor_addNumService_addRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.joinRoomParams,
 *   !proto.proto.joinResult>}
 */
const methodDescriptor_addNumService_joinRoom = new grpc.web.MethodDescriptor(
  '/proto.addNumService/joinRoom',
  grpc.web.MethodType.UNARY,
  proto.proto.joinRoomParams,
  proto.proto.joinResult,
  /**
   * @param {!proto.proto.joinRoomParams} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.joinResult.deserializeBinary
);


/**
 * @param {!proto.proto.joinRoomParams} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.joinResult)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.joinResult>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServiceClient.prototype.joinRoom =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.addNumService/joinRoom',
      request,
      metadata || {},
      methodDescriptor_addNumService_joinRoom,
      callback);
};


/**
 * @param {!proto.proto.joinRoomParams} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.joinResult>}
 *     Promise that resolves to the response
 */
proto.proto.addNumServicePromiseClient.prototype.joinRoom =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.addNumService/joinRoom',
      request,
      metadata || {},
      methodDescriptor_addNumService_joinRoom);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.proto.Null,
 *   !proto.proto.roomList>}
 */
const methodDescriptor_addNumService_getRooms = new grpc.web.MethodDescriptor(
  '/proto.addNumService/getRooms',
  grpc.web.MethodType.UNARY,
  proto.proto.Null,
  proto.proto.roomList,
  /**
   * @param {!proto.proto.Null} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.proto.roomList.deserializeBinary
);


/**
 * @param {!proto.proto.Null} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.proto.roomList)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.proto.roomList>|undefined}
 *     The XHR Node Readable Stream
 */
proto.proto.addNumServiceClient.prototype.getRooms =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/proto.addNumService/getRooms',
      request,
      metadata || {},
      methodDescriptor_addNumService_getRooms,
      callback);
};


/**
 * @param {!proto.proto.Null} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.proto.roomList>}
 *     Promise that resolves to the response
 */
proto.proto.addNumServicePromiseClient.prototype.getRooms =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/proto.addNumService/getRooms',
      request,
      metadata || {},
      methodDescriptor_addNumService_getRooms);
};


module.exports = proto.proto;

