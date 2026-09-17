package core

type snake struct {
	body          []point
	direction     point
	nextDirection point
	growing       bool
}

func newSnake(startX, startY int) *snake {
	return &snake{
		body: []point{
			{startX, startY},
			{startX - 1, startY},
			{startX - 2, startY},
		},
		direction:     point{1, 0},
		nextDirection: point{1, 0},
		growing:       false,
	}
}

func (s *snake) head() point {
	return s.body[0]
}

func (s *snake) move() {
	s.direction = s.nextDirection
	newHead := s.head().add(s.direction)
	s.body = append([]point{newHead}, s.body...)

	if !s.growing {
		s.body = s.body[:len(s.body)-1]
	} else {
		s.growing = false
	}
}

func (s *snake) changeDirection(dir point) {
	if s.direction.x+dir.x == 0 && s.direction.y+dir.y == 0 {
		return
	}
	if s.nextDirection.x+dir.x == 0 && s.nextDirection.y+dir.y == 0 {
		return
	}
	s.nextDirection = dir
}

func (s *snake) grow() {
	s.growing = true
}

func (s *snake) hitsItself() bool {
	head := s.head()
	for i := 1; i < len(s.body); i++ {
		if head.equals(s.body[i]) {
			return true
		}
	}
	return false
}

func (s *snake) contains(p point) bool {
	for _, segment := range s.body {
		if segment.equals(p) {
			return true
		}
	}
	return false
}
