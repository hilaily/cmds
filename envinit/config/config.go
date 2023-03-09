package config

import "os"

var (
	InCN = os.Getenv("INCN") == "1"
)
