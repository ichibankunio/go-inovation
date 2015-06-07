package main

type GameMode int

const (
	GAMEMODE_NORMAL GameMode = iota
	GAMEMODE_LUNKER
)

var clearFlagItems = [...]FieldType{
	FIELD_ITEM_FUJI,
	FIELD_ITEM_TAKA,
	FIELD_ITEM_NASU,
}

func IsItemForClear(it FieldType) bool {
	for _, e := range clearFlagItems {
		if e == it {
			return true
		}
	}
	return false
}

type GameData struct {
	itemGetFlags [FIELD_ITEM_MAX]bool
	time         int
	jumpMax      int
	lifeMax      int
	lunkerMode   bool
}

func NewGameData(gameMode GameMode) *GameData {
	g := &GameData{}
	switch gameMode {
	case GAMEMODE_NORMAL:
		g.lifeMax = 3
		g.lunkerMode = false
	case GAMEMODE_LUNKER:
		g.lifeMax = 1
		g.lunkerMode = true
		g.jumpMax = 1
	}
	return g
}

func (g *GameData) Update() {
	g.time++
}

func (g *GameData) TimeInSecond() int {
	return g.time / 60
}

func (g *GameData) IsGameClear() bool {
	for _, e := range clearFlagItems {
		if !g.itemGetFlags[e] {
			return false
		}
	}
	return true
}

func (g *GameData) IsGetOmega() bool {
	return g.itemGetFlags[FIELD_ITEM_OMEGA]
}

func (g *GameData) GetItemCount() int {
	f := 0
	for _, b := range g.itemGetFlags {
		if b {
			f++
		}
	}
	return f
}
