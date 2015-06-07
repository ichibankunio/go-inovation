package inovation5

type View struct {
	position PositionF
}

func NewView(position PositionF) *View {
	return &View{
		position: position,
	}
}

func (v *View) ToScreenPosition(p PositionF) PositionF {
	x := p.X - v.position.X + g_width/2
	y := p.Y - v.position.Y + g_height/2
	return PositionF{x, y}
}

func (v *View) GetPosition() PositionF {
	return v.position
}

func (v *View) Update(position, speed PositionF) {
	const VIEW_DIRECTION_OFFSET = 30.0
	v.position.X = v.position.X*0.95 + (position.X+speed.X*VIEW_DIRECTION_OFFSET)*0.05
	v.position.Y = v.position.Y*0.95 + position.Y*0.05
}
