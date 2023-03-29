package cron

import (
	"fmt"
	"github.com/one-go/one-web/models"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func InitCron() {
	// 设置时区
	local, _ := time.LoadLocation("Asia/Shanghai")
	c := cron.New(cron.WithLocation(local), cron.WithSeconds())
	spec := "*/60 * * * * *"
	// 添加一个任务，每 10s 执行一次
	intervalId, err := c.AddFunc(spec, generateMetricHierarchyData)
	if err != nil {
		logrus.Error("cron task error, intervalId:{%d} err:%s", intervalId, err)
		return

	}
	// 开始执行（每个任务会在自己的 goroutine 中执行）
	c.Start()
	//t1 := time.NewTimer(time.Second * 10)
	//for {
	//	select {
	//	case <-t1.C:
	//		t1.Reset(time.Second * 10)
	//	}
	//}
}

func generateMetricHierarchyData() {
	r := rand.Intn(10)
	var serviceName = fmt.Sprintf("%s%d", "one-go-", r)
	data := &models.MetricHierarchy{
		ServiceName:  serviceName,
		DataKind:     r,
		MetricParent: fmt.Sprintf("%s%d", "one_parent_", r),
		MetricChild:  fmt.Sprintf("%s%d", "one_child_", r),
		//HashCode:     strconv.Itoa(r),
		HashCode:     serviceName,
		UseFlag:      true,
		UpdateTime:   time.Now(),
		CreationTime: time.Now(),
	}
	err := models.InsertMetricHierarchy(data)
	if err != nil {
		logrus.Error(err)
		time.Sleep(time.Duration(1) * time.Minute)
	}
	logrus.Infof("generate metric data, time = %s", time.Now().Format("2006-01-02 15:04:05.000 -0700"))
}
