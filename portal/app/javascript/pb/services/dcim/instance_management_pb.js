// source: services/dcim/instance_management.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var resources_contestant_instance_pb = require('../../resources/contestant_instance_pb.js');
goog.object.extend(proto, resources_contestant_instance_pb);
goog.exportSymbol('proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest', null, global);
goog.exportSymbol('proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.displayName = 'proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.displayName = 'proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.toObject = function(opt_includeInstance) {
  return proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.toObject = function(includeInstance, msg) {
  var f, obj = {
    token: jspb.Message.getFieldWithDefault(msg, 1, ""),
    instance: (f = msg.getInstance()) && resources_contestant_instance_pb.ContestantInstance.toObject(includeInstance, f)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest;
  return proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setToken(value);
      break;
    case 2:
      var value = new resources_contestant_instance_pb.ContestantInstance;
      reader.readMessage(value,resources_contestant_instance_pb.ContestantInstance.deserializeBinaryFromReader);
      msg.setInstance(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getToken();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getInstance();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      resources_contestant_instance_pb.ContestantInstance.serializeBinaryToWriter
    );
  }
};


/**
 * optional string token = 1;
 * @return {string}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.getToken = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest} returns this
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.setToken = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional isuxportal.proto.resources.ContestantInstance instance = 2;
 * @return {?proto.isuxportal.proto.resources.ContestantInstance}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.getInstance = function() {
  return /** @type{?proto.isuxportal.proto.resources.ContestantInstance} */ (
    jspb.Message.getWrapperField(this, resources_contestant_instance_pb.ContestantInstance, 2));
};


/**
 * @param {?proto.isuxportal.proto.resources.ContestantInstance|undefined} value
 * @return {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest} returns this
*/
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.setInstance = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest} returns this
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.clearInstance = function() {
  return this.setInstance(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateRequest.prototype.hasInstance = function() {
  return jspb.Message.getField(this, 2) != null;
};





if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.prototype.toObject = function(opt_includeInstance) {
  return proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.toObject = function(includeInstance, msg) {
  var f, obj = {

  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse;
  return proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.isuxportal.proto.services.dcim.InformInstanceStateUpdateResponse.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};


goog.object.extend(exports, proto.isuxportal.proto.services.dcim);
