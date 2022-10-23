package _0221023

import "time"

func haveConflict(event1 []string, event2 []string) bool {
	ts1 := []int64{}
	ts2 := []int64{}

	loc, _ := time.LoadLocation("Local") //获取当地时区

	ts, err := time.ParseInLocation("2006-01-02 15:04:05", "2022-10-21 "+event1[0]+":00", loc)
	if err != nil {
		panic(err)
	}
	ts1 = append(ts1, ts.Unix())

	ts, err = time.ParseInLocation("2006-01-02 15:04:05", "2022-10-21 "+event1[1]+":00", loc)
	if err != nil {
		panic(err)
	}
	ts1 = append(ts1, ts.Unix())

	ts, err = time.ParseInLocation("2006-01-02 15:04:05", "2022-10-21 "+event2[0]+":00", loc)
	if err != nil {
		panic(err)
	}
	ts2 = append(ts2, ts.Unix())

	ts, err = time.ParseInLocation("2006-01-02 15:04:05", "2022-10-21 "+event2[1]+":00", loc)
	if err != nil {
		panic(err)
	}
	ts2 = append(ts2, ts.Unix())

	if ts1[0] >= ts2[0] && ts1[0] <= ts2[1] {
		return false
	}

	if ts1[1] >= ts2[0] && ts1[1] <= ts2[1] {
		return false
	}

	return true

}
