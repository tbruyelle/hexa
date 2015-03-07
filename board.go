package main

import (
	"github.com/bmatsuo/hexgrid"
	"github.com/bmatsuo/hexgrid/hexcoords"
	"github.com/tbruyelle/fsm"
	"log"
)

const (
	Sqrt3     = 1.73205080757
	hexRadius = 25
	hexHeight = hexRadius * 2
	hexWidth  = Sqrt3 / 2 * hexHeight
)

func NewBoard() {
	grid := hexgrid.NewGrid(5, 5, hexRadius, nil, nil, nil)
	xoffset, yoffset := float32(hexWidth*float32(grid.NumCols()-1)/2), float32(hexHeight*float32(grid.NumRows()-1)/2)
	for u := grid.ColMin(); u <= grid.ColMax(); u++ {
		for v := grid.RowMin(); v <= grid.RowMax(); v++ {
			var (
				c    = hexcoords.Hex{u, v}
				tile = grid.GetTile(c)
				n    = &fsm.Object{
					X:      float32(tile.Pos.Y) + yoffset,
					Y:      float32(tile.Pos.X) + xoffset,
					Width:  hexWidth,
					Height: hexHeight,
					Sprite: hexSprite,
				}
			)
			n.Register(scene, eng)
			tile.Value = n
			log.Printf("tile %d,%d coord %f,%f offset %f,%f", u, v, tile.Pos.X, tile.Pos.Y, n.X, n.Y)
		}
	}
}
