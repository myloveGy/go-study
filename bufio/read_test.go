package bufio

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTimeFormat(t *testing.T) {
	now := time.Now()
	fmt.Println(TimeFormat(now))
	assert.Equal(t, now.Format("2006-01-02 15:04:05"), TimeFormat(now))
}

func TestStringToTask(t *testing.T) {
	task, err := StringToTask("")
	assert.Error(t, err)
	assert.Nil(t, task)

	task, err = StringToTask("1,name,2,jinxing.liu,2021-03-18 11:20:30,2021-03-18 11:20:30")
	assert.NoError(t, err)
	assert.Equal(t, 1, task.Id)
}

func TestWriteFileToTask(t *testing.T) {
	now := time.Now()
	end := now.Add(time.Hour * 24)
	err := WriteFileToTask("./format.txt", []*Task{
		{Id: 1, Name: "name", Status: 1, User: "jinxing.liu1", StartTime: now, EndTime: end},
		{Id: 2, Name: "name", Status: 1, User: "jinxing.liu2", StartTime: now, EndTime: end},
		{Id: 3, Name: "name", Status: 1, User: "jinxing.liu3", StartTime: now, EndTime: end},
	})

	assert.NoError(t, err)
	os.Remove("./format.txt")
}

func TestReadFileToTask(t *testing.T) {
	now := time.Now()
	end := now.Add(time.Hour * 24)

	err := WriteFileToTask("./format.txt", []*Task{
		{Id: 1, Name: "name", Status: 1, User: "jinxing.liu1", StartTime: now, EndTime: end},
		{Id: 2, Name: "name", Status: 1, User: "jinxing.liu2", StartTime: now, EndTime: end},
		{Id: 3, Name: "name", Status: 1, User: "jinxing.liu3", StartTime: now, EndTime: end},
	})

	assert.NoError(t, err)
	taskList, err := ReadFileToTask("./format.txt")
	assert.NoError(t, err)
	assert.Equal(t, 3, len(taskList))
	os.Remove("./format.txt")
}

