package model

import "math"

type Metric struct {
	NP  int32   `json:"np"`
	IF  float64 `json:"if"`
	TSS int32   `json:"tss"`
}

func CalcMetric(tp *[]UnitData, FTP int32) (m Metric) {
	NP := CalcNP(tp)
	IF := float64(NP) / float64(FTP)
	TSS := int32(float64(len(*tp)) * float64(NP) * IF * 100.0 / (float64(FTP) * 3600.0))
	m = Metric{NP: NP, IF: IF, TSS: TSS}
	return
}

/*
	Normalized Powerは以下の方法で求めることができる
	1. 30秒ごとのブロックで平均出力値を求める
	2. 求められた出力値を4乗する
	3. 得られた4乗値を平均し，4乗根をとる
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
	NP計算の並行処理版(60*60くらいのサイズのデータでは遅い．．．)
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
