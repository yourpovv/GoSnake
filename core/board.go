package core

import (
	"strings"

	"github.com/charmbracelet/lipgloss"

)

type board struct {
	width  int
	height int
}

func newBoard(width, height int) *board {
	return &board{
		width:  width,
		height: height,
	}
}

func (b *board) render(snake *snake, food *food) string {
	var sb strings.Builder

	greenStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00"))
	topBottom := greenStyle.Render("╔" + strings.Repeat("═", b.width) + "╗")

	sb.WriteString(topBottom + "\n")

	for y := 0; y < b.height; y++ {
		sb.WriteString(greenStyle.Render("║"))

		for x := 0; x < b.width; x++ {
			p := point{x, y}

			if snake.head().equals(p) {
				sb.WriteString(lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FFD700")).
					Bold(true).
					Render("@"))
			} else if snake.contains(p) {
				sb.WriteString(lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FFD700")).
					Render("o"))
			} else if food.position.equals(p) {
				sb.WriteString(lipgloss.NewStyle().
					Foreground(lipgloss.Color("#FF0000")).
					Bold(true).
					Render("*"))
			} else {
				sb.WriteString(" ")
			}
		}

		sb.WriteString(greenStyle.Render("║") + "\n")
	}

	sb.WriteString(greenStyle.Render("╚"+strings.Repeat("═", b.width)+"╝") + "\n")

	return sb.String()
}

func (b *board) isOutOfBounds(p point) bool {
	return p.x < 0 || p.x >= b.width || p.y < 0 || p.y >= b.height
}
