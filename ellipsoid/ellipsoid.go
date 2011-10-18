// This is a Go conversion of an existing Perl conversion 
// of existing Fortran code (see ACKNOWLEDGEMENTS) and the 
// author of this class makes no claims of originality. Nor 
// can he even vouch for the results of the calculations, 
// although they do seem to work for him and have been 
// tested against other methods.

// ellipsoid performs latitude and longitude calculations 
// on the surface of an ellipsoid.
package ellipsoid

// Version 1.0 based on Geo::Ellipsoid Version 1.12.

/*

SYNOPSIS

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
	// direction 45.0 deg. for 2000 meters.
	lon3, lat3 := geo1.At(lat1, lon1, 2000.0, 45.0)
	fmt.Printf("lat3 = %v lon3 = %v\n", lat3, lon3)
}


DESCRIPTION

ellipsoid performs geometrical calculations on the surface of 
an ellipsoid. An ellipsoid is a three-dimension object formed from 
the rotation of an ellipse about one of its axes. The approximate 
shape of the earth is an ellipsoid, so ellipsoid can accurately
calculate distance and bearing between two widely-separated locations 
on the earth's surface.

The shape of an ellipsoid is defined by the lengths of its
semi-major and semi-minor axes. The shape may also be specifed by
the flattening ratio f as:

    f = ( semi-major - semi-minor ) / semi-major

which, since f is a small number, is normally given as the reciprocal 
of the flattening 1/f.

The shape of the earth has been surveyed and estimated differently
at different times over the years. The two most common sets of values
used to describe the size and shape of the earth in the United States
are 'NAD27', dating from 1927, and 'WGS84', from 1984. United States
Geological Survey topographical maps, for example, use one or the
other of these values, and commonly-available Global Positioning
System (GPS) units can be set to use one or the other.
See "DEFINED ELLIPSOIDS" below for the ellipsoid survey values
that may be selected for use by ellipsoid.

*/

import "math"
import "fmt"

const (
	pi                      = math.Pi
	twopi                   = math.Pi * 2.0
	max_loop_count          = 20
	eps                     = 1.0e-23
	debug                   = false
	Meter                   = 0 //    1.0    meter
	Foot                    = 1 //    0.3048 meter are a foot
	Kilometer               = 2 // 1000.0    meter are a kilometer
	Mile                    = 3 // 1609.344  meter are a mile
	Nm                      = 4 // 1852.0    meter are a nautical mile, 
	Degrees                 = iota
	Radians                 = iota
	Longitude_is_symmetric  = true
	Longitude_not_symmetric = false
	Bearing_is_symmetric    = true
	Bearing_not_symmetric   = false
)

type Ellipsoid struct {
	Ellipse             ellipse
	Units               int
	Distance_units      int
	Longitude_symmetric bool
	Bearing_symmetry    bool
	Distance_factor     float64
	// Having the Distance_factor AND the Distance_units in this struct is redundant
	// but it looks nicer in the code.
}

type ellipse struct {
	Equatorial     float64
	Inv_flattening float64
}

func deg2rad(d float64) (r float64) {
	return d * pi / 180.0
}
func rad2deg(d float64) (r float64) {
	return d * 180.0 / pi
}

/* Init

The Init constructor must be called with a list of parameters to set 
the value of the ellipsoid to be used, the value of the units to be 
used for angles and distances, and whether or not the output range 
of longitudes and bearing angles should be symmetric around zero 
or always greater than zero. There is no default constructor, all 
arguments are required; they may not be abbreviated.

Example:

	geo := ellipsoid.Init(
		"WGS84",  // for possible values see below.
		ellipsoid.Degrees, // possible values: Degrees or Radians 
		ellipsoid.Meter,   // possible values: Meter, Kilometer, 
				   // Foot, Nm, Mile
		ellipsoid.Longitude_is_symmetric, // possible values 
						  // Longitude_is_symmetric or 
						  // Longitude_not_symmetric
		ellipsoid.Bearing_is_symmetric    // possible 
						  // values Bearing_is_symmetric or 
						  // Bearing_not_symmetric
	)

*/
func Init(name string, units int, dist_units int, long_sym bool, bear_sym bool) (e Ellipsoid) {
	m := map[string]ellipse{
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
		"WGS84":              ellipse{6378137.0, 298.257223563}}

	e2, ok := m[name]
	if !ok {
		fmt.Printf("ellipsoid.go: Warning: Invalid ellipse type '%v'\n", name)
	}

	//                      m    ft      km      mi        nm
	conversion := []float64{1.0, 0.3048, 1000.0, 1609.344, 1852.0}
	ellipsoid := Ellipsoid{e2, units, dist_units, long_sym, bear_sym, conversion[dist_units]}
	return ellipsoid
}

/* To

Returns range, bearing between two specified locations.

    dist, theta  = geo.To( lat1, lon1, lat2, lon2 )

*/
func (ellipsoid Ellipsoid) To(lat1, lon1, lat2, lon2 float64) (distance, bearing float64) {

	if ellipsoid.Units == Degrees {
		lat1 = deg2rad(lat1)
		lon1 = deg2rad(lon1)
		lat2 = deg2rad(lat2)
		lon2 = deg2rad(lon2)
	}

	distance, bearing = ellipsoid.calculateBearing(lat1, lon1, lat2, lon2)
	if ellipsoid.Units == Degrees {
		bearing = rad2deg(bearing)
	}

	distance /= ellipsoid.Distance_factor

	return
}

/* At

Returns the list latitude,longitude in degrees or radians that is a
specified range and bearing from a given location.

    lat2, lon2  = geo.At( lat1, lon1, range, bearing )

*/
func (ellipsoid Ellipsoid) At(lat1, lon1, distance, bearing float64) (lat2, lon2 float64) {

	if ellipsoid.Units == Degrees {
		lat1 = deg2rad(lat1)
		lon1 = deg2rad(lon1)
		bearing = deg2rad(bearing)
	}

	lat2, lon2 = ellipsoid.calculateTargetlocation(lat1, lon1, distance, bearing)

	if ellipsoid.Longitude_symmetric == Longitude_is_symmetric {
		if lon2 > pi {
			lon2 -= twopi
		}
	}
	if ellipsoid.Longitude_symmetric == Longitude_not_symmetric {
		if lon2 < 0.0 {
			lon2 += twopi
		}
	}

	if ellipsoid.Units == Degrees {
		lat2 = rad2deg(lat2)
		lon2 = rad2deg(lon2)
	}

	return
}

func (ellipsoid Ellipsoid) calculateTargetlocation(lat1, lon1, distance, bearing float64) (lat2, lon2 float64) {

	if debug == true {
		fmt.Printf("_forward(lat1=%v,lon1=%v,range=%v,bearing=%v)\n", lat1, lon1, distance, bearing)
	}

	eps := 0.5e-13

	a := ellipsoid.Ellipse.Equatorial
	f := 1 / ellipsoid.Ellipse.Inv_flattening
	r := 1.0 - f

	clat1 := math.Cos(lat1)
	if clat1 == 0 {
		fmt.Printf("WARNING: Division by Zero in ellipsoid.go.\n")
		return 0.0, 0.0
	}
	tu := r * math.Sin(lat1) / clat1
	faz := bearing

	s := ellipsoid.Distance_factor * distance

	sf := math.Sin(faz)
	cf := math.Cos(faz)

	baz := 0.0
	if cf != 0.0 {
		baz = 2.0 * math.Atan2(tu, cf)
	}

	cu := 1.0 / math.Sqrt(1.0+tu*tu)
	su := tu * cu
	sa := cu * sf
	c2a := 1.0 - (sa * sa)
	x := 1.0 + math.Sqrt((((1.0/(r*r))-1.0)*c2a)+1.0)
	x = (x - 2.0) / x
	c := 1.0 - x
	c = (((x * x) / 4.0) + 1.0) / c
	d := x * ((0.375 * x * x) - 1.0)
	tu = ((s / r) / a) / c
	y := tu

	if debug == true {
		fmt.Printf("r=%.8f, tu=%.8f, faz=%.8f\n", r, tu, faz)
		fmt.Printf("baz=%.8f, sf=%.8f, cf=%.8f\n", baz, sf, cf)
		fmt.Printf("cu=%.8f, su=%.8f, sa=%.8f\n", cu, su, sa)
		fmt.Printf("x=%.8f, c=%.8f, y=%.8f\n", x, c, y)
	}

	var cy, cz, e, sy float64
	for true {
		sy = math.Sin(y)
		cy = math.Cos(y)
		cz = math.Cos(baz + y)
		e = (2.0 * cz * cz) - 1.0
		c = y
		x = e * cy
		y = (2.0 * e) - 1.0
		y = (((((((((sy * sy * 4.0) - 3.0) * y * cz * d) / 6.0) + x) * d) / 4.0) - cz) * sy * d) + tu

		if math.Fabs(y-c) <= eps {
			break
		}
	}
	baz = (cu * cy * cf) - (su * sy)
	c = r * math.Sqrt((sa*sa)+(baz*baz))
	d = su*cy + cu*sy*cf
	lat2 = math.Atan2(d, c)
	c = cu*cy - su*sy*cf
	x = math.Atan2(sy*sf, c)
	c = (((((-3.0 * c2a) + 4.0) * f) + 4.0) * c2a * f) / 16.0
	d = ((((e * cy * c) + cz) * sy * c) + y) * sa
	lon2 = lon1 + x - (1.0-c)*d*f

	if debug == true {
		fmt.Printf("returns(lat2=%v,lon2=%v)\n", lat2, lon2)
	}
	return lat2, lon2
}

func (ellipsoid Ellipsoid) calculateBearing(lat1, lon1, lat2, lon2 float64) (distance, bearing float64) {
	a := ellipsoid.Ellipse.Equatorial
	f := 1 / ellipsoid.Ellipse.Inv_flattening

	if lon1 < 0 {
		lon1 += twopi
	}
	if lon2 < 0 {
		lon2 += twopi
	}

	r := 1.0 - f
	clat1 := math.Cos(lat1)
	if clat1 == 0 {
		fmt.Printf("WARNING: Division by Zero in ellipsoid.go.\n")
		return 0.0, 0.0
	}
	clat2 := math.Cos(lat2)
	if clat2 == 0 {
		fmt.Printf("WARNING: Division by Zero in ellipsoid.go.\n")
		return 0.0, 0.0
	}
	tu1 := r * math.Sin(lat1) / clat1
	tu2 := r * math.Sin(lat2) / clat2
	cu1 := 1.0 / (math.Sqrt((tu1 * tu1) + 1.0))
	su1 := cu1 * tu1
	cu2 := 1.0 / (math.Sqrt((tu2 * tu2) + 1.0))
	s := cu1 * cu2
	baz := s * tu2
	faz := baz * tu1
	dlon := lon2 - lon1

	if debug == true {
		fmt.Printf("a=%v, f=%v\n", a, f)
		fmt.Printf("lat1=%v, lon1=%v\n", lat1, lon1)
		fmt.Printf("lat2=%v, lon2=%v\n", lat2, lon2)

		fmt.Printf("r=%v, tu1=%v, tu2=%v\n", r, tu1, tu2)
		fmt.Printf("faz=%.8f, dlon=%.8f, su1=%v\n", faz, dlon, su1)
	}

	x := dlon
	cnt := 0

	var c2a, c, cx, cy, cz, d, del, e, sx, sy, y float64
	// This originally was a do-while loop. Exit condition is at end of loop.
	for true {
		if debug == true {
			fmt.Printf("  x=%.8f\n", x)
		}
		sx = math.Sin(x)
		cx = math.Cos(x)
		tu1 = cu2 * sx
		tu2 = baz - (su1 * cu2 * cx)

		if debug == true {
			fmt.Printf("    sx=%.8f, cx=%.8f, tu1=%.8f, tu2=%.8f\n", sx, cx, tu1, tu2)
		}

		sy = math.Sqrt(tu1*tu1 + tu2*tu2)
		cy = s*cx + faz
		y = math.Atan2(sy, cy)
		var sa float64
		if sy == 0.0 {
			sa = 1.0
		} else {
			if sy == 0 {
				fmt.Printf("WARNING: Division by Zero in ellipsoid.go.\n")
				return 0.0, 0.0
			}
			sa = (s * sx) / sy
		}

		if debug == true {
			fmt.Printf("    sy=%.8f, cy=%.8f, y=%.8f, sa=%.8f\n", sy, cy, y, sa)
		}

		c2a = 1.0 - (sa * sa)
		cz = faz + faz
		if c2a > 0.0 {
			cz = ((-cz) / c2a) + cy
		}
		e = (2.0 * cz * cz) - 1.0
		c = (((((-3.0 * c2a) + 4.0) * f) + 4.0) * c2a * f) / 16.0
		d = x
		x = ((e*cy*c+cz)*sy*c + y) * sa
		x = (1.0-c)*x*f + dlon
		del = d - x

		if debug == true {
			fmt.Printf("    c2a=%.8f, cz=%.8f\n", c2a, cz)
			fmt.Printf("    e=%.8f, d=%.8f\n", e, d)
			fmt.Printf("    (d-x)=%.8g\n", del)
		}
		if math.Fabs(del) <= eps {
			break
		}
		cnt++
		if cnt > max_loop_count {
			break
		}

	}

	faz = math.Atan2(tu1, tu2)
	baz = math.Atan2(cu1*sx, (baz*cx-su1*cu2)) + pi
	x = math.Sqrt(((1.0/(r*r))-1.0)*c2a+1.0) + 1.0
	x = (x - 2.0) / x
	c = 1.0 - x
	c = ((x*x)/4.0 + 1.0) / c
	d = ((0.375 * x * x) - 1.0) * x
	x = e * cy

	if debug == true {
		fmt.Printf("e=%.8f, cy=%.8f, x=%.8f\n", e, cy, x)
		fmt.Printf("sy=%.8f, c=%.8f, d=%.8f\n", sy, c, d)
		fmt.Printf("cz=%.8f, a=%.8f, r=%.8f\n", cz, a, r)
	}

	s = 1.0 - e - e
	s = ((((((((sy * sy * 4.0) - 3.0) * s * cz * d / 6.0) - x) * d / 4.0) + cz) * sy * d) + y) * c * a * r

	if debug == true {
		fmt.Printf("s=%.8f\n", s)
	}

	// adjust azimuth to (0,360) or (-180,180) as specified
	if ellipsoid.Bearing_symmetry == Bearing_is_symmetric {
		if faz < -(pi) {
			faz += twopi
		}
		if faz >= pi {
			faz -= twopi
		}
	} else {
		if faz < 0 {
			faz += twopi
		}
		if faz >= twopi {
			faz -= twopi
		}
	}

	distance, bearing = s, faz
	return
}
/*
 DEFINED ELLIPSOIDS

The following ellipsoids are defined in Geo::Ellipsoid, with the 
semi-major axis in meters and the reciprocal flattening as shown.


    Ellipsoid        Semi-Major Axis (m.)     1/Flattening
    ---------        -------------------     ---------------
    AIRY                 6377563.396         299.3249646
    AIRY-MODIFIED        6377340.189         299.3249646
    AUSTRALIAN           6378160.0           298.25
    BESSEL-1841          6377397.155         299.1528128
    CLARKE-1880          6378249.145         293.465
    EVEREST-1830         6377276.345         290.8017
    EVEREST-MODIFIED     6377304.063         290.8017
    FISHER-1960          6378166.0           298.3
    FISHER-1968          6378150.0           298.3
    GRS80                6378137.0           298.25722210088
    HOUGH-1956           6378270.0           297.0
    HAYFORD              6378388.0           297.0
    IAU76                6378140.0           298.257
    KRASSOVSKY-1938      6378245.0           298.3
    NAD27                6378206.4           294.9786982138
    NWL-9D               6378145.0           298.25
    SOUTHAMERICAN-1969   6378160.0           298.25
    SOVIET-1985          6378136.0           298.257
    WGS72                6378135.0           298.26
    WGS84                6378137.0           298.257223563

 LIMITATIONS

The methods should not be used on points which are too near the poles
(above or below 89 degrees), and should not be used on points which
are antipodal, i.e., exactly on opposite sides of the ellipsoid. The
methods will not return valid results in these cases.

The Go-version does not support all features of the Perl module. If you
need advanced features, like defining your own ellipses at runtime, 
calculating x,y dislocations, etc, please refer to the package on CPAN

Geo::Ellipsoid

http://search.cpan.org/~jgibson/Geo-Ellipsoid-1.12/lib/Geo/Ellipsoid.pm

FIXME: Add more checks for div by 0.

 ACKNOWLEDGEMENTS

The conversion algorithms used here are Perl translations of Fortran
routines written by LCDR L. Pfeifer NGS Rockville MD that implement
T. Vincenty's Modified Rainsford's method with Helmert's elliptical
terms as published in "Direct and Inverse Solutions of Ellipsoid on
the Ellipsoid with Application of Nested Equations", T. Vincenty,
Survey Review, April 1975.

The Fortran source code files inverse.for and forward.for
may be obtained from

    ftp://ftp.ngs.noaa.gov/pub/pcsoft/for_inv.3d/source/

 AUTHOR

Jim Gibson, <Jim@Gibson.org> (Perl version)
Stefan Schroeder <ondekoza@gmail.com> (Port from Perl to Golang)

 BUGS

See LIMITATIONS, above.

Please report any bugs or feature requests to
the author.

COPYRIGHT & LICENSE

Copyright 2005-2008 Jim Gibson, all rights reserved.

This program is free software; you can redistribute it and/or modify it
under the same terms as Perl.

*/