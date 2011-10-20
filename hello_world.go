package main

import "fmt"
import "./ellipsoid"

func main() {
    lat1, lon1 := 37.619002, -122.374843 //SFO
    lat2, lon2 := 33.942536, -118.408074 //LAX

    // Create Ellipsoid object with WGS84-ellipsoid, 
    // angle units are degrees, distance units are meter.
    geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)

    // Calculate the distance and bearing from SFO to LAX.
    distance, bearing := geo1.To(lat1, lon1, lat2, lon2)
    fmt.Printf("Distance = %v Bearing = %v\n", distance, bearing)

    // Calculate where you are when going from SFO in 
    // direction 45.0 deg. for 20000 meters.
    lon3, lat3 := geo1.At(lat1, lon1, 20000.0, 45.0)
    fmt.Printf("lat3 = %v lon3 = %v\n", lat3, lon3)

    _, _, v := geo1.Intermediate(lat1, lon1, lat2, lon2, 5)
    for i := 0; i < 6 ; i++ {
        fmt.Printf("%v %v # %v\n",v[i*2], v[i*2+1], i)
    }
}
