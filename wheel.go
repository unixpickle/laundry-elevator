package main

import (
	"github.com/unixpickle/model3d/model3d"
	"github.com/unixpickle/model3d/toolbox3d"
)

const (
	WheelRadius                  = 15.0
	WheelThickness               = 10.0
	WheelRodRadius               = 3.0
	WheelRodScrewLength          = 4.0
	WheelRodScrewRadius          = 3.0
	WheelRodScrewGrooveSize      = 1.0
	WheelRodCutoutSlack          = 0.4
	WheelRodNutRadius            = 4.0
	WheelRodNutSlack             = 0.2
	WheelCutoutSlack             = 0.4
	WheelSideSlack               = 1
	WheelSideThickness           = 3.0
	WheelMountDistance           = 15.0
	WheelMountCutoutSize         = 7.0
	WheelMountCutoutSide         = 2.0
	WheelMountScrewBaseSize      = BasketHoleSize + 4.0
	WheelMountScrewBaseThickness = BasketThickness
	WheelMountScrewShaftSize     = BasketHoleSize - 0.2
	WheelMountScrewShaftLength   = BasketThickness
	WheelMountScrewRadius        = 3.0
	WheelMountScrewLength        = WheelRodScrewLength
)

func WheelBracketSolid() model3d.Solid {
	base := model3d.Subtract(
		model3d.NewRect(
			model3d.XYZ(
				-(WheelMountDistance+WheelMountCutoutSize/2+WheelMountCutoutSide),
				-(WheelMountCutoutSize/2+WheelMountCutoutSide),
				-WheelSideThickness,
			),
			model3d.XYZ(
				WheelSideSlack*2+WheelSideThickness*2+WheelThickness,
				WheelMountCutoutSize/2+WheelMountCutoutSide,
				0,
			),
		),
		model3d.NewRect(
			model3d.XYZ(
				-(WheelMountDistance+WheelMountCutoutSize/2),
				-WheelMountCutoutSize/2,
				-WheelSideThickness,
			),
			model3d.XYZ(
				-WheelMountDistance+WheelMountCutoutSize/2,
				WheelMountCutoutSize/2,
				0,
			),
		),
	)

	leftSide := model3d.NewRect(
		model3d.XYZ(
			0,
			-(WheelMountCutoutSize/2+WheelMountCutoutSide),
			0,
		),
		model3d.XYZ(
			WheelSideThickness,
			(WheelMountCutoutSize/2+WheelMountCutoutSide),
			WheelSideSlack+WheelRadius+WheelCutoutSlack*2+WheelRodRadius+WheelMountCutoutSide,
		),
	)
	rightSide := model3d.NewRect(
		model3d.XYZ(
			WheelSideThickness+WheelSideSlack*2+WheelThickness,
			-(WheelMountCutoutSize/2+WheelMountCutoutSide),
			0,
		),
		model3d.XYZ(
			WheelSideSlack*2+WheelThickness+WheelSideThickness*2,
			(WheelMountCutoutSize/2+WheelMountCutoutSide),
			WheelSideSlack+WheelRadius+WheelCutoutSlack*2+WheelRodRadius+WheelMountCutoutSide,
		),
	)

	rod := WheelRodSolid(0)
	sides := model3d.Subtract(model3d.JoinedSolid{leftSide, rightSide}, rod)

	return model3d.JoinedSolid{base, sides}
}

func WheelRodSolid(outset float64) model3d.Solid {
	rodZ := WheelSideSlack + WheelRadius
	rod := &model3d.Cylinder{
		P1:     model3d.XYZ(WheelSideThickness/2, 0, rodZ),
		P2:     model3d.XYZ(WheelSideThickness+WheelSideSlack*2+WheelThickness, 0, rodZ),
		Radius: WheelRodRadius + outset,
	}
	squareCutout := model3d.NewRect(
		model3d.XYZ(
			WheelSideThickness+WheelSideSlack*2+WheelThickness,
			-WheelRodRadius,
			rodZ-WheelRodRadius,
		),
		model3d.XYZ(
			WheelSideThickness+WheelSideSlack*2+WheelThickness+WheelSideThickness,
			WheelRodRadius,
			rodZ+WheelRodRadius,
		),
	)
	screw := WheelRodScrew()
	return model3d.JoinedSolid{rod, squareCutout, screw}
}

func WheelSolid() model3d.Solid {
	midZ := WheelSideSlack + WheelRadius
	wheel := &model3d.Cylinder{
		P1:     model3d.XYZ(WheelSideThickness+WheelSideSlack, 0, midZ),
		P2:     model3d.XYZ(WheelSideThickness+WheelSideSlack+WheelThickness, 0, midZ),
		Radius: WheelRadius,
	}
	rod := WheelRodSolid(WheelRodCutoutSlack)
	return model3d.Subtract(
		wheel,
		rod,
	)
}

func WheelRodScrewNutSolid() model3d.Solid {
	screw := WheelRodScrew()
	screw.Radius += WheelRodNutSlack
	surround := &model3d.Cylinder{
		P1:     screw.P1,
		P2:     screw.P2,
		Radius: WheelRodNutRadius,
	}
	return model3d.Subtract(surround, screw)
}

func WheelRodScrew() *toolbox3d.ScrewSolid {
	rodZ := WheelSideSlack + WheelRadius
	return &toolbox3d.ScrewSolid{
		P1: model3d.XYZ(
			WheelSideThickness+WheelSideSlack*2+WheelThickness+WheelSideThickness,
			0,
			rodZ,
		),
		P2: model3d.XYZ(
			WheelSideThickness+WheelSideSlack*2+WheelThickness+WheelSideThickness+WheelRodScrewLength,
			0,
			rodZ,
		),
		Radius:     WheelRodScrewRadius,
		GrooveSize: WheelRodScrewGrooveSize,
	}
}

func WheelBracketScrewSolid() model3d.Solid {
	return model3d.StackSolids(
		model3d.NewRect(
			model3d.XYZ(-WheelMountScrewBaseSize/2, -WheelMountScrewBaseSize/2, 0),
			model3d.XYZ(WheelMountScrewBaseSize/2, WheelMountScrewBaseSize/2, WheelMountScrewBaseThickness),
		),
		model3d.NewRect(
			model3d.XYZ(-WheelMountScrewShaftSize/2, -WheelMountScrewShaftSize/2, 0),
			model3d.XYZ(WheelMountScrewShaftSize/2, WheelMountScrewShaftSize/2, WheelMountScrewShaftLength),
		),
		model3d.NewRect(
			model3d.XYZ(-WheelMountCutoutSize/2, -WheelMountCutoutSize/2, 0),
			model3d.XYZ(WheelMountCutoutSize/2, WheelMountCutoutSize/2, WheelSideThickness),
		),
		&toolbox3d.ScrewSolid{
			P1:         model3d.Z(0),
			P2:         model3d.Z(WheelMountScrewLength),
			Radius:     WheelMountScrewRadius,
			GrooveSize: WheelRodScrewGrooveSize,
		},
	)
}
