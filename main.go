package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

const (
	direita  uint8 = 0
	esquerda uint8 = 1
	cima uint8 = 2
	baixo uint8 = 3
)
type Cobra struct{
	tamQuadrado float64
	posX        float64
	posY        float64
	tamTelaX    float64
	tamTelaY    float64
	direcao     uint8
}

type Game struct {
	tamTela float64
  cobra *Cobra
}
func (g *Cobra) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.direcao = direita
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.direcao = esquerda
	} else if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.direcao = cima
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.direcao = baixo
	}

	switch g.direcao {
	case direita:
		if g.posX+g.tamQuadrado < g.tamTelaX {
			g.posX += 5
		}
	case esquerda:
		if g.posX > 0 {
			g.posX -= 5
		}
	case cima:
		if g.posY > 0 {
			g.posY -= 5
		}
	case baixo:
		if g.posY+g.tamQuadrado < g.tamTelaY {
			g.posY += 5
		}
	}
	return nil
}

func (g *Cobra) Draw(screen *ebiten.Image) {
	col := color.RGBA{255, 0, 127, 255}
	ebitenutil.DrawRect(screen, g.posX, g.posY, g.tamQuadrado, g.tamQuadrado, col)
}



func (g *Game) Update() error {
	return g.cobra.Update()
}
func (g *Game) Draw(screen *ebiten.Image) {
	g.cobra.Draw(screen)
}

func (g *Cobra) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.tamTelaX), int(g.tamTelaY)
}

func main() {
	cobra := &Cobra{
		tamQuadrado: 50,
		posX:        100,
		posY:        100,
		tamTelaX:    800,
		tamTelaY:    600,
		direcao:     cima, 
	}
	ebiten.SetWindowSize(int(cobra.tamTelaX), int(cobra.tamTelaY))
	ebiten.SetTPS(60)
	ebiten.SetWindowTitle("Mover Quadrado com Setas")
	if err := ebiten.RunGame(cobra); err != nil {
		panic(err)
	}
}
