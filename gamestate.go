package inovation5

type GameStateMsg int

const (
	GAMESTATE_MSG_NONE GameStateMsg = iota
	GAMESTATE_MSG_REQ_TITLE
	GAMESTATE_MSG_REQ_GAME
	GAMESTATE_MSG_REQ_OPENING
	GAMESTATE_MSG_REQ_ENDING
	GAMESTATE_MSG_REQ_SECRET1
	GAMESTATE_MSG_REQ_SECRET2
)

type GameState struct {
	msg GameStateMsg
}

func (g *GameState) GetMsg() GameStateMsg {
	return g.msg
}

func (g *GameState) SetMsg(m GameStateMsg) {
	g.msg = m
}
