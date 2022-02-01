package service

import "one-go/models"

func GetMetricHierarchyList(params *models.HierarchyParams) ([]models.MetricHierarchy, error) {
	return models.SelectMetricHierarchyList(params)
}
