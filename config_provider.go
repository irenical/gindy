package gindy

// ConfigProvider interface
type ConfigProvider interface {
	SetString(key string, value string)
	GetString(key string, defaultValue string) (string, error)
	GetStringMandatory(key string) (string, error)

	SetInt(key string, value int64)
	GetInt(key string, defaultValue int64) (int64, error)
	GetIntMandatory(key string) (int64, error)

	SetBoolean(key string, value bool)
	GetBoolean(key string, defaultValue bool) (bool, error)
	GetBooleanMandatory(key string) (bool, error)

	Listen(callback func(key string, value string))
	UnListen()
}
