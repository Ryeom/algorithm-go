package _3_16

import (
	"algorithm/util"
	"testing"
)

/*
🔥 문제: 최대 수익 구하기 (그리디 알고리즘)
문제 설명
한 회사에서는 하루에 주식을 사고 팔 수 있습니다.
주식의 가격은 매일 다르고, 주식의 가격이 배열로 주어집니다.
주식은 하루에 한 번만 사고 팔 수 있습니다. 단, 주식을 팔기 전에 반드시 사야 하며, 가장 큰 이익을 얻기 위해서는 언제 주식을 사고 팔아야 할지 구해야 합니다.
입력
prices: 주식의 가격을 나타내는 배열 (주식 가격은 항상 양의 정수로 주어짐)
출력
최대 수익을 구하세요.
*/
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
