package service

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"one-go/models"
	"one-go/storage"
	"sync"
	"time"
)

type task struct {
	key string
	val []models.MetricHierarchy
}

var taskChan chan task

// use a WaitGroup
var wg sync.WaitGroup

func GetMetricHierarchyList(params *models.HierarchyParams) ([]models.MetricHierarchy, error) {
	var val []models.MetricHierarchy
	key := fmt.Sprintf("%s-%d", params.ServiceName, params.DataKind)
	v, err := storage.Redis.Get(context.Background(), key).Bytes()
	if err != nil && err != redis.Nil {
		logrus.Errorf("get data from redis error, key: %s, err: %#v", key, err)
	}
	if v == nil {
		val, err = models.SelectMetricHierarchyList(params)
		if len(val) == 0 {
			logrus.Info("data from mysql")
			return val, nil
		}
		t := task{
			key: key,
			val: val,
		}
		if !TryEnqueue(t) {
			logrus.Warnf("enqueue error, key:%s", key)
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

func init() {
	taskChan = make(chan task, 10)
	// pass the channel to the workers, let them wait on it
	for i := 0; i < 10; i++ {
		// increment the WaitGroup before starting the worker
		wg.Add(1)
		go consumeTask(i)
	}
}

// TryEnqueue tries to enqueue a job to the given job channel. Returns true if
// the operation was successful, and false if enqueuing would not have been
// possible without blocking. Job is not enqueued in the latter case.
// 非阻塞入队
//如果想尝试入队，在需要阻塞的时候返回 fail 怎么办？这种方式能够获取提交任务的失败状态返回 503。关键在于使用 select 的 default 语句：
func TryEnqueue(t task) bool {
	select {
	case taskChan <- t:
		return true
	default:
		return false
	}
}

func consumeTask(id int) {
	defer wg.Done()
	for t := range taskChan {
		logrus.Infof("process task, id = %d", id)
		var buffer bytes.Buffer
		ecoder := gob.NewEncoder(&buffer)
		err := ecoder.Encode(t.val)
		if err != nil {
			logrus.Errorf("encode data error, key: %s, err: %#v", t.key, err)
		}
		err = storage.Redis.Set(context.Background(), t.key, buffer.Bytes(), 0).Err()
		if err != nil {
			logrus.Errorf("failed to write data to redis error, key: %s, err: %#v", t.key, err)
		}
	}
	logrus.Info("close task channel...")
	// now use the WaitTimeout instead of wg.Wait()
	WaitTimeout(&wg, 5*time.Second)
}

// WaitTimeout does a Wait on a sync.WaitGroup object but with a specified
// timeout. Returns true if the wait completed without timing out, false
// otherwise.
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	go func() {
		wg.Wait()
		close(taskChan)
	}()
	select {
	case <-taskChan:
		return true
	case <-time.After(timeout):
		return false
	}
}
