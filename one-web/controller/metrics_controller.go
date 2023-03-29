package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/one-go/one-web/internal"
	"github.com/one-go/one-web/models"
	"github.com/one-go/one-web/service"
	"github.com/sirupsen/logrus"
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
	logrus.Printf("metric hierarchy params: %#v", params)
	metricHierarchyList, err := service.GetMetricHierarchyList(&params)
	if err != nil {
		internal.APIResponse(c, err, nil)
		logrus.Errorf("get metric hierarchy error, cause = %v", err)
		return
	}
	internal.APIResponse(c, nil, metricHierarchyList)
}
