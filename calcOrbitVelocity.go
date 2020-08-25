package calcSat

import (
	"math"
)

const gravitationalConstant = 398600.4 // 重力定数
const earthRadius = 6378.137 // 地球の半径

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
func CircleOrbitVelocity(radius float64)(velocity float64){
	velocity = math.Sqrt(gravitationalConstant / (radius + earthRadius))
	return
}

// 円軌道の周期
func CirclePeriod(radius float64)(period float64){
	period = 2 * math.Pi * math.Sqrt(math.Pow(radius + earthRadius, 3) / gravitationalConstant)
	return
}

// 放物線軌道の速度



// 放物線軌道の周期



// 双曲線軌道の速度



// 双曲線軌道の周期


