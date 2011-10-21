// These are the 'Intermediate' tests
package main

import "fmt"
import "./ellipsoid"

var pass_counter, fail int

func main() {

	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.Longitude_not_symmetric, ellipsoid.Bearing_not_symmetric)
	lat1, lon1 := 37.619002, -122.374843 //SFO
	lat2, lon2 := 33.942536, -118.408074 //LAX

	number_of_hops := 5
	// Go from SFO to LAX with 5 hops, while ignoring range and bearing.
	// The array v is used to store the coordinates for all intermediate
	// points (including the start and the end point). 
	_, _, v := geo1.Intermediate(lat1, lon1, lat2, lon2, number_of_hops)
	fmt.Printf("-------------------------\n")
	for i := 0; i < len(v)/2; i++ {
		fmt.Printf("%v %v # %v\n", v[i*2], v[i*2+1], i)
	}
	fmt.Printf("-------------------------\n")

}
