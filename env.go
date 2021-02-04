package conf

import "os"

type EnvProvider struct{}

func (p *EnvProvider) Param(key string) Param {
	v := os.Getenv(key)
	if v == "" {
		return Param{Err: ErrMissingKey}
	}

	return Param{Value: v}
}

func NewEnvProvider() *EnvProvider {
	return &EnvProvider{}
}
