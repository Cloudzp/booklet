package main_test

import (
	"testing"
	"time"
)

func TestTimeInt642String(t *testing.T) {
	//t, _ := time.Parse("2006-01-02 15:04:05","2019-05-01 15:59:59")

	time := time.Unix(1561331734, 0).Format("2006-01-02 15:04:05")
	t.Log(time)
}
