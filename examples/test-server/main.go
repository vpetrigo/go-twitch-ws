package main

import (
	"os"

	"github.com/sirupsen/logrus"

	"github.com/vpetrigo/go-twitch-ws"
)

const websocketTwitchTestServer = "ws://127.0.0.1:8080/ws"

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	c := twitchws.NewClient(
		websocketTwitchTestServer,
		twitchws.WithOnWelcome(onWelcomeEvent))

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
