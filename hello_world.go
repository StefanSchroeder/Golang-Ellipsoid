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
    fmt.Printf("----- Testing To\n");
    fmt.Printf("Distance = %v Bearing = %v\n", distance, bearing)

    // Calculate where you are when going from SFO in 
    // direction 45.0 deg. for 20000 meters.
    fmt.Printf("----- Testing At\n");
    lat3, lon3 := geo1.At(lat1, lon1, 20000.0, 45.0)
    fmt.Printf("lat3 = %v lon3 = %v\n", lat3, lon3)

    // Convert Lat-Lon-Alt to ECEF.
    lat4, lon4, alt4 := 39.197807, -77.108574 , 55.0 // Some location near Baltimore
    // that the author of the Perl module geo-ecef used. I reused the coords of the tests.
    fmt.Printf("----- Testing ToECEF\n");
    x, y, z := geo1.ToECEF(lat4, lon4, alt4)
    fmt.Printf("x = %v \ny = %v \nz = %v\n", x, y, z)

    // Convert ECEF to Lat-Lon-Alt.
    x1, y1, z1 := 1.1042590709397183e+06, -4.824765955871677e+06, 4.0093940281868847e+06
    fmt.Printf("----- Testing ToLLA\n");
    lat5, lon5, alt5 := geo1.ToLLA(x1, y1, z1)
    fmt.Printf("lat5 = %v lon5 = %v alt3 = %v\n", lat5, lon5, alt5)

    fmt.Printf("----- Done ---\n");
}



