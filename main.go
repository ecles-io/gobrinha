package main

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	direita  uint8 = 0
	esquerda uint8 = 1
	cima     uint8 = 2
	baixo    uint8 = 3
	tamTelaX = 800.0
	tamTelaY = 600.0
	tamQuadradoCobra = 50.0
	tamQuadradoAleatorio = 25.0
)

type Segmento struct {
	posX float64
	posY float64
}

type Cobra struct {
	segmentos   []Segmento
	tamQuadrado float64
	direcao     uint8
	tamTelaX    float64
	tamTelaY    float64
}

type Game struct {
	cobra             *Cobra
	quadradoAleatX    float64
	quadradoAleatY    float64
	tamQuadradoAleatorio float64
	gameOver          bool
}

func (g *Cobra) Update() {
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
		g.segmentos[0].posX += 5
	case esquerda:
		g.segmentos[0].posX -= 5
	case cima:
		g.segmentos[0].posY -= 5
	case baixo:
		g.segmentos[0].posY += 5
	}

	ultimoSeg := g.segmentos[0]
	for i:= 1; i < len(g.segmentos); i++ {
		ultimaPosSeg := g.segmentos[i]
		g.segmentos[i] = ultimaPos
		ultimaPos = ultimaPosSeg 
	}
}

func (g *Cobra) Crescer() {
	ultimo := g.segmentos[len(g.segmentos)-1]
	novoSegmento := Segmento{posX:0, posY:0}
	switch g.direcao{
	case direita:
		novoSegmento.posY = ultimo.posY
		novoSegmento.posX = ultimo.posX - tamQuadradoCobra
	case esquerda:
		novoSegmento.posY = ultimo.posY
		novoSegmento.posX = ultimo.posX + tamQuadradoCobra
	case cima:
		novoSegmento.posX = ultimo.posX
		novoSegmento.posY = ultimo.posY + tamQuadradoCobra
	case baixo:
		novoSegmento.posX = ultimo.posX
		novoSegmento.posY = ultimo.posY - tamQuadradoCobra
	}
	g.segmentos = append(g.segmentos, novoSegmento)
}

func (g *Cobra) Draw(screen *ebiten.Image) {
	for i, segmento := range g.segmentos {
		r := uint8(255 * (len(g.segmentos) - i) / len(g.segmentos))
		gColor := uint8(255 * (len(g.segmentos) - i) / len(g.segmentos))
		b := uint8(255)

		col := color.RGBA{
			R: r,
			G: gColor,
			B: b,
			A: 255,
		}
		ebitenutil.DrawRect(screen, segmento.posX, segmento.posY, g.tamQuadrado, g.tamQuadrado, col)
	}
}

func (g *Game) Update() error {
	if g.gameOver {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			x, y := ebiten.CursorPosition()
			if x >= 350 && x <= 450 && y >= 300 && y <= 350 {
				g.restartGame()
			}
		}
		return nil
	}
	cabeca := g.cobra.segmentos[0]
	if cabeca.posX < 0 || cabeca.posX+g.cobra.tamQuadrado > g.cobra.tamTelaX || cabeca.posY < 0 || cabeca.posY+g.cobra.tamQuadrado > g.cobra.tamTelaY {
		g.gameOver = true
	}

	if cabeca.posX < g.quadradoAleatX+g.tamQuadradoAleatorio &&
		cabeca.posX+g.cobra.tamQuadrado > g.quadradoAleatX &&
		cabeca.posY < g.quadradoAleatY+g.tamQuadradoAleatorio &&
		cabeca.posY+g.cobra.tamQuadrado > g.quadradoAleatY {
		g.quadradoAleatX = float64(rand.Intn(int(g.cobra.tamTelaX - g.tamQuadradoAleatorio)))
		g.quadradoAleatY = float64(rand.Intn(int(g.cobra.tamTelaY - g.tamQuadradoAleatorio)))

		g.cobra.Crescer()
	}
	g.cobra.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.cobra.Draw(screen)
	col := color.RGBA{0, 255, 0, 255}
	ebitenutil.DrawRect(screen, g.quadradoAleatX, g.quadradoAleatY, g.tamQuadradoAleatorio, g.tamQuadradoAleatorio, col)
	if g.gameOver {
		ebitenutil.DebugPrint(screen, "Game Over! Pressione o botão para jogar novamente")
		ebitenutil.DrawRect(screen, 350, 300, 100, 50, color.RGBA{255, 0, 0, 255})
		ebitenutil.DebugPrintAt(screen, "Jogar Novamente", 355, 320)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.cobra.tamTelaX), int(g.cobra.tamTelaY)
}

func (g *Game) restartGame() {
	g.gameOver = false
	g.cobra = &Cobra{
		segmentos:   []Segmento{{posX: 100, posY: 100}},
		tamQuadrado: g.cobra.tamQuadrado,
		tamTelaX:    g.cobra.tamTelaX,
		tamTelaY:    g.cobra.tamTelaY,
		direcao:     direita,
	}
	g.quadradoAleatX = float64(rand.Intn(int(g.cobra.tamTelaX - g.tamQuadradoAleatorio)))
	g.quadradoAleatY = float64(rand.Intn(int(g.cobra.tamTelaY - g.tamQuadradoAleatorio)))
}

func main() {

	rand.Seed(time.Now().UnixNano())

	quadradoAleatX := float64(rand.Intn(int(tamTelaX - tamQuadradoAleatorio)))
	quadradoAleatY := float64(rand.Intn(int(tamTelaY - tamQuadradoAleatorio)))

	cobra := &Cobra{
		segmentos:   []Segmento{{posX: 100, posY: 100}},
		tamQuadrado: tamQuadradoCobra,
		tamTelaX:    tamTelaX,
		tamTelaY:    tamTelaY,
		direcao:     direita,
	}

	game := &Game{
		cobra:             cobra,
		quadradoAleatX:    quadradoAleatX,
		quadradoAleatY:    quadradoAleatY,
		tamQuadradoAleatorio: tamQuadradoAleatorio,
		gameOver:          false,
	}

	ebiten.SetWindowSize(int(tamTelaX), int(tamTelaY))
	ebiten.SetTPS(20)
	ebiten.SetWindowTitle("Cobra: Transição de Branco para Azul")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
