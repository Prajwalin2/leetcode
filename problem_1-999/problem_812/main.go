package main

import "math"

func main() {
}

func largestTriangleArea(points [][]int) float64 {
	var mx float64
	for i := range points {
		for j := i + 1; j < len(points); j++ {
			for k := j + 1; k < len(points); k++ {
				x, y, z := points[i], points[j], points[k]
				double := math.Abs(float64(
					(x[0]*(y[1]-z[1]) +
						(y[0]*(z[1]-x[1]) +
							(z[0] * (x[1] - y[1]))))))
				mx = max(mx, double)
			}
		}
	}
	return mx / 2
}
