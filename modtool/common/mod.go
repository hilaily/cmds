package common

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hilaily/kit/pathx"
	"github.com/sirupsen/logrus"
	"golang.org/x/mod/modfile"
)

// FindModFile ...
func FindModFile() (string, error) {
	cur, err := os.Getwd()
	if err != nil {
		return "", err
	}
	logrus.Debug("current dir: ", cur)
	exist := false
	for {
		f := filepath.Join(cur, "go.mod")
		exist = pathx.IsExist(f)
		logrus.Debugf("file: %s, exist: %v", f, exist)
		if exist {
			return f, nil
		}
		if cur == "/" {
			break
		}
		cur = filepath.Dir(cur)
	}
	return cur, fmt.Errorf("can not find go.mod")
}

// GetModName ...
func GetModName(modFilePath string) (string, error) {
	en, err := os.ReadFile(modFilePath)
	if err != nil {
		return "", err
	}
	f, err := modfile.Parse(modFilePath, en, nil)
	if err != nil {
		return "", err
	}
	return f.Module.Mod.Path, nil
}
