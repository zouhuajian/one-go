package controller

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"one-go/internal"
	"one-go/models"
	"one-go/service"
)

// GetMetricHierarchy
// @Router /api/v1/metrics/hierarchy
func GetMetricHierarchy(c *gin.Context) {
	var params models.HierarchyParams
	if err := c.BindQuery(&params); err != nil {
		internal.APIResponse(c, err, nil)
		return
	}
	//params.ServiceName, _ = c.GetQuery("service_name")
	log.Printf("metric hierarchy params: %#v", params)
	metricHierarchyList, err := service.GetMetricHierarchyList(&params)
	if err != nil {
		internal.APIResponse(c, err, nil)
		return
	}
	internal.APIResponse(c, nil, metricHierarchyList)
}
