package goft

import "sync"

type TaskFunc func(params ...interface{})

var taskList chan *TaskExecutor
var once sync.Once

func init() {
	chlist := getTaskList()
	go func() {
		for t := range chlist {
			t.Exec()
		}
	}()
}

type TaskExecutor struct {
	f      TaskFunc
	params []interface{}
}

func getTaskList() chan *TaskExecutor {
	once.Do(func() {
		taskList = make(chan *TaskExecutor)
	})
	return taskList
}

func NewTaskExecutor(f TaskFunc, params []interface{}) *TaskExecutor {
	return &TaskExecutor{f: f, params: params}
}
func (this *TaskExecutor) Exec() {
	this.f(this.params...)
}
func Task(f TaskFunc, params ...interface{}) {
	getTaskList() <- NewTaskExecutor(f, params)
}
