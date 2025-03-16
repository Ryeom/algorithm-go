package _3_16

import (
	"algorithm/util"
	"testing"
)

// 최대 수익을 구하는 함수
func maxProfit(prices []int) int {
	// TODO: 여기에 코드를 작성하세요.
	return 0 // 임시값
}

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		index  int
		prices []int
		want   int
	}{
		{
			index:  1,
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5, // 1에 사고 6에 팔아서 5의 이익
		},
		{
			index:  2,
			prices: []int{7, 6, 4, 3, 1},
			want:   0, // 계속 하락하므로 이익이 없음
		},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("MaxProfit", maxProfit, tt.prices).(int)
		if got != tt.want {
			t.Errorf("❌ 실패: 테스트 %d - maxProfit(%v) = %v; want %v", tt.index, tt.prices, got, tt.want)
		} else {
			t.Logf("✅ 성공: 테스트 %d - got = %v", tt.index, got)
		}
	}
}
