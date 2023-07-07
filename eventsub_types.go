package twitchws

type eventSubScope struct {
	Version       string
	MsgType       interface{}
	ConditionType interface{}
}

var (
	eventSubTypes = map[string]eventSubScope{
		"channel.update":                                         {Version: "beta", MsgType: &ChannelUpdateEvent{}, ConditionType: &ChannelUpdateCondition{}},
		"channel.follow":                                         {Version: "2", MsgType: &ChannelFollowEvent{}, ConditionType: &ChannelFollowCondition{}},
		"channel.subscribe":                                      {Version: "1", MsgType: &ChannelSubscribeEvent{}, ConditionType: &ChannelSubscribeCondition{}},
		"channel.subscription.end":                               {Version: "1", MsgType: &ChannelSubscriptionEndEvent{}, ConditionType: &ChannelSubscriptionEndCondition{}},
		"channel.subscription.gift":                              {Version: "1", MsgType: &ChannelSubscriptionGiftEvent{}, ConditionType: &ChannelSubscriptionGiftCondition{}},
		"channel.subscription.message":                           {Version: "1", MsgType: &ChannelSubscriptionMessageEvent{}, ConditionType: &ChannelSubscriptionMessageCondition{}},
		"channel.cheer":                                          {Version: "1", MsgType: &ChannelCheerEvent{}, ConditionType: &ChannelCheerCondition{}},
		"channel.raid":                                           {Version: "1", MsgType: &ChannelRaidEvent{}, ConditionType: &ChannelRaidCondition{}},
		"channel.ban":                                            {Version: "1", MsgType: &ChannelBanEvent{}, ConditionType: &ChannelBanCondition{}},
		"channel.unban":                                          {Version: "1", MsgType: &ChannelUnbanEvent{}, ConditionType: &ChannelUnbanCondition{}},
		"channel.moderator.add":                                  {Version: "1", MsgType: &ChannelModeratorAddEvent{}, ConditionType: &ChannelModeratorAddCondition{}},
		"channel.moderator.remove":                               {Version: "1", MsgType: &ChannelModeratorRemoveEvent{}, ConditionType: &ChannelModeratorRemoveCondition{}},
		"channel.guest_star_session.begin":                       {Version: "beta", MsgType: &ChannelGuestStarSessionBeginEvent{}, ConditionType: &ChannelGuestStarSessionBeginCondition{}},
		"channel.guest_star_session.end":                         {Version: "beta", MsgType: &ChannelGuestStarSessionEndEvent{}, ConditionType: &ChannelGuestStarSessionEndCondition{}},
		"channel.guest_star_guest.update":                        {Version: "beta", MsgType: &ChannelGuestStarGuestUpdateEvent{}, ConditionType: &ChannelGuestStarGuestUpdateCondition{}},
		"channel.guest_star_slot.update":                         {Version: "beta", MsgType: &ChannelGuestStarSlotUpdateEvent{}, ConditionType: &ChannelGuestStarSlotUpdateCondition{}},
		"channel.guest_star_settings.update":                     {Version: "beta", MsgType: &ChannelGuestStarSettingsUpdateEvent{}, ConditionType: &ChannelGuestStarSettingsUpdateCondition{}},
		"channel.channel_points_custom_reward.add":               {Version: "1", MsgType: &ChannelChannelPointsCustomRewardAddEvent{}, ConditionType: &ChannelChannelPointsCustomRewardAddCondition{}},
		"channel.channel_points_custom_reward.update":            {Version: "1", MsgType: &ChannelChannelPointsCustomRewardUpdateEvent{}, ConditionType: &ChannelChannelPointsCustomRewardUpdateCondition{}},
		"channel.channel_points_custom_reward.remove":            {Version: "1", MsgType: &ChannelChannelPointsCustomRewardRemoveEvent{}, ConditionType: &ChannelChannelPointsCustomRewardRemoveCondition{}},
		"channel.channel_points_custom_reward_redemption.add":    {Version: "1", MsgType: &ChannelChannelPointsCustomRewardRedemptionAddEvent{}, ConditionType: &ChannelChannelPointsCustomRewardRedemptionAddCondition{}},
		"channel.channel_points_custom_reward_redemption.update": {Version: "1", MsgType: &ChannelChannelPointsCustomRewardRedemptionUpdateEvent{}, ConditionType: &ChannelChannelPointsCustomRewardRedemptionUpdateCondition{}},
		"channel.poll.begin":                                     {Version: "1", MsgType: &ChannelPollBeginEvent{}, ConditionType: &ChannelPollBeginCondition{}},
		"channel.poll.progress":                                  {Version: "1", MsgType: &ChannelPollProgressEvent{}, ConditionType: &ChannelPollProgressCondition{}},
		"channel.poll.end":                                       {Version: "1", MsgType: &ChannelPollEndEvent{}, ConditionType: &ChannelPollEndCondition{}},
		"channel.prediction.begin":                               {Version: "1", MsgType: &ChannelPredictionBeginEvent{}, ConditionType: &ChannelPredictionBeginCondition{}},
		"channel.prediction.progress":                            {Version: "1", MsgType: &ChannelPredictionProgressEvent{}, ConditionType: &ChannelPredictionProgressCondition{}},
		"channel.prediction.lock":                                {Version: "1", MsgType: &ChannelPredictionLockEvent{}, ConditionType: &ChannelPredictionLockCondition{}},
		"channel.prediction.end":                                 {Version: "1", MsgType: &ChannelPredictionEndEvent{}, ConditionType: &ChannelPredictionEndCondition{}},
		"channel.charity_campaign.donate":                        {Version: "1", MsgType: &ChannelCharityCampaignDonateEvent{}, ConditionType: &ChannelCharityCampaignDonateCondition{}},
		"channel.charity_campaign.start":                         {Version: "1", MsgType: &ChannelCharityCampaignStartEvent{}, ConditionType: &ChannelCharityCampaignStartCondition{}},
		"channel.charity_campaign.progress":                      {Version: "1", MsgType: &ChannelCharityCampaignProgressEvent{}, ConditionType: &ChannelCharityCampaignProgressCondition{}},
		"channel.charity_campaign.stop":                          {Version: "1", MsgType: &ChannelCharityCampaignStopEvent{}, ConditionType: &ChannelCharityCampaignStopCondition{}},
		"drop.entitlement.grant":                                 {Version: "1", MsgType: &DropEntitlementGrantEvent{}, ConditionType: &DropEntitlementGrantCondition{}},
		"extension.bits_transaction.create":                      {Version: "1", MsgType: &ExtensionBitsTransactionCreateEvent{}, ConditionType: &ExtensionBitsTransactionCreateCondition{}},
		"channel.goal.begin":                                     {Version: "1", MsgType: &ChannelGoalBeginEvent{}, ConditionType: &ChannelGoalBeginCondition{}},
		"channel.goal.progress":                                  {Version: "1", MsgType: &ChannelGoalProgressEvent{}, ConditionType: &ChannelGoalProgressCondition{}},
		"channel.goal.end":                                       {Version: "1", MsgType: &ChannelGoalEndEvent{}, ConditionType: &ChannelGoalEndCondition{}},
		"channel.hype_train.begin":                               {Version: "1", MsgType: &ChannelHypeTrainBeginEvent{}, ConditionType: &ChannelHypeTrainBeginCondition{}},
		"channel.hype_train.progress":                            {Version: "1", MsgType: &ChannelHypeTrainProgressEvent{}, ConditionType: &ChannelHypeTrainProgressCondition{}},
		"channel.hype_train.end":                                 {Version: "1", MsgType: &ChannelHypeTrainEndEvent{}, ConditionType: &ChannelHypeTrainEndCondition{}},
		"channel.shield_mode.begin":                              {Version: "1", MsgType: &ChannelShieldModeBeginEvent{}, ConditionType: &ChannelShieldModeBeginCondition{}},
		"channel.shield_mode.end":                                {Version: "1", MsgType: &ChannelShieldModeEndEvent{}, ConditionType: &ChannelShieldModeEndCondition{}},
		"channel.shoutout.create":                                {Version: "1", MsgType: &ChannelShoutoutCreateEvent{}, ConditionType: &ChannelShoutoutCreateCondition{}},
		"channel.shoutout.receive":                               {Version: "1", MsgType: &ChannelShoutoutReceiveEvent{}, ConditionType: &ChannelShoutoutReceiveCondition{}},
		"stream.online":                                          {Version: "1", MsgType: &StreamOnlineEvent{}, ConditionType: &StreamOnlineCondition{}},
		"stream.offline":                                         {Version: "1", MsgType: &StreamOfflineEvent{}, ConditionType: &StreamOfflineCondition{}},
		"user.authorization.grant":                               {Version: "1", MsgType: &UserAuthorizationGrantEvent{}, ConditionType: &UserAuthorizationGrantCondition{}},
		"user.authorization.revoke":                              {Version: "1", MsgType: &UserAuthorizationRevokeEvent{}, ConditionType: &UserAuthorizationRevokeCondition{}},
		"user.update":                                            {Version: "1", MsgType: &UserUpdateEvent{}, ConditionType: &UserUpdateCondition{}},
	}
)
