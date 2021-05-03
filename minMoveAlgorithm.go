package 2DSP

func BottomLeftFirst() {
	// var (
	// 	top  float32 = -1
	// 	topx float32

	// 	bottom  float32 = 10000
	// 	bottomx float32

	// 	left  float32 = 10000
	// 	lefty float32

	// 	right  float32 = -1
	// 	righty float32
	// )

}
func moveToInitialPosition(srtip *Strip)float64{
	// 首先判断条带内是否摆放待排件
	
}
func getShiftY(strip *Strip, pol *Polygon) float64 {
	// 首先将待排件平移到strip外
	nums := len(pol.Vertices)
	for i := 0; i < int(nums); i++ {
		polXmax := MaxFloat64(pol.Vertices[i].X)
		polXmin := MinFloat64(pol.Vertices[i].X)

		polYmax := MaxFloat64(pol.Vertices[i].Y)
		polYmin := MinFloat64(pol.Vertices[i].Y)

		if polXmax < strip.Height && polXmin >= 0 && polYmax <= strip.Width && polYmin >= 0 {

		}
	}

}

func getShiftX(strip *Strip, pol *Polygon) float64 {

}

func getMovedCordinate(strip *Strip, pol *Polygon) (dist *Point) {

}

func insertPolygon(strip *Strip, pol *Polygon, dist *Point) {

}

func MINMoveDist() {

}
