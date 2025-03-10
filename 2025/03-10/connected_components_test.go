package _3_10

import (
	"algorithm/util"
	"reflect"
	"testing"
)

/*
🔥 문제: 연결된 영역(Connected Components) 개수 구하기
난이도: 중급
문제 유형: #DFS #그래프 탐색
권장 풀이 시간: ⏳ 30~40분

📝 문제 설명
n x m 크기의 2D 격자가 주어진다.
각 칸에는 0(빈 칸) 또는 1(채워진 칸)이 존재한다.
상하좌우로 연결된 1들을 하나의 영역(Connected Component)로 간주한다.
전체 영역(Connected Component)의 개수를 반환하는 함수를 작성하라.

*/

// 영역의 개수를 찾는 함수 (직접 구현)
func countConnectedComponents(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}
	var dfs func(i, j int)
	n, m := len(grid), len(grid[0])
	dfs = func(i, j int) {
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] == 0 {
			return
		}

		grid[i][j] = 0
		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)
	}

	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}

// 테스트 코드
func TestCountConnectedComponents(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{1, 1, 0, 0, 1},
				{1, 0, 0, 1, 1},
				{0, 0, 1, 0, 0},
				{1, 1, 0, 0, 1},
			},
			5, // 정답: 5개의 영역이 존재
		},
		{
			[][]int{
				{1, 1, 0},
				{0, 1, 1},
				{1, 0, 0},
			},
			2, // 정답: 2개의 영역이 존재
		},
		{
			[][]int{
				{1, 0, 1, 0, 1},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 0, 1},
			},
			6, // 정답: 6개의 영역이 존재
		},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("ConnectedComponents", countConnectedComponents, tt.grid).(int)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("❌ 실패: countConnectedComponents(%v) = %v; want %v", tt.grid, got, tt.want)
		} else {
			t.Logf("✅ 성공: got = %v", got)
		}
	}
}
