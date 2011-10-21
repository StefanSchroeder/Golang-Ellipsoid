//#package ellipsoid
package main

import "fmt"
import "math"
import "./ellipsoid"

var pass, fail int

func delta_within(v float64, t float64, d float64) {
	diff := math.Fabs(v - t)
	if diff < d {
		fmt.Printf("Test OK\n")
		pass++
	} else {
		fmt.Printf("Test FAIL\n")
		fail++
		fmt.Printf("v - t = %f - %f = %f\n", v, t, diff)
	}
	fmt.Printf("---\n")
}

type testobject struct {
	t string
	lat1 float64
	lon1 float64
	lat2 float64
	lon2 float64
	delta float64
}

func main() {
	epsilon := 50.0
        all_tests := []testobject{
		{"Istanbul to Delhi",41.1, 29, 28.67, 77.21, 4550},
		{"Tokyo to New_York",35.67, 139.77, 40.67, -73.94,  10846 },
		{"Lima to Cairo",-12.07, -77.05, 30.06, 31.25,  12423 },
		{"Lahore to Rio_de_Janeiro",31.56, 74.35, -22.91, -43.2,  13845 },
		{"Santiago to Kaifeng",19.48, -70.69, 34.85, 114.35,  19529 },
		{"Calcutta to Toronto",22.57, 88.36, 43.65, -79.38,  12540 },
		{"Rangoon to Sydney",16.79, 96.15, -33.87, 151.21,  8104 },
		{"Madras to Riyadh",13.09, 80.27, 24.65, 46.77,  3739 },
		{"Chongqing to Chengdu",29.57, 106.58, 30.67, 104.07,  268 },
		{"Tianjin to Melbourne",39.13, 117.2, -37.81, 144.96,  9015 },
		{"Pusan to Abidjan",35.11, 129.03, 5.33, -4.03,  13362 },
		{"Yokohama to Ibadan",35.47, 139.62, 7.38, 3.93,  13373 },
		{"Singapore to Ankara",1.3, 103.85, 39.93, 32.85,  8303 },
		{"Berlin to Montreal",52.52, 13.38, 45.52, -73.57,  6001 },
		{"Pyongyang to Lanzhou",39.02, 125.75, 36.05, 103.68,  1959 },
		{"Guangzhou to Casablanca",23.12, 113.25, 33.6, -7.62,  11132 },
		{"Durban to Madrid",-29.87, 30.99, 40.42, -3.71,  8593 },
		{"Nanjing to Kabul",32.05, 118.78, 34.53, 69.17,  4571 },
		{"Pune to Surat",18.53, 73.84, 21.2, 72.82,  312 },
		{"Jiddah to Chicago",21.5, 39.17, 41.84, -87.68,  11102 },
		{"Kanpur to Luanda",26.47, 80.33, -8.82, 13.24,  8228 },
		{"Taiyuan to Salvador",37.87, 112.55, -12.97, -38.5,  16033 },
		{"Taegu to Rome",35.87, 128.6, 41.89, 12.5,  9202 },
		{"Changchun to Kiev",43.87, 125.35, 50.43, 30.52,  6701 },
		{"Faisalabad to Izmir",31.41, 73.11, 38.43, 27.15,  4215 }}

	e := ellipsoid.Init("WGS84", ellipsoid.Degrees, ellipsoid.Kilometer, ellipsoid.Longitude_is_symmetric, ellipsoid.Bearing_is_symmetric)
	for _, v := range all_tests {
		fmt.Printf("Going from %v\n", v.t)
		r, _ := e.To(v.lat1, v.lon1, v.lat2, v.lon2)
		delta_within(r, v.delta, epsilon)
	}

	fmt.Printf("Summary: pass=%v fail=%v\n", pass, fail)
}
