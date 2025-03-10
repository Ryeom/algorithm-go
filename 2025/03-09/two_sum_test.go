package _3_09

import (
	"algorithm/util"
	"reflect"
	"testing"
)

/*
📌 문제 1: 배열에서 두 수의 합 찾기 (Two Sum)
난이도: 초급
문제 유형: #해시맵 #배열 #투포인터
권장 풀이 시간: ⏳ 20분
출처: 직접 만든 문제

📝 문제 설명
정수 배열 nums와 정수 target이 주어졌을 때, 배열 내에서 두 개의 수를 더해서 target이 되는 인덱스를 찾는 문제이다.
같은 원소를 두 번 사용할 수 없으며, 반드시 하나의 정답만 존재한다고 가정한다.

📥 입력
nums: 정수 배열 (len(nums) ≥ 2)
target: 정수 (target을 만족하는 두 수가 반드시 존재함)
📤 출력
두 수의 인덱스를 담은 []int 슬라이스를 반환 (오름차순 정렬 필요 없음)
*/

// 브루트포스(완전 탐색) 방식 -> 비효율적! 두 개의 for 문을 사용하면 O(N²) 입력 크기가 커질수록 실행시간이 제곱만큼 증가

func twoSum(nums []int, target int) []int {
	// TODO: 여기에 코드를 작성하세요.

	var index []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				index = append(index, i, j)

				return index
			}
		}
	}

	return nil
}

//해시맵(map) 사용 : 시간 복잡도는 O(N) 입력 크기가 커져도 실행 시간이 거의 선형적으로 증가

func twoSumBetter(nums []int, target int) []int {
	numMap := make(map[int]int)
	for i, num := range nums {
		m := target - num
		if idx, ok := numMap[m]; ok {
			return []int{idx, i}
		}
		numMap[num] = i
	}
	return nil
}

/*
🔹 해시맵(O(N))이 느릴 수도 있는 경우
1️⃣ 정렬이 필요할 때

해시맵은 내부적으로 순서를 보장하지 않음
예: "가장 작은 값을 찾거나 정렬해야 하는 경우"
대신 정렬이 필요하면 정렬된 배열을 쓰는 게 유리 (예: 이분 탐색 O(logN))
2️⃣ 메모리 사용이 중요한 경우

해시맵은 배열보다 추가적인 메모리를 사용함
예: "배열을 그냥 정렬하고 탐색하는 방법"이 메모리를 적게 씀
map을 사용하면 추가적인 공간 O(N)이 필요함 (배열은 O(1) 추가 공간)
3️⃣ 충돌(Hash Collision) 발생 시 느려질 수 있음

해시맵 내부에서 해시 충돌이 많아지면 O(1) 대신 최악의 경우 O(N)이 될 수도 있음
하지만 Go의 map은 충돌을 잘 해결하는 구조라서 심각한 문제는 거의 없음
*/

// 테스트 코드
func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("TwoSum", twoSum, tt.nums, tt.target).([]int)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("fail 함수명(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		} else {
			t.Logf("sucess~! got = %v", got)
		}
	}
}
