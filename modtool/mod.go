package main

import "strings"

func newMod() (*mod, error) {
	m := &mod{}
	f, err := m.findModFile()
	if err != nil {
		return nil, err
	}
	m.modName, err = m.getModName(f)
	return m, err
}

type mod struct {
	modName string
}

func (m *mod) getModPrefix(repoPath, modName string) string {
	// https://github.com/hilaily/cmds.git
	// module github.com/hilaily/cmds/gotool
	return strings.ReplaceAll(modName, repoPath+"/", "")
}

func (m *mod) getImportPath(gitRepoURL string) string {
	if strings.HasPrefix(gitRepoURL, "http") {
		// https://github.com/hilaily/cmds.git
		u := strings.TrimSuffix(gitRepoURL, ".git")
		u = strings.TrimPrefix(u, "https://")
		u = strings.TrimPrefix(u, "http://")
		return u
	}
	// git@github.com:hilaily/cmds.git
	u := strings.TrimSuffix(gitRepoURL, ".git")
	i := strings.Index(u, "@")
	u = u[i+1:]
	u = strings.ReplaceAll(u, ":", "/")
	return u
}
