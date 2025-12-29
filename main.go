package main

import "github.com/unixpickle/model3d/model3d"

const (
	SpoolDelta = 0.5
)

func main() {
	spool := model3d.DualContour(SpoolSolid(), SpoolDelta, true, false)
	spool = spool.EliminateCoplanar(1e-5)
	spool.SaveGroupedSTL("spool.stl")

}
