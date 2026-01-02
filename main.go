package main

import "github.com/unixpickle/model3d/model3d"

const (
	SpoolDelta  = 0.5
	BasketDelta = 2.0
	WheelDelta  = 0.25
)

func main() {
	wheelBracket := model3d.DualContour(WheelBracketSolid(), WheelDelta, true, false)
	wheelBracket = wheelBracket.EliminateCoplanar(1e-5)
	wheelBracket.SaveGroupedSTL("wheel_bracket.stl")

	wheelBracketScrew := model3d.DualContour(WheelBracketScrewSolid(), WheelDelta, true, false)
	wheelBracketScrew = wheelBracketScrew.EliminateCoplanar(1e-5)
	wheelBracketScrew.SaveGroupedSTL("wheel_bracket_screw.stl")

	wheelRod := model3d.DualContour(WheelRodSolid(0), WheelDelta, true, false)
	wheelRod = wheelRod.EliminateCoplanar(1e-5)
	wheelRod.SaveGroupedSTL("wheel_rod.stl")

	wheelRodNut := model3d.DualContour(WheelRodScrewNutSolid(), WheelDelta, true, false)
	wheelRodNut = wheelRodNut.EliminateCoplanar(1e-5)
	wheelRodNut.SaveGroupedSTL("wheel_rod_nut.stl")

	wheel := model3d.DualContour(WheelSolid(), WheelDelta, true, false)
	wheel = wheel.EliminateCoplanar(1e-5)
	wheel.SaveGroupedSTL("wheel.stl")

	basket := model3d.DualContour(BasketSolid(), BasketDelta, true, false)
	basket = basket.EliminateCoplanar(1e-5)
	basket.SaveGroupedSTL("basket.stl")

	spool := model3d.DualContour(SpoolSolid(), SpoolDelta, true, false)
	spool = spool.EliminateCoplanar(1e-5)
	spool.SaveGroupedSTL("spool.stl")
}
