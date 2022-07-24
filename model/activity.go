package model

type Activity struct {
	Trackpoint []struct {
		Time      string  `xml:"Time" json:"time"`
		HeartRate int64   `xml:"HeartRateBpm>Value" json:"heart_rate"`
		Cadence   int64   `xml:"Cadence" json:"cadence"`
		Speed     float64 `xml:"Extensions>TPX>Speed" json:"speed"`
		Watts     int64   `xml:"Extensions>TPX>Watts" json:"watts"`
	} `xml:"Activities>Activity>Lap>Track>Trackpoint" json:"trackpoint"`
}
