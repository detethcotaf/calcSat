package calcSat

import (
	"errors"
	"math"
)

// init - last の数値をintervalごとにリストにして返す
func CreateNumList(init, last, interval float64)(nums []float64, err error){
	if interval < 0.01 {
		return nums, errors.New("[Error] interval must be 0.01 or higher")
	}

	for i := init; i < last; {
		nums = append(nums, i)
		i = math.Round((i + interval)*100) / 100
	}
	err = nil
	return nums, err
}

// 数値のリストの中から最大値と最小値を返す
func GetMaxMin(inputList []float64)(max, min float64, err error){
	if len(inputList) == 0 {
		return 0, 0, errors.New("input list is empty")
	}
	max = inputList[0]
	min = inputList[0]
	for _, p := range inputList{
		max = math.Max(max, p)
		min = math.Min(min, p)
	}
	err = nil
	return
}

// thetaデータの生成
func GetThetaData() (thetaData []float64){
	for angle := 0.0; angle <= 360; {
		theta := angle * math.Pi / 180
		thetaData = append(thetaData, theta)
		angle = angle + 10
	}
	return
}

// 座標変換 直交座標 → 極座標
func Cartesian2polar(x, y float64)(r, theta float64) {
	r = math.Hypot(x, y) // => math.Sqrt(x**2 + y**2)
	theta = math.Atan2(y, x)
	return
}

// 座標変換 極座標 → 直交座標
func Polar2cartesian(r, theta float64)(x, y float64){
	x = r * math.Cos(theta)
	y = r * math.Sin(theta)
	return
}

