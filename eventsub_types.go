package twitchws

type eventSubScope struct {
	Version       string
	MsgType       interface{}
	ConditionType interface{}
}

var (
	eventSubTypes = map[string]eventSubScope{
		"channel.update":                                         {Version: "2", MsgType: &ChannelUpdateEvent{}, ConditionType: &ChannelUpdateEventCondition{}},
		"channel.follow":                                         {Version: "2", MsgType: &ChannelFollowEvent{}, ConditionType: &ChannelFollowEventCondition{}},
		"channel.subscribe":                                      {Version: "1", MsgType: &ChannelSubscribeEvent{}, ConditionType: &ChannelSubscribeEventCondition{}},
		"channel.subscription.end":                               {Version: "1", MsgType: &ChannelSubscriptionEndEvent{}, ConditionType: &ChannelSubscriptionEndEventCondition{}},
		"channel.subscription.gift":                              {Version: "1", MsgType: &ChannelSubscriptionGiftEvent{}, ConditionType: &ChannelSubscriptionGiftEventCondition{}},
		"channel.subscription.message":                           {Version: "1", MsgType: &ChannelSubscriptionMessageEvent{}, ConditionType: &ChannelSubscriptionMessageEventCondition{}},
		"channel.cheer":                                          {Version: "1", MsgType: &ChannelCheerEvent{}, ConditionType: &ChannelCheerEventCondition{}},
		"channel.raid":                                           {Version: "1", MsgType: &ChannelRaidEvent{}, ConditionType: &ChannelRaidEventCondition{}},
		"channel.ban":                                            {Version: "1", MsgType: &ChannelBanEvent{}, ConditionType: &ChannelBanEventCondition{}},
		"channel.unban":                                          {Version: "1", MsgType: &ChannelUnbanEvent{}, ConditionType: &ChannelUnbanEventCondition{}},
		"channel.moderator.add":                                  {Version: "1", MsgType: &ChannelModeratorAddEvent{}, ConditionType: &ChannelModeratorAddEventCondition{}},
		"channel.moderator.remove":                               {Version: "1", MsgType: &ChannelModeratorRemoveEvent{}, ConditionType: &ChannelModeratorRemoveEventCondition{}},
		"channel.guest_star_session.begin":                       {Version: "beta", MsgType: &ChannelGuestStarSessionBeginEvent{}, ConditionType: &ChannelGuestStarSessionBeginEventCondition{}},
		"channel.guest_star_session.end":                         {Version: "beta", MsgType: &ChannelGuestStarSessionEndEvent{}, ConditionType: &ChannelGuestStarSessionEndEventCondition{}},
		"channel.guest_star_guest.update":                        {Version: "beta", MsgType: &ChannelGuestStarGuestUpdateEvent{}, ConditionType: &ChannelGuestStarGuestUpdateEventCondition{}},
		"channel.guest_star_slot.update":                         {Version: "beta", MsgType: &ChannelGuestStarSlotUpdateEvent{}, ConditionType: &ChannelGuestStarSlotUpdateEventCondition{}},
		"channel.guest_star_settings.update":                     {Version: "beta", MsgType: &ChannelGuestStarSettingsUpdateEvent{}, ConditionType: &ChannelGuestStarSettingsUpdateEventCondition{}},
		"channel.channel_points_custom_reward.add":               {Version: "1", MsgType: &ChannelPointsCustomRewardAddEvent{}, ConditionType: &ChannelPointsCustomRewardAddEventCondition{}},
		"channel.channel_points_custom_reward.update":            {Version: "1", MsgType: &ChannelPointsCustomRewardUpdateEvent{}, ConditionType: &ChannelPointsCustomRewardUpdateEventCondition{}},
		"channel.channel_points_custom_reward.remove":            {Version: "1", MsgType: &ChannelPointsCustomRewardRemoveEvent{}, ConditionType: &ChannelPointsCustomRewardRemoveEventCondition{}},
		"channel.channel_points_custom_reward_redemption.add":    {Version: "1", MsgType: &ChannelPointsCustomRewardRedemptionAddEvent{}, ConditionType: &ChannelPointsCustomRewardRedemptionAddEventCondition{}},
		"channel.channel_points_custom_reward_redemption.update": {Version: "1", MsgType: &ChannelPointsCustomRewardRedemptionUpdateEvent{}, ConditionType: &ChannelPointsCustomRewardRedemptionUpdateEventCondition{}},
		"channel.poll.begin":                                     {Version: "1", MsgType: &ChannelPollBeginEvent{}, ConditionType: &ChannelPollBeginEventCondition{}},
		"channel.poll.progress":                                  {Version: "1", MsgType: &ChannelPollProgressEvent{}, ConditionType: &ChannelPollProgressEventCondition{}},
		"channel.poll.end":                                       {Version: "1", MsgType: &ChannelPollEndEvent{}, ConditionType: &ChannelPollEndEventCondition{}},
		"channel.prediction.begin":                               {Version: "1", MsgType: &ChannelPredictionBeginEvent{}, ConditionType: &ChannelPredictionBeginEventCondition{}},
		"channel.prediction.progress":                            {Version: "1", MsgType: &ChannelPredictionProgressEvent{}, ConditionType: &ChannelPredictionProgressEventCondition{}},
		"channel.prediction.lock":                                {Version: "1", MsgType: &ChannelPredictionLockEvent{}, ConditionType: &ChannelPredictionLockEventCondition{}},
		"channel.prediction.end":                                 {Version: "1", MsgType: &ChannelPredictionEndEvent{}, ConditionType: &ChannelPredictionEndEventCondition{}},
		"channel.charity_campaign.donate":                        {Version: "1", MsgType: &CharityDonationEvent{}, ConditionType: &CharityDonationEventCondition{}},
		"channel.charity_campaign.start":                         {Version: "1", MsgType: &CharityCampaignStartEvent{}, ConditionType: &CharityCampaignStartEventCondition{}},
		"channel.charity_campaign.progress":                      {Version: "1", MsgType: &CharityCampaignProgressEvent{}, ConditionType: &CharityCampaignProgressEventCondition{}},
		"channel.charity_campaign.stop":                          {Version: "1", MsgType: &CharityCampaignStopEvent{}, ConditionType: &CharityCampaignStopEventCondition{}},
		"drop.entitlement.grant":                                 {Version: "1", MsgType: &DropEntitlementGrantEvent{}, ConditionType: &DropEntitlementGrantEventCondition{}},
		"extension.bits_transaction.create":                      {Version: "1", MsgType: &ExtensionBitsTransactionCreateEvent{}, ConditionType: &ExtensionBitsTransactionCreateEventCondition{}},
		"channel.goal.begin":                                     {Version: "1", MsgType: &GoalsEvent{}, ConditionType: &GoalsEventCondition{}},
		"channel.goal.progress":                                  {Version: "1", MsgType: &GoalsEvent{}, ConditionType: &GoalsEventCondition{}},
		"channel.goal.end":                                       {Version: "1", MsgType: &GoalsEvent{}, ConditionType: &GoalsEventCondition{}},
		"channel.hype_train.begin":                               {Version: "1", MsgType: &HypeTrainBeginEvent{}, ConditionType: &HypeTrainBeginEventCondition{}},
		"channel.hype_train.progress":                            {Version: "1", MsgType: &HypeTrainProgressEvent{}, ConditionType: &HypeTrainProgressEventCondition{}},
		"channel.hype_train.end":                                 {Version: "1", MsgType: &HypeTrainEndEvent{}, ConditionType: &HypeTrainEndEventCondition{}},
		"channel.shield_mode.begin":                              {Version: "1", MsgType: &ShieldModeBeginEvent{}, ConditionType: &ShieldModeBeginEventCondition{}},
		"channel.shield_mode.end":                                {Version: "1", MsgType: &ShieldModeEndEvent{}, ConditionType: &ShieldModeEndEventCondition{}},
		"channel.shoutout.create":                                {Version: "1", MsgType: &ShoutoutCreateEvent{}, ConditionType: &ShoutoutCreateEventCondition{}},
		"channel.shoutout.receive":                               {Version: "1", MsgType: &ShoutoutReceivedEvent{}, ConditionType: &ShoutoutReceivedEventCondition{}},
		"stream.online":                                          {Version: "1", MsgType: &StreamOnlineEvent{}, ConditionType: &StreamOnlineEventCondition{}},
		"stream.offline":                                         {Version: "1", MsgType: &StreamOfflineEvent{}, ConditionType: &StreamOfflineEventCondition{}},
		"user.authorization.grant":                               {Version: "1", MsgType: &UserAuthorizationGrantEvent{}, ConditionType: &UserAuthorizationGrantEventCondition{}},
		"user.authorization.revoke":                              {Version: "1", MsgType: &UserAuthorizationRevokeEvent{}, ConditionType: &UserAuthorizationRevokeEventCondition{}},
		"user.update":                                            {Version: "1", MsgType: &UserUpdateEvent{}, ConditionType: &UserUpdateEventCondition{}},
	}
)
