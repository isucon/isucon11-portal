// package: isuxportal.proto.services.common
// file: isuxportal/services/common/me.proto

import * as jspb from "google-protobuf";
import * as isuxportal_resources_team_pb from "../../../isuxportal/resources/team_pb";
import * as isuxportal_resources_contestant_pb from "../../../isuxportal/resources/contestant_pb";

export class GetCurrentSessionRequest extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCurrentSessionRequest.AsObject;
  static toObject(includeInstance: boolean, msg: GetCurrentSessionRequest): GetCurrentSessionRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetCurrentSessionRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCurrentSessionRequest;
  static deserializeBinaryFromReader(message: GetCurrentSessionRequest, reader: jspb.BinaryReader): GetCurrentSessionRequest;
}

export namespace GetCurrentSessionRequest {
  export type AsObject = {
  }
}

export class GetCurrentSessionResponse extends jspb.Message {
  hasTeam(): boolean;
  clearTeam(): void;
  getTeam(): isuxportal_resources_team_pb.Team | undefined;
  setTeam(value?: isuxportal_resources_team_pb.Team): void;

  hasContestant(): boolean;
  clearContestant(): void;
  getContestant(): isuxportal_resources_contestant_pb.Contestant | undefined;
  setContestant(value?: isuxportal_resources_contestant_pb.Contestant): void;

  getDiscordInvitationUrl(): string;
  setDiscordInvitationUrl(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetCurrentSessionResponse.AsObject;
  static toObject(includeInstance: boolean, msg: GetCurrentSessionResponse): GetCurrentSessionResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: GetCurrentSessionResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetCurrentSessionResponse;
  static deserializeBinaryFromReader(message: GetCurrentSessionResponse, reader: jspb.BinaryReader): GetCurrentSessionResponse;
}

export namespace GetCurrentSessionResponse {
  export type AsObject = {
    team?: isuxportal_resources_team_pb.Team.AsObject,
    contestant?: isuxportal_resources_contestant_pb.Contestant.AsObject,
    discordInvitationUrl: string,
  }
}

