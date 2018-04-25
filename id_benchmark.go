package chinaid

import "testing"

// BenchmarkIDNum 解析基准测试
func BenchmarkIDNum(b *testing.B) {
	testID := "420683199006041237"
	for i := 0; i < b.N; i++ {
		if _, err := IDCard(testID).Decode(); err != nil {
			panic("INVALID_ID_NUM")
		}
	}
}
