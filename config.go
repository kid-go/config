package config

import (
	"github.com/spf13/cast"
	"sync"
	"time"
)

type Config interface {
	Getter
	SetProvider(provider Provider)
}

type Provider interface {
	Get(key string) interface{}
	OnWatch(func()) error
	ReadConfig() error
	Unmarshal(key string, receiver interface{}) error
}

type Getter interface {
	GetString(key string) string
	GetStringSlice(key string) []string
	GetStringMap(key string) map[string]string
	GetStringMapStringSlice(string) map[string][]string
	GetInt(key string) int
	GetIntSlice(key string) []int
	GetInt64(key string) int64
	GetBool(key string) bool
	GetFloat(key string) float64
	GetTime(key string) time.Time
	GetDuration(key string) time.Duration
}

type kidConfig struct {
	provider Provider
}

func (conf *kidConfig) Create() Config {
	return &kidConfig{}
}

func (conf *kidConfig) SetProvider(provider Provider) {
	conf.provider = provider
}

func (conf *kidConfig) get(key string) interface{} {
	return conf.provider.Get(key)
}

func (conf *kidConfig) GetString(key string) string {
	return cast.ToString(conf.get(key))
}

func (conf *kidConfig) GetStringSlice(key string) []string {
	return cast.ToStringSlice(conf.get(key))
}

func (conf *kidConfig) GetStringMap(key string) map[string]string {
	return cast.ToStringMapString(conf.get(key))
}

func (conf *kidConfig) GetStringMapStringSlice(key string) map[string][]string {
	return cast.ToStringMapStringSlice(conf.get(key))
}

func (conf *kidConfig) GetInt(key string) int {
	return cast.ToInt(conf.get(key))
}

func (conf *kidConfig) GetIntSlice(key string) []int {
	return cast.ToIntSlice(conf.get(key))
}

func (conf *kidConfig) GetInt64(key string) int64 {
	return cast.ToInt64(conf.get(key))
}

func (conf *kidConfig) GetBool(key string) bool {
	return cast.ToBool(conf.get(key))
}

func (conf *kidConfig) GetFloat(key string) float64 {
	return cast.ToFloat64(conf.get(key))
}

func (conf *kidConfig) GetTime(key string) time.Time {
	return cast.ToTime(conf.get(key))
}

func (conf *kidConfig) GetDuration(key string) time.Duration {
	return cast.ToDuration(conf.get(key))
}

var (
	_defaultConfig Config
	_defaultOnce   sync.Once
)

func New() Config {
	_defaultOnce.Do(func() {
		_defaultConfig = &kidConfig{}
	})
	return _defaultConfig
}

func SetProvider(provider Provider) {
	_defaultConfig.SetProvider(provider)
}

func GetString(key string) string {
	return _defaultConfig.GetString(key)
}

func GetStringSlice(key string) []string {
	return _defaultConfig.GetStringSlice(key)
}

func GetStringMap(key string) map[string]string {
	return _defaultConfig.GetStringMap(key)
}

func GetStringMapStringSlice(key string) map[string][]string {
	return _defaultConfig.GetStringMapStringSlice(key)
}

func GetInt(key string) int {
	return _defaultConfig.GetInt(key)
}

func GetIntSlice(key string) []int {
	return _defaultConfig.GetIntSlice(key)
}

func GetInt64(key string) int64 {
	return _defaultConfig.GetInt64(key)
}

func GetBool(key string) bool {
	return _defaultConfig.GetBool(key)
}

func GetFloat(key string) float64 {
	return _defaultConfig.GetFloat(key)
}

func GetTime(key string) time.Time {
	return _defaultConfig.GetTime(key)
}

func GetDuration(key string) time.Duration {
	return _defaultConfig.GetDuration(key)
}
