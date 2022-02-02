package service

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/sirupsen/logrus"
	"one-go/models"
	"one-go/storage"
)

func GetMetricHierarchyList(params *models.HierarchyParams) ([]models.MetricHierarchy, error) {
	var val []models.MetricHierarchy
	key := fmt.Sprintf("%s-%d", params.ServiceName, params.DataKind)
	ctx := context.Background()
	v, err := storage.Redis.Get(ctx, key).Bytes()
	if err != nil {
		logrus.Errorf("get data from redis error, key: %s, err: %#v", key, err)
	}
	if v == nil {
		val, err = models.SelectMetricHierarchyList(params)
		var buffer bytes.Buffer
		ecoder := gob.NewEncoder(&buffer)
		err = ecoder.Encode(val)
		if err != nil {
			logrus.Errorf("encode data error, key: %s, err: %#v", key, err)
		}
		err = storage.Redis.Set(ctx, key, buffer.Bytes(), 0).Err()
		if err != nil {
			logrus.Errorf("failed to write data to redis error, key: %s, err: %#v", key, err)
		}
		logrus.Info("data from mysql")
	} else {
		reader := bytes.NewReader(v)
		dec := gob.NewDecoder(reader)
		err = dec.Decode(&val)
		if err != nil {
			logrus.Errorf("decode data error, key: %s, err: %#v", key, err)
		}
		logrus.Info("data from redis")
	}
	return val, nil
}
