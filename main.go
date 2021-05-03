package main

func main() {

	srtip := &Strip{}
	polygons := []*Polygon{}
	for _, pol := range polygons {
		dist := getMovedCordinate(srtip, pol)

		insertPolygon(srtip, pol, dist)
	}
}
