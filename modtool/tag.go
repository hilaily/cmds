package main

import (
	"fmt"
	"strings"

	"github.com/hilaily/kit/listx"
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

func (t *tag) do(cmd string, args ...string) {
	switch cmd {
	case "show":
		t.show()
	case "newtag":
		t.newTag(verType(args[0]))
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
	logrus.Debugf("mod prefix: %s", prefix)
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

func (t *tag) newTag(typ verType) {
	tags, err := t.getAllTags()
	if err != nil {
		pRed(err.Error())
		return
	}
	removePrefixTags := make([]string, 0, len(tags))
	modPrefix, _ := t.getModPrefix()
	for _, v := range tags {
		removePrefixTags = append(removePrefixTags,
			strings.ReplaceAll(v, modPrefix+"/", ""),
		)
	}
	logrus.Debugf("remove prefix tags: %v", removePrefixTags)

	preReleasePrefix := string(typ)
	v, err := t.newTagVersion(removePrefixTags, typ, preReleasePrefix)
	if err != nil {
		pRed(err.Error())
		return
	}
	pNomarl(modPrefix + "/v" + v.String())
}

func (t *tag) help() {
	logrus.Info("this is a help")
}

func (t *tag) newTagVersion(tags []string, typ verType, preReleasePrefix string) (*version, error) {
	res, has, err := max(tags, preReleasePrefix)
	if err != nil {
		return nil, err
	}
	var ver *version
	if !has {
		ver = firstVersion(typ, preReleasePrefix)
		return ver, nil
	}
	ver = res.inc(typ, preReleasePrefix)
	return ver, nil
}

func (t *tag) getModPrefix() (string, error) {
	if t._modPrefix == "" {
		u, err := t.git.getRepoURL()
		if err != nil {
			return "", fmt.Errorf("get repo url fail, %w", err)
		}
		logrus.Debugf("get repo url: %s, mod name: %s", u, t.mod.modName)
		p := t.mod.getRepoPath(u)
		logrus.Debug("get repo path: ", p)
		t._modPrefix = t.mod.getModPrefix(p, t.mod.modName)
	}
	return t._modPrefix, nil
}

func (t *tag) getAllTags() ([]string, error) {
	p, err := t.getModPrefix()
	if err != nil {
		return nil, err
	}
	res, err := t.git.getRemoteTags(p)
	if err != nil {
		return nil, err
	}
	res2, err := t.git.getLocalTags(p)
	if err != nil {
		return nil, err
	}
	res = append(res, res2...)
	res = listx.Dedup(res)
	return res, nil
}
