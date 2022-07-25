package model

import (
	"encoding/xml"
	"io/ioutil"
	"math"
	"snowyuki31/training-dashboard-api/utils"
)

type UnitData struct {
	Time      string  `xml:"Time" json:"time"`
	HeartRate int32   `xml:"HeartRateBpm>Value" json:"heart_rate"`
	Cadence   int32   `xml:"Cadence" json:"cadence"`
	Speed     float64 `xml:"Extensions>TPX>Speed" json:"speed"`
	Watts     int32   `xml:"Extensions>TPX>Watts" json:"watts"`
}

type Activity struct {
	Id         string     `xml:"Activities>Activity>Id" json:"id"`
	Trackpoint []UnitData `xml:"Activities>Activity>Lap>Track>Trackpoint" json:"trackpoint"`
}

type Metric struct {
	Np int32 `json:"np"`
}

func LoadData(id string) (data Activity) {
	raw, _ := ioutil.ReadFile("data/activity_" + id + ".tcx")

	err := xml.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}

	return
}

func (a Activity) CalcMean() UnitData {
	tp := a.Trackpoint

	var hr, cd, wt int32
	var sp float64
	for _, v := range tp {
		hr += v.HeartRate
		cd += v.Cadence
		sp += v.Speed
		wt += v.Watts
	}
	l := int32(len(tp))

	return UnitData{Time: "-", HeartRate: hr / l, Cadence: cd / l, Speed: sp / (float64(l)), Watts: wt / l}
}

func (a Activity) CalcMax() UnitData {
	tp := a.Trackpoint

	var hr, cd, wt int32
	var sp float64
	for _, v := range tp {
		utils.Chmax(&hr, v.HeartRate)
		utils.Chmax(&cd, v.Cadence)
		utils.Chmax(&sp, v.Speed)
		utils.Chmax(&wt, v.Watts)
	}

	return UnitData{Time: "-", HeartRate: hr, Cadence: cd, Speed: sp, Watts: wt}
}

func (a Activity) CalcMetric() Metric {
	tp := a.Trackpoint

	// Calculate NP (Normalized Power)
	np := CalcNP(&tp)

	return Metric{Np: np}
}

/*
	Normalized Powerは以下の方法で求めることができる
	1. 30秒ごとのブロックで平均出力値を求める
	2. 求められた出力値を4乗する
	3. 4乗値を平均し，4乗根をとる
*/
func CalcNP(tp *[]UnitData) int32 {
	n := len(*tp)

	var ret, bs float64
	var tmp, m int32
	for i := 0; i < n; i++ {
		tmp += (*tp)[i].Watts
		bs++

		if (i+1)%30 == 0 || (i+1) == n {
			ret += math.Pow(float64(tmp)/bs, 4)
			tmp, bs = 0, 0
			m++
		}
	}

	ret = math.Pow(ret/float64(m), 0.25)

	return int32(math.Round(ret))
}

/*
	並行処理版
*/
func ParallelCalcNP(tp *[]UnitData, con int) int32 {
	// 30の倍数の長さのブロックで分割
	n := len(*tp)
	blocksize := n / con
	blocksize -= blocksize % 30
	c := make(chan float64, con)

	var start, end int
	for t := 0; t < con; t++ {
		end += blocksize
		if t == con-1 {
			end = n
		}

		// 30秒ごとの平均の4乗値の和をチャンネルに入れてく
		go func(tp *[]UnitData, start, end int) {
			var ret, bs float64
			var tmp int32
			for i := start; i < end; i++ {
				tmp += (*tp)[i].Watts
				bs++

				if (i+1)%30 == 0 || (i+1) == n {
					ret += math.Pow(float64(tmp)/bs, 4)
					tmp, bs = 0, 0
				}
			}
			c <- ret
		}(tp, start, end)

		start += blocksize
	}

	var np float64
	for i := 0; i < con; i++ {
		np += <-c
	}

	ret := math.Pow(np/math.Ceil(float64(n)/30.0), 0.25)

	return int32(math.Round(ret))
}
