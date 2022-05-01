package errors

import (
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ReadFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, errors.Wrap(err, "can not open file")
	}
	defer f.Close()

	buf, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "can not read file")
	}
	return buf, nil
}

func ReadConfig() ([]byte, error) {
	home := os.Getenv("HOME")

	config, err := ReadFile(filepath.Join(home, ".setting.xml"))

	return config, errors.WithMessage(err, "can not read config setting.xml")
}
