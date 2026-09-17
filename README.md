<div align="center">
  
# GoSnake

**Terminal snake game built with Go, Bubbletea and Lipgloss.**

[![GitHub](https://img.shields.io/github/stars/yourpovv/GoSnake?style=social)](https://github.com/yourpovv/GoSnake)

https://github.com/user-attachments/assets/26910b70-7c5d-4c22-bbfa-f19bc4ea7357

</div>

## Requirements

- Go 1.21+

## Running it

```bash
go run .
```

or build it first:

```bash
go build -o gosnake.exe
.\gosnake.exe
```

## Controls

- WASD, Arrow keys, or Vim keys (hjkl) to move
- Space to pause
- Q to quit
- R to restart after game over

## How it works

The snake moves around a 30x20 grid eating food to grow longer. Game speeds up every 50 points. Don't hit the walls or yourself

The game ticks every 150ms initially, getting faster as you score. Direction changes are buffered to prevent reversing into yourself between ticks

## License

[MIT](LICENSE) © [YourPOVV](https://github.com/yourpovv)
