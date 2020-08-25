package calcSat

import "math"

// 放物線を計算するためのクロージャ（直交座標x, y）
// 焦点距離：q
func CalcParabola(q float64) func(x float64) float64 {
	return func(x float64) float64 {
		return math.Sqrt(4 * q * x)
	}
}

// 楕円


// 双曲線



// 極座標による円錐曲線を計算するためのクロージャ(極座標θ, r)
// 離心率: e, 半直弦: l
func CalcConicPolar(e, l float64) func(theta float64) float64 {
	return func(theta float64) float64 {
		return l / (1 + e * math.Cos(theta))
	}
}



// それぞれの座標を変換する関数
// r2=x2+y2,tanθ=yx

//func convertCoordinate(s string, a, b float64) (float64, float64){
//	if s == "xy" {
//		x := a
//		y := b
//		r := math.Sqrt((x*x + y*y))
//		t := math.Atan( y / x )
//		return r, t
//	} else if s == "rt" {
//		r := a
//		t := b
//		//x :=
//	}
//}



