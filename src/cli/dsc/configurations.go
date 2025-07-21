package dsc

import (
	"path/filepath"
	"strings"

	"github.com/jandedobbeleer/oh-my-posh/src/config"
)

type Configurations []*Configuration

type Configuration struct {
	Data   *config.Config `json:"data,omitempty"`
	Path   string         `json:"path,omitempty"`
	Format string         `json:"format,omitempty"`
}

func (c *Configurations) exists(configPath string) bool {
	for _, cfg := range *c {
		if cfg.Path == configPath {
			return true
		}
	}

	return false
}

func (c *Configurations) Add(configPath string) {
	if configPath == "" || strings.HasPrefix(configPath, "http") {
		return
	}

	if c.exists(configPath) {
		return
	}

	cfg := &Configuration{
		Path:   configPath,
		Format: strings.TrimPrefix(filepath.Ext(configPath), "."),
	}

	*c = append(*c, cfg)
}
