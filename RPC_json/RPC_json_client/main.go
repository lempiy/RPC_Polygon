package main

import (
	"fmt"
	"net/rpc/jsonrpc"
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

func main() {
	address := ":1200"
	client, err := jsonrpc.Dial("tcp", address)
	if err != nil {
		fmt.Println(err)
	}
	poly := []Point{
		Point{X: 100, Y: 100},
		Point{X: 150, Y: 50},
		Point{X: 200, Y: 100},
		Point{X: 150, Y: 150}}
	target := Point{X: 125, Y: 125}
	args := Args{
		Target:  target,
		Polygon: poly}

	result := new(Result)
	err = client.Call("PolygonMath.IsPointInsidePolygon", args, result)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Is point inside poly? - %v\n", result.IsPiP)

	result2 := new(Result)
	err = client.Call("PolygonMath.GetPolygonArea", args, result2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The total area of polygon is - %v\n", result2.TotalArea)
}
