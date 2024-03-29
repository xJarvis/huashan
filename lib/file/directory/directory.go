package directory

import (
	"os"
	"path/filepath"
)


func GetWorkDir() (string, error) {
	wd, err := os.Getwd()
	return wd, err
}


func GetExecDir() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	ed := filepath.Dir(execPath)
	return ed, nil
}

func IsAbs(path string) bool {
	return len(path) > 0 && path[0] == '/'
}