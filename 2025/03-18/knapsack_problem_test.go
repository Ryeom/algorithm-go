package _3_18

import (
	"algorithm/util"
	"testing"
)

// 배낭 문제를 풀기 위한 함수
func knapsack(weights []int, values []int, capacity int) int {
	// TODO: 여기에 코드를 작성하세요.
	return 0 // 임시값
}
func TestKnapsack(t *testing.T) {
	tests := []struct {
		index    int
		weights  []int
		values   []int
		capacity int
		want     int
	}{
		{
			index:    1,
			weights:  []int{10, 20, 30},
			values:   []int{60, 100, 120},
			capacity: 50,
			want:     220, // 20kg과 30kg 아이템을 선택하여 220의 가치 얻음
		},
		{
			index:    2,
			weights:  []int{50, 60, 70},
			values:   []int{60, 100, 120},
			capacity: 50,
			want:     60, // 50kg짜리 아이템만 선택
		},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("Knapsack", knapsack, tt.weights, tt.values, tt.capacity).(int)
		if got != tt.want {
			t.Errorf("❌ 실패: 테스트 %d - knapsack(%v, %v, %d) = %v; want %v", tt.index, tt.weights, tt.values, tt.capacity, got, tt.want)
		} else {
			t.Logf("✅ 성공: 테스트 %d - got = %v", tt.index, got)
		}
	}
}
