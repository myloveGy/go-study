package bufio

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Id        int
	Name      string
	Status    int
	User      string
	StartTime time.Time
	EndTime   time.Time
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func (t *Task) String() string {
	return fmt.Sprintf(
		"%d,%s,%d,%s,%s,%s",
		t.Id,
		t.Name,
		t.Status,
		t.User,
		TimeFormat(t.StartTime),
		TimeFormat(t.EndTime),
	)
}

func StringToTask(line string) (*Task, error) {
	// 拆分字符串
	arr := strings.Split(line, ",")
	if len(arr) != 6 {
		return nil, errors.New("数据存在问题")
	}

	// 一步一步
	// 解析id
	id, err := strconv.Atoi(arr[0])
	if err != nil {
		return nil, err
	}

	// 解析状态
	status, err := strconv.Atoi(arr[2])
	if err != nil {
		return nil, err
	}

	// 解析时间
	time1, err := time.Parse("2006-01-02 15:04:05", arr[4])
	if err != nil {
		return nil, err
	}

	time2, err := time.Parse("2006-01-02 15:04:05", arr[5])
	if err != nil {
		return nil, err
	}

	return &Task{
		Id:        id,
		Name:      arr[1],
		Status:    status,
		User:      arr[3],
		StartTime: time1,
		EndTime:   time2,
	}, nil
}

func ReadFileToTask(filename string) ([]*Task, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	taskList := make([]*Task, 0)
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		if task, err := StringToTask(reader.Text()); err == nil {
			taskList = append(taskList, task)
		}
	}

	return taskList, nil
}

func WriteFileToTask(filename string, taskList []*Task) error {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
	if err != nil {
		return err
	}

	defer f.Close()

	for _, v := range taskList {
		if _, err := f.WriteString(v.String() + "\n"); err != nil {
			return err
		}
	}

	return nil
}
