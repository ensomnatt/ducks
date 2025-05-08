package logger

import (
	"os"

	"github.com/ensomnatt/ducks/internal/config"
	"log/slog"
)

var Log *slog.Logger

func init() {
	env := config.GetConfig().Env

	var handler slog.Handler

	switch env {
	case "local", "", "dev":
		handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		})
	case "prod":
		handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelInfo,
		})
	}

	Log = slog.New(handler)
	slog.SetDefault(Log)
}
