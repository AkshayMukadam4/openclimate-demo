package oracle

import (

	"github.com/robfig/cron"
	"log"
	"github.com/pkg/errors"

)

func GetAndCommitDaily() {

	dailyNoaaData, err := GetNoaaDailyCO2()
	if err != nil {
		log.Fatal(errors.Wrap(err, "GetAndCommitDaily() failed"))
	}

	err = Verify("Earth", "Earth", 1, dailyNoaaData)
	if err != nil {
		log.Fatal(errors.Wrap(err, "GetAndCommitDaily() failed"))
	}
}

func ScheduleNoaaCO2() {

	c := cron.New()
	c.AddFunc("@daily", GetAndCommitDaily)
	// c.AddFunc("@monthly", GetNoaaMonthlyCO2)

	c.Start()
}