package steamworks

import (
	_ "embed"
)

//go:embed libsteam_api.so
var libSteamAPI []byte
