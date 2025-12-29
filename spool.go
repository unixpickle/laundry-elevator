package main

import "github.com/unixpickle/model3d/model3d"

const (
	SpoolWidth         = 30
	SpoolRadiusSmall   = 40
	SpoolRadiusLarge   = 60
	SpoolSideThickness = 5

	MotorDRadius     = 8.0 / 2
	MotorDCutoff     = 7.0 / 2
	MotorCutoutDepth = 20.0

	HolderRadius      = 10
	HolderCutoutDepth = 15

	TieOffWidth = 5
)

func SpoolSolid() model3d.Solid {
	spoolBody := model3d.StackSolids(
		&model3d.Cylinder{P2: model3d.Z(SpoolSideThickness), Radius: SpoolRadiusLarge},
		&model3d.Cylinder{P2: model3d.Z(SpoolWidth), Radius: SpoolRadiusSmall},
		&model3d.Cylinder{P2: model3d.Z(SpoolSideThickness), Radius: SpoolRadiusLarge},
	)
	tieSection := &model3d.Cylinder{
		P1:     model3d.XYZ(0, SpoolRadiusSmall+TieOffWidth*2, SpoolSideThickness/2),
		P2:     model3d.XYZ(0, SpoolRadiusSmall-TieOffWidth/2, SpoolWidth/2+SpoolSideThickness),
		Radius: TieOffWidth / 2,
	}
	return model3d.Subtract(
		model3d.JoinedSolid{spoolBody, tieSection},
		model3d.JoinedSolid{MotorCutout(), HolderCutout()},
	)
}

func MotorCutout() model3d.Solid {
	cylinder := &model3d.Cylinder{P2: model3d.Z(MotorCutoutDepth), Radius: MotorDRadius}
	box := model3d.NewRect(
		model3d.XYZ(-MotorDRadius, -MotorDRadius, 0),
		model3d.XYZ(MotorDRadius, MotorDCutoff, MotorCutoutDepth),
	)
	return model3d.IntersectedSolid{cylinder, box}
}

func HolderCutout() model3d.Solid {
	height := float64(SpoolSideThickness*2 + SpoolWidth)
	return &model3d.Cylinder{
		P1:     model3d.Z(height - HolderCutoutDepth),
		P2:     model3d.Z(height + 1e-5),
		Radius: HolderRadius,
	}
}
