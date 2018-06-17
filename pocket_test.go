package main

import (
	"os"
	"testing"
)

func TestPocketList(t *testing.T) {
	pocket := newPocket()
	key := os.Getenv("CONSUMER_KEY")
	token := os.Getenv("ACCESS_TOKEN")

	if err := pocket.GetList(key, token); err != nil {
		panic(err)
	}
}
