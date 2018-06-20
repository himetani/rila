package pocket

import (
	"os"
	"testing"
)

func ___TestPocketList(t *testing.T) {
	pocket := NewPocket()
	key := os.Getenv("CONSUMER_KEY")
	token := os.Getenv("ACCESS_TOKEN")

	if err := pocket.GetList(key, token); err != nil {
		panic(err)
	}
}
