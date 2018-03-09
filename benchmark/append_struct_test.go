package main

import (
	"testing"
	"time"
)

func structSlice(count int) {
	arr1 := make([]AvCharge, count)
	for i := 0; i < count; i++ {
		adc := AvCharge{
			ID: 1111111, AvID: 222222, MID: 333333, TagID: 4, SubTagID: 5, IsOriginal: 1, DanmakuCount: 10, CommentCount: 11,
			CollectCount: 22, CoinCount: 11, ShareCount: 11, ElecPlayCount: 11, TotalPlayCount: 11, WebPlayCount: 11, AppPlayCount: 333,
			H5PlayCount: 111, LvUnknown: 100, Lv_1: 11, Lv_2: 22, Lv_3: 33, Lv_4: 44, Lv_5: 55, Lv_6: 66,
			VScore: 33, IncCharge: 100, TotalCharge: 100, IsDeleted: 0, Date: time.Now(), UploadTime: time.Now(),
			CTime: time.Now(), MTime: time.Now(),
		}
		arr1[i] = adc
	}
}

func structSlice1(count int) {
	arr := []AvCharge{}
	for i := 0; i < count; i++ {
		adc := AvCharge{
			ID: 1111111, AvID: 222222, MID: 333333, TagID: 4, SubTagID: 5, IsOriginal: 1, DanmakuCount: 10, CommentCount: 11,
			CollectCount: 22, CoinCount: 11, ShareCount: 11, ElecPlayCount: 11, TotalPlayCount: 11, WebPlayCount: 11, AppPlayCount: 333,
			H5PlayCount: 111, LvUnknown: 100, Lv_1: 11, Lv_2: 22, Lv_3: 33, Lv_4: 44, Lv_5: 55, Lv_6: 66,
			VScore: 33, IncCharge: 100, TotalCharge: 100, IsDeleted: 0, Date: time.Now(), UploadTime: time.Now(),
			CTime: time.Now(), MTime: time.Now(),
		}
		arr = append(arr, adc)
	}
}

func BenchmarkStructSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice(10000000)
	}
}

func BenchmarkStructSlice1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		slice1(10000000)
	}
}
