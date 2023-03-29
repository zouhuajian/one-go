package models

import (
	"github.com/one-go/one-web/storage"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

type MetricHierarchy struct {
	HierarchyId  int64     `json:"hierarchy_id"`
	ServiceName  string    `json:"service_name"`
	DataKind     int       `json:"data_kind"`
	MetricParent string    `json:"metric_parent"`
	MetricChild  string    `json:"metric_child"`
	HashCode     string    `json:"hash_code"`
	UseFlag      bool      `json:"use_flag"`
	UpdateTime   time.Time `json:"update_time"`
	CreationTime time.Time `json:"creation_time"`
}

type HierarchyParams struct {
	ServiceName string `form:"service_name" binding:"required"`
	DataKind    int    `form:"data_kind" binding:"required"`
}

const (
	MetricHierarchyTable = "monitor_metric_hierarchy"
)

func SelectMetricHierarchyList(params *HierarchyParams) ([]MetricHierarchy, error) {
	var metricHierarchyList []MetricHierarchy
	//result := models.DB.Table(MetricHierarchyTable).Where("service_name=?", params.ServiceName).Find(&metricHierarchyList)
	result := storage.MySQL.Table(MetricHierarchyTable).Where(*params).Find(&metricHierarchyList).Limit(200)
	if result.Error != nil {
		logrus.Errorf("select metric hierarchy:{%#v} err:%s", params, result.Error.Error())
		return []MetricHierarchy{}, result.Error
	}
	return metricHierarchyList, nil
}

func InsertMetricHierarchy(metric *MetricHierarchy) error {
	result := storage.MySQL.Table(MetricHierarchyTable).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "hash_code"}},
		DoUpdates: clause.Assignments(map[string]interface{}{"update_time": metric.UpdateTime}),
	}).Create(metric)
	if result.Error != nil {
		logrus.Errorf("insert metric:{%#v} err:%s", metric, result.Error.Error())
		return result.Error
	}

	return nil
}
