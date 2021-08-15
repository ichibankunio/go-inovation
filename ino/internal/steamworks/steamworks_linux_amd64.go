package steamworks

import (
	_ "embed"
)

//go:embed libsteam_api64.so
var libSteamAPI []byte
