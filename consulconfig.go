package gindy

import (
	"strconv"
)

// ConsulConfig type to hold consul info
type ConsulConfig struct {
	kv       map[string]string
	callback func(key string, value string)
}

// SetString sets the value for the given key
func (cc *ConsulConfig) SetString(key string, value string) {
	ensure(cc)
	cc.kv[key] = value
}

// GetString returns the value associated with key
// if key is not present, the defaultValue is returned
func (cc ConsulConfig) GetString(key string, defaultValue string) (string, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return defaultValue, nil
	}
	return value, nil
}

// GetStringMandatory returns the value associated with key
// if key is not present, panics
func (cc ConsulConfig) GetStringMandatory(key string) (string, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		panic("mandatory string <" + key + "> is not present")
	}
	return value, nil
}

// SetInt sets the given value into the given key
func (cc *ConsulConfig) SetInt(key string, value int64) {
	ensure(cc)
	cc.kv[key] = string(value)
}

// GetInt returns the value associated with key
// if key is not present, the defaultValue is returned
func (cc ConsulConfig) GetInt(key string, defaultValue int64) (int64, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return defaultValue, nil
	}
	return strconv.ParseInt(value, 10, 64)
}

// GetIntMandatory returns the value associated with key
// if key is not present, panics
func (cc ConsulConfig) GetIntMandatory(key string) (int64, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		panic("mandatory int <" + key + "> is not present")
	}
	return strconv.ParseInt(value, 10, 64)
}

// SetBoolean sets the given value into the given key
func (cc *ConsulConfig) SetBoolean(key string, value bool) {
	ensure(cc)
	cc.kv[key] = strconv.FormatBool(value)
}

// GetBoolean returns the value associated with key
// if key is not present, the defaultValue is returned
func (cc ConsulConfig) GetBoolean(key string, defaultValue bool) (bool, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return defaultValue, nil
	}
	return strconv.ParseBool(value)
}

// GetBooleanMandatory returns the value associated with key
// if key is not present, panics
func (cc ConsulConfig) GetBooleanMandatory(key string) (bool, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		panic("mandatory bool <" + key + "> is not present")
	}
	return strconv.ParseBool(value)
}

// Listen to start listening for KV changes
func (cc *ConsulConfig) Listen(callback func(key string, value string)) {
	cc.callback = callback
}

// UnListen to stop listening for KV changes
func (cc *ConsulConfig) UnListen() {
	cc.callback = nil
}

func ensure(config *ConsulConfig) {
	if config.kv == nil {
		config.kv = make(map[string]string)
	}
}
