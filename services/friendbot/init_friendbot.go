package main

import (
	"net/http"

	"github.com/test/go/clients/testhorizon"
	"github.com/test/go/services/friendbot/internal"
	"github.com/test/go/strkey"
)

func initFriendbot(
	friendbotSecret string,
	networkPassphrase string,
	testhorizonURL string,
	startingBalance string,
) *internal.Bot {

	if friendbotSecret == "" || networkPassphrase == "" || testhorizonURL == "" || startingBalance == "" {
		return nil
	}

	// ensure its a seed if its not blank
	strkey.MustDecode(strkey.VersionByteSeed, friendbotSecret)

	return &internal.Bot{
		Secret: friendbotSecret,
		TestHorizon: &testhorizon.Client{
			URL:  testhorizonURL,
			HTTP: http.DefaultClient,
		},
		Network:           networkPassphrase,
		StartingBalance:   startingBalance,
		SubmitTransaction: internal.AsyncSubmitTransaction,
	}
}
