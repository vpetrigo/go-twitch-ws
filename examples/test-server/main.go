package main

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"

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
		twitchws.WithOnDisconnect(onDisconnect))

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

func onWelcomeEvent(metadata twitchws.Metadata, payload twitchws.Payload) {
	logrus.Debugf("Metadata: %+v", metadata)
	logrus.Debugf("Payload: %+v", payload)
}

func onNotificationEvent(metadata twitchws.Metadata, payload twitchws.Payload) {
	notification := payload.Payload.(twitchws.Notification)

	switch event := notification.Event.(type) {
	case *twitchws.ChannelFollowEvent:
		logrus.Debugf("Channel follow: %+v", event)
	}
}

func onConnect() {
	logrus.Debugf("Connected: %s", time.Now().Format(time.RFC3339))
}

func onDisconnect() {
	logrus.Debugf("Disconnected: %s", time.Now().Format(time.RFC3339))
}
