package main

import (
	"container/list"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type Object struct {
	Hight float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

type Line struct {
	P1, P2 *Point
}

func NewLine(p1, p2 *Point) *Line {
	return &Line{p1, p2}
}

// GetSin 获取l与l1的sin值
func (l *Line) GetSin(l1 *Line) float64 {
	vecl := NewPoint(l.P2.X-l.P1.X, l.P2.Y-l.P1.Y)
	vecl1 := NewPoint(l1.P2.X-l1.P1.X, l1.P2.Y-l1.P1.Y)
	product := vecl.X*vecl.Y - vecl1.X*vecl.Y
	mold := math.Sqrt(vecl.X*vecl.X+vecl.Y*vecl.Y) + math.Sqrt(vecl1.X*vecl1.X+vecl1.Y*vecl1.Y)
	return product / mold
}

// Direction 返回l与l1的方向。
//  true正方向，false 负方向
func (l *Line) Direction(l1 *Line) bool {
	if l.GetSin(l1) > 0 {
		return true
	}
	return false
}

type Polygon struct {
	Vertices []*Point
	Holes    *list.List
}

//Area 计算多边形面积
func (p *Polygon) Area() float64 {
	var area float64
	num := len(p.Vertices)
	for i := 0; i < num-1; i++ {
		area += p.Vertices[i].X*p.Vertices[i+1].Y - p.Vertices[i+1].X*p.Vertices[i].Y
	}
	return area / 2
}

func (p *Polygon) MinX() float64 {
	min := p.Vertices[0].X
	for _, v := range p.Vertices {
		if min > v.X {
			min = v.X
		}
	}
	return min
}

func (p *Polygon) MinY() float64 {
	min := p.Vertices[0].Y
	for _, v := range p.Vertices {
		if min > v.Y {
			min = v.Y
		}
	}
	return min
}

func (p *Polygon) MaxX() float64 {
	max := p.Vertices[0].X
	for _, v := range p.Vertices {
		if max < v.X {
			max = v.X
		}
	}
	return max
}

func (p *Polygon) MaxY() float64 {
	max := p.Vertices[1].Y
	for _, v := range p.Vertices {
		if max < v.Y {
			max = v.Y
		}
	}
	return max
}

type Strip struct {
	Height  float64 //长度
	Width   float64
	Sp      *Point
	Obejcts []*Polygon
}

// MaxX 获取多边形最右值点的x值
func (s *Strip) MaxX() float64 {

	max := s.Obejcts[0].Vertices[0].X
	for _, v := range s.Obejcts[0].Vertices {
		if max < v.X {
			max = v.X
		}
	}
	return max
}
