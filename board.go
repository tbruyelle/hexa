package main

import (
	"github.com/bmatsuo/hexgrid"
	"github.com/bmatsuo/hexgrid/hexcoords"
	"github.com/tbruyelle/fsm"
	"log"
)

const (
	Sqrt3             = 1.73205080757
	hexRadius         = 24
	hexHeight         = hexRadius * 2
	hexVerticalOffset = hexHeight * .75
	hexWidth          = Sqrt3 / 2 * hexHeight
)

func NewBoard() {
	log.Printf("hex width %v, height %v", hexWidth, hexHeight)
	grid := hexgrid.NewGrid(9, 9, hexRadius, nil, nil, nil)
	xoffset, yoffset := float32(hexWidth*float32(grid.NumCols()-1)/2), float32(hexVerticalOffset*float32(grid.NumRows()-1)/2)
	log.Printf("offsets %f,%f", xoffset, yoffset)
	for u := grid.ColMin(); u <= grid.ColMax(); u++ {
		for v := grid.RowMin(); v <= grid.RowMax(); v++ {
			var (
				c    = hexcoords.Hex{u, v}
				tile = grid.GetTile(c)
				n    = &fsm.Object{
					X:      float32(u)*hexWidth + xoffset,
					Y:      float32(v)*hexVerticalOffset + yoffset,
					Width:  hexWidth,
					Height: hexHeight,
					Sprite: hexSprite,
				}
			)
			if hexcoords.ColumnIsHigh(v) {
				n.X += hexWidth / 2
			}
			n.Register(scene, eng)
			tile.Value = n
			log.Printf("tile %d,%d coord %f,%f offset %f,%f", u, v, tile.Pos.X, tile.Pos.Y, n.X, n.Y)
		}
	}
}
