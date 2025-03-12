package _3_13

import (
	"algorithm/util"
	"testing"
)

// 거스름돈을 최소 화폐 단위로 거슬러주는 함수
func coinChange(coins []int, amount int) int {
	// 금액이 0일 경우에는 0 반환
	if amount == 0 {
		return 0
	}

	count := 0
	// 각 화폐 단위로 금액을 거슬러줌

	for _, coin := range coins {
		// 해당 화폐 단위로 거슬러줄 수 있는 개수 계산
		n := amount / coin
		// 해당 화폐 단위로 거슬러준 금액만큼 차감
		amount -= n * coin
		// 사용된 화폐 개수만큼 증가
		count += n

		// 금액이 0이면 더 이상 계산할 필요 없음
		if amount == 0 {
			break
		}
	}

	// 금액이 정확히 거슬러지지 않으면 -1 반환
	if amount > 0 {
		return -1
	}

	return count
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		index  int
		coins  []int
		amount int
		want   int
	}{
		{
			index:  1,
			coins:  []int{500, 100, 50, 10, 5, 1},
			amount: 1260,
			want:   6, // 500*2, 100*1, 50*1, 10*1, 5*1
		},
		{
			index:  2,
			coins:  []int{500, 100, 50, 10, 5, 1},
			amount: 130,
			want:   4, // 100*1, 10*3
		},
		{
			index:  3,
			coins:  []int{25, 10, 5, 1},
			amount: 63,
			want:   6, // 25*2, 10*1, 5*1, 1*3
		},
		{
			index:  4,
			coins:  []int{25, 10, 5, 1},
			amount: 11,
			want:   2, // 10*1, 1*1
		},
		{
			index:  5,
			coins:  []int{25, 10, 5, 1},
			amount: 30,
			want:   2, // 25*1, 5*1
		},
		{
			index:  6,
			coins:  []int{5, 3, 2},
			amount: 11,
			want:   -1, // 5*2, 2*1
		},
		{
			index:  7,
			coins:  []int{5, 2, 1},
			amount: 7,
			want:   2, // 5*1, 2*1
		},
		{
			index:  8,
			coins:  []int{25, 10, 5, 1},
			amount: 99,
			want:   9, // 25*3, 10*2, 5*0, 1*4
		},
		{
			index:  9,
			coins:  []int{3, 2, 1},
			amount: 4,
			want:   2, // 3*1, 1*1
		},
		{
			index:  10,
			coins:  []int{5, 2, 1},
			amount: 11,
			want:   3, // 5*2, 1*1
		},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("CoinChange", coinChange, tt.coins, tt.amount).(int)
		if got != tt.want {
			t.Errorf("❌ 실패: 테스트 %d - coinChange(%v, %d) = %v; want %v", tt.index, tt.coins, tt.amount, got, tt.want)
		} else {
			t.Logf("✅ 성공: 테스트 %d - got = %v", tt.index, got)
		}
	}
}
