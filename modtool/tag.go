package main

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

func newTag() (*tag, error) {
	m, err := newMod()
	if err != nil {
		return nil, err
	}
	g := newGit()
	return &tag{
		mod: m,
		git: g,
	}, nil
}

type tag struct {
	mod        *mod
	git        *git
	_modPrefix string
}

func (t *tag) do(cmd string) {
	switch cmd {
	case "show":
		t.show()
	default:
		t.help()
	}

}

func (t *tag) show() {
	prefix, err := t.getModPrefix()
	if err != nil {
		pRed(err.Error())
		return
	}
	pNomarl("remote tags: ")
	r, err := t.git.getRemoteTags(prefix)
	if err != nil {
		pRed("get remote tags err: %s", err.Error())
		return
	}
	pNomarl(strings.Join(r, "\t"))
	pNomarl("local tags: ")
	r, err = t.git.getLocalTags(prefix)
	if err != nil {
		pRed("get remote tags err: %s", err.Error())
		return
	}
	pNomarl(strings.Join(r, "\t"))
}

func (t *tag) help() {
	logrus.Info("this is a help")
}

func (t *tag) getModPrefix() (string, error) {
	if t._modPrefix == "" {
		u, err := t.git.getRepoURL()
		if err != nil {
			return "", fmt.Errorf("get repo url fail, %w", err)
		}
		p := t.mod.getImportPath(u)
		t._modPrefix = t.mod.getModPrefix(p, t.mod.modName)
	}
	return t._modPrefix, nil
}
