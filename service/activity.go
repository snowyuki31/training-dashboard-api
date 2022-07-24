package service

import (
	"encoding/xml"
	"io/ioutil"
	"snowyuki31/training-dashboard-api/model"
)

type ActivityService struct{}

func (ActivityService) LoadData(id string) model.Activity {
	var data model.Activity
	raw, _ := ioutil.ReadFile("data/activity_" + id + ".tcx")

	err := xml.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}

	return data
}
