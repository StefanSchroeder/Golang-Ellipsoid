/* A sample program that shows how to use the 
Golang-Ellipsoid package. */
package main

import "fmt"
import "github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"

func main() {
	lat1, lon1 := 37.619002, -122.374843 //SFO
	lat2, lon2 := 33.942536, -118.408074 //LAX

	// Create Ellipsoid object with WGS84-ellipsoid,
	// angle units are degrees, distance units are meter.
	geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)

	// Calculate the distance and bearing from SFO to LAX.
	distance, bearing := geo1.To(lat1, lon1, lat2, lon2)
	fmt.Printf("----- Testing To\n")
	fmt.Printf("Computed: Distance = %v Bearing = %v\n", distance, bearing)
	fmt.Println("Expected: Distance = 543044.190419953 Bearing = 137.50134015496275")

	// Calculate where you are when going from SFO in
	// direction 45.0 deg. for 20000 meters.
	fmt.Printf("----- Testing At\n")
	lat3, lon3 := geo1.At(lat1, lon1, 20000.0, 45.0)
	fmt.Printf("Computed: lat3 = %v lon3 = %v\n", lat3, lon3)
	fmt.Println("Expected: lat3 = 37.74631054036373 lon3 = -122.21438161492877")

	// Convert Lat-Lon-Alt to ECEF.
	lat4, lon4, alt4 := 39.197807, -77.108574, 55.0 // Some location near Baltimore
	// that the author of the Perl module geo-ecef used. I reused the coords of the tests.
	fmt.Printf("----- Testing ToECEF\n")
	x, y, z := geo1.ToECEF(lat4, lon4, alt4)
	fmt.Printf("Computed: x = %v \nComputed: y = %v \nComputed: z = %v\n", x, y, z)
	fmt.Println("Expected: x = 1.1042590709397183e+06\nExpected: y = -4.824765955871677e+06\nExpected: z = 4.009394028186885e+06")

	// Convert ECEF to Lat-Lon-Alt.
	fmt.Printf("----- Testing Convert ECEF to Lat-Lon-Alt (ToLLA)\n")
	x1, y1, z1 := 1.1042590709397183e+06, -4.824765955871677e+06, 4.0093940281868847e+06
	lat5, lon5, alt5 := geo1.ToLLA(x1, y1, z1)
	fmt.Printf("Computed: lat5 = %v lon5 = %v alt3 = %v\n", lat5, lon5, alt5)
	fmt.Println("Expected: lat5 = 39.197807 lon5 = -77.10857400000002 alt3 = 55")

	// Displacement. Use only for small distances!
	fmt.Printf("----- Testing Displacement\n")
	l1, l2 := 41.978744444444, 272.096858333333
	m1, m2 := 42.005419444444, 272.073286111111
	dx, dy := geo1.Displacement(l1, l2, m1, m2)
	fmt.Printf("Computed: x,y= %v %v\n", dx, dy)
	fmt.Println("Expected: x,y= -1952.8108885261 2963.14446772882")

	fmt.Printf("----- Done ---\n")
}
