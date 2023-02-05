package utils

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

func GetBuildConfig(path string) string {
	data, _ := ReadFile(path)

	t := make(map[interface{}]interface{})

	err := yaml.Unmarshal(data, &t)

	if err != nil {
		log.Fatal(err)
	}

	res, err := raymond.Render(string(data[:]), t)

	if err != nil {
		log.Fatal(err)
	}

	return res
}

func GetRuntimeConfig(path string) string {
	data, _ := ReadFile(path)

	output := raymond.MustRender(string(data[:]), viper.AllSettings())

	ReadInConfig(output)

	return output
}

func ReadInConfig(data string) {
	if err := viper.ReadConfig(bytes.NewBuffer([]byte([]byte(data)))); err != nil {
		log.Fatal(err)
	}
}

func SetEnvKey(t interface{}) {
	v := reflect.ValueOf(t)
	if v.Kind() == reflect.Map {
		keys := make(map[string]interface{})
		for _, key := range v.MapKeys() {
			strct := v.MapIndex(key)
			keys[key.String()] = strct.Interface()
		}
	}
}

func ReplaceEnvVars() {
	for _, k := range viper.AllKeys() {
		value := viper.GetString(k)
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}") {
			env := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}")
			viper.Set(k, getRequiredEnv(env))
		}
		if strings.HasPrefix(value, "${") && strings.HasSuffix(value, "}?") {
			env := strings.TrimSuffix(strings.TrimPrefix(value, "${"), "}?")
			viper.Set(k, getOptionalEnv(env))
		}
	}
}

func getRequiredEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		log.Fatal("Environment variable required: " + key)
	}
	return value
}

func getOptionalEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return ""
	}
	return value
}
