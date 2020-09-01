# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: isuxportal/misc/bypass_token.proto

require 'google/protobuf'

Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("isuxportal/misc/bypass_token.proto", :syntax => :proto3) do
    add_message "isuxportal.proto.misc.BypassTokenPayload" do
      optional :filler, :string, 1
      optional :expiry, :int64, 2
      repeated :usages, :enum, 3, "isuxportal.proto.misc.BypassTokenPayload.Usage"
    end
    add_enum "isuxportal.proto.misc.BypassTokenPayload.Usage" do
      value :CREATE_TEAM, 0
      value :JOIN_TEAM, 1
    end
  end
end

module Isuxportal
  module Proto
    module Misc
      BypassTokenPayload = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.misc.BypassTokenPayload").msgclass
      BypassTokenPayload::Usage = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.misc.BypassTokenPayload.Usage").enummodule
    end
  end
end