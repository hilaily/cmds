package main

import "strings"

type git struct {
}

func (g *git) getRepoPath(gitRepoURL string) string {
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
