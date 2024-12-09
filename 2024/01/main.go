package main

import (
	"bufio"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type LocationList []int

type LocationParser struct {
}

func (lb *LocationParser) Result() int {
	return 0
}

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

	var listA LocationList
	var listB LocationList

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

		a, err := strconv.Atoi(result[0])
		if err != nil {
			slog.Error("malformed file: could not convert to int")
			os.Exit(1)
		}
		listA = append(listA, a)

		b, err := strconv.Atoi(result[1])
		if err != nil {
			slog.Error("malformed file: could not convert to int")
			os.Exit(1)
		}
		listB = append(listB, b)

		slog.Info("a" + result[0])
		slog.Info("b" + result[1])
	}
}
