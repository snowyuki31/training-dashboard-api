package model

import (
	"encoding/xml"
	"io/ioutil"
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

	return Metric{NP: np}
}
