package timer

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var job = myJob{}

type myJob struct{}

func (job myJob) Run() {
	myFunc()
}

func myFunc() {
	time.Sleep(time.Second)
	fmt.Println("1s...")
}

func TestNewTimerTask(t *testing.T) {
	tm := NewTimerTask()
	_tm := tm.(*timer)

	{
		_, err := tm.AddTaskByFunc("func", "@every 1s", myFunc)
		assert.Nil(t, err)
		_, ok := _tm.taskList["func"]
		if !ok {
			t.Error("no find func")
		}
	}

	{
		_, err := tm.AddTaskByJob("job", "@every 1s", job)
		assert.Nil(t, err)
		_, ok := _tm.taskList["job"]
		if !ok {
			t.Error("no find job")
		}
	}

	{
		_, ok := tm.FindCron("func")
		if !ok {
			t.Error("no find func")
		}
		_, ok = tm.FindCron("job")
		if !ok {
			t.Error("no find job")
		}
		_, ok = tm.FindCron("none")
		if ok {
			t.Error("find none")
		}
	}
	{
		tm.Clear("func")
		_, ok := tm.FindCron("func")
		if ok {
			t.Error("find func")
		}
	}
}
