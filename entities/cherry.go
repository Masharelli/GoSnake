package entities

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

// Cherry : Object which snakes eats

type Cherry struct {
	game   *Game
	cherry ebiten.Image
	xLimit int
	yLimit int
	xPos   float64
	yPos   float64
	eaten  bool
}

// CreateCherry : Generates a Cherry
func CreateCherry(g *Game) *Cherry {
	c := Cherry{
		game:   g,
		xLimit: 30,
		yLimit: 30,
		eaten:  false,
	}

	cherry, _, _ := ebitenutil.NewImageFromFile("assets/cherry.png", ebiten.FilterDefault)

	c.cherry = *cherry

	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	c.xPos = float64(random.Intn(c.xLimit) * 20)
	c.yPos = float64(random.Intn(c.yLimit) * 20)
	fmt.Print(c.xPos)
	fmt.Print(" ")
	fmt.Println(c.yPos)

	return &c
}

// Update : Logical update of the snake
func (c *Cherry) Update(dotTime int) error {
	if c.eaten == false {
		return nil // Return aviso de que ya se la comieron
	}
	return nil
}

// Draw the cherry
func (c *Cherry) Draw(screen *ebiten.Image, dotTime int) error {
	characterDO = &ebiten.DrawImageOptions{}
	characterDO.GeoM.Translate(c.xPos, c.yPos)
	screen.DrawImage(&c.cherry, characterDO)

	return nil
}
