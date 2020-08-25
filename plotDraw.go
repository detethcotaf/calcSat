package calcSat

import (
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
	"log"
	"math"
)

func PlotDrawLiner(graphTitle, savePath string, data [][]float64){
	// 図の作成
	p, err := plot.New()
	if err != nil{
		log.Fatalln("error: ", err)
	}

	// グラフの装飾の設定
	p.Title.Text = graphTitle
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid()) // 補助線

	// 描画するデータの取得
	pts := make(plotter.XYs, len(data))
	var xDataList []float64
	var yDataList []float64

	for i, pair := range data {
		pts[i].X = pair[0]
		xDataList = append(xDataList, pair[0])
		pts[i].Y = pair[1]
		yDataList = append(yDataList, pair[1])
	}

	// 折れ線グラフの描画
	err = plotutil.AddLinePoints(p, pts)
	if err != nil {
		log.Fatalln("failed to draw liner: ", err)
	}

	// 座標範囲の設定
	xMax, xMin, err := GetMaxMin(xDataList) // common.goの関数
	if err != nil {
		log.Fatalln("failed to get max and min from xDataList: ", err)
	}
	yMax, yMin, err := GetMaxMin(yDataList) // common.goの関数
	if err != nil {
		log.Fatalln("failed to get max and min from yDataList: ", err)
	}
	p.X.Min = xMin - 1
	p.X.Max = xMax + 1
	p.Y.Min = yMin - 1
	p.Y.Max = yMax + 1

	// 保存
	if err := p.Save(6*vg.Inch, 6*vg.Inch, savePath); err != nil {
		log.Fatalln("Failed to save plot:", err)
	}
}

func Draw(xData []float64, fn func(float64) float64, graphTitle, savePath string){
	var dataList [][]float64

	for _, x := range xData{
		y := fn(x)
		pair := []float64{x, y}
		dataList = append(dataList, pair)
	}
	PlotDrawLiner(graphTitle, savePath, dataList)
	log.Println("success draw and save fig")
}

func DrawConic(e, l float64, thetaList []float64, graphTitle, savePath string){
	c := CalcConicPolar(e, l)

	// radius算出
	var radiusList []float64
	for _, v := range thetaList{
		r := c(v)
		radiusList = append(radiusList, r)
	}

	// 座標変換
	var xys [][]float64
	for i, theta := range thetaList{
		radius := radiusList[i]
		xy := make([]float64, 2)
		x, y := Polar2cartesian(radius, theta)
		xy[0] = math.Round(x*100) / 100
		xy[1] = math.Round(y*100) / 100
		xys = append(xys, xy)
	}

	// 描画＆保存
	PlotDrawLiner(graphTitle, savePath, xys)
}
