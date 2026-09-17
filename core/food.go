package core

import (
	"math/rand"
	"time"
)

type food struct {
	position point
	rng      *rand.Rand
}

func newFood(width, height int, snake *snake) *food {
	f := &food{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	f.respawn(width, height, snake)
	return f
}

func (f *food) spawn(width, height int) {
	f.position = point{
		x: f.rng.Intn(width),
		y: f.rng.Intn(height),
	}
}

func (f *food) respawn(width, height int, snake *snake) {
	for {
		f.spawn(width, height)
		if !snake.contains(f.position) {
			break
		}
	}
}
