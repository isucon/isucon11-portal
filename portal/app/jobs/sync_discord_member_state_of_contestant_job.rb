class SyncDiscordMemberStateOfContestantJob < ApplicationJob
  self.log_arguments = false
  def perform(contestant, force_is_member = nil)
    Rails.logger.info "Syncing discord guild member state of discord user #{contestant.discord_tag} (#{contestant.discord_id}), contestant_id=#{contestant.id}"

    is_member = force_is_member
    if is_member.nil?
      is_member = begin
        Discordrb::API::Server.resolve_member(
          discord_token,
          server_id,
          contestant.discord_id
        )
      rescue RestClient::NotFound => e
        false
      end
    end

    contestant.update_attributes!(
      is_discord_guild_member: is_member,
    )
  end

  private def discord_token
    @discord_token ||= "Bot #{Rails.application.config.x.discord.bot_token}"
  end

  private def server_id
    Rails.application.config.x.discord.server_id
  end
end
