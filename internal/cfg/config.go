package cfg

import (
	"context"

	"github.com/spf13/viper"
	"github.com/umanit/toggl-redmine/internal/app"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type configKey int

const (
	_ configKey = iota
	key
)

type ApiConfig struct {
	Key string `json:"key"`
	Url string `json:"url"`
}

type Config struct {
	Toggl   *ApiConfig `json:"toggl"`
	Redmine *ApiConfig `json:"redmine"`
}

func ContextWithConfig(ctx context.Context) context.Context {
	appDir := app.GetAppDir()
	viper.AddConfigPath(appDir)
	viper.SetConfigType("json")
	viper.SetConfigName("config")
	viper.SetDefault("toggl.url", "https://api.track.toggl.com/api/v9")
	viper.SetDefault("redmine.url", "https://suivi.umanit.fr")
	_ = viper.SafeWriteConfig()

	if err := viper.ReadInConfig(); err != nil {
		runtime.LogFatalf(ctx, "cannot read config file: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		runtime.LogFatalf(ctx, "cannot parse config file: %v", err)
	}

	return context.WithValue(ctx, key, cfg)
}

func ConfigFromContext(ctx context.Context) Config {
	cfg, ok := ctx.Value(key).(Config)
	if !ok {
		runtime.LogFatal(ctx, "can't load config")
	}

	return cfg
}

func (c *Config) AllValuesFilled() bool {
	return c.Toggl != nil && c.Redmine != nil && c.Redmine.Key != "" && c.Redmine.Url != "" &&
		c.Toggl.Key != "" && c.Toggl.Url != ""
}

func (c *Config) Save(n Config) error {
	// Sauvegarde de la configuration viper
	viper.Set("toggl.key", n.Toggl.Key)
	viper.Set("toggl.url", n.Toggl.Url)
	viper.Set("redmine.key", n.Redmine.Key)
	viper.Set("redmine.url", n.Redmine.Url)

	// Sauvegarde de la configuration « live »
	c.Toggl.Key = n.Toggl.Key
	c.Toggl.Url = n.Toggl.Url
	c.Redmine.Key = n.Redmine.Key
	c.Redmine.Url = n.Redmine.Url

	return viper.WriteConfig()
}
