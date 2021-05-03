package 2DSP

func MaxFloat64(nums ...float64) float64 {
	m := nums[0]
	for _, num := range nums {
		if num > m {
			m = num
		}
	}
	return m
}

func MinFloat64(nums ...float64) float64 {
	m := nums[0]
	for _, num := range nums {
		if num < m {
			m = num
		}
	}
	return m
}

func IsEqual(a float32, b float32) bool {
	if ((a - b) < 0.01) && ((a - b) > -0.01) {
		return true
	}
	return false
}

// func Insert(h, d *Node) bool {
// 	if h.Next == nil {
// 		h.Next = d
// 		return true
// 	}
// 	n := h
// 	for n.Next != nil {
// 		n = n.Next
// 	}
// 	d.Next = n.Next
// 	n.Next = d

// 	return true
// }

// func rectset(value1 float32, value2 float32) Node { //位置摆放
// 	var h Node
// 	h.Next = nil

// 	var d Node
// 	d.Data.X = value1
// 	d.Data.Y = value2
// 	Insert(&h, &d)

// 	return h
// }

// 向量叉积AxB=|A||B|sinα
// func LineDirection(l1, l2, l3, l4 *Line) (sin float64) {
// 	sinA := ((l1.P2.X-l1.P1.X)*(l4.P2.Y-l3.P1.Y) - (l4.P2.X-l4.P1.X)*(l1.P2.Y-l1.P1.Y)) /
// 		math.Sqrt((l1.P2.X-l1.P1.X)*(l1.P2.X-l1.P1.X)+(l1.P2.Y-l1.P1.Y)*(l1.P2.Y-l1.P1.Y)) *
// 		math.Sqrt((l4.P2.Y-l3.P1.Y)*(l4.P2.Y-l3.P1.Y)+(l4.P2.X-l4.P1.X)*(l4.P2.X-l4.P1.X))
// 	sin = sinA
// 	return
// }

// func minCrashDis() {}

// func UptoDownMove(head []Header, notpack int, packed int, height float32) (float32, float32, float32, float32) {

// 	fmt.Println("")
// 	return
// }
