package service

import (
	"snowyuki31/training-dashboard-api/model"
)

type StatisticsService struct{}

func (StatisticsService) GetStatistics() model.Statistics {
	// Temporarily writing naive codes here
	data := ActivityService{}.LoadData()

	var tp [2000]int64

	for _, val := range data.Trackpoint {
		tp[val.Watts] += 1
	}

	return model.Statistics{TimeInPowers: tp}
}
