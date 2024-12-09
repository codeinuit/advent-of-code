package main

import (
	"bufio"
	"log/slog"
	"os"
	"slices"
	"strconv"
	"strings"
)

type LocationList []int

type LocationParser struct {
	left  LocationList
	right LocationList
}

func (lb *LocationParser) AddLine(l, r int) {
	lb.left = append(lb.left, l)
	lb.right = append(lb.right, r)
}

func (lb *LocationParser) Result() int {
	slices.Sort(lb.left)
	slices.Sort(lb.right)

	res := 0

	for i := 0; i != len(lb.left); i += 1 {
		tmp := lb.left[i] - lb.right[i]

		if tmp < 0 {
			tmp = -tmp
		}

		res += tmp
	}

	return res
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

	lp := LocationParser{}

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

		b, err := strconv.Atoi(result[1])
		if err != nil {
			slog.Error("malformed file: could not convert to int")
			os.Exit(1)
		}

		lp.AddLine(a, b)
	}

	slog.With("result", lp.Result()).Info("finished")
}
