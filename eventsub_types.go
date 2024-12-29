package twitchws

import "github.com/vpetrigo/go-twitch-ws/pkg/eventsub"

type eventSubScope struct {
	Version       string
	MsgType       interface{}
	ConditionType interface{}
}

var (
	eventSubTypes = map[string][]eventSubScope{
		"automod.message.hold": {
			{Version: "1", MsgType: nil},
			{Version: "2", MsgType: nil},
		},
		"automod.message.update": {
			{Version: "1", MsgType: nil},
			{Version: "2", MsgType: nil},
		},
		"automod.settings.update": {
			{Version: "1", MsgType: &eventsub.AutomodSettingsUpdateEvent{}},
		},
		"automod.terms.update": {
			{Version: "1", MsgType: &eventsub.AutomodTermsUpdateEvent{}},
		},
		"channel.ad_break.begin": {
			{Version: "1", MsgType: &eventsub.ChannelAdBreakBeginEvent{}},
		},
		"channel.ban": {
			{Version: "1", MsgType: &eventsub.ChannelBanEvent{}},
		},
		"channel.channel_points_automatic_reward_redemption.add": {
			{Version: "1", MsgType: &eventsub.ChannelPointsAutomaticRewardRedemptionAddEvent{}},
		},
		"channel.channel_points_custom_reward.add": {
			{Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardAddEvent{}},
		},
		"channel.channel_points_custom_reward.remove": {
			{Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardRemoveEvent{}},
		},
		"channel.channel_points_custom_reward.update": {
			{Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardUpdateEvent{}},
		},
		"channel.channel_points_custom_reward_redemption.add": {
			{Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardRedemptionAddEvent{}},
		},
		"channel.channel_points_custom_reward_redemption.update": {
			{Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardRedemptionUpdateEvent{}},
		},
		"channel.charity_campaign.donate": {
			{Version: "1", MsgType: &eventsub.CharityDonationEvent{}},
		},
		"channel.charity_campaign.progress": {
			{Version: "1", MsgType: &eventsub.CharityCampaignProgressEvent{}},
		},
		"channel.charity_campaign.start": {
			{Version: "1", MsgType: &eventsub.CharityCampaignStartEvent{}},
		},
		"channel.charity_campaign.stop": {
			{Version: "1", MsgType: &eventsub.CharityCampaignStopEvent{}},
		},
		"channel.chat.clear": {
			{Version: "1", MsgType: &eventsub.ChannelChatClearEvent{}},
		},
		"channel.chat.clear_user_messages": {
			{Version: "1", MsgType: &eventsub.ChannelChatClearUserMessagesEvent{}},
		},
		"channel.chat.message": {
			{Version: "1", MsgType: nil},
		},
		"channel.chat.message_delete": {
			{Version: "1", MsgType: &eventsub.ChannelChatMessageDeleteEvent{}},
		},
		"channel.chat.notification": {
			{Version: "1", MsgType: nil},
		},
		"channel.chat.user_message_hold": {
			{Version: "1", MsgType: nil},
		},
		"channel.chat.user_message_update": {
			{Version: "1", MsgType: nil},
		},
		"channel.chat_settings.update": {
			{Version: "1", MsgType: &eventsub.ChannelChatSettingsUpdateEvent{}},
		},
		"channel.cheer": {
			{Version: "1", MsgType: &eventsub.ChannelCheerEvent{}},
		},
		"channel.follow": {
			{Version: "2", MsgType: &eventsub.ChannelFollowEvent{}},
		},
		"channel.goal.begin": {
			{Version: "1", MsgType: &eventsub.GoalsEvent{}},
		},
		"channel.goal.end": {
			{Version: "1", MsgType: &eventsub.GoalsEvent{}},
		},
		"channel.goal.progress": {
			{Version: "1", MsgType: &eventsub.GoalsEvent{}},
		},
		"channel.guest_star_guest.update": {
			{Version: "beta", MsgType: &eventsub.ChannelGuestStarGuestUpdateEvent{}},
		},
		"channel.guest_star_session.begin": {
			{Version: "beta", MsgType: &eventsub.ChannelGuestStarSessionBeginEvent{}},
		},
		"channel.guest_star_session.end": {
			{Version: "beta", MsgType: &eventsub.ChannelGuestStarSessionEndEvent{}},
		},
		"channel.guest_star_settings.update": {
			{Version: "beta", MsgType: &eventsub.ChannelGuestStarSettingsUpdateEvent{}},
		},
		"channel.hype_train.begin": {
			{Version: "1", MsgType: &eventsub.HypeTrainBeginEvent{}},
		},
		"channel.hype_train.end": {
			{Version: "1", MsgType: &eventsub.HypeTrainEndEvent{}},
		},
		"channel.hype_train.progress": {
			{Version: "1", MsgType: &eventsub.HypeTrainProgressEvent{}},
		},
		"channel.moderate": {
			{Version: "1", MsgType: &eventsub.ChannelModerateEvent{}},
			{Version: "2", MsgType: nil},
		},
		"channel.moderator.add": {
			{Version: "1", MsgType: &eventsub.ChannelModeratorAddEvent{}},
		},
		"channel.moderator.remove": {
			{Version: "1", MsgType: &eventsub.ChannelModeratorRemoveEvent{}},
		},
		"channel.poll.begin": {
			{Version: "1", MsgType: &eventsub.ChannelPollBeginEvent{}},
		},
		"channel.poll.end": {
			{Version: "1", MsgType: &eventsub.ChannelPollEndEvent{}},
		},
		"channel.poll.progress": {
			{Version: "1", MsgType: &eventsub.ChannelPollProgressEvent{}},
		},
		"channel.prediction.begin": {
			{Version: "1", MsgType: &eventsub.ChannelPredictionBeginEvent{}},
		},
		"channel.prediction.end": {
			{Version: "1", MsgType: &eventsub.ChannelPredictionEndEvent{}},
		},
		"channel.prediction.lock": {
			{Version: "1", MsgType: &eventsub.ChannelPredictionLockEvent{}},
		},
		"channel.prediction.progress": {
			{Version: "1", MsgType: &eventsub.ChannelPredictionProgressEvent{}},
		},
		"channel.raid": {
			{Version: "1", MsgType: &eventsub.ChannelRaidEvent{}},
		},
		"channel.shared_chat.begin": {
			{Version: "1", MsgType: &eventsub.ChannelSharedChatSessionBeginEvent{}},
		},
		"channel.shared_chat.end": {
			{Version: "1", MsgType: &eventsub.ChannelSharedChatSessionEndEvent{}},
		},
		"channel.shared_chat.update": {
			{Version: "1", MsgType: &eventsub.ChannelSharedChatSessionUpdateEvent{}},
		},
		"channel.shield_mode.begin": {
			{Version: "1", MsgType: &eventsub.ShieldModeEvent{}},
		},
		"channel.shield_mode.end": {
			{Version: "1", MsgType: &eventsub.ShieldModeEvent{}},
		},
		"channel.shoutout.create": {
			{Version: "1", MsgType: &eventsub.ShoutoutCreateEvent{}},
		},
		"channel.shoutout.receive": {
			{Version: "1", MsgType: &eventsub.ShoutoutReceivedEvent{}},
		},
		"channel.subscribe": {
			{Version: "1", MsgType: &eventsub.ChannelSubscribeEvent{}},
		},
		"channel.subscription.end": {
			{Version: "1", MsgType: &eventsub.ChannelSubscriptionEndEvent{}},
		},
		"channel.subscription.gift": {
			{Version: "1", MsgType: &eventsub.ChannelSubscriptionGiftEvent{}},
		},
		"channel.subscription.message": {
			{Version: "1", MsgType: &eventsub.ChannelSubscriptionMessageEvent{}},
		},
		"channel.suspicious_user.message": {
			{Version: "1", MsgType: &eventsub.ChannelSuspiciousUserMessageEvent{}},
		},
		"channel.suspicious_user.update": {
			{Version: "1", MsgType: &eventsub.ChannelSuspiciousUserUpdateEvent{}},
		},
		"channel.unban": {
			{Version: "1", MsgType: &eventsub.ChannelUnbanEvent{}},
		},
		"channel.unban_request.create": {
			{Version: "1", MsgType: &eventsub.ChannelUnbanRequestCreateEvent{}},
		},
		"channel.unban_request.resolve": {
			{Version: "1", MsgType: &eventsub.ChannelUnbanRequestResolveEvent{}},
		},
		"channel.update": {
			{Version: "2", MsgType: &eventsub.ChannelUpdateEvent{}},
		},
		"channel.vip.add": {
			{Version: "1", MsgType: &eventsub.ChannelVIPAddEvent{}},
		},
		"channel.vip.remove": {
			{Version: "1", MsgType: &eventsub.ChannelVIPRemoveEvent{}},
		},
		"channel.warning.acknowledge": {
			{Version: "1", MsgType: nil},
		},
		"channel.warning.send": {
			{Version: "1", MsgType: &eventsub.ChannelWarningSendEvent{}},
		},
		"conduit.shard.disabled": {
			{Version: "1", MsgType: &eventsub.ConduitShardDisabledEvent{}},
		},
		"drop.entitlement.grant": {
			{Version: "1", MsgType: &eventsub.DropEntitlementGrantEvent{}},
		},
		"extension.bits_transaction.create": {
			{Version: "1", MsgType: &eventsub.ExtensionBitsTransactionCreateEvent{}},
		},
		"stream.offline": {
			{Version: "1", MsgType: &eventsub.StreamOfflineEvent{}},
		},
		"stream.online": {
			{Version: "1", MsgType: &eventsub.StreamOnlineEvent{}},
		},
		"user.authorization.grant": {
			{Version: "1", MsgType: &eventsub.UserAuthorizationGrantEvent{}},
		},
		"user.authorization.revoke": {
			{Version: "1", MsgType: &eventsub.UserAuthorizationRevokeEvent{}},
		},
		"user.update": {
			{Version: "1", MsgType: &eventsub.UserUpdateEvent{}},
		},
		"user.whisper.message": {
			{Version: "1", MsgType: &eventsub.WhisperReceivedEvent{}},
		},
	}
)
