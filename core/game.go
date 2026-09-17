package core

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type tickMsg time.Time

type gameState int

const (
	playing gameState = iota
	paused
	gameOver
)

type Game struct {
	board *board
	snake *snake
	food  *food
	score int
	state gameState
	speed time.Duration
}

func NewGame() *Game {
	width, height := 30, 20
	snake := newSnake(width/2, height/2)
	return &Game{
		board: newBoard(width, height),
		snake: snake,
		food:  newFood(width, height, snake),
		score: 0,
		state: playing,
		speed: 150 * time.Millisecond,
	}
}

func (g *Game) Init() tea.Cmd {
	return tick(g.speed)
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return g, tea.Quit
		case "r":
			if g.state == gameOver {
				newG := NewGame()
				return newG, tick(newG.speed)
			}
		case "p", " ":
			switch g.state {
			case playing:
				g.state = paused
			case paused:
				g.state = playing
			}
		case "up", "w", "k":
			g.snake.changeDirection(point{0, -1})
		case "down", "s", "j":
			g.snake.changeDirection(point{0, 1})
		case "left", "a", "h":
			g.snake.changeDirection(point{-1, 0})
		case "right", "d", "l":
			g.snake.changeDirection(point{1, 0})
		}
	case tickMsg:
		if g.state == playing {
			g.snake.move()

			if g.board.isOutOfBounds(g.snake.head()) || g.snake.hitsItself() {
				g.state = gameOver
				return g, nil
			}

			if g.snake.head().equals(g.food.position) {
				g.snake.grow()
				g.food.respawn(g.board.width, g.board.height, g.snake)
				g.score += 10

				if g.score%50 == 0 && g.speed > 50*time.Millisecond {
					g.speed -= 10 * time.Millisecond
				}
			}
		}
		return g, tick(g.speed)
	}

	return g, nil
}

func (g *Game) View() string {
	title := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFD700")).
		Bold(true).
		MarginBottom(1).
		Render("🐍 GoSnake")

	scoreText := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#00FFFF")).
		Render(fmt.Sprintf("Score: %d | Length: %d", g.score, len(g.snake.body)))

	gameBoard := g.board.render(g.snake, g.food)

	controls := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#888888")).
		MarginTop(1).
		Render("WASD/Arrows: Move | Space: Pause | Q: Quit")

	var statusBar string
	switch g.state {
	case paused:
		statusBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FFFF00")).
			Bold(true).
			Render("⏸ PAUSED")
	case gameOver:
		statusBar = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#FF0000")).
			Bold(true).
			Render("💀 GAME OVER | Press R to restart")
	default:
		statusBar = ""
	}

	view := fmt.Sprintf("%s\n%s\n\n%s\n%s", title, scoreText, gameBoard, controls)

	if statusBar != "" {
		view = fmt.Sprintf("%s\n\n%s", view, statusBar)
	}

	return lipgloss.NewStyle().
		MarginLeft(2).
		MarginTop(1).
		Render(view)
}

func tick(d time.Duration) tea.Cmd {
	return tea.Tick(d, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
