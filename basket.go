package main

import (
	"math"

	"github.com/unixpickle/model3d/model3d"
)

const (
	BasketWidth           = 200
	BasketDepth           = 200
	BasketHeight          = 200
	BasketHoleSize        = 20.0
	BasketHoleSpace       = 8.0
	BasketThickness       = 5
	BasketBottomHoleWidth = 5
)

func BasketSolid() model3d.Solid {
	rect := model3d.Subtract(
		model3d.NewRect(
			model3d.XYZ(-BasketWidth/2, -BasketDepth/2, 0),
			model3d.XYZ(BasketWidth/2, BasketDepth/2, BasketHeight),
		),
		model3d.NewRect(
			model3d.XYZ(-BasketWidth/2+BasketThickness, -BasketDepth/2+BasketThickness, BasketThickness),
			model3d.XYZ(BasketWidth/2-BasketThickness, BasketDepth/2-BasketThickness, BasketHeight),
		),
	)
	sideHoles := model3d.JoinedSolid{}

	numHoles := math.Floor((BasketWidth - BasketHoleSpace) / (BasketHoleSize + BasketHoleSpace))
	holeWidth := float64(numHoles)*(BasketHoleSpace+BasketHoleSize) + BasketHoleSpace

	for x := -holeWidth/2 + BasketHoleSpace; x+BasketHoleSize < BasketWidth/2; x += BasketHoleSize + BasketHoleSpace {
		for z := BasketHoleSpace + BasketThickness; z+BasketHoleSpace+BasketHoleSize < BasketHeight; z += BasketHoleSize + BasketHoleSpace {
			sideHoles = append(sideHoles, model3d.NewRect(
				model3d.XYZ(x, -1000, z),
				model3d.XYZ(x+BasketHoleSize, 1000, z+BasketHoleSize),
			))
		}
	}

	numHoles = math.Floor((BasketDepth - BasketHoleSpace) / (BasketHoleSize + BasketHoleSpace))
	holeWidth = float64(numHoles)*(BasketHoleSpace+BasketHoleSize) + BasketHoleSpace
	for y := -holeWidth/2 + BasketHoleSpace; y+BasketHoleSpace+BasketHoleSize < BasketDepth/2; y += BasketHoleSize + BasketHoleSpace {
		for z := BasketHoleSpace + BasketThickness; z+BasketHoleSpace+BasketHoleSize < BasketHeight; z += BasketHoleSize + BasketHoleSpace {
			sideHoles = append(sideHoles, model3d.NewRect(
				model3d.XYZ(-1000, y, z),
				model3d.XYZ(1000, y+BasketHoleSize, z+BasketHoleSize),
			))
		}
	}
	return model3d.Subtract(rect, sideHoles.Optimize())
}
