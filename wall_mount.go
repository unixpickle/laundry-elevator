package main

import (
	"github.com/unixpickle/model3d/model3d"
	"github.com/unixpickle/model3d/toolbox3d"
)

const (
	WallMountBackHeight    = 50.0
	WallMountSideHeight    = 20.0
	WallMountEdgeThickness = 5.0
	WallMountLength        = 30.0
	WallMountExtraSpace    = 5.0
	WallMountThickness     = MountBoardThickness
	WallMountScrewLength   = MountBoardThickness + WallMountThickness + MountBoardNutThickness
	WallMountBackCutZ      = WallMountBackHeight / 2
	WallMountBackCutSize   = 7.0
	WallMountBackCutMargin = 20.0
)

func WallMountSolid() model3d.Solid {
	depth := MountBoardDepth + WallMountExtraSpace*2 + WallMountEdgeThickness*2
	board := model3d.NewRect(
		model3d.XYZ(0, 0, -WallMountThickness),
		model3d.XYZ(WallMountLength, depth, 0),
	)
	edges := model3d.JoinedSolid{
		model3d.NewRect(
			model3d.XYZ(0, 0, -WallMountThickness),
			model3d.XYZ(WallMountEdgeThickness, depth, WallMountBackHeight),
		),
		model3d.NewRect(
			model3d.XYZ(0, 0, -WallMountThickness),
			model3d.XYZ(WallMountLength, WallMountEdgeThickness, WallMountSideHeight),
		),
		model3d.NewRect(
			model3d.XYZ(0, depth-WallMountEdgeThickness, -WallMountThickness),
			model3d.XYZ(WallMountLength, depth, WallMountSideHeight),
		),
	}
	center := depth / 2

	y1 := center - MountBoardTrackSpace/2 - MountBoardTrackSize/2
	y2 := center - MountBoardTrackSpace/2 + MountBoardTrackSize/2
	y3 := center + MountBoardTrackSpace/2 - MountBoardTrackSize/2
	y4 := center + MountBoardTrackSpace/2 + MountBoardTrackSize/2
	cuts := model3d.JoinedSolid{
		model3d.NewRect(
			model3d.XYZ(MountBoardTrackOuterMargin, y1, -WallMountThickness),
			model3d.XYZ(WallMountLength, y2, 0),
		),
		model3d.NewRect(
			model3d.XYZ(MountBoardTrackOuterMargin, y3, -WallMountThickness),
			model3d.XYZ(WallMountLength, y4, 0),
		),
		// Horizontal back cut
		model3d.NewRect(
			model3d.XYZ(0, WallMountBackCutMargin, WallMountBackCutZ-WallMountBackCutSize/2),
			model3d.XYZ(WallMountEdgeThickness, depth-WallMountBackCutMargin, WallMountBackCutZ+WallMountBackCutSize/2),
		),
	}

	return model3d.Subtract(model3d.JoinedSolid{board, edges}, cuts)
}

func WallMountScrewSolid() model3d.Solid {
	return model3d.StackSolids(
		&model3d.Cylinder{
			P2:     model3d.Z(MountBoardScrewBaseThickness),
			Radius: MountBoardScrewBaseRadius,
		},
		&toolbox3d.ScrewSolid{
			P2:         model3d.Z(WallMountScrewLength),
			Radius:     MountBoardScrewRadius,
			GrooveSize: MountBoardScrewGrooveSize,
		},
	)
}
