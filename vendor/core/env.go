package core

import (
	"fmt"
	"os"
)

type Environtment struct {
	configs map[string]interface{}
}

func (e *Environtment) Get(key string) string {
	textEnv := os.Getenv(key)

	if textEnv == "" {
		return fmt.Sprint(e.configs[key])
	} else {
		return textEnv
	}
}

func Env(path string) Environtment {
	var environment Environtment
	environment = Environtment{Config(path)}

	return environment
}
