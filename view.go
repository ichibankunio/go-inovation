package inovation5

type View struct {
	position PositionF
}

func (v *View) ToScreenPosition(p PositionF) PositionF {
	x := p.X - v.position.X + g_width/2
	y := p.Y - v.position.Y + g_height/2
	return PositionF{x, y}
}

func (v *View) GetPosition() PositionF {
	return v.position
}

func (v *View) SetPosition(p PositionF) {
	v.position = p
}
