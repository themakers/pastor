package tmp

import (
	"log/slog"
	"os"
	"path/filepath"
)

type tmp struct {
	Path string
}

func (t tmp) Purge() {
	if err := os.RemoveAll(t.Path); err != nil {
		slog.Warn("failed to remove tmp file/dir",
			slog.String("path", t.Path),
			slog.Any("error", err),
		)
	}
}

func File(name string) tmp {
	f, err := os.CreateTemp("", name)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	return tmp{Path: f.Name()}
}

func Dir(name string) tmp {
	path, err := os.MkdirTemp("", name)
	if err != nil {
		panic(err)
	}
	return tmp{Path: filepath.Clean(path)}
}
