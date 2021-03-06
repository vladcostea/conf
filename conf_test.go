package conf

import (
	"os"
	"testing"
)

func TestDefaultProvider(t *testing.T) {
	c := New(func(cfg *Config) {
		cfg.AddProvider(DefaultProvider{
			"PARAM_1": "Value1",
			"PARAM-2": "Value2",
		}, map[string][]string{
			"Param1": {"PARAM_1"},
			"Param2": {"PARAM_2", "PARAM-2"},
		})
	})

	v1 := c.Param("Param1").String()
	if v1 != "Value1" {
		t.Errorf("wanted Value1 got %s", v1)
	}

	v2 := c.Param("Param2").String()
	if v2 != "Value2" {
		t.Errorf("wanted Value2 got %s", v2)
	}

	v3 := c.Param("Param3")
	if v3.Err != ErrMissingKey {
		t.Errorf("wanted ErrMissingKey got %v", v3.Err)
	}
}

func TestEnvWithDefault(t *testing.T) {
	c := New(func(cfg *Config) {
		cfg.AddProvider(DefaultProvider{
			"PARAM_1": "Value1",
		}, map[string][]string{
			"Param1": {"PARAM_1"},
		})

		cfg.AddProvider(NewEnvProvider(), map[string][]string{"Param2": {"PARAM_2"}})
	})

	os.Setenv("PARAM_2", "Value2")

	v1 := c.Param("Param1").String()
	if v1 != "Value1" {
		t.Errorf("wanted Value1 got %s", v1)
	}

	v2 := c.Param("Param2").String()
	if v2 != "Value2" {
		t.Errorf("wanted Value2 got %s", v1)
	}
}
