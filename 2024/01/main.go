package main

import (
	"bufio"
	"log/slog"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		slog.Error("missing filename in arguments")
		os.Exit(1)
	}

	filename := os.Args[1]

	f, err := os.Open(filename)
	if err != nil {
		slog.With("message", err.Error()).Error("could not read file")
		os.Exit(1)
	}

	defer func() {
		if err = f.Close(); err != nil {
			slog.With("message", err.Error()).Error("could not close file properly")
		}
	}()

	r := bufio.NewReader(f)

	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		result := strings.Split(string(line), "   ")
		if len(result) != 2 {
			slog.Error("malformed file")
			os.Exit(1)
		}

		slog.Info("a" + result[0])
		slog.Info("b" + result[1])
	}
}
