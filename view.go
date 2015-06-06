package inovation5

type View struct {
	position PositionF
}

func (v *View) ToScreenPosition(p Position) Position {
	x := p.X - int(v.position.X) + g_width/2
	y := p.Y - int(v.position.Y) + g_height/2
	return Position{x, y}
}

func (v *View) GetPosition() PositionF {
	return v.position
}

func (v *View) SetPosition(p PositionF) {
	v.position = p
}
