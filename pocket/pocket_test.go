package pocket

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

func TestPocketGetOldestItem(t *testing.T) {
	pocket := NewPocket()
	key := os.Getenv("CONSUMER_KEY")
	token := os.Getenv("ACCESS_TOKEN")

	item, err := pocket.GetOldestItem(key, token, 0)
	if err != nil {
		t.Error(err)
	}

	actual, err := json.Marshal(&item)
	if err != nil {
		t.Error(err)
	}

	golden := filepath.Join("testdata", "test.golden")
	if *update {
		ioutil.WriteFile(golden, actual, 0644)
	}

	expected, _ := ioutil.ReadFile(golden)
	if !bytes.Equal(actual, expected) {
		t.Errorf("expected=%s, got=%s", string(expected), string(actual))
	}

}
