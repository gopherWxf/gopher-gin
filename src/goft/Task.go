package goft

import (
	"github.com/robfig/cron/v3"
	"sync"
)

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor
var once sync.Once

var taskCron *cron.Cron
var onceCron sync.Once

func init() {
	chlist := getTaskList()
	go func() {
		for t := range chlist {
			doTask(t)
		}
	}()
}
func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}
func getCronTask() *cron.Cron {
	onceCron.Do(func() {
		taskCron = cron.New(cron.WithSeconds())
	})
	return taskCron
}
func doTask(t *TaskExecutor) {
	go func() {
		defer func() {
			if t.callback != nil {
				t.callback()
			}
		}()
		t.Exec()
	}()
}

type TaskExecutor struct {
	f        TaskFunc
	params   []interface{}
	callback func()
}

func NewTaskExecutor(f TaskFunc, params []interface{}, callback func()) *TaskExecutor {
	return &TaskExecutor{f: f, params: params, callback: callback}
}

func (this *TaskExecutor) Exec() {
	this.f(this.params...)
}
func Task(f TaskFunc, callback func(), params ...interface{}) {
	if f == nil {
		return
	}
	go func() {
		getTaskList() <- NewTaskExecutor(f, params, callback)
	}()
}
