package steamworks

import (
	_ "embed"
)

//go:embed steam_api.dll
var steamAPIDLL []byte
