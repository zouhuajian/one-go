package models

import (
	log "github.com/sirupsen/logrus"
	"one-go/storage"
	"time"
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
	ServiceName string `form:"service_name" binding: "require"`
	DataKind    int    `form:"data_kind"`
}

const (
	MetricHierarchyTable = "monitor_metric_hierarchy"
)

func SelectMetricHierarchyList(params *HierarchyParams) ([]MetricHierarchy, error) {
	var metricHierarchyList []MetricHierarchy
	//result := models.DB.Table(MetricHierarchyTable).Where("service_name=?", params.ServiceName).Find(&metricHierarchyList)
	result := storage.MySQL.Table(MetricHierarchyTable).Where(*params).Find(&metricHierarchyList)
	log.Printf("row = %d, error = %#v\n", result.RowsAffected, result.Error)
	return metricHierarchyList, nil
}
