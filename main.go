package slog4

import "golang.org/x/exp/slog"

var Log *slog.Logger

func Print(msg string) {
	Log.Info("print", "msg", msg)
}
