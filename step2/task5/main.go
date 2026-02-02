package main

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

var ErrBadTimePeriod = errors.New("bad time period")
var ErrBadLogExtraction = errors.New("bad log extraction")
var ErrBadLogData = errors.New("bad log data")

func parseDate(line string, layout string) (time.Time, error) {
	if len(line) < 10 {
		return time.Time{}, ErrBadLogData
	}

	sDate := line[:10]
	date, err := time.Parse(layout, sDate)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func ExtractLog(fileName string, start time.Time, end time.Time) ([]string, error) {
	log := make([]string, 0)

	isGoodTimePeriod := start.Equal(end) || start.Before(end)
	if !isGoodTimePeriod {
		return nil, ErrBadTimePeriod
	}

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		date, err := parseDate(line, "02.01.2006")
		if err != nil {
			continue
		}

		isInPeriod := date.Equal(start) || date.Equal(end) || (date.After(start) && date.Before(end))
		if isInPeriod {
			log = append(log, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	if len(log) == 0 {
		return nil, ErrBadLogExtraction
	}

	return log, nil
}
