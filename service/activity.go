package service

import (
	"encoding/xml"
	"io/ioutil"
	"snowyuki31/training-dashboard-api/model"
)

type ActivityService struct{}

func (ActivityService) LoadData() model.Activity {
	var data model.Activity
	raw, _ := ioutil.ReadFile(`data/activity_9246957399.tcx`)

	err := xml.Unmarshal(raw, &data)
	if err != nil {
		panic(err)
	}

	return data
}
