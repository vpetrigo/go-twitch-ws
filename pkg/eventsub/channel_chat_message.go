package eventsub

type ChannelChatMessage struct {
	BroadcasterUserID           string         `json:"broadcaster_user_id"`
	BroadcasterUserName         string         `json:"broadcaster_user_name"`
	BroadcasterUserLogin        string         `json:"broadcaster_user_login"`
	ChatterUserID               string         `json:"chatter_user_id"`
	ChatterUserName             string         `json:"chatter_user_name"`
	ChatterUserLogin            string         `json:"chatter_user_login"`
	MessageID                   string         `json:"message_id"`
	Message                     ChatMessage    `json:"message"`
	MessageType                 string         `json:"message_type"`
	Badges                      []ChatBadge    `json:"badges"`
	Cheer                       *CheerMetadata `json:"cheer,omitempty"`
	Color                       string         `json:"color"`
	Reply                       *ReplyMetadata `json:"reply,omitempty"`
	ChannelPointsCustomRewardID *string        `json:"channel_points_custom_reward_id,omitempty"`
	SourceBroadcasterUserID     *string        `json:"source_broadcaster_user_id,omitempty"`
	SourceBroadcasterUserName   *string        `json:"source_broadcaster_user_name,omitempty"`
	SourceBroadcasterUserLogin  *string        `json:"source_broadcaster_user_login,omitempty"`
	SourceMessageID             *string        `json:"source_message_id,omitempty"`
	SourceBadges                *ChatBadge     `json:"source_badges,omitempty"`
}

type ChatMessage struct {
	Text      string            `json:"text"`
	Fragments []MessageFragment `json:"fragments"`
}

type MessageFragment struct {
	Type      string             `json:"type"`
	Text      string             `json:"text"`
	Cheermote *CheermoteMetadata `json:"cheermote,omitempty"`
	Emote     *EmoteMetadata     `json:"emote,omitempty"`
	Mention   *MentionMetadata   `json:"mention,omitempty"`
}

type CheermoteMetadata struct {
	Prefix string `json:"prefix"`
	Bits   int    `json:"bits"`
	Tier   int    `json:"tier"`
}

type EmoteMetadata struct {
	ID         string   `json:"id"`
	EmoteSetID string   `json:"emote_set_id"`
	OwnerID    string   `json:"owner_id"`
	Format     []string `json:"format"`
}

type MentionMetadata struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	UserLogin string `json:"user_login"`
}

type ChatBadge struct {
	SetID string `json:"set_id"`
	ID    string `json:"id"`
	Info  string `json:"info"`
}

type CheerMetadata struct {
	Bits int `json:"bits"`
}

type ReplyMetadata struct {
	ParentMessageID   string `json:"parent_message_id"`
	ParentMessageBody string `json:"parent_message_body"`
	ParentUserID      string `json:"parent_user_id"`
	ParentUserName    string `json:"parent_user_name"`
	ParentUserLogin   string `json:"parent_user_login"`
	ThreadMessageID   string `json:"thread_message_id"`
	ThreadUserID      string `json:"thread_user_id"`
	ThreadUserName    string `json:"thread_user_name"`
	ThreadUserLogin   string `json:"thread_user_login"`
}

// type ChannelChatMessage struct {
// 	BroadcasterUserID          string            `json:"broadcaster_user_id"`
// 	BroadcasterUserName        string            `json:"broadcaster_user_name"`
// 	BroadcasterUserLogin       string            `json:"broadcaster_user_login"`
// 	ChatterUserID              string            `json:"chatter_user_id"`
// 	ChatterUserName            string            `json:"chatter_user_name"`
// 	ChatterIsAnonymous         bool              `json:"chatter_is_anonymous"`
// 	Color                      string            `json:"color"`
// 	Badges                     []Badge           `json:"badges"`
// 	SystemMessage              string            `json:"system_message"`
// 	MessageID                  string            `json:"message_id"`
// 	Message                    ChatMessage       `json:"message"`
// 	NoticeType                 string            `json:"notice_type"`
// 	Sub                        *Sub              `json:"sub,omitempty"`
// 	Resub                      *Resub            `json:"resub,omitempty"`
// 	SubGift                    *SubGift          `json:"sub_gift,omitempty"`
// 	CommunitySubGift           *CommunitySubGift `json:"community_sub_gift,omitempty"`
// 	GiftPaidUpgrade            *GiftPaidUpgrade  `json:"gift_paid_upgrade,omitempty"`
// 	PrimePaidUpgrade           *PrimePaidUpgrade `json:"prime_paid_upgrade,omitempty"`
// 	PayItForward               *PayItForward     `json:"pay_it_forward,omitempty"`
// 	Raid                       *Raid             `json:"raid,omitempty"`
// 	Unraid                     *json.RawMessage  `json:"unraid,omitempty"`
// 	Announcement               *Announcement     `json:"announcement,omitempty"`
// 	BitsBadgeTier              *BitsBadgeTier    `json:"bits_badge_tier,omitempty"`
// 	CharityDonation            *CharityDonation  `json:"charity_donation,omitempty"`
// 	SourceBroadcasterUserID    *string           `json:"source_broadcaster_user_id,omitempty"`
// 	SourceBroadcasterUserName  *string           `json:"source_broadcaster_user_name,omitempty"`
// 	SourceBroadcasterUserLogin *string           `json:"source_broadcaster_user_login,omitempty"`
// 	SourceMessageID            *string           `json:"source_message_id,omitempty"`
// 	SourceBadges               *Badge            `json:"source_badges,omitempty"`
// 	SharedChatSub              *Sub              `json:"shared_chat_sub,omitempty"`
// 	SharedChatResub            *Resub            `json:"shared_chat_resub,omitempty"`
// 	SharedChatSubGift          *SubGift          `json:"shared_chat_sub_gift,omitempty"`
// 	SharedChatCommunitySubGift *CommunitySubGift `json:"shared_chat_community_sub_gift,omitempty"`
// 	SharedChatGiftPaidUpgrade  *GiftPaidUpgrade  `json:"shared_chat_gift_paid_upgrade,omitempty"`
// 	SharedChatPrimePaidUpgrade *PrimePaidUpgrade `json:"shared_chat_prime_paid_upgrade,omitempty"`
// 	SharedChatPayItForward     *PayItForward     `json:"shared_chat_pay_it_forward,omitempty"`
// 	SharedChatRaid             *Raid             `json:"shared_chat_raid,omitempty"`
// 	SharedChatAnnouncement     *Announcement     `json:"shared_chat_announcement,omitempty"`
// }
//
// type Badge struct {
// 	SetID string `json:"set_id"`
// 	ID    string `json:"id"`
// 	Info  string `json:"info"`
// }
//
// type ChatMessage struct {
// 	Text      string     `json:"text"`
// 	Fragments []Fragment `json:"fragments"`
// }
//
// type Fragment struct {
// 	Type      string     `json:"type"`
// 	Text      string     `json:"text"`
// 	Cheermote *Cheermote `json:"cheermote,omitempty"`
// 	Emote     *Emote     `json:"emote,omitempty"`
// 	Mention   *Mention   `json:"mention,omitempty"`
// }
//
// type Cheermote struct {
// 	Prefix Prefix `json:"prefix"`
// 	Bits   int    `json:"bits"`
// 	Tier   int    `json:"tier"`
// }
//
// type Prefix struct {
// 	Prefix string `json:"prefix"`
// }
//
// type Emote struct {
// 	ID         string   `json:"id"`
// 	EmoteSetID string   `json:"emote_set_id"`
// 	OwnerID    string   `json:"owner_id"`
// 	Format     []string `json:"format"`
// }
//
// type Mention struct {
// 	UserID    string `json:"user_id"`
// 	UserName  string `json:"user_name"`
// 	UserLogin string `json:"user_login"`
// }
//
// type CharityDonation struct {
// 	CharityName string `json:"charity_name"`
// 	Amount      Amount `json:"amount"`
// }
