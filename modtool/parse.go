package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/hilaily/kit/pathx"
	"github.com/sirupsen/logrus"
	"golang.org/x/mod/modfile"
)

func findModFile() (string, error) {
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
			return cur, nil
		}
		if cur == "/" {
			break
		}
		cur = filepath.Dir(cur)
	}
	return cur, fmt.Errorf("can not find go.mod")
}

func getModName(modPath string) (string, error) {
	en, err := os.ReadFile(modPath)
	if err != nil {
		return "", err
	}
	f, err := modfile.Parse(modPath, en, nil)
	if err != nil {
		return "", err
	}
	return f.Module.Mod.Path, nil
}
