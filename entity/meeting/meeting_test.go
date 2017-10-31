package meeting

import (
	"testing"
	"time"
)

func TestCreateMeeting(t *testing.T) {
	var parti []string
	parti = append(parti, "aa")
	parti = append(parti, "bb")
	t1, _ := time.Parse("2006-01-02 15:04:05", "2017-10-29 07:37:18")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2017-10-29 08:37:28")
	err := CreateMeeting("meet1", parti, t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}

func TestQueryMeeting(t *testing.T) {
	t1, _ := time.Parse("2006-01-02 15:04:05", "2017-10-29 07:37:18")
	t2, _ := time.Parse("2006-01-02 15:04:05", "2017-10-29 08:37:28")
	err := QueryMeeting(t1, t2)
	if err != nil {
		t.Errorf("error:%s", err)
	}
}
