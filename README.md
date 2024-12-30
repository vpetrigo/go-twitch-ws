[![lint](https://github.com/vpetrigo/go-twitch-ws/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/vpetrigo/go-twitch-ws/actions/workflows/golangci-lint.yml)
[![tests](https://github.com/vpetrigo/go-twitch-ws/actions/workflows/tests.yml/badge.svg)](https://github.com/vpetrigo/go-twitch-ws/actions/workflows/tests.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/vpetrigo/go-twitch-ws.svg)](https://pkg.go.dev/github.com/vpetrigo/go-twitch-ws)
[![Go Report Card](https://goreportcard.com/badge/github.com/vpetrigo/go-twitch-ws)](https://goreportcard.com/report/github.com/vpetrigo/go-twitch-ws)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/vpetrigo/go-twitch-ws)
![GitHub Release](https://img.shields.io/github/v/release/vpetrigo/go-twitch-ws)

# Twitch WebSocket EventSub Client

----------------------------------

This repo contains of two packages:

- `twitchws`: Twitch WebSocket EventSub client
- [`eventsub`](pkg/eventsub): EventSub messages declared by Twitch

## Package `twitchws`

Implements Twitch WebSocket EventSub protocol requirements
described [here](https://dev.twitch.tv/docs/eventsub/handling-websocket-events/). The client is able to handle:

- EventSub notifications
- keep-alive messages
- reconnection requests
- revocation requests

Usage example with [Twitch CLI](https://github.com/twitchdev/twitch-cli):

```go
package main

import (
	"fmt"
	
	"github.com/vpetrigo/go-twitch-ws"
)

const websocketTwitchTestServer = "ws://127.0.0.1:8080/ws"

func main() {
	messageHandler := func(m *twitchws.Metadata, p *twitchws.Payload) {
		fmt.Printf("Metadata: %+v\n", m)
		fmt.Printf("Payload: %+v\n", p)
	}
	stateHandler := func(state string) func() {
		return func() {
			fmt.Printf("Event: %s\n", state)
		}
	}

	c := twitchws.NewClient(
		websocketTwitchTestServer,
		twitchws.WithOnWelcome(messageHandler),
		twitchws.WithOnNotification(messageHandler),
		twitchws.WithOnConnect(stateHandler("Connect")),
		twitchws.WithOnDisconnect(stateHandler("Disconnect")),
		twitchws.WithOnRevocation(messageHandler))

	err := c.Connect()

	if err != nil {
		fmt.Println(err)
	}

	err = c.Wait()

	if err != nil {
		fmt.Println(err)
	}

	err = c.Close()

	if err != nil {
		fmt.Println(err)
	}
}

```

## Package `eventsub`

It is an attempt to automatically
parse [Twitch official API reference](https://dev.twitch.tv/docs/eventsub/eventsub-reference/) in order to generate
respective types. There are some issues that have to be resolved before all types will be properly
parsed.

# Examples

All examples are available in the [`examples`](examples) directory

# Contribution

Contributions are always welcome! If you have an idea, it's best to float it by me before working on it to ensure no
effort is wasted. If there's already an open issue for it, knock yourself out. See the
[**contributing section**](CONTRIBUTING.md) for additional details

## Design

- [Implementation status](docs/SUPPORTED.md)
