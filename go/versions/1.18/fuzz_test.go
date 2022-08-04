package __18

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func Reverse(s string) string {
	if !utf8.ValidString(s) {
		return s
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// go test -v -run TestReverse fuzz_test.go
func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{"Hello, world", "dlrow ,olleH"},
		{" ", " "},
		{"!12345", "54321!"},
	}
	for _, tc := range testcases {
		rev := Reverse(tc.in)
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) { // go test -fuzz=Fuzz
	// go test -fuzz=Fuzz -fuzztime 30s
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Use f.Add to provide a seed corpus
	}
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
		if utf8.ValidString(orig) && !utf8.ValidString(rev) {
			t.Errorf("Reverse produced invalid UTF-8 string %q", rev)
		}
	})
}

// go test -v -benchtime=5s -benchmem -bench Benchmark_Add fuzz_test.go
// ns/op 表示每一个操作耗费多少时间（纳秒）
func Benchmark_Add(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}

// go test -v -benchtime=2s -benchmem -bench Benchmark_Alloc fuzz_test.go
// B/op 每次调用分配多少byte
// allocs/op 表示每一次调用有几次分配
func Benchmark_Alloc(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("%d", i)
		m[i] = struct{}{}
	}
}

// go test -v -benchtime=5s -benchmem -bench Benchmark_Add_TimerControl fuzz_test.go
func Benchmark_Add_TimerControl(b *testing.B) {
	// 重置计时器
	b.ResetTimer()
	// 停止计时器
	b.StopTimer()
	// 开始计时器
	b.StartTimer()
	var n int
	for i := 0; i < b.N; i++ {
		n++
	}
}
