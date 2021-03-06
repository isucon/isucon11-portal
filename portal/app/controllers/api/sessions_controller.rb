require 'isuxportal/services/common/me_pb'

class Api::SessionsController < Api::ApplicationController
  skip_before_action :require_staff_when_always_required # for broadcast view

  def show
    render protobuf: Isuxportal::Proto::Services::Common::GetCurrentSessionResponse.new(
      contestant: current_contestant&.to_pb(detail: true),
      team: current_team&.to_pb,
      discord_server_id: current_contestant ? Rails.application.config.x.discord.server_id : "",
      contest: Contest.to_pb(team: current_team),
      contestant_instances: (Contest.contest_open_for_team?(team: current_team) ? current_team&.contestant_instances&.order(number: :asc)&.map(&:to_pb) : nil) || [],
      push_vapid_key: Rails.application.config.x.webpush.vapid_key_public,
    )
  end
end
