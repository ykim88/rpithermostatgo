package driver

import (
	"bytes"
	"errors"
	"os"
	"strconv"
	"strings"
)

const errorValue float64 = -56000

func NewDriver(path string) *fsdriver {
	return &fsdriver{sysfspath: path}
}

type fsdriver struct {
	sysfspath string
}

func (d *fsdriver) Read() (float64, error) {
	raw, err := readFile(d.sysfspath)
	if err != nil {
		return 0.0, err
	}

	indexTemp := strings.LastIndex(raw, "t=")
	if indexTemp == -1 {
		return errorValue, errors.New("failed to read sensor temperature")
	}

	temperatureValue, err := strconv.ParseFloat(raw[indexTemp+2:len(raw)-1], 64)
	if err != nil {
		return errorValue, err
	}

	return temperatureValue, nil
}

func readFile(path string) (string, error) {
	fd, err := os.Open(path)
	defer fd.Close()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.ReadFrom(fd)
	return string(buffer.Bytes()), nil
}
