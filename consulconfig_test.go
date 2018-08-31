package gindy

import (
	"testing"
)

func TestGetZeroString(t *testing.T) {
	config := ConsulConfig{}
	v, err := config.GetString("myKey", "")
	if err != nil {
		t.Error(err)
	}
	if v != "" {
		t.Errorf("expecting an zero value for key %s, got %s\n", "myKey", v)
	}
}

func TestGetExistingString(t *testing.T) {
	config := ConsulConfig{}
	config.SetString("myKey", "myValue")

	v, err := config.GetString("myKey", "")
	if err != nil {
		t.Error(err)
	}
	if v != "myValue" {
		t.Errorf("expecting <%s> for key <%s>, got %s\n", "myValue", "myKey", v)
	}
}
