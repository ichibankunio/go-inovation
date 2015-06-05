package inovation5

const (
	GAMESTATE_MSG_NONE= iota
	GAMESTATE_MSG_REQ_TITLE
	GAMESTATE_MSG_REQ_GAME
	GAMESTATE_MSG_REQ_OPENING
	GAMESTATE_MSG_REQ_ENDING
	GAMESTATE_MSG_REQ_SECRET1
	GAMESTATE_MSG_REQ_SECRET2
)

type GameState struct {
	game *Game
	msg  int
}

func (g *GameState) GetMsg() int {
	return g.msg
}

func (g *GameState) SetMsg(m int) {
	g.msg = m
}

func (g *GameState) Draw() {
	print("GameState draw")
}

func (g *GameState) Update() {
	print("GameState update")
}

