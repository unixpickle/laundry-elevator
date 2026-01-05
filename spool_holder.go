package main

import "github.com/unixpickle/model3d/model3d"

const (
	SpoolHolderRunway       = 50.0
	SpoolHolderThickness    = 2.0
	SpoolHolderCenterHeight = 29.0
	SpoolHolderSlack        = 2.0
	SpoolHolderRailSize     = 6.0
	SpoolHolderRailMargin   = 5.0
	SpoolHolderRodJutLength = 1.0
	SpoolHolderRodLength    = 10.0
)

func SpoolHolderSolid() model3d.Solid {
	end := MountBoardTrackSpace/2 + SpoolHolderRailSize/2 + SpoolHolderRailMargin
	positive := model3d.JoinedSolid{
		model3d.NewRect(
			model3d.XYZ(-SpoolHolderRunway, -end, 0),
			model3d.XYZ(0, end, SpoolHolderThickness),
		),
		model3d.NewRect(
			model3d.XYZ(0, -end, 0),
			model3d.XYZ(SpoolHolderThickness, end, SpoolHolderCenterHeight+HolderRadius+SpoolHolderSlack*2),
		),
		&model3d.Cylinder{
			P1:     model3d.XYZ(SpoolHolderThickness, 0, SpoolHolderCenterHeight),
			P2:     model3d.XYZ(SpoolHolderThickness+SpoolHolderRodJutLength, 0, SpoolHolderCenterHeight),
			Radius: HolderRadius + SpoolHolderSlack,
		},
		&model3d.Cylinder{
			P1:     model3d.XYZ(SpoolHolderThickness, 0, SpoolHolderCenterHeight),
			P2:     model3d.XYZ(SpoolHolderThickness+SpoolHolderRodLength, 0, SpoolHolderCenterHeight),
			Radius: HolderRadius - SpoolHolderSlack,
		},
	}

	spoolY := MountBoardTrackSpace/2 - SpoolHolderRailSize/2
	rails := model3d.JoinedSolid{
		model3d.NewRect(
			model3d.XYZ(-SpoolHolderRunway+SpoolHolderRailMargin, spoolY, 0),
			model3d.XYZ(-SpoolHolderRailMargin, spoolY+SpoolHolderRailSize, SpoolHolderThickness),
		),
		model3d.NewRect(
			model3d.XYZ(-SpoolHolderRunway+SpoolHolderRailMargin, -(spoolY+SpoolHolderRailSize), 0),
			model3d.XYZ(-SpoolHolderRailMargin, -spoolY, SpoolHolderThickness),
		),
	}
	return model3d.Subtract(positive, rails)
}
