package main

import (
	"log"

	"github.com/unixpickle/model3d/model3d"
)

const (
	SpoolDelta           = 0.5
	BasketDelta          = 2.0
	WheelDelta           = 0.25
	MountBoardDelta      = 2.0
	MountBoardScrewDelta = 0.1
)

func main() {
	DumpMesh(MountBoardSolid(), MountBoardDelta, "mount_board.stl")
	DumpMesh(MountBoardNutSolid(), MountBoardScrewDelta, "mount_board_nut.stl")
	DumpMesh(MountBoardScrewSolid(), MountBoardScrewDelta, "mount_board_screw.stl")
	DumpMesh(WheelBracketSolid(), WheelDelta, "wheel_bracket.stl")
	DumpMesh(WheelBracketScrewSolid(), WheelDelta, "wheel_bracket_screw.stl")
	DumpMesh(WheelRodSolid(0), WheelDelta, "wheel_rod.stl")
	DumpMesh(WheelRodScrewNutSolid(), WheelDelta, "wheel_rod_nut.stl")
	DumpMesh(WheelSolid(), WheelDelta, "wheel.stl")
	DumpMesh(BasketSolid(), BasketDelta, "basket.stl")

}

func DumpMesh(solid model3d.Solid, delta float64, filename string) {
	log.Printf("Generating %s ...", filename)
	mesh := model3d.DualContour(solid, delta, true, false)
	mesh = mesh.EliminateCoplanar(1e-5)
	mesh.SaveGroupedSTL(filename)
}
