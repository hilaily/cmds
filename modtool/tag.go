package main

import (
	"fmt"
	"strings"

	"github.com/hilaily/kit/listx"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func tagCommand() *cli.Command {
	return &cli.Command{
		Name:  "tag",
		Usage: "for create or show git tag of this go module",
		Subcommands: []*cli.Command{
			tagShowCommand(),
			tagNewCommand(),
		},
	}
}

func tagShowCommand() *cli.Command {
	return &cli.Command{
		Name:  "show",
		Usage: "show git tag of this go module",
		Action: func(*cli.Context) error {
			t, err := newTag()
			if err != nil {
				pRed(err.Error())
				return err
			}
			t.show()
			return nil
		},
	}
}

func tagNewCommand() *cli.Command {
	return &cli.Command{
		Name:  "new",
		Usage: "create a new git tag of this go module",
		Description: `For example:
modtool tag new minor
modtool tag new patch 
modtool tag new alpha
modtool tag new -p=false beta`,
		ArgsUsage: "specify type of version, like major/minor/patch and pre release version prefix",
		Action: func(c *cli.Context) error {
			t, err := newTag()
			if err != nil {
				pRed(err.Error())
				return err
			}
			typ := c.Args().Get(0)
			if typ == "" {
				return fmt.Errorf("you should specify what type of version(mojor/minor/pre)\nlike modtool tag new patch")
			}
			t.newTag(verType(typ), c.Bool("push"))
			return nil
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "push",
				Aliases: []string{"p"},
				Value:   true,
				Usage:   "get new tag name, add tag and push to git remote",
			},
		},
	}
}

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

func (t *tag) show() {
	prefix, err := t.getModPrefix()
	if err != nil {
		pRed(err.Error())
		return
	}
	logrus.Debugf("mod prefix: %s", prefix)
	r, err := t.git.getRemoteTags(prefix)
	if err != nil {
		pRed("get remote tags err: %s", err.Error())
		return
	}
	r, err = sortVer(prefix, r)
	if err != nil {
		pRed("get remote tags err: %s", err.Error())
		return
	}
	logrus.Debugf("remote tags: %v", r)
	pBlue("remote tags: ")
	pTable(r)
	//pNomarl(strings.Join(r, "\t\t"))
	pBlue("\nlocal tags: ")
	r, err = t.git.getLocalTags(prefix)
	if err != nil {
		pRed("get remote tags err: %s", err.Error())
		return
	}
	r, err = sortVer(prefix, r)
	if err != nil {
		pRed("get remote tags err: %s", err.Error())
		return
	}
	pTable(r)
	//pNomarl(strings.Join(r, "\t\t"))
}

func (t *tag) newTag(typ verType, push bool) {
	logrus.Debugf("get push: %v", push)
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
	v, latest, err := t.newTagVersion(removePrefixTags, typ, preReleasePrefix)
	if err != nil {
		pRed(err.Error())
		return
	}
	newTag := v.String(modPrefix)
	if latest == nil {
		pBlue("there is no old tags")
	} else {
		pBlue("latset tag is:\n")
		pNomarl("    %s\n", latest.String(modPrefix))
	}
	pBlue("will create tag:\n")
	pNomarl("    %s\n", newTag)
	if push {
		pBlue("create tag local")
		_, err := t.git.createTagLocal(newTag)
		if err != nil {
			pRed(err.Error())
			return
		}
		pBlue("push tag to remote")
		res, err := t.git.pushTagRemote(newTag)
		if err != nil {
			pRed(err.Error())
			return
		}
		pNomarl(res)
	}
}

// return 1: the new tag
// return 2: latest tag
func (t *tag) newTagVersion(tags []string, typ verType, preReleasePrefix string) (*version, *version, error) {
	res, has, err := latest(tags, preReleasePrefix)
	if err != nil {
		return nil, nil, err
	}
	var ver *version
	if !has {
		ver = firstVersion(typ, preReleasePrefix)
		return ver, nil, nil
	}
	ver = res.inc(typ, preReleasePrefix)
	return ver, res, nil
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
