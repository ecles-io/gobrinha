package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	"log"
	"math/rand"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
	snakeSize    = 10
	tickRate     = 5 // Define a velocidade da cobra (quanto maior, mais lenta)
)

type Point struct {
	X, Y int
}

type Game struct {
	snake     []Point
	direction Point
	food      Point
	started   bool
	tickCount int
}

func (g *Game) Update() error {
	// Controle da direção
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) && g.direction.Y != 1 {
		g.direction = Point{X: 0, Y: -1}
		g.started = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) && g.direction.Y != -1 {
		g.direction = Point{X: 0, Y: 1}
		g.started = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) && g.direction.X != 1 {
		g.direction = Point{X: -1, Y: 0}
		g.started = true
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) && g.direction.X != -1 {
		g.direction = Point{X: 1, Y: 0}
		g.started = true
	}

	// Se o jogo não começou, não atualiza a posição da cobra
	if !g.started {
		return nil
	}

	// Controlar a velocidade da cobra
	g.tickCount++
	if g.tickCount < tickRate {
		return nil
	}
	g.tickCount = 0

	// Atualizar a posição da cobra
	head := g.snake[0]
	head.X += g.direction.X * snakeSize
	head.Y += g.direction.Y * snakeSize

	// Inserir nova cabeça da cobra
	g.snake = append([]Point{head}, g.snake...)

	// Verificar se a cobra comeu a comida
	if head.X == g.food.X && head.Y == g.food.Y {
		g.food = g.newFoodPosition()
	} else {
		// Remover a cauda se a cobra não comeu
		g.snake = g.snake[:len(g.snake)-1]
	}

	// Verificar colisão com a parede
	if head.X < 0 || head.Y < 0 || head.X >= screenWidth || head.Y >= screenHeight {
		return ebiten.Termination
	}

	// Verificar colisão com o próprio corpo
	for _, s := range g.snake[1:] {
		if head.X == s.X && head.Y == s.Y {
			return ebiten.Termination
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Desenhar a cobra
	for _, s := range g.snake {
		ebitenutil.DrawRect(screen, float64(s.X), float64(s.Y), snakeSize, snakeSize, color.RGBA{0, 255, 0, 255})
	}

	// Desenhar a comida
	ebitenutil.DrawRect(screen, float64(g.food.X), float64(g.food.Y), snakeSize, snakeSize, color.RGBA{255, 0, 0, 255})
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) newFoodPosition() Point {
	return Point{
		X: rand.Intn(screenWidth/snakeSize) * snakeSize,
		Y: rand.Intn(screenHeight/snakeSize) * snakeSize,
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	game := &Game{
		snake: []Point{
			{X: screenWidth / 2, Y: screenHeight / 2},
		},
		direction: Point{X: 0, Y: 0}, // Cobra começa parada
		food:      Point{},
	}

	game.food = game.newFoodPosition()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Snake Game")

		if err := ebiten.RunGame(game); err != nil && err != ebiten.Termination {
			log.Fatal(err)
		}
	}