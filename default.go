package conf

import "errors"

var ErrMissingDefaultParam = errors.New("missing default")

type DefaultProvider map[string]string

func (p DefaultProvider) Param(key string) Param {
	v, ok := p[key]
	if !ok {
		return Param{Err: ErrMissingDefaultParam}
	}

	return Param{Value: v}
}
