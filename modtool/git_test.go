package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetLocalTags(t *testing.T) {
	g := newGit()
	_, err := g.getLocalTags("")
	assert.NoError(t, err)
}

func TestGetRemoteTags(t *testing.T) {
	g := newGit()
	_, err := g.getRemoteTags("")
	assert.NoError(t, err)
}
