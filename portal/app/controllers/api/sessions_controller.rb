require 'isuxportal/services/common/me_pb'

class Api::SessionsController < Api::ApplicationController
  def show
    render protobuf: Isuxportal::Proto::Services::Common::GetCurrentSessionResponse.new(
      contestant: current_contestant&.to_pb,
      team: current_team&.to_pb,
      discord_server_id: current_contestant ? Rails.application.config.x.discord.server_id : "",
    )
  end
end