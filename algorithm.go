package main

import "math"

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
func moveToInitialPosition(srtip *Strip, pol *Polygon) float64 {
	// 首先判断条带内是否摆放待排件
	polTempX := []float64{}
	polTempY := []float64{}
	if len(strip.Obejcts) == 0 {
		// strip中未排放多边形
		var (
			disX float64
			disY float64
		)
		disX = math.Abs(MinX(pol.Vertices[0].X) - 0)
		disY = math.Abs(MaxY(pol.Vertices[0].Y) - 0)
		for i := 0; i < len(pol.Vertices); i++ {

			polTempX = append(polTempX, pol.Vertices[i].X-disX)
			polTempY = append(polTempY, pol.Vertices[i].Y-disY)
		}

		polNewX := []float64{}
		polNewY := []float64{}
		for i := 0; i < len(pol.Vertices); i++ {
			polNewX = append(polNewX, polTempX[i])
			polNewY = append(polNewY, polTempY[i]+strip.Width)
		}

	}
	// 在当前列中，进行待排件的摆放
	if len(strip.Obejcts) == 1 {
		// 判断待排件移动时 与正反边的接触点的距离矢量差能否放置待排件

	}
	// 当前列排放满，无法再进行后续的待排件的插入

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
