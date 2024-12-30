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
package main

import (
	"log/slog"
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
	c := twitchws.NewClient(
		websocketTwitchTestServer,
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
	helixClient, err := helix.NewClient(&helix.Options{
		ClientID:     "pm8irdclcqvj3it9ev4h6hccrk7ebv",
		ClientSecret: "fpak1tegsg57hdp4t6ltvjzv7f8akq",
		APIBaseURL:   helixTwitchTestServer,
	})

	if err != nil {
		log.Error("helix client error", "err", err)
		return
	}

	session, _ := payload.Payload.(twitchws.Session)
	response, err := helixClient.CreateEventSubSubscription(&helix.EventSubSubscription{
		Type:    helix.EventSubTypeChannelFollow,
		Version: "2",
		Transport: helix.EventSubTransport{
			Method:    "websocket",
			SessionID: session.ID,
		},
	})
	log.Debug("helix response", "response", response)
}

func onNotificationEvent(metadata *twitchws.Metadata, payload *twitchws.Payload) {
	notification := payload.Payload.(twitchws.Notification)
	log.Debug("Metadata:", "metadata", metadata)
	log.Debug("Notification:", "notification", notification)

	if event, ok := notification.Event.(*eventsub.ChannelFollowEvent); ok {
		log.Debug("", "event", event)
		log.Debug("", "condition", notification.Subscription.Condition)
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
