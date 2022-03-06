package config

import "github.com/spf13/viper"

func GetStringOrDefault(viper *viper.Viper, key, value string) string {
	viper.SetDefault(key, value)
	return viper.GetString(key)
}
