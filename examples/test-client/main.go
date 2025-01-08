// This is the test client that relies on Twitch CLI - https://github.com/twitchdev/twitch-cli to demonstrate
// WebSocket client operation.
//
// Run the WebSocket server:
//
//	twitch event websocket start-server
//
// Start the `test-client` application and send different events to verify it works.
// Test events can be sent like that:
//
//	twitch event trigger -T websocket channel.follow
//
// Required subscription types for the User App Token:
//
//	user:read:email user:read:follows moderator:read:followers user:write:chat user:read:chat
package main

import (
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/nicklaw5/helix/v2"

	"github.com/vpetrigo/go-twitch-ws"
	"github.com/vpetrigo/go-twitch-ws/pkg/eventsub"
)

const helixTwitchTestServer = "http://127.0.0.1:8080"
const websocketTwitchTestServer = "ws://127.0.0.1:8080/ws"

var log *slog.Logger

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	log = slog.Default()

	log.Debug("Starting the test client...")
	c := twitchws.NewClientDefault(
		twitchws.WithOnWelcome(onWelcomeEvent),
		twitchws.WithOnNotification(onNotificationEvent),
		twitchws.WithOnConnect(onConnect),
		twitchws.WithOnDisconnect(onDisconnect),
		twitchws.WithOnRevocation(onRevocationEvent),
		twitchws.WithOnReconnect(onReconnect))

	err := c.Connect()

	if err != nil {
		log.Error("connect error", "err", err)
	}

	err = c.Wait()

	if err != nil {
		log.Error("wait error", "err", err)
	}

	err = c.Close()

	if err != nil {
		log.Error("close error", "err", err)
	}
}

func onWelcomeEvent(metadata *twitchws.Metadata, payload *twitchws.Payload) {
	log.Debug("Welcome message:", "metadata", metadata)
	log.Debug("Payload:", "payload", payload)

	clientID := os.Getenv("HELIX_CLIENT_ID")
	clientSecret := os.Getenv("HELIX_CLIENT_SECRET")
	userAccessToken := os.Getenv("HELIX_USER_ACCESS_TOKEN")
	refreshToken := os.Getenv("HELIX_REFRESH_TOKEN")

	if clientID == "" || clientSecret == "" || userAccessToken == "" || refreshToken == "" {
		panic("HELIX_CLIENT_ID, HELIX_CLIENT_SECRET, HELIX_USER_ACCESS_TOKEN, HELIX_REFRESH_TOKEN must be set")
	}

	helixClient, err := helix.NewClient(&helix.Options{
		ClientID:        clientID,
		ClientSecret:    clientSecret,
		UserAccessToken: userAccessToken,
		RefreshToken:    refreshToken,
		APIBaseURL:      helix.DefaultAPIBaseURL,
	})

	if err != nil {
		log.Error("helix client error", "err", err)
		return
	}

	session, _ := payload.Payload.(twitchws.Session)

	refresh, err := helixClient.RefreshUserAccessToken(helixClient.GetRefreshToken())

	if err != nil {
		log.Error("helix client error", "err", err)
		return
	}

	log.Debug("token refresh", "response", refresh)
	helixClient.SetUserAccessToken(refresh.Data.AccessToken)
	helixClient.SetRefreshToken(refresh.Data.RefreshToken)

	moderatorUserId := os.Getenv("HELIX_MOD_USER_ID")
	broadcasterUserId := os.Getenv("HELIX_BROADCASTER_USER_ID")
	userId := os.Getenv("HELIX_USER_ID")

	if moderatorUserId == "" || broadcasterUserId == "" || userId == "" {
		panic("HELIX_MOD_USER_ID, HELIX_BROADCASTER_USER_ID, HELIX_USER_ID must be set")
	}

	transport := helix.EventSubTransport{
		Method:    "websocket",
		SessionID: session.ID,
	}
	response, err := helixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:      helix.EventSubTypeChannelFollow,
		Version:   "2",
		Condition: helix.EventSubCondition{ModeratorUserID: moderatorUserId, BroadcasterUserID: broadcasterUserId},
		Transport: transport,
	})
	log.Debug("EventSub channel.follow", "response", response)

	response, err = helixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:      helix.EventSubTypeChannelChatMessage,
		Version:   "1",
		Condition: helix.EventSubCondition{BroadcasterUserID: broadcasterUserId, UserID: userId},
		Transport: transport,
	})
	log.Debug("EventSub channel.chat.message", "response", response)
}

func onNotificationEvent(metadata *twitchws.Metadata, payload *twitchws.Payload) {
	notification := payload.Payload.(twitchws.Notification)
	log.Debug("Metadata:", "metadata", metadata)
	log.Debug("Notification:", "notification", notification)

	switch event := notification.Event.(type) {
	case *eventsub.ChannelFollowEvent:
		log.Info("", "event", event)
		log.Info("", "condition", notification.Subscription.Condition)
	case *eventsub.ChannelChatMessage:
		log.Info("message", "message", event.Message.Text, "from", event.ChatterUserName, "from_id", event.ChatterUserID)

		clientID := os.Getenv("HELIX_CLIENT_ID")
		clientSecret := os.Getenv("HELIX_CLIENT_SECRET")
		userAccessToken := os.Getenv("HELIX_USER_ACCESS_TOKEN")
		refreshToken := os.Getenv("HELIX_REFRESH_TOKEN")

		if clientID == "" || clientSecret == "" || userAccessToken == "" || refreshToken == "" {
			panic("HELIX_CLIENT_ID, HELIX_CLIENT_SECRET, HELIX_USER_ACCESS_TOKEN, HELIX_REFRESH_TOKEN must be set")
		}

		helixClient, err := helix.NewClient(&helix.Options{
			ClientID:        clientID,
			ClientSecret:    clientSecret,
			UserAccessToken: userAccessToken,
			RefreshToken:    refreshToken,
			APIBaseURL:      helix.DefaultAPIBaseURL,
		})

		if err != nil {
			log.Error("helix client error", "err", err)
			return
		}

		if event.Reply != nil {
			return
		}

		moderatorUserId := os.Getenv("HELIX_MOD_USER_ID")

		response, err := helixClient.SendChatMessage(&helix.SendChatMessageParams{
			BroadcasterID:        event.BroadcasterUserID,
			Message:              strings.ToUpper(event.Message.Text),
			SenderID:             moderatorUserId,
			ReplyParentMessageID: event.MessageID,
		})

		log.Info("message sent", "response", response, "err", err)
	}
}

func onReconnect(metadata *twitchws.Metadata, payload *twitchws.Payload) {
	log.Debug("Reconnect:", "metadata", metadata, "payload", payload)
}

func onRevocationEvent(_ *twitchws.Metadata, payload *twitchws.Payload) {
	log.Debug("Revocation:", payload)
}

func onConnect() {
	log.Debug("Connected:", "time", time.Now().Format(time.RFC3339))
}

func onDisconnect() {
	log.Debug("Disconnected:", "time", time.Now().Format(time.RFC3339))
}
