package twitchws

import (
	"github.com/vpetrigo/go-twitch-ws/pkg/eventsub"
)

type eventSubScope struct {
	Version       string
	MsgType       interface{}
	ConditionType interface{}
}

var (
	eventSubTypes = map[string][]eventSubScope{
		"automod.message.hold": {
			{Version: "1", MsgType: nil, ConditionType: nil},
			{Version: "2", MsgType: nil, ConditionType: nil},
		},
		"automod.message.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
			{Version: "2", MsgType: nil, ConditionType: nil},
		},
		"automod.settings.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"automod.terms.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.ad_break.begin": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.ban": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.channel_points_automatic_reward_redemption.add": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.channel_points_custom_reward.add": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.channel_points_custom_reward.remove": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.channel_points_custom_reward.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.channel_points_custom_reward_redemption.add": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.channel_points_custom_reward_redemption.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.charity_campaign.donate": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.charity_campaign.progress": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.charity_campaign.start": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.charity_campaign.stop": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat.clear": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat.clear_user_messages": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat.message": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat.message_delete": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat.notification": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat.user_message_hold": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat.user_message_update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.chat_settings.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.cheer": {
			{Version: "1", MsgType: &eventsub.ChannelCheerEvent{}, ConditionType: &EventsubCondition{}},
		},
		"channel.follow": {
			{Version: "2", MsgType: &eventsub.ChannelFollowEvent{}, ConditionType: &EventsubCondition{}},
		},
		"channel.goal.begin": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.goal.end": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.goal.progress": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.guest_star_guest.update": {
			{Version: "beta", MsgType: nil, ConditionType: nil},
		},
		"channel.guest_star_session.begin": {
			{Version: "beta", MsgType: nil, ConditionType: nil},
		},
		"channel.guest_star_session.end": {
			{Version: "beta", MsgType: nil, ConditionType: nil},
		},
		"channel.guest_star_settings.update": {
			{Version: "beta", MsgType: nil, ConditionType: nil},
		},
		"channel.hype_train.begin": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.hype_train.end": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.hype_train.progress": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.moderate": {
			{Version: "1", MsgType: nil, ConditionType: nil},
			{Version: "2", MsgType: nil, ConditionType: nil},
		},
		"channel.moderator.add": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.moderator.remove": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.poll.begin": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.poll.end": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.poll.progress": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.prediction.begin": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.prediction.end": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.prediction.lock": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.prediction.progress": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.raid": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.shared_chat.begin": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.shared_chat.end": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.shared_chat.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.shield_mode.begin": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.shield_mode.end": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.shoutout.create": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.shoutout.receive": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.subscribe": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.subscription.end": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.subscription.gift": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.subscription.message": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.suspicious_user.message": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.suspicious_user.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.unban": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.unban_request.create": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.unban_request.resolve": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.update": {
			{Version: "2", MsgType: nil, ConditionType: nil},
		},
		"channel.vip.add": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.vip.remove": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.warning.acknowledge": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"channel.warning.send": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"conduit.shard.disabled": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"drop.entitlement.grant": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"extension.bits_transaction.create": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"stream.offline": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"stream.online": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"user.authorization.grant": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"user.authorization.revoke": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"user.update": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
		"user.whisper.message": {
			{Version: "1", MsgType: nil, ConditionType: nil},
		},
	}
)
