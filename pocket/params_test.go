package pocket

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestString(t *testing.T) {
	bytes, err := ioutil.ReadFile(filepath.Join(os.Getenv("GOPATH"), "src", "github.com", "himetani", "rila", "pocket", "testdata", "params.json"))
	if err != nil {
		t.Error("Unexpected error happend")
	}

	expected := string(bytes)

	params := &Params{ConsumerKey: "consumer_key", AccessToken: "access_token"}
	actual := params.String()

	if expected != actual {
		t.Errorf("expected %s, but got = %s", expected, actual)
	}
}
