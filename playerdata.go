package inovation5

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

type PlayerData struct {
	itemGetFlags [FIELD_ITEM_MAX]bool
	time         int
	jumpMax      int
	lifeMax      int
	lunkerMode   bool
}

func NewPlayerData(gameMode GameMode) *PlayerData {
	p := &PlayerData{}
	switch gameMode {
	case GAMEMODE_NORMAL:
		p.lifeMax = 3
		p.lunkerMode = false
	case GAMEMODE_LUNKER:
		p.lifeMax = 1
		p.lunkerMode = true
		p.jumpMax = 1
	}
	return p
}

func (p *PlayerData) Update() {
	p.time++
}

func (p *PlayerData) TimeInSecond() int {
	return p.time / 60
}

func (p *PlayerData) IsGameClear() bool {
	for _, e := range clearFlagItems {
		if !p.itemGetFlags[e] {
			return false;
		}
        }
        return true;
}

func (p *PlayerData) IsItemForClear(it FieldType) bool {
	for _, e := range clearFlagItems {
		if e == it {
			return true;
		}
        }
        return false;
}

func (p *PlayerData) IsGetOmega() bool {
	return p.itemGetFlags[FIELD_ITEM_OMEGA]
}


func (p *PlayerData) GetItemCount() int {
	f := 0
	for _, b := range p.itemGetFlags {
		if b {
			f++
		}
	}
	return f
}
