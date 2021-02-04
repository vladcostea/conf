package conf

import "errors"

var ErrMissingKey = errors.New("param not configured")
var ErrParam = errors.New("failed to load param")

type Config struct {
	keys map[string]keyMapping
}

type keyMapping struct {
	keys     []string
	provider Provider
}

func New(opt func(c *Config)) *Config {
	c := &Config{keys: map[string]keyMapping{}}
	opt(c)
	return c
}

func (c *Config) AddProvider(pr Provider, mapping map[string][]string) {
	if c.keys == nil {
		c.keys = map[string]keyMapping{}
	}

	for k, keys := range mapping {
		c.keys[k] = keyMapping{keys: keys, provider: pr}
	}
}

func (c *Config) Param(key string) Param {
	m, ok := c.keys[key]
	if !ok {
		return Param{Err: ErrMissingKey}
	}

	for _, k := range m.keys {
		p := m.provider.Param(k)
		if p.Err == nil {
			return p
		}
	}

	return Param{Err: ErrParam}
}

type Provider interface {
	Param(key string) Param
}

// Param represents a generic configuration value.
// It's default representation is as a String but it can be converted to other primitive types.
type Param struct {
	// Err holds any error that might occur when fetching the param from it's provider
	Err error
	// Value holds the string value for a param
	Value string
}

func (p Param) String() string {
	return p.Value
}
