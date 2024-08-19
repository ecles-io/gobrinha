package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
)

const (
	direita uint8 = 0
	esquerda uint8 = 1
)

type Game struct{
	tamQuadrado float64
	posicaoDoQuadrado float64
	tamTela float64
	direcao uint8
}

func (g *Game) Update() error {
	if g.direcao == direita {
		if (g.posicaoDoQuadrado + g.tamQuadrado) < g.tamTela {
			g.posicaoDoQuadrado += 5
		
		} else {
			g.direcao = esquerda
		}
	} else if g.direcao == esquerda {
		if g.posicaoDoQuadrado > 0 {
			g.posicaoDoQuadrado -= 5
		} else {
			g.direcao = direita
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	col := color.RGBA{255, 1, 127, 255} 

	y:= 100  
	ebitenutil.DrawRect(screen, g.posicaoDoQuadrado, float64(y), g.tamQuadrado, g.tamQuadrado, col)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 640, 480
}

func main() {
	game := &Game{posicaoDoQuadrado:0, tamTela: 640.0, direcao: direita, tamQuadrado: 200.0}
	ebiten.SetWindowSize(640, 480)
	ebiten.SetTPS(60)
	ebiten.SetWindowTitle("Quadrado com BitEngine")
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
