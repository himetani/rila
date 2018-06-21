package pocket

import (
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	bytes, err := ioutil.ReadFile(filepath.Join("testdata", "params.json"))
	if err != nil {
		t.Error("Unexpected error happend")
	}

	expected := strings.TrimSuffix(string(bytes), "\n")

	params := &Params{ConsumerKey: "consumer_key", AccessToken: "access_token"}
	actual := params.String()

	if expected != actual {
		t.Errorf("expected %s, but got = %s", expected, actual)
	}
}
