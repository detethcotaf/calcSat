package calcSat

import (
	"fmt"
	"math"
)



// 楕円軌道の飛行時間
// 半長径a = (peri + apo)/2
func EllipseFlightTime(peri, apo, e, theta float64)float64{
	//return math.Sqrt(math.Pow((peri+apo)/2, 3)/gravitationalConstant) *
	//	(2*math.Atan(math.Sqrt((1-e)/(1+e)) * math.Tan(theta/2)) -
	//		e*math.Sqrt(1-math.Pow(e,2))*math.Sin(theta)/(1+e*math.Cos(theta)))

	a := math.Sqrt(math.Pow((peri+apo)/2, 3)/gravitationalConstant)
	b := 2*math.Atan(math.Sqrt((1-e)/(1+e)) * math.Tan(theta/2))
	c := e*math.Sqrt(1-math.Pow(e,2))*math.Sin(theta)/(1+e*math.Cos(theta))
	fmt.Println("a: ",a)
	fmt.Println("b: ",b)
	fmt.Println("c: ",c)
	//
	//fmt.Println(math.Sqrt((1-e)/(1+e)) * math.Tan(theta/2))
	//fmt.Println("mark: ", math.Sqrt((1-e)/(1+e)))
	return a * (b - c)
}

// 双曲線軌道の飛行時間
func HyperbolaFlightTime(peri, apo, e, theta float64)float64{
	a := math.Sqrt(math.Pow((peri+apo)/2, 3)/gravitationalConstant)
	b := e * math.Sqrt(math.Pow(e,2)-1) * math.Sin(theta) / (1+e*math.Cos(theta))
	c := math.Log( (math.Sqrt(e+1) + math.Sqrt(e-1)*math.Tan(theta/2)) /
		(math.Sqrt(e+1) - math.Sqrt(e-1)*math.Tan(theta/2)) )

	return a * (b-c)
}

// 楕円軌道の速度
func EllipseOrbitVelocity(peri, apo float64)(periVelocity, apoVelocity float64){
	periVelocity = math.Sqrt(
		2 * gravitationalConstant * (1 / (earthRadius + peri) - 1 / (earthRadius * 2 + peri + apo)),
		)
	apoVelocity = math.Sqrt(
		2 * gravitationalConstant * (1 / (earthRadius + apo) - 1 / (earthRadius * 2 + peri + apo)),
		)
	return
}

// 楕円軌道の周期
func EllipsePeriod(peri, apo float64)(period float64){
	period = math.Pi * math.Sqrt(math.Pow(2 * earthRadius + peri + apo, 3) / (2 * gravitationalConstant))
	return
}


// 円軌道の速度
func CircleOrbitVelocity(altitude float64)(velocity float64){
	velocity = math.Sqrt(gravitationalConstant / (altitude + earthRadius))
	return
}

// 円軌道の周期
func CirclePeriod(altitude float64)(period float64){
	period = 2 * math.Pi * math.Sqrt(math.Pow(altitude + earthRadius, 3) / gravitationalConstant)
	return
}


// 双曲線余剰速度
// a = 半長径 (semi-major axis)
func calcHEV(a float64)(hev float64){
	// hev = hyperbolicExcessVelocity
	hev = math.Sqrt(gravitationalConstant/a)
	return
}

// 地球と太陽の距離を半径とした円軌道の速度
func CircleOrbitVelocitySun()float64{
	return math.Sqrt(gravitationalConstantSun / (distanceEarthSun))
}

// 地球が止まっていた場合の太陽からの脱出速度
func escapeSunFromEarthVelocity()float64{
	return math.Sqrt(2) * CircleOrbitVelocitySun()
}

// 地球にいた場合の太陽からの脱出速度
// altitudeは地表からの高度
// 地球が止まっていた場合の太陽からの脱出速度 - 地球の公転速度 を双曲線余剰速度として計算している
func EscapeSunVelocity(altitude float64)float64{
	return math.Sqrt(2*gravitationalConstant/(earthRadius + altitude) +
		math.Pow(escapeSunFromEarthVelocity() - earthOrbitalSpeed, 2))
}


