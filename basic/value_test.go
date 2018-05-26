package basic

import (
	"testing"
	"encoding/json"
)

func TestInitialize(t *testing.T) {
	bs, err := json.Marshal(AllCardMap)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(string(bs))
}
