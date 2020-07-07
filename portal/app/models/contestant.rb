require 'isuxportal/resources/contestant_pb'

class Contestant < ApplicationRecord
  belongs_to :team
  has_many :ssh_public_keys

  validates :name, presence: true
  validates :github_id, presence: true, uniqueness: true
  validates :github_login, presence: true
  validates :discord_id, presence: true, uniqueness: true
  validates :discord_tag, presence: true
  validates :avatar_url, presence: true

  validate :validate_number_of_team_members

  def validate_number_of_team_members
    return if persisted?
    unless team&.joinable?
      errors.add :team_id, 'チームメンバーの上限に達しています'
    end
  end

  def to_pb(detail: false)
    Isuxportal::Proto::Resources::Contestant.new(
      id: id,
      team_id: team_id,
      name: name,
      contestant_detail: !detail ? nil : Isuxportal::Proto::Resources::Contestant::ContestantDetail.new(
        avatar_url: avatar_url,
        github_login: github_login,
        discord_tag: discord_tag,
        is_student: student,
      ),
    )
  end
end