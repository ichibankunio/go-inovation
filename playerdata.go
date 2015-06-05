package inovation5

const (
	GAMEMODE_NORMAL = iota
	GAMEMODE_LUNKER
)

var clearFlagItems = [...]int{
	FIELD_ITEM_FUJI,
	FIELD_ITEM_TAKA,
	FIELD_ITEM_NASU,
}

type PlayerData struct {
	itemGetFlags [FIELD_ITEM_MAX]bool
	playtime     int
	jumpMax      int
	lifeMax      int
	lunkerMode   bool
}

func NewPlayerData(gameMode int) *PlayerData {
	p := &PlayerData{}
	switch gameMode {
	default:
		fallthrough
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

func (p *PlayerData) IsGameClear() bool {
	for _, e := range clearFlagItems {
		if !p.itemGetFlags[e] {
			return false;
		}
        }
        return true;
}

func (p *PlayerData) IsItemForClear(it int) bool {
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
