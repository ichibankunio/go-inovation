package steamworks

import (
	_ "embed"
)

//go:embed steam_api64.dll
var steamAPIDLL []byte
