# ellipsoid.go

NOTE: The package has not been made go get-friendly yet. Some of my forkers did that, I'll catch up some day.
Also the compilation instructions in this readme have not been updated for Go-1.0.

ellipsoid.go performs latitude and longitude calculations 
on the surface of an ellipsoid. And converts ECEF to LLA and
vice-versa.

This is a Go conversion of an existing Perl conversion 
of existing Fortran code (see ACKNOWLEDGEMENTS) and the 
author of this class makes no claims of originality. Nor 
can he even vouch for the results of the calculations, 
although they do seem to work for him and have been 
tested against other methods.

## Overview

* Calculating distance and bearing when two locations with longitude and latitude are are given.
* Calculate target location when one location with longitude and latitude and distance and bearing are given.
* Supports several ellipsoids (incl. WGS84) out of the box.
* Convert cartesian ECEF-coordinates to longitude, latitude, altitude and vice versa.

## Installation

Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html). 

## Example
    
	package main

	import "fmt"
	import "ellipsoid"

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
		lat3, lon3 := geo1.At(lat1, lon1, 20000.0, 45.0)
		fmt.Printf("lat3 = %v lon3 = %v\n", lat3, lon3)
		
		// Convert Lat-Lon-Alt to ECEF.
		lat4, lon4, alt4 := 39.197807, -77.108574 , 55.0 // Some location near Baltimore
		// that the author of the Perl module geo-ecef used. I reused the coords of the tests.
		x, y, z := geo1.ToECEF(lat4, lon4, alt4)
		fmt.Printf("x = %v \ny = %v \nz = %v\n", x, y, z)

		// Convert ECEF to Lat-Lon-Alt.
		x1, y1, z1 := 1.1042590709397183e+06, -4.824765955871677e+06, 4.0093940281868847e+06
		lat5, lon5, alt5 := geo1.ToLLA(x1, y1, z1)
		fmt.Printf("lat5 = %v lon5 = %v alt3 = %v\n", lat5, lon5, alt5)
	}

To run the application, put the code in a file called hello.go and run:

    go run hello.go

This should print:

		Distance = 543044.190419953 Bearing = 137.50134015496275
		lat3 = 37.74631054036373 lon3 = -122.21438161492877
		x = 1.1042590709397183e+06
		y = -4.824765955871677e+06
		z = 4.0093940281868847e+06
		lat5 = 39.197807 lon5 = -77.10857400000002 alt3 = 55

### Parameters

## Init

The first argument is an ellipsoid from this list:

	"AIRY":               ellipse{6377563.396, 299.3249646},
	"AIRY-MODIFIED":      ellipse{6377340.189, 299.3249646},
	"AUSTRALIAN":         ellipse{6378160.0, 298.25},
	"BESSEL-1841":        ellipse{6377397.155, 299.1528128},
	"CLARKE-1880":        ellipse{6378249.145, 293.465},
	"EVEREST-1830":       ellipse{6377276.345, 300.8017},
	"EVEREST-MODIFIED":   ellipse{6377304.063, 300.8017},
	"FISHER-1960":        ellipse{6378166.0, 298.3},
	"FISHER-1968":        ellipse{6378150.0, 298.3},
	"GRS80":              ellipse{6378137.0, 298.25722210088},
	"HOUGH-1956":         ellipse{6378270.0, 297.0},
	"HAYFORD":            ellipse{6378388.0, 297.0},
	"IAU76":              ellipse{6378140.0, 298.257},
	"KRASSOVSKY-1938":    ellipse{6378245.0, 298.3},
	"NAD27":              ellipse{6378206.4, 294.9786982138},
	"NWL-9D":             ellipse{6378145.0, 298.25},
	"SOUTHAMERICAN-1969": ellipse{6378160.0, 298.25},
	"SOVIET-1985":        ellipse{6378136.0, 298.257},
	"WGS72":              ellipse{6378135.0, 298.26},
	"WGS84":              ellipse{6378137.0, 298.257223563}

The second argument is either 

	Degrees or Radians
	
The third argument is either

	Longitude_is_symmetric or Longitude_not_symmetric

That's internally a boolean, true or false.

The fourth argument is either

	Bearing_is_symmetric or Bearing_not_symmetric

That's also internally a boolean, true or false.

## To and At

See examples.

## Documentation

Read the code or google for Geo::Ellipsoid, that is the Perl package
on CPAN that this package is a port of.

## About and Acknowledgments

This package was ported from Perl to Go by Stefan Schroeder.

Thank you to Jim Gibson for writing the Perl module Geo::Ellipsoid.
And to the authors of the Fortran module that he ported it from.

This package has no other website other than github.

## Bugs and Limitations

Only To and At-functions are implemented from the Geo-Ellipsoid-package.


