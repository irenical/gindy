package gindy

import (
	"os"
	"strings"
	"time"

	consul "github.com/hashicorp/consul/api"
)

// WatchedConsulConfigFactory a factory of consul config with kv auto updates
//
// Usage:
//		config := WatchedConsulConfigFactory{address: "localhost:8500"}.create()
type WatchedConsulConfigFactory struct {
	AppID   string
	Address string
}

// Create creates a new instance of ConsulConfig
func (ccf WatchedConsulConfigFactory) Create() *ConsulConfig {
	appID := ccf.AppID
	if appID == "" {
		appID = readAppIDFromEnv()
	}

	config := new(ConsulConfig)

	// consul details
	consulConfig := consul.DefaultConfig()
	// consulConfig.Address = readConsulAddressFromEnv()
	consulConfig.Address = ccf.Address

	// create a consul api client (for kv info)
	client, err := consul.NewClient(consulConfig)
	if err != nil {
		panic(err)
	}

	kvClient := client.KV()

	// get initial kv map synchronously on startup
	config.kv = getKVMap(kvClient, appID)

	// execute kvpairs polling in the background to keep local map updated and notify of changes
	go pollKVs(kvClient, config, appID)

	return config
}

func readAppIDFromEnv() string {
	appID := os.Getenv("applicationId")
	if appID == "" {
		appID = os.Getenv("appId")
	}
	if appID == "" {
		appID = os.Getenv("@appId")
	}

	if appID == "" {
		panic("mandatory appId property not set")
	}
	return appID
}

// aux consul config specific methods
func requestKVPairs(kvClient *consul.KV, prefix string) (consul.KVPairs, error) {
	// TODO: configurable consul QueryOptions (timeouts, etc)
	queryOptions := &consul.QueryOptions{}

	kvPairs, _, err := kvClient.List(prefix, queryOptions)
	return kvPairs, err
}

func convertKVToMap(kvPairs consul.KVPairs, prefix string) map[string]string {
	raw := make(map[string]string)
	for _, p := range kvPairs {
		paths := strings.Split(p.Key, prefix+"/")
		if len(paths) > 0 {
			raw[paths[1]] = string(p.Value)
		} else {
			raw[paths[0]] = string(p.Value)
		}
	}
	return raw
}

func checkForChanges(currentKV map[string]string, newKV map[string]string, callback func(key string, value string)) {
	// if no old kvmap or callback is not defined, skip
	if len(currentKV) <= 0 || callback == nil {
		return
	}

	for k, v := range newKV {
		oldValue := currentKV[k]

		// for each key in new map, check for changes in 'current' config - if diff, notify
		if oldValue != newKV[k] {
			callback(k, v)
		}
	}
}

func getKVMap(kvClient *consul.KV, prefix string) map[string]string {
	// request KV pairs from consul
	kvPairs, err := requestKVPairs(kvClient, prefix)
	if err != nil {
		panic(err)
	}

	// convert KV pairs to map
	kvMap := convertKVToMap(kvPairs, prefix)

	return kvMap
}

func pollKVs(kvClient *consul.KV, config *ConsulConfig, prefix string) {
	for {
		kvMap := getKVMap(kvClient, prefix)

		// check for value changes and notify
		checkForChanges(config.kv, kvMap, config.callback)

		// update local map with new KV pairs
		config.kv = kvMap

		// TODO: configurable polling interval
		time.Sleep(time.Second)
	}
}
