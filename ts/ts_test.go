package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestA(t *testing.T) {
	res, err := transfer("1669189503")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	t.Log(res)

	res, err = transfer("2022-11-22 13:00:00")
	assert.NoError(t, err)
	assert.NotEmpty(t, res)
	t.Log(res)
}
