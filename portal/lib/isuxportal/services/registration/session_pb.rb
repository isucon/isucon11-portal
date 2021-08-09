# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: isuxportal/services/registration/session.proto

require 'google/protobuf'

require 'isuxportal/resources/team_pb'
require 'isuxportal/resources/coupon_pb'
Google::Protobuf::DescriptorPool.generated_pool.build do
  add_file("isuxportal/services/registration/session.proto", :syntax => :proto3) do
    add_message "isuxportal.proto.services.registration.GetRegistrationSessionQuery" do
      optional :team_id, :int64, 1
      optional :invite_token, :string, 2
      optional :bypass_token, :string, 16
    end
    add_message "isuxportal.proto.services.registration.GetRegistrationSessionResponse" do
      optional :team, :message, 1, "isuxportal.proto.resources.Team"
      optional :status, :enum, 2, "isuxportal.proto.services.registration.GetRegistrationSessionResponse.Status"
      optional :github_login, :string, 3
      optional :github_avatar_url, :string, 4
      optional :discord_tag, :string, 5
      optional :discord_avatar_url, :string, 6
      optional :member_invite_url, :string, 7
      optional :discord_server_id, :string, 8
      optional :coupon, :message, 9, "isuxportal.proto.resources.Coupon"
    end
    add_enum "isuxportal.proto.services.registration.GetRegistrationSessionResponse.Status" do
      value :CLOSED, 0
      value :NOT_JOINABLE, 1
      value :NOT_LOGGED_IN, 2
      value :CREATABLE, 3
      value :JOINABLE, 4
      value :JOINED, 5
      value :DISQUALIFIED, 6
    end
    add_message "isuxportal.proto.services.registration.UpdateRegistrationRequest" do
      optional :team_name, :string, 1
      optional :name, :string, 2
      optional :email_address, :string, 3
      optional :is_student, :bool, 4
    end
    add_message "isuxportal.proto.services.registration.UpdateRegistrationResponse" do
    end
    add_message "isuxportal.proto.services.registration.DeleteRegistrationRequest" do
    end
    add_message "isuxportal.proto.services.registration.DeleteRegistrationResponse" do
    end
  end
end

module Isuxportal
  module Proto
    module Services
      module Registration
        GetRegistrationSessionQuery = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.services.registration.GetRegistrationSessionQuery").msgclass
        GetRegistrationSessionResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.services.registration.GetRegistrationSessionResponse").msgclass
        GetRegistrationSessionResponse::Status = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.services.registration.GetRegistrationSessionResponse.Status").enummodule
        UpdateRegistrationRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.services.registration.UpdateRegistrationRequest").msgclass
        UpdateRegistrationResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.services.registration.UpdateRegistrationResponse").msgclass
        DeleteRegistrationRequest = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.services.registration.DeleteRegistrationRequest").msgclass
        DeleteRegistrationResponse = ::Google::Protobuf::DescriptorPool.generated_pool.lookup("isuxportal.proto.services.registration.DeleteRegistrationResponse").msgclass
      end
    end
  end
end
