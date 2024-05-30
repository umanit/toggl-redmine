package cfg

import (
	"context"

	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

type configKey int

const (
	_ configKey = iota
	key
)

type api struct {
	Key string
	Url string
}

type Config struct {
	Toggl   *api
	Redmine *api
}

func ContextWithConfig(ctx context.Context) context.Context {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		SkipDefaults: true,
		SkipEnv:      true,
		SkipFlags:    true,
		Files:        []string{"config.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		panic(err)
	}

	return context.WithValue(ctx, key, cfg)
}

func ConfigFromContext(ctx context.Context) (Config, bool) {
	cfg, ok := ctx.Value(key).(Config)
	return cfg, ok
}

func (c *Config) AllFill() bool {
	return c.Toggl != nil && c.Redmine != nil && c.Redmine.Key != "" && c.Redmine.Url != "" && c.Toggl.Key != "" && c.Toggl.Url != ""
}
