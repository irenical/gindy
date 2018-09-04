package gindy

import (
	"fmt"
	"testing"
)

func TestConsulConnection(t *testing.T) {
	if !testing.Short() {
		t.Skip("skipping test of consul connection")
	}

	config := ConsulConfigFactory{Address: "localhost:8500", AppID: "myAppID"}.Create()

	config.Listen(func(key string, value string) {
		fmt.Printf("key <%s> changed to <%s>\n", key, value)
	})

	quitChannel := make(chan interface{})

	<-quitChannel
}
