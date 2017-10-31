package meeting

import (
	"fmt"
	"testing"
	"time"
)

var t1 = time.Now()
var t2 = t1.Add(time.Second)
var meetName = t1.Format("2006-01-02 15:04:05")

func TestCreateMeeting(t *testing.T) {
	//测试添加会议
	fmt.Println(t1, t2)
	var parti []string
	parti = append(parti, "aa")
	parti = append(parti, "bb")
	err := CreateMeeting(meetName, parti, t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestQueryMeeting(t *testing.T) {
	err := QueryMeeting(t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}
