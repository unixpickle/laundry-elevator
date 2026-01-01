package main

import "github.com/unixpickle/model3d/model3d"

const (
	SpoolDelta  = 0.5
	BasketDelta = 2.0
)

func main() {
	basket := model3d.DualContour(BasketSolid(), BasketDelta, true, false)
	basket = basket.EliminateCoplanar(1e-5)
	basket.SaveGroupedSTL("basket.stl")

	spool := model3d.DualContour(SpoolSolid(), SpoolDelta, true, false)
	spool = spool.EliminateCoplanar(1e-5)
	spool.SaveGroupedSTL("spool.stl")
}
