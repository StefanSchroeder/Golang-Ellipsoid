package main

// This short script is used for comparison of the golang-package results
// with those of the Perl module Geo::Ellipsoid.
// It is obsolete and superseded by the builtin tests and only
// here for reference.
import "github.com/StefanSchroeder/Golang-Ellipsoid/ellipsoid"
import "fmt"

func main() {
	lat1, lon1 := 37.619002, -122.374843
	lon2, lat2 := 33.942536, -118.408074
	// To
	{
		geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("1 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("2 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Radians, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingNotSymmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("3 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Radians, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("4 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Kilometer, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("5 dist = %v bear = %v\n", dist, bear)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Foot, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		dist, bear := geo1.To(lat1, lon1, lon2, lat2)
		fmt.Printf("6 dist = %v bear = %v\n", dist, bear)
	}
	// At
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Foot, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		lat2, lon2 := geo1.At(lat1, lon1, 2000.0, 45.0)
		fmt.Printf("7 lat  = %v lon = %v\n", lat2, lon2)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		lat2, lon2 := geo1.At(lat1, lon1, 2000.0, 45.0)
		fmt.Printf("8 lat  = %v lon = %v\n", lat2, lon2)
	}
	{
		geo1 := ellipsoid.Init("AIRY", ellipsoid.Degrees, ellipsoid.Meter, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		lat2, lon2 := geo1.At(90.0, 90.0, 1000.0, 90.0)
		fmt.Printf("9 lat  = %v lon = %v\n", lat2, lon2)
	}
	// To
	{
		geo1 := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Nm, ellipsoid.LongitudeIsSymmetric, ellipsoid.BearingIsSymmetric)
		lat3, lon3 := 73.06, 19.11 // Mumbai
		lat4, lon4 := 4.89, 52.37  // Amsterdam
		dist, bear := geo1.To(lat3, lon3, lat4, lon4)
		fmt.Printf("10 dist  = %v bear = %v\n", dist, bear)
	}
}
