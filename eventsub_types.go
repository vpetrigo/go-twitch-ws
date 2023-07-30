package twitchws

import "github.com/vpetrigo/go-twitch-ws/pkg/eventsub"

type eventSubScope struct {
	Version       string
	MsgType       interface{}
	ConditionType interface{}
}

var (
	eventSubTypes = map[string]eventSubScope{
		"channel.update":                                         {Version: "2", MsgType: &eventsub.ChannelUpdateEvent{}, ConditionType: &eventsub.ChannelUpdateEventCondition{}},
		"channel.follow":                                         {Version: "2", MsgType: &eventsub.ChannelFollowEvent{}, ConditionType: &eventsub.ChannelFollowEventCondition{}},
		"channel.subscribe":                                      {Version: "1", MsgType: &eventsub.ChannelSubscribeEvent{}, ConditionType: &eventsub.ChannelSubscribeEventCondition{}},
		"channel.subscription.end":                               {Version: "1", MsgType: &eventsub.ChannelSubscriptionEndEvent{}, ConditionType: &eventsub.ChannelSubscriptionEndEventCondition{}},
		"channel.subscription.gift":                              {Version: "1", MsgType: &eventsub.ChannelSubscriptionGiftEvent{}, ConditionType: &eventsub.ChannelSubscriptionGiftEventCondition{}},
		"channel.subscription.message":                           {Version: "1", MsgType: &eventsub.ChannelSubscriptionMessageEvent{}, ConditionType: &eventsub.ChannelSubscriptionMessageEventCondition{}},
		"channel.cheer":                                          {Version: "1", MsgType: &eventsub.ChannelCheerEvent{}, ConditionType: &eventsub.ChannelCheerEventCondition{}},
		"channel.raid":                                           {Version: "1", MsgType: &eventsub.ChannelRaidEvent{}, ConditionType: &eventsub.ChannelRaidEventCondition{}},
		"channel.ban":                                            {Version: "1", MsgType: &eventsub.ChannelBanEvent{}, ConditionType: &eventsub.ChannelBanEventCondition{}},
		"channel.unban":                                          {Version: "1", MsgType: &eventsub.ChannelUnbanEvent{}, ConditionType: &eventsub.ChannelUnbanEventCondition{}},
		"channel.moderator.add":                                  {Version: "1", MsgType: &eventsub.ChannelModeratorAddEvent{}, ConditionType: &eventsub.ChannelModeratorAddEventCondition{}},
		"channel.moderator.remove":                               {Version: "1", MsgType: &eventsub.ChannelModeratorRemoveEvent{}, ConditionType: &eventsub.ChannelModeratorRemoveEventCondition{}},
		"channel.guest_star_session.begin":                       {Version: "beta", MsgType: &eventsub.ChannelGuestStarSessionBeginEvent{}, ConditionType: &eventsub.ChannelGuestStarSessionBeginEventCondition{}},
		"channel.guest_star_session.end":                         {Version: "beta", MsgType: &eventsub.ChannelGuestStarSessionEndEvent{}, ConditionType: &eventsub.ChannelGuestStarSessionEndEventCondition{}},
		"channel.guest_star_guest.update":                        {Version: "beta", MsgType: &eventsub.ChannelGuestStarGuestUpdateEvent{}, ConditionType: &eventsub.ChannelGuestStarGuestUpdateEventCondition{}},
		"channel.guest_star_slot.update":                         {Version: "beta", MsgType: &eventsub.ChannelGuestStarSlotUpdateEvent{}, ConditionType: &eventsub.ChannelGuestStarSlotUpdateEventCondition{}},
		"channel.guest_star_settings.update":                     {Version: "beta", MsgType: &eventsub.ChannelGuestStarSettingsUpdateEvent{}, ConditionType: &eventsub.ChannelGuestStarSettingsUpdateEventCondition{}},
		"channel.channel_points_custom_reward.add":               {Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardAddEvent{}, ConditionType: &eventsub.ChannelPointsCustomRewardAddEventCondition{}},
		"channel.channel_points_custom_reward.update":            {Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardUpdateEvent{}, ConditionType: &eventsub.ChannelPointsCustomRewardUpdateEventCondition{}},
		"channel.channel_points_custom_reward.remove":            {Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardRemoveEvent{}, ConditionType: &eventsub.ChannelPointsCustomRewardRemoveEventCondition{}},
		"channel.channel_points_custom_reward_redemption.add":    {Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardRedemptionAddEvent{}, ConditionType: &eventsub.ChannelPointsCustomRewardRedemptionAddEventCondition{}},
		"channel.channel_points_custom_reward_redemption.update": {Version: "1", MsgType: &eventsub.ChannelPointsCustomRewardRedemptionUpdateEvent{}, ConditionType: &eventsub.ChannelPointsCustomRewardRedemptionUpdateEventCondition{}},
		"channel.poll.begin":                                     {Version: "1", MsgType: &eventsub.ChannelPollBeginEvent{}, ConditionType: &eventsub.ChannelPollBeginEventCondition{}},
		"channel.poll.progress":                                  {Version: "1", MsgType: &eventsub.ChannelPollProgressEvent{}, ConditionType: &eventsub.ChannelPollProgressEventCondition{}},
		"channel.poll.end":                                       {Version: "1", MsgType: &eventsub.ChannelPollEndEvent{}, ConditionType: &eventsub.ChannelPollEndEventCondition{}},
		"channel.prediction.begin":                               {Version: "1", MsgType: &eventsub.ChannelPredictionBeginEvent{}, ConditionType: &eventsub.ChannelPredictionBeginEventCondition{}},
		"channel.prediction.progress":                            {Version: "1", MsgType: &eventsub.ChannelPredictionProgressEvent{}, ConditionType: &eventsub.ChannelPredictionProgressEventCondition{}},
		"channel.prediction.lock":                                {Version: "1", MsgType: &eventsub.ChannelPredictionLockEvent{}, ConditionType: &eventsub.ChannelPredictionLockEventCondition{}},
		"channel.prediction.end":                                 {Version: "1", MsgType: &eventsub.ChannelPredictionEndEvent{}, ConditionType: &eventsub.ChannelPredictionEndEventCondition{}},
		"channel.charity_campaign.donate":                        {Version: "1", MsgType: &eventsub.CharityDonationEvent{}, ConditionType: &eventsub.CharityDonationEventCondition{}},
		"channel.charity_campaign.start":                         {Version: "1", MsgType: &eventsub.CharityCampaignStartEvent{}, ConditionType: &eventsub.CharityCampaignStartEventCondition{}},
		"channel.charity_campaign.progress":                      {Version: "1", MsgType: &eventsub.CharityCampaignProgressEvent{}, ConditionType: &eventsub.CharityCampaignProgressEventCondition{}},
		"channel.charity_campaign.stop":                          {Version: "1", MsgType: &eventsub.CharityCampaignStopEvent{}, ConditionType: &eventsub.CharityCampaignStopEventCondition{}},
		"drop.entitlement.grant":                                 {Version: "1", MsgType: &eventsub.DropEntitlementGrantEvent{}, ConditionType: &eventsub.DropEntitlementGrantEventCondition{}},
		"extension.bits_transaction.create":                      {Version: "1", MsgType: &eventsub.ExtensionBitsTransactionCreateEvent{}, ConditionType: &eventsub.ExtensionBitsTransactionCreateEventCondition{}},
		"channel.goal.begin":                                     {Version: "1", MsgType: &eventsub.GoalsEvent{}, ConditionType: &eventsub.GoalsEventCondition{}},
		"channel.goal.progress":                                  {Version: "1", MsgType: &eventsub.GoalsEvent{}, ConditionType: &eventsub.GoalsEventCondition{}},
		"channel.goal.end":                                       {Version: "1", MsgType: &eventsub.GoalsEvent{}, ConditionType: &eventsub.GoalsEventCondition{}},
		"channel.hype_train.begin":                               {Version: "1", MsgType: &eventsub.HypeTrainBeginEvent{}, ConditionType: &eventsub.HypeTrainBeginEventCondition{}},
		"channel.hype_train.progress":                            {Version: "1", MsgType: &eventsub.HypeTrainProgressEvent{}, ConditionType: &eventsub.HypeTrainProgressEventCondition{}},
		"channel.hype_train.end":                                 {Version: "1", MsgType: &eventsub.HypeTrainEndEvent{}, ConditionType: &eventsub.HypeTrainEndEventCondition{}},
		"channel.shield_mode.begin":                              {Version: "1", MsgType: &eventsub.ShieldModeBeginEvent{}, ConditionType: &eventsub.ShieldModeBeginEventCondition{}},
		"channel.shield_mode.end":                                {Version: "1", MsgType: &eventsub.ShieldModeEndEvent{}, ConditionType: &eventsub.ShieldModeEndEventCondition{}},
		"channel.shoutout.create":                                {Version: "1", MsgType: &eventsub.ShoutoutCreateEvent{}, ConditionType: &eventsub.ShoutoutCreateEventCondition{}},
		"channel.shoutout.receive":                               {Version: "1", MsgType: &eventsub.ShoutoutReceivedEvent{}, ConditionType: &eventsub.ShoutoutReceivedEventCondition{}},
		"stream.online":                                          {Version: "1", MsgType: &eventsub.StreamOnlineEvent{}, ConditionType: &eventsub.StreamOnlineEventCondition{}},
		"stream.offline":                                         {Version: "1", MsgType: &eventsub.StreamOfflineEvent{}, ConditionType: &eventsub.StreamOfflineEventCondition{}},
		"user.authorization.grant":                               {Version: "1", MsgType: &eventsub.UserAuthorizationGrantEvent{}, ConditionType: &eventsub.UserAuthorizationGrantEventCondition{}},
		"user.authorization.revoke":                              {Version: "1", MsgType: &eventsub.UserAuthorizationRevokeEvent{}, ConditionType: &eventsub.UserAuthorizationRevokeEventCondition{}},
		"user.update":                                            {Version: "1", MsgType: &eventsub.UserUpdateEvent{}, ConditionType: &eventsub.UserUpdateEventCondition{}},
	}
)
