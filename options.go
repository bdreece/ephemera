package ephemera

import (
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/bdreece/ephemera/internal/security"
	"github.com/bdreece/ephemera/pkg/database"
	"github.com/bdreece/ephemera/pkg/storage"
)

type Option func(p *Params) error

func WithPort(port int) Option {
	return func(p *Params) error {
		p.Addr = net.JoinHostPort("", fmt.Sprint(port))
		return nil
	}
}

func WithEnvironment(env Env) Option {
	return func(p *Params) error {
		p.Environment = env
		return nil
	}
}

func WithLogLevel(level slog.Leveler) Option {
	return func(p *Params) error {
		handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: level,
		})

		p.Logger = slog.New(handler)
		return nil
	}
}

func WithJwtSecret(secret security.Key) Option {
	return func(p *Params) error {
		p.JWTSecret = secret
		return nil
	}
}

func WithSqliteDSN(dsn database.DSN) Option {
	return func(p *Params) error {
		db, err := database.OpenSQLite(dsn)
		if err != nil {
			return err
		}

		p.DB = db
		return nil
	}
}

func WithStorageRoot(path string) Option {
	return func(p *Params) error {
		root, err := os.OpenRoot(path)
		if err != nil {
			return err
		}

		p.StorageProvider = &storage.RootProvider{
			Root: root,
		}
		return nil
	}
}
