package ephemera

import (
	"errors"
	"flag"
	"log/slog"

	"github.com/bdreece/ephemera/internal/security"
	"github.com/bdreece/ephemera/pkg/database"
	"github.com/kelseyhightower/envconfig"
)

type JWTConfig struct {
	Secret   security.Key `envconfig:"jwt_secret"`
	Audience string       `envconfig:"jwt_audience"`
	Issuer   string       `envconfig:"jwt_issuer"`
}

type Config struct {
	JWTConfig

	Port        int          `envconfig:"port"`
	Environment Env          `envconfig:"env"`
	LogLevel    LogLevel     `envconfig:"log_level"`
	SqliteDSN   database.DSN `envconfig:"sqlite_dsn"`
	StorageRoot string       `envconfig:"storage_root"`
}

var defaultJWTConfig = JWTConfig{
	Audience: "ephemera.bdreece.dev",
	Issuer:   "ephemera.bdreece.dev",
}

var DefaultConfig = Config{
	Port:        8080,
	StorageRoot: "/var/lib/ephemera/media",
	JWTConfig:   defaultJWTConfig,
	SqliteDSN: database.DSN{
		Path: "/var/lib/ephemera/db.sqlite3",
	},
}

func (cfg *Config) UnmarshalEnvVars() {
	envconfig.MustProcess("ephemera", cfg)
}

func (cfg *Config) UnmarshalFlags(set *flag.FlagSet) {
	flag.IntVar(&cfg.Port, "port", cfg.Port, "http port")
	flag.Var(&cfg.Environment, "env", "environment (dev|prod)")
	flag.Var(&cfg.SqliteDSN, "dsn", "sqlite dsn")
	flag.StringVar(&cfg.Audience, "jwt.aud", cfg.Audience, "jwt audience")
	flag.StringVar(&cfg.Issuer, "jwt.iss", cfg.Issuer, "jwt issuer")
	flag.Var(&cfg.LogLevel, "level", "log level (DEBUG|WARN|INFO|ERROR)")
	flag.StringVar(&cfg.StorageRoot, "root", cfg.StorageRoot, "path to storage root")
	flag.Parse()
}

func WithConfig(cfg *Config) Option {
	return func(p *Params) error {
		_ = WithPort(cfg.Port)(p)
		_ = WithEnvironment(cfg.Environment)(p)
		_ = WithLogLevel(cfg.LogLevel)(p)
		_ = WithJwtSecret(cfg.Secret)(p)

		if err := WithSqliteDSN(cfg.SqliteDSN)(p); err != nil {
			return err
		}

		if err := WithStorageRoot(cfg.StorageRoot)(p); err != nil {
			return err
		}

		return nil
	}
}

//go:generate go tool stringer -type Env -linecomment
type Env int

const (
	EnvDevelopment Env = iota // dev
	EnvProduction             // prod
)

var ErrParseEnv = errors.New("ephemera: failed to parse env")

// UnmarshalText implements [encoding.TextUnmarshaler].
func (env *Env) UnmarshalText(text []byte) error {
	switch string(text) {
	case EnvDevelopment.String():
		*env = EnvDevelopment
	case EnvProduction.String():
		*env = EnvProduction
	default:
		return ErrParseEnv
	}

	return nil
}

func (env Env) IsDevelopment() bool { return env == EnvDevelopment }

func (env Env) IsProduction() bool { return env == EnvProduction }

// Set implements [flag.Value].
func (env *Env) Set(s string) error {
	return env.UnmarshalText([]byte(s))
}

// LogValue implements [slog.LogValuer].
func (env Env) LogValue() slog.Value { return slog.StringValue(env.String()) }

type LogLevel struct{ val slog.Level }

// String implements [flag.Value].
func (level LogLevel) String() string { return level.val.String() }

// Level implements [slog.Leveler].
func (level LogLevel) Level() slog.Level { return level.val }

// UnmarshalText implements [encoding.TextUnmarshaler].
func (level *LogLevel) UnmarshalText(text []byte) error {
	return level.val.UnmarshalText(text)
}

// Set implements [flag.Value].
func (level *LogLevel) Set(s string) error {
	return level.val.UnmarshalText([]byte(s))
}
