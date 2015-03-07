package main

import (
	"github.com/bmatsuo/hexgrid"
	"github.com/bmatsuo/hexgrid/hexcoords"
	"github.com/tbruyelle/fsm"
	"log"
)

const (
	Sqrt3     = 1.73205080757
	hexRadius = 20
	hexWidth  = hexRadius * 2
	hexHeight = Sqrt3 / 2 * hexWidth
)

func NewBoard() {
	grid := hexgrid.NewGrid(5, 5, hexRadius, nil, nil, nil)
	xoffset, yoffset := float32(hexWidth*5/2), float32(hexHeight*5/2)
	for u := grid.ColMin(); u <= grid.ColMax(); u++ {
		for v := grid.RowMin(); v <= grid.RowMax(); v++ {
			var (
				c    = hexcoords.Hex{u, v}
				tile = grid.GetTile(c)
				n    = &fsm.Object{
					X:      float32(tile.Pos.X) + xoffset,
					Y:      float32(tile.Pos.Y) + yoffset,
					Width:  hexWidth,
					Height: hexHeight,
					Sprite: hexSprite,
				}
			)
			n.Register(scene, eng)
			tile.Value = n
			log.Printf("tile %d,%d coord %f,%f", u, v, n.X, n.Y)
		}
	}

}
