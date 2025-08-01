package ephemera

import (
	"database/sql"
	"html/template"
	"io/fs"

	"github.com/bdreece/ephemera/pkg/storage"
)

type Config struct {
	Port        int
	DB          *sql.DB
	Assets      []fs.FS
	Environment string
	Storage     storage.Provider
	Templates   *template.Template
}

type Option interface {
	apply(cfg *Config) error
}

func (cfg *Config) apply(other *Config) error {
	*other = *cfg
	return nil
}

type option func(cfg *Config) error

func (fn option) apply(cfg *Config) error {
	return fn(cfg)
}

func WithPort(port int) Option {
	return option(func(cfg *Config) error {
		cfg.Port = port
		return nil
	})
}

func WithSqlite(dsn string) Option {
	return nil
}
