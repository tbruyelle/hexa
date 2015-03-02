package main

import (
	"github.com/bmatsuo/hexgrid"
	"github.com/bmatsuo/hexgrid/hexcoords"
	"github.com/tbruyelle/fsm"
)

const (
	Sqrt3     = 1.73205080757
	hexRadius = 20
	hexWidth  = hexRadius * 2
	hexHeight = Sqrt3 / 2 * hexWidth
)

func NewBoard() {
	grid := hexgrid.NewGrid(5, 5, hexRadius, nil, nil, nil)
	for u := grid.ColMin(); u <= grid.ColMax(); u++ {
		for v := grid.RowMin(); v <= grid.RowMax(); v++ {
			var (
				c    = hexcoords.Hex{u, v}
				tile = grid.GetTile(c)
				n    = &fsm.Object{
					X:      hexWidth * float32(u),
					Y:      hexHeight * float32(v),
					Width:  hexWidth,
					Height: hexHeight,
					Sprite: hexSprite,
				}
			)
			n.Register(scene, eng)
			tile.Value = n
		}
	}

}
