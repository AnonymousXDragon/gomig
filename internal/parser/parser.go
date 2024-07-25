package parser

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type Migration struct {
	Version int
	Name    string
	Up      string
	Down    string
}

func ParseMigFile(path string) (*Migration, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	var upContent, downContent strings.Builder
	var isUp bool

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line := scanner.Text()

		if strings.Contains(line, "-- =migrate:up") {
			isUp = true
			continue
		}

		if strings.Contains(line, "-- =migrate:down") {
			isUp = false
			continue
		}

		if isUp {
			upContent.WriteString(line + "\n")

		} else {
			downContent.WriteString(line + "\n")
		}

	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	fileName := filepath.Base(path)
	version, name, _ := strings.Cut(fileName, "_")

	versionInt, err := strconv.Atoi(version)
	if err != nil {
		return nil, err
	}

	mig := &Migration{
		Version: versionInt,
		Name:    strings.TrimSuffix(name, ".sql"),
		Up:      upContent.String(),
		Down:    downContent.String(),
	}

	return mig, nil
}
