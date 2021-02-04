package conf

import "errors"

var ErrMissingDefaultParam = errors.New("missing default")

type DefaultProvider struct {
	defaults map[string]string
}

func NewDefaultProvider(d map[string]string) *DefaultProvider {
	return &DefaultProvider{
		defaults: d,
	}
}

func (p *DefaultProvider) Param(key string) Param {
	v, ok := p.defaults[key]
	if !ok {
		return Param{Err: ErrMissingDefaultParam}
	}

	return Param{Value: v}
}
