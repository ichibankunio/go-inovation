package assets

import (
	"embed"
)

//go:embed resources/*
var Assets embed.FS
