package main

import (
	"testing"
	"time"
)

type AvCharge struct {
	ID             int64
	AvID           int64
	MID            int64
	TagID          int64
	SubTagID       int64
	IsOriginal     int
	DanmakuCount   int64
	CommentCount   int64
	CollectCount   int64
	CoinCount      int64
	ShareCount     int64
	ElecPlayCount  int64
	TotalPlayCount int64
	WebPlayCount   int64
	AppPlayCount   int64
	H5PlayCount    int64
	LvUnknown      int64
	Lv_1           int64
	Lv_2           int64
	Lv_3           int64
	Lv_4           int64
	Lv_5           int64
	Lv_6           int64
	VScore         int64
	IncCharge      int64
	TotalCharge    int64
	IsDeleted      int
	Date           time.Time
	UploadTime     time.Time
	CTime          time.Time
	MTime          time.Time
}

func pointS(count int) {
	a := []*AvCharge{}
	for i := 0; i < 150; i++ {
		a1 := pointSlice(count)
		a = append(a, a1...)
	}
}

func pointS1(count int) {
	a := []*AvCharge{}
	for i := 0; i < 150; i++ {
		a1 := pointSlice1(count)
		a = append(a, a1...)
	}
}

func pointSlice(count int) []*AvCharge {
	arr1 := make([]*AvCharge, count)
	var adc *AvCharge
	for i := 0; i < count; i++ {
		adc = &AvCharge{
			ID: 1111111, AvID: 222222, MID: 333333, TagID: 4, SubTagID: 5, IsOriginal: 1, DanmakuCount: 10, CommentCount: 11,
		}
		arr1[i] = adc
	}
	return arr1
}

func pointSlice1(count int) []*AvCharge {
	arr := []*AvCharge{}
	for i := 0; i < count; i++ {
		adc := &AvCharge{
			ID: 1111111, AvID: 222222, MID: 333333, TagID: 4, SubTagID: 5, IsOriginal: 1, DanmakuCount: 10, CommentCount: 11,
		}
		arr = append(arr, adc)
	}
	return arr
}

func BenchmarkPointSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointS(2000)
	}
}

func BenchmarkPointSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pointS1(2000)
	}
}
