package core

type point struct {
	x, y int
}

func (p point) equals(other point) bool {
	return p.x == other.x && p.y == other.y
}

func (p point) add(other point) point {
	return point{p.x + other.x, p.y + other.y}
}
