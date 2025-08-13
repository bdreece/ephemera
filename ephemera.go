package ephemera

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/bdreece/ephemera/internal/security"
	"github.com/bdreece/ephemera/pkg/database"
	"github.com/bdreece/ephemera/pkg/storage"
)

type Params struct {
	Addr            string
	Environment     Env
	DB              database.DB
	JWTSecret       security.Key
	StorageProvider storage.Provider
	Logger          *slog.Logger
}

type App struct {
	env Env
	srv http.Server
}

var (
	version   string
	metadata  string
	gitTag    string
	gitSha    string
	gitCommit string
)

func DebugInfo() (info struct {
	Version   string
	Metadata  string
	GitTag    string
	GitSHA    string
	GitCommit string
}) {
	info.Version = version
	info.Metadata = metadata
	info.GitTag = gitTag
	info.GitSHA = gitSha
	info.GitCommit = gitCommit
	return
}

func New(opts ...Option) *App {
	p := new(Params)
	for _, fn := range opts {
		fn(p)
	}

	if p.Logger == nil {
		_ = WithLogLevel(DefaultConfig.LogLevel)(p)
	}
	if p.DB == nil {
		_ = WithSqliteDSN(DefaultConfig.SqliteDSN)
	}
	if p.StorageProvider == nil {
		_ = WithStorageRoot(DefaultConfig.StorageRoot)
	}

	slog.SetDefault(p.Logger)

	app := App{
		env: p.Environment,
		srv: http.Server{
			Addr:    p.Addr,
			Handler: newRouter(p),
		},
	}

	return &app
}

func (app *App) Run(ctx context.Context) error {
	go app.start()
	slog.Info(
		"http server listening",
		"addr", app.srv.Addr,
		"environment", app.env,
		"version", version,
		"metadata", metadata,
		"tag", gitTag,
		"sha", gitSha,
		"commit", gitCommit,
	)

	<-ctx.Done()
	return app.stop()
}

func (app *App) start() {
	if err := app.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		slog.Error("server closed unexpectedly", "error", err)
	}
}

func (app *App) stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.srv.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
