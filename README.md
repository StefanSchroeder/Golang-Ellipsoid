# ellipsoid.go

ellipsoid.go performs latitude and longitude calculations 
on the surface of an ellipsoid.

This is a Go conversion of an existing Perl conversion 
of existing Fortran code (see ACKNOWLEDGEMENTS) and the 
author of this class makes no claims of originality. Nor 
can he even vouch for the results of the calculations, 
although they do seem to work for him and have been 
tested against other methods.

## Overview

* Calculating distance and bearing when two locations with longitude and latitude are are given.
* Calculate target location when one location with longitude and latitude and distance and bearing are given.
* Supports several ellipsoid (incl. WGS84) out of the box.

## Installation

Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html). 

## Example
    
	package main

	import "fmt"
	import "geo/ellipsoid"

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
	}

To run the application, put the code in a file called hello.go and run:

    8g hello.go && 8l -o hello hello.8 && ./hello

This should print:

FIXME.

### Parameters

FIXME.

## Documentation

This package is open source. You know what to do.

## About and Acknowledgments

This package was ported from Perl to Go by  Stefan Schroeder.

Thank you to Jim Gibson for writing the Perl module Geo::Ellipsis.
And to the authors of the Fortran module that he ported it from.

This package has no other website other than github.
