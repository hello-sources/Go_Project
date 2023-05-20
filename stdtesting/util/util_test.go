package util

import (
	"fmt"
	"testing"
	"time"
)

var commTestData []commStruct

type commStruct struct {
	Group         string
	SizeStr       string
	ExpectSize    int64
	ExpectSizeStr string
}

// 功能测试
func TestParseSize(t *testing.T) {
	testData := commTestData
	for _, data := range testData {
		fmt.Println(data)
		size, sizeStr := ParseSize(data.SizeStr)
		if size != data.ExpectSize || sizeStr != data.ExpectSizeStr {
			t.Errorf("测试结果不符合预期：%+v", data)
		}
	}
}

// 功能测试子测试、并发测试
func TestParseSizeSub(t *testing.T) {
	if testing.Short() {
		t.Skip("跳过测试用例 TestParseSizeSub")
	}

	testData := make(map[string][]commStruct)
	for _, item := range commTestData {
		group := item.Group
		_, ok := testData[group]
		if !ok {
			testData[group] = make([]commStruct, 0)
		}
		testData[group] = append(testData[group], item)
	}

	for k, _ := range testData {
		t.Run("mb", func(t *testing.T) {
			t.Parallel()
			for _, data := range testData[k] {
				size, sizeStr := ParseSize(data.SizeStr)
				if size != data.ExpectSize || sizeStr != data.ExpectSizeStr {
					t.Errorf("测试结果不符合预期：%+v", data)
				}
			}
		})
	}
}

// 模糊测试
func FuzzParseSize(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		size, sizeStr := ParseSize(a)
		if size == 0 || sizeStr == "" {
			t.Errorf("输入异常导致parsesize 没拿到正确结果")
		}
	})
}

// 模糊测试
func FuzzParseSize2(f *testing.F) {
	f.Fuzz(func(t *testing.T, a string) {
		size, sizeStr := ParseSize(a)
		if size == 0 || sizeStr == "" {
			t.Errorf("输入异常导致parsesize 没拿到正确结果")
		}
	})
}

// 基准测试(性能测试)
func BenchmarkParseSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParseSize("1MB")
	}
}

func BenchmarkParseSizeSub(b *testing.B) {
	testData := make(map[string][]commStruct)
	for _, item := range commTestData {
		group := item.Group
		_, ok := testData[group]
		if !ok {
			testData[group] = make([]commStruct, 0)
		}
		testData[group] = append(testData[group], item)
	}
	//b.ResetTimer()
	for k, _ := range testData {
		b.Run(k, func(b *testing.B) {
			preBenchMark()
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				preBenchMark1()
				b.StartTimer()
				ParseSize(testData[k][0].SizeStr)
			}
		})
	}
}

func preBenchMark() {
	time.Sleep(time.Second * 10)
}

func preBenchMark1() {
	time.Sleep(time.Nanosecond * 500)
}

// 测试用例的入口函数
func TestMain(m *testing.M) {
	initCommData()
	m.Run()
}

func initCommData() {
	commTestData = []commStruct{
		{"B", "1b", B, "1B"},
		{"B", "100b", 100 * B, "100B"},
		{"KB", "1kb", KB, "1KB"},
		{"KB", "100KB", 100 * KB, "100KB"},
		{"MB", "1Mb", MB, "1MB"},

		{"GB", "1Gb", GB, "1GB"},
		{"GB", "10Gb", 10 * GB, "10GB"},
		{"TB", "1tb", TB, "1TB"},
		{"PB", "10Pb", 10 * PB, "10PB"},
		{"unknown", "1G", 100 * MB, "100MB"},
	}
}
