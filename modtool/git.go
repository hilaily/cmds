package main

import (
	"fmt"
	"strings"

	"github.com/hilaily/kit/listx"
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
		return nil, fmt.Errorf("get remote tags, output: %s, %w", string(res), err)
	}
	arr := strings.Split(string(res), "\n")
	ret := make([]string, 0, len(arr))
	for _, v := range arr {
		if len(v) == 0 {
			continue
		}
		vv := strings.Fields(v)
		t := strings.ReplaceAll(vv[1], "refs/tags/", "")
		t = strings.TrimSuffix(t, "^{}")
		logrus.Debugf("remote res: %v", t)
		if prefix == "" && !strings.Contains(t, "/") ||
			prefix != "" && strings.HasPrefix(t, prefix) {
			logrus.Debugf("remote res in: %v", t)
			ret = append(ret, t)
		}
	}
	ret = listx.Dedup(ret)
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
		if prefix == "" && !strings.Contains(v, "/") ||
			prefix != "" && strings.HasPrefix(v, prefix) {
			ret = append(ret, v)
		}
	}
	ret = listx.Dedup(ret)
	logrus.Debugf("local tags: %#+v", ret)
	return ret, nil
}

func (g *git) createTagLocal(tag string) (string, error) {
	res, err := exec("git tag " + tag)
	if err != nil {
		return "", fmt.Errorf("new tag fail %w", err)
	}
	return string(res), nil
}

func (g *git) pushTagRemote(tag string) (string, error) {
	res2, err := exec(fmt.Sprintf("git push %s %s", g.upstream, tag))
	if err != nil {
		return "", fmt.Errorf("push tag fail %w", err)
	}
	return string(res2), nil
}
