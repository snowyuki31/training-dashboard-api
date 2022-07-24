package model

type Statistics struct {
	TimeInPowers [2000]int64 `json:"time_in_power"`
}

func GetStatistics() (s Statistics) {
	// Temporarily writing naive codes here
	data := LoadData("activity_9217977872.tcx")

	var tp [2000]int64

	for _, val := range data.Trackpoint {
		tp[val.Watts] += 1
	}

	s.TimeInPowers = tp

	return
}
