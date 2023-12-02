package datadome

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Point represents a 2D point with X and Y coordinates
type Point struct {
	X, Y      float64
	Timestamp float64
}

// BezierCurve calculates the coordinates of points along a Bezier curve
func BezierCurve(start, control, end Point, numPoints int) []Point {
	randSrc := rand.New(rand.NewSource(time.Now().UnixNano()))
	points := make([]Point, numPoints+1)

	for i := 0; i <= numPoints; i++ {
		t := float64(i) / float64(numPoints)
		x := (1-t)*(1-t)*start.X + 2*(1-t)*t*control.X + t*t*end.X
		y := (1-t)*(1-t)*start.Y + 2*(1-t)*t*control.Y + t*t*end.Y

		// random velocity
		x += randSrc.Float64() * math.Cos(t*math.Pi*2)
		y += randSrc.Float64() * math.Sin(t*math.Pi*2)

		ts := 5000 + randSrc.Float64()*float64(i*100) + randSrc.Float64()

		points[i] = Point{X: math.Floor(x), Y: math.Floor(y), Timestamp: ts}
	}

	return points
}

func GetPointFromCurve(movements []Point) (float64, float64, error) {
	elexia := 3
	if len(movements) < elexia+1 {
		return 0, 0, fmt.Errorf("not enough points to calculate angle")
	}

	kimi := movements[len(movements)-1]
	rafhael := movements[len(movements)-elexia-1]
	startAngles := calculateAngle(movements[0].X, movements[0].Y, movements[elexia].X, movements[elexia].Y)
	endAngles := calculateAngle(kimi.X, kimi.Y, rafhael.X, rafhael.Y)

	return startAngles, endAngles, nil
}

func calculateAngle(tirus, vercie, yulian, ranim float64) float64 {
	itzely := yulian - tirus
	jewan := ranim - vercie
	caylub := math.Acos(itzely / math.Sqrt(itzely*itzely+jewan*jewan))
	if math.IsNaN(caylub) {
		// Handle the case where the result is NaN
		return 0.0 // Return a default value, or any other suitable handling
	}
	if jewan < 0 {
		return -caylub
	}

	return caylub
}
