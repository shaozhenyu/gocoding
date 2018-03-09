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

func pointSlice(count int) {
	arr1 := make([]*AvCharge, count)
	for i := 0; i < count; i++ {
		adc := &AvCharge{
			ID: 1111111, AvID: 222222, MID: 333333, TagID: 4, SubTagID: 5, IsOriginal: 1, DanmakuCount: 10, CommentCount: 11,
			CollectCount: 22, CoinCount: 11, ShareCount: 11, ElecPlayCount: 11, TotalPlayCount: 11, WebPlayCount: 11, AppPlayCount: 333,
			H5PlayCount: 111, LvUnknown: 100, Lv_1: 11, Lv_2: 22, Lv_3: 33, Lv_4: 44, Lv_5: 55, Lv_6: 66,
			VScore: 33, IncCharge: 100, TotalCharge: 100, IsDeleted: 0, Date: time.Now(), UploadTime: time.Now(),
			CTime: time.Now(), MTime: time.Now(),
		}
		arr1[i] = adc
	}
}

func pointSlice1(count int) {
	arr := []*AvCharge{}
	for i := 0; i < count; i++ {
		adc := &AvCharge{
			ID: 1111111, AvID: 222222, MID: 333333, TagID: 4, SubTagID: 5, IsOriginal: 1, DanmakuCount: 10, CommentCount: 11,
			CollectCount: 22, CoinCount: 11, ShareCount: 11, ElecPlayCount: 11, TotalPlayCount: 11, WebPlayCount: 11, AppPlayCount: 333,
			H5PlayCount: 111, LvUnknown: 100, Lv_1: 11, Lv_2: 22, Lv_3: 33, Lv_4: 44, Lv_5: 55, Lv_6: 66,
			VScore: 33, IncCharge: 100, TotalCharge: 100, IsDeleted: 0, Date: time.Now(), UploadTime: time.Now(),
			CTime: time.Now(), MTime: time.Now(),
		}
		arr = append(arr, adc)
	}
}

func BenchmarkPointSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice(10000000)
	}
}

func BenchmarkPointSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1(10000000)
	}
}
