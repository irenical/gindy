package gindy

import (
	"errors"
	"fmt"
	"strconv"
)

// Config type to hold consul info
type Config struct {
	kv       map[string]string
	callback func(key string, value string)
}

// SetString sets the value for the given key
func (cc *Config) SetString(key string, value string) {
	ensure(cc)
	cc.kv[key] = value
}

// GetString returns the value associated with key
// if key is not present, the defaultValue is returned
func (cc Config) GetString(key string, defaultValue string) (string, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return defaultValue, nil
	}
	return value, nil
}

// GetStringMandatory returns the value associated with key
// if key is not present, panics
func (cc Config) GetStringMandatory(key string) (string, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return "", errors.New(fmt.Sprintf("mandatory string <%s> is not present", key))
	}
	return value, nil
}

// SetInt sets the given value into the given key
func (cc *Config) SetInt(key string, value int64) {
	ensure(cc)
	cc.kv[key] = string(value)
}

// GetInt returns the value associated with key
// if key is not present, the defaultValue is returned
func (cc Config) GetInt(key string, defaultValue int64) (int64, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return defaultValue, nil
	}
	return strconv.ParseInt(value, 10, 64)
}

// GetIntMandatory returns the value associated with key
// if key is not present, panics
func (cc Config) GetIntMandatory(key string) (int64, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return 0, errors.New(fmt.Sprintf("mandatory int <%s> is not present", key))
	}
	return strconv.ParseInt(value, 10, 64)
}

// SetBoolean sets the given value into the given key
func (cc *Config) SetBoolean(key string, value bool) {
	ensure(cc)
	cc.kv[key] = strconv.FormatBool(value)
}

// GetBoolean returns the value associated with key
// if key is not present, the defaultValue is returned
func (cc Config) GetBoolean(key string, defaultValue bool) (bool, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return defaultValue, nil
	}
	return strconv.ParseBool(value)
}

// GetBooleanMandatory returns the value associated with key
// if key is not present, panics
func (cc Config) GetBooleanMandatory(key string) (bool, error) {
	value, isPresent := cc.kv[key]
	if !isPresent {
		return false, errors.New(fmt.Sprintf("mandatory bool <%s> is not present", key))
	}
	return strconv.ParseBool(value)
}

// Listen to start listening for KV changes
func (cc *Config) Listen(callback func(key string, value string)) {
	cc.callback = callback
}

// UnListen to stop listening for KV changes
func (cc *Config) UnListen() {
	cc.callback = nil
}

func ensure(config *Config) {
	if config.kv == nil {
		config.kv = make(map[string]string)
	}
}
