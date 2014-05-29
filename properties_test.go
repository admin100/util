package properties

import (
	"testing"
)

func Test_Get(t *testing.T) {
	p, err := Load("db.properties")
	if err != nil {
		t.Error("Load")
	}
	value := p.Get("url")
	if value == "" {
		t.Error("Get")
	}
	t.Log(value)
}
