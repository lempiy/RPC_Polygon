package main

import (
	"fmt"
	"math"
	"net/http"
	"net/rpc"
)

//Args is shared arguments for RPC
type Args struct {
	Target  Point
	Polygon []Point
}

//Point is 2d point in catesian space
type Point struct {
	X int
	Y int
}

//Result is shared result for RPC
type Result struct {
	IsPiP     bool
	TotalArea int
}

//PolygonMath is base class for Remote calling
type PolygonMath struct{}

//IsPointInsidePolygon checks weather or not Args.Target is inside Args.Polygon
func (p *PolygonMath) IsPointInsidePolygon(dt *Args, answer *Result) error {
	numbeOfVert := len(dt.Polygon)
	j := 0
	for i := 1; i < numbeOfVert; i++ {
		if ((dt.Polygon[i].Y > dt.Target.Y) != (dt.Polygon[j].Y > dt.Target.Y)) &&
			(dt.Target.X < (dt.Polygon[j].X-dt.Polygon[i].X)*(dt.Target.Y-dt.Polygon[i].Y)/
				(dt.Polygon[j].Y-dt.Polygon[i].Y)+dt.Polygon[i].X) {
			answer.IsPiP = !answer.IsPiP
		}
		j = i
	}
	return nil
}

//GetPolygonArea calculates the total area of Args.Polygon
func (p *PolygonMath) GetPolygonArea(dt *Args, answer *Result) error {
	total := 0
	for i := range dt.Polygon {
		addX, j := dt.Polygon[i].X, 0
		if i != len(dt.Polygon)-1 {
			j = 1
		}
		addY := dt.Polygon[j].Y
		subX := dt.Polygon[j].X
		subY := dt.Polygon[i].Y

		total += (addX * addY / 2)
		total -= (subX * subY / 2)
	}
	answer.TotalArea = int(math.Abs(float64(total)))
	return nil
}

func main() {
	polyMath := new(PolygonMath)
	rpc.Register(polyMath)
	rpc.HandleHTTP()

	err := http.ListenAndServe(":1200", nil)
	if err != nil {
		fmt.Println(err)
	}
}
