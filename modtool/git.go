package main

import (
	"strings"

	"github.com/sirupsen/logrus"
)

func newGit() *git {
	return &git{
		upstream: "origin",
	}
}

type git struct {
	upstream string
}

func (g *git) getRepoURL() (string, error) {
	res, err := exec("git remote get-url --all " + g.upstream)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func (g *git) getRemoteTags(prefix string) ([]string, error) {
	res, err := exec("git ls-remote --tags " + g.upstream)
	if err != nil {
		return nil, err
	}
	arr := strings.Split(string(res), "\n")
	ret := make([]string, 0, len(arr))
	for _, v := range arr {
		if len(v) == 0 {
			continue
		}
		vv := strings.Fields(v)
		if prefix == "" || strings.HasPrefix(vv[1], prefix) {
			ret = append(ret, strings.ReplaceAll(vv[1], "refs/tags/", ""))
		}
	}
	logrus.Debugf("tags: %v", ret)
	return ret, nil
}

func (g *git) getLocalTags(prefix string) ([]string, error) {
	res, err := exec("git tag")
	if err != nil {
		return nil, err
	}
	arr := strings.Split(string(res), "\n")
	ret := make([]string, 0, len(arr))
	for _, v := range arr {
		if len(v) == 0 {
			continue
		}
		if prefix == "" || strings.HasPrefix(v, prefix) {
			ret = append(ret, v)
		}
	}
	logrus.Debugf("local tags: %#+v", ret)
	return ret, nil
}
