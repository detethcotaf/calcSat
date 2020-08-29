package calcSat

import (
	"fmt"
	"math"
)

// 平均近点離角
func MeanAnomaly(a, deltaT float64)float64{
	return math.Sqrt(gravitationalConstant/math.Pow(a, 3)) * deltaT
}

// ケプラー方程式
func KeplerEquation(e, M float64) func(Ebefore float64) float64{
	// Ebefore: n回目の離心近点離角
	// M: 平均近点離角
	// e: 離心率
	// Eafter: n+1回目の離心近点離角

	return func(Ebefore float64)float64{
		FE := Ebefore - e*math.Sin(Ebefore) - M
		Eafter := Ebefore - FE/(1-e*math.Cos(Ebefore))
		return Eafter
	}

}

// ケプラー方程式をニュートン-ラフソン法で解く
func NewtonRaphson(e, before, ae float64) float64 {
	// M: 平均近点離角
	// e: 離心率
	// a: allowable error 許容誤差

	equation := KeplerEquation(e, before)

	var after float64
	err := 100.0 // 許容誤差の初期化（0だとforが回らないため）
	count := 0

	for ; err > ae; {
		after = equation(before)
		err = math.Abs(after - before)
		fmt.Println("許容誤差: ", err)

		before = after
		count = count + 1
		fmt.Println("繰り返し：", count, "回目")
		fmt.Println("----")
	}
	return after
}

// 宇宙機や天体から中心星までの距離r
func DistanceSat2Star(a, e, E float64)float64{
	return a * (1 - e * math.Cos(E))
}

// 経路角を求める
func PathAngle(e, theta float64)float64{
	return math.Atan(e*math.Sin(theta)/(1+e*math.Cos(theta)))
}

// 軌道遷移による速度増分を算出する
func OrbitalTransitionSpeed(initialAltitude, targetAltitude, a1, a2 float64)(float64, float64){
	initialVelocity := CircleOrbitVelocity(initialAltitude)
	periVelocity, apoVelocity := EllipseOrbitVelocity(initialAltitude, targetAltitude)
	targetVelocity := CircleOrbitVelocity(targetAltitude)

	fmt.Println("initial velocity: ", initialVelocity)
	fmt.Println("peri velocity: ", periVelocity)
	fmt.Println("apo velocity: ", apoVelocity)
	fmt.Println("target velocity: ", targetVelocity)

	deltaInitialSpeed := math.Sqrt(math.Pow(periVelocity, 2) + math.Pow(initialVelocity,2) - 2*periVelocity*initialVelocity*math.Cos(a1))
	deltaTargetSpeed := math.Sqrt(math.Pow(targetVelocity, 2) + math.Pow(apoVelocity,2) - 2*targetVelocity*apoVelocity*math.Cos(a2))

	return deltaInitialSpeed, deltaTargetSpeed
}