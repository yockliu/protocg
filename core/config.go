package core

import (
	"fmt"
	"reflect"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
)

func parseConfig() *Generator {
	g := new(Generator)
	opt := viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(StringToRemoteHookFunc()))
	err := viper.Unmarshal(g, opt)
	if err != nil {
		fmt.Println(err)
	}
	return g
}

func StringToRemoteHookFunc() mapstructure.DecodeHookFuncType {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {

		if f.Kind() != reflect.String {
			return data, nil
		}

		if t != reflect.TypeOf(Remote{}) {
			return data, nil
		}

		raw := data.(string)

		return Remote{Repo: raw}, nil
	}
}
