package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/vpetrigo/go-twitch-ws/pkg/eventsub"

	"github.com/vpetrigo/go-twitch-ws"
)

const websocketTwitchTestServer = "ws://127.0.0.1:8080/ws"

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	c := twitchws.NewClient(
		websocketTwitchTestServer,
		twitchws.WithOnWelcome(onWelcomeEvent),
		twitchws.WithOnNotification(onNotificationEvent),
		twitchws.WithOnConnect(onConnect),
		twitchws.WithOnDisconnect(onDisconnect),
		twitchws.WithOnRevocation(onRevocationEvent))

	err := c.Connect()

	if err != nil {
		logrus.Fatal(err)
	}

	err = c.Wait()

	if err != nil {
		logrus.Fatal(err)
	}

	err = c.Close()

	if err != nil {
		logrus.Fatal(err)
	}
}

func onWelcomeEvent(metadata *twitchws.Metadata, payload *twitchws.Payload) {
	logrus.Debugf("Metadata: %+v", metadata)
	logrus.Debugf("Payload: %+v", payload)
}

func onNotificationEvent(_ *twitchws.Metadata, payload *twitchws.Payload) {
	notification := payload.Payload.(twitchws.Notification)
	logrus.Debugf("Notification: %+v", notification)

	if event, ok := notification.Event.(*eventsub.ChannelFollowEvent); ok {
		condition := notification.Subscription.Condition.(*eventsub.ChannelFollowEventCondition)
		logrus.Debugf("Channel follow: %+v", event)
		logrus.Debugf("Condition: %+v", condition)
	}
}

func onRevocationEvent(_ *twitchws.Metadata, payload *twitchws.Payload) {
	logrus.Debugf("Revocation: %+v", payload)
}

func onConnect() {
	logrus.Debugf("Connected: %s", time.Now().Format(time.RFC3339))
}

func onDisconnect() {
	logrus.Debugf("Disconnected: %s", time.Now().Format(time.RFC3339))
}
