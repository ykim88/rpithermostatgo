package driver

import (
	"errors"
	"strconv"
	"strings"
	"syscall"
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

	temperatureValue, err := strconv.ParseFloat(raw[indexTemp+2:len(raw)], 64)
	if err != nil {
		return errorValue, err
	}

	return temperatureValue, nil
}

func readFile(path string) (string, error) {
	fd, err := syscall.Open(path, syscall.O_RDONLY|syscall.O_CLOEXEC|syscall.O_NONBLOCK, 0644)
	defer syscall.Close(fd)
	if err != nil {
		return "", err
	}
	stat := syscall.Stat_t{}
	err = syscall.Stat(path, &stat)
	if err != nil {
		return "", err
	}

	buffer := make([]byte, stat.Size)
	_, err = syscall.Read(fd, buffer)
	if err != nil {
		return "", err
	}

	return string(buffer), nil
}
