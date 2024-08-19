// package main

// import (
// 	"log"

// 	"github.com/hajimehoshi/ebiten/v2"
// 	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
// )

// type Game struct{}

// func (g *Game) Update() error {
// 	return nil
// }

// func (g *Game) Draw(screen *ebiten.Image) {
// 	ebitenutil.DebugPrint(screen, "Hello, World!")
// }

// func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
// 	return 320, 240
// }

// func main() {
// 	ebiten.SetWindowSize(640, 480)
// 	ebiten.SetWindowTitle("Hello, World!")
// 	if err := ebiten.RunGame(&Game{}); err != nil {
// 		log.Fatal(err)
// 	}
// }
// HLine draws a horizontal line

package main

import (
    "github.com/fogleman/gg"
)

func main() {
    // Define o tamanho da imagem
    const width = 500
    const height = 500

    // Cria um novo contexto de desenho
    dc := gg.NewContext(width, height)

    // Define a cor de fundo como branco
    dc.SetRGB(1, 1, 1)
    dc.Clear()

    // Define a cor do quadrado como vermelho
    dc.SetRGB(1, 0, 0)

    // Desenha um quadrado no centro da imagem
    size := 200.0
    x := (width - size) / 2
    y := (height - size) / 2
    dc.DrawRectangle(x, y, size, size)

    // Preenche o quadrado com a cor definida
    dc.Fill()

    // Salva a imagem em um arquivo PNG
    dc.SavePNG("square.png")
}