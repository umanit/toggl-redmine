package cfg

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/viper"
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
	home, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Errorf("cannot find home directory: %v", err))
	}

	viper.AddConfigPath(home)
	viper.SetConfigType("json")
	viper.SetConfigName(".toggl-redmine")
	viper.SetDefault("toggl.url", "https://suivi.umanit.fr")
	viper.SetDefault("redmine.url", "https://api.track.toggl.com/api/v9")
	_ = viper.SafeWriteConfig()

	if err = viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("cannot read config file: %v", err))
	}

	var cfg Config
	if err = viper.Unmarshal(&cfg); err != nil {
		panic(fmt.Errorf("cannot parse config file: %v", err))
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
