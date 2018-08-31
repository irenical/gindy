package gindy

import (
	"fmt"
	"testing"
)

func TestConsulConnection(t *testing.T) {
	if !testing.Short() {
		t.Skip("skipping test of consul connection")
	}

	config := WatchedConsulConfigFactory{address: "localhost:8500", appID: "myAppID"}.create()

	config.Listen(func(key string, value string) {
		fmt.Printf("key <%s> changed to <%s>\n", key, value)
	})

	quitChannel := make(chan interface{})

	<-quitChannel
}
