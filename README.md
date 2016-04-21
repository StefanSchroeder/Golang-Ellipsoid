[![GoDoc](https://godoc.org/github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid?status.png)](https://godoc.org/github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid)
[![Build Status](https://travis-ci.org/StefanSchroeder/Golang-Ellipsoid.svg?branch=master)](https://travis-ci.org/StefanSchroeder/Golang-Ellipsoid)
[![Go Report Card](http://goreportcard.com/badge/StefanSchroeder/Golang-Ellipsoid)](http://goreportcard.com/report/StefanSchroeder/Golang-Ellipsoid) 


# ellipsoid.go

ellipsoid.go performs latitude and longitude calculations 
on the surface of an ellipsoid. And converts ECEF to LLA and
vice-versa.

This is a Go conversion of an existing Perl conversion 
of existing Fortran code (To and At-functions; see ACKNOWLEDGEMENTS) and the 
author of this package makes no claims of originality. Nor 
can he even vouch for the results of the calculations, 
although they do seem to work for him and have been 
tested against other methods.

It has been successfully tested with Go 1.3, 1.4, 1.5, 1.6.

## Overview

* Calculating distance and bearing when two locations with longitude and latitude are given (To).
* Calculate target location when one location with longitude and latitude and distance and bearing are given (At).
* Supports several ellipsoids (incl. WGS84) out of the box.
* Convert cartesian ECEF-coordinates to longitude, latitude, altitude (*ToLLA*) and vice versa (ToECEF).
* Supports computation of lat-lon conversion to cartesian x,y (Displacement and Location). Handle with care.

## Installation

Make sure you have the a working Go environment. See the [install instructions](http://golang.org/doc/install.html). 

Install package with

	go get github.com/StefanSchroeder/Golang-Ellipsoid

## Example
    
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

To run the application, put the code in a file called hello-wgs84.go and run:

    go run hello-wgs84.go

This should print:

		Distance = 543044.190419953 Bearing = 137.50134015496275
		lat3 = 37.74631054036373 lon3 = -122.21438161492877
		x = 1.1042590709397183e+06
		y = -4.824765955871677e+06
		z = 4.0093940281868847e+06
		lat5 = 39.197807 lon5 = -77.10857400000002 alt3 = 55

## Parameters

### Init

The first argument is an ellipsoid from this list:

        "AIRY":                  {6377563.396, 299.3249646},
        "AIRY-MODIFIED":         {6377340.189, 299.3249646},
        "AUSTRALIAN":            {6378160.0, 298.25},
        "BESSEL-1841":           {6377397.155, 299.1528128},
        "BESSEL-1841-NAMIBIA":   {6377483.865, 299.152813},
        "CLARKE-1866":           {6378206.400, 294.978698},
        "CLARKE-1880":           {6378249.145, 293.465},
        "EVEREST-1830":          {6377276.345, 300.8017},
        "EVEREST-1948":          {6377304.063, 300.8017},
        "EVEREST-SABAH-SARAWAK": {6377298.556, 300.801700},
        "EVEREST-1956":          {6377301.243, 300.801700},
        "EVEREST-1969":          {6377295.664, 300.801700},
        "FISHER-1960":           {6378166.0, 298.3},
        "FISCHER-1960-MODIFIED": {6378155.000, 298.300000},
        "FISHER-1968":           {6378150.0, 298.3},
        "GRS80":                 {6378137.0, 298.25722210088},
        "HELMERT-1906":          {6378200.000, 298.300000},
        "HOUGH-1956":            {6378270.0, 297.0},
        "HAYFORD":               {6378388.0, 297.0},
        "IAU76":                 {6378140.0, 298.257},
        "INTERNATIONAL":         {6378388.000, 297.000000},
        "KRASSOVSKY-1938":       {6378245.0, 298.3},
        "NAD27":                 {6378206.4, 294.9786982138},
        "NWL-9D":                {6378145.0, 298.25},
        "SGS85":                 {6378136.000, 298.257000},
        "SOUTHAMERICAN-1969":    {6378160.0, 298.25},
        "SOVIET-1985":           {6378136.0, 298.257},
        "WGS60":                 {6378165.000, 298.300000},
        "WGS66":                 {6378145.000, 298.250000},
        "WGS72":                 {6378135.0, 298.26},
        "WGS84":                 {6378137.0, 298.257223563},

The second argument is either 

	Degrees or Radians
	
This parameter applies to input- and output-parameters.
	
The third argument is either

	LongitudeIsSymmetric or LongitudeNotSymmetric

That's internally a boolean, true or false. If the longitude is symmetric the
longitude in the result of the At-function will we be in the range [-180..180],
else (not-symmetric) the range will be [0..360].

The fourth argument is either

	BearingIsSymmetric or BearingNotSymmetric

That's also internally a boolean, true or false.

### To

The To-Function computes the distance in the units provided to Init as a Float64 and the bearing in degrees [0...360]
or Radian. Input parameters are the latitude and longitude of the starting point and
latitude and longitude of the destination. All parameters are Float64. The bearing is 
the direction when standing on the starting point and looking towards the destination point.
Obviously the compass direction is not too meaningful near the poles.

	distance, bearing := geo1.To(lat1, lon1, lat2, lon2)

### At

Returns the list latitude,longitude in degrees or radians that is a
specified range and bearing from a given location.

	lat2, lon2  = geo.At( lat1, lon1, range, bearing )

### Intermediate

Takes two coordinates with longitude and latitude; and a step count and 
returns range and bearing and an array with the lons and lats of intermediate
points on a straight line (whatever that is on an ellipsoid), INCLUDING the
start and the endpoint.

So if you put in point1 and point2 with step count 4, the output will be

	point1
	i1
	i2
	i3
	point2

Each point is two float64 values, lat and lon, thus you have an array
with 4x2 + 2 = 5x2 cells.

Steps cannot be 0.

I have not tested the upper limit for steps.

### ToECEF

The ToECEF-Function computes the ECEF tripel for a set of 
latitude, longitude and altitude for the given ellipsoid object.

	x, y, z := geo.ToECEF(lat, lon, alt)

The inverse function is ToLLA.

### ToLLA

The ToLLA-Function computes latitude, langitude and elevation for an ECEF
tripel.

	lat, lon, alt := geo.ToLLA(x, y, z)

The inverse function is ToECEF.

### Displacement

Returns the (x,y) displacement in distance units between the two specified 
locations.

    x, y  = geo.Displacement( lat1, lon1, lat2, lon2 )
    
NOTE: The x and y displacements are only approximations and only valid
between two locations that are fairly near to each other. Beyond 10 kilometers
or more, the concept of X and Y on a curved surface loses its meaning.


### Location

Returns the list (latitude,longitude) of a location at a given (x,y)
displacement from a given location.

	lat2, lon2 = geo.Location( lat1, lon1, x, y )

The note from Displacement applies.


### Notes

If you need background information read the code or go to Geo::Ellipsoid or Geo::ECEF, these are the Perl modules
on CPAN that this package is a port of. 

## About and Acknowledgments

This package was ported from Perl to Go by Stefan Schroeder.

Thank you to Jim Gibson for writing the Perl module Geo::Ellipsoid.
And to the authors of the Fortran module that he ported it from.

This package has no website other than github.

## Bugs and Limitations

Not all functions are implemented from the Geo-Ellipsoid-package. 



