package main

import (
	"github.com/unixpickle/model3d/model3d"
	"github.com/unixpickle/model3d/toolbox3d"
)

const (
	MountBoardThickness          = 10.0
	MountBoardWidth              = 220.0
	MountBoardDepth              = SpoolRadiusLarge*2 + 40.0
	MountBoardSpoolSlack         = 10.0
	MountBoardTrackSpace         = 45.0
	MountBoardTrackSize          = 8.0
	MountBoardTrackOuterMargin   = 10.0
	MountBoardTrackInnerMargin   = 4.0
	MountBoardScrewBaseRadius    = 5.0
	MountBoardScrewBaseThickness = 4.0
	MountBoardScrewRadius        = 5.0 / 2
	MountBoardScrewGrooveSize    = 0.7
	MountBoardScrewLength        = MountBoardThickness + 2.0 + MountBoardNutThickness
	MountBoardScrewSlack         = 0.2
	MountBoardNutThickness       = 4.0
	MountBoardNutRadius          = 5.0
)

func MountBoardSolid() model3d.Solid {
	base := model3d.NewRect(
		model3d.XYZ(0, 0, 0),
		model3d.XYZ(MountBoardWidth, MountBoardDepth, MountBoardThickness),
	)

	spoolHalf := (SpoolWidth+SpoolSideThickness*2)/2 + MountBoardSpoolSlack
	spoolCutoutSpace := (MountBoardDepth - SpoolRadiusLarge*2) / 2
	spoolCutout := model3d.NewRect(
		model3d.XYZ(MountBoardWidth/2-spoolHalf, spoolCutoutSpace, 0),
		model3d.XYZ(MountBoardWidth/2+spoolHalf, MountBoardDepth-spoolCutoutSpace, MountBoardThickness),
	)

	midY := MountBoardDepth / 2

	x1 := MountBoardTrackOuterMargin
	x2 := spoolCutout.Min().X - MountBoardTrackInnerMargin
	x3 := spoolCutout.Max().X + MountBoardTrackInnerMargin
	x4 := MountBoardWidth - MountBoardTrackOuterMargin
	y1 := midY - MountBoardTrackSpace/2 - MountBoardTrackSize/2
	y2 := midY - MountBoardTrackSpace/2 + MountBoardTrackSize/2
	y3 := midY + MountBoardTrackSpace/2 - MountBoardTrackSize/2
	y4 := midY + MountBoardTrackSpace/2 + MountBoardTrackSize/2
	railCutouts := model3d.JoinedSolid{
		model3d.NewRect(
			model3d.XYZ(x1, y1, 0),
			model3d.XYZ(x2, y2, MountBoardThickness),
		),
		model3d.NewRect(
			model3d.XYZ(x3, y1, 0),
			model3d.XYZ(x4, y2, MountBoardThickness),
		),
		model3d.NewRect(
			model3d.XYZ(x1, y3, 0),
			model3d.XYZ(x2, y4, MountBoardThickness),
		),
		model3d.NewRect(
			model3d.XYZ(x3, y3, 0),
			model3d.XYZ(x4, y4, MountBoardThickness),
		),
	}

	cutouts := model3d.JoinedSolid{spoolCutout, railCutouts}

	return model3d.Subtract(base, cutouts)
}

func MountBoardNutSolid() model3d.Solid {
	return model3d.Subtract(
		&model3d.Cylinder{
			P2:     model3d.Z(MountBoardNutThickness),
			Radius: MountBoardNutRadius,
		},
		&toolbox3d.ScrewSolid{
			P2:         model3d.Z(1000),
			Radius:     MountBoardScrewRadius + MountBoardScrewSlack,
			GrooveSize: MountBoardScrewGrooveSize,
		},
	)
}

func MountBoardScrewSolid() model3d.Solid {
	return model3d.StackSolids(
		&model3d.Cylinder{
			P2:     model3d.Z(MountBoardScrewBaseThickness),
			Radius: MountBoardScrewBaseRadius,
		},
		&toolbox3d.ScrewSolid{
			P2:         model3d.Z(MountBoardScrewLength),
			Radius:     MountBoardScrewRadius,
			GrooveSize: MountBoardScrewGrooveSize,
		},
	)
}
