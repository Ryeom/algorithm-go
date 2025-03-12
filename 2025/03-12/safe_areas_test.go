package _3_12

import (
	"algorithm/util"
	"reflect"
	"testing"
)

// DFS로 안전 영역 개수 찾기 (직접 구현)
func countSafeAreas(grid [][]int, waterLevel int) int {
	if len(grid) == 0 {
		return 0
	}
	n, m := len(grid), len(grid[0])
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, m)
	}

	var dfs func(i, j int)
	dfs = func(i, j int) {
		if i < 0 || i >= n || j < 0 || j >= m {
			return
		}
		if visited[i][j] {
			return
		}
		if grid[i][j] <= waterLevel {
			return
		}
		visited[i][j] = true

		dfs(i+1, j)
		dfs(i-1, j)
		dfs(i, j+1)
		dfs(i, j-1)

		return
	}
	count := 0
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] > waterLevel && !visited[i][j] {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}

// 테스트 코드
func TestCountSafeAreas(t *testing.T) {
	tests := []struct {
		grid       [][]int
		waterLevel int
		want       int
	}{
		{
			[][]int{
				{5, 3, 7, 2},
				{6, 2, 4, 1},
				{7, 3, 5, 6},
			},
			3,
			2,
		},
		{
			[][]int{
				{1, 2, 3, 4, 5},
				{5, 4, 3, 2, 1},
				{1, 2, 3, 4, 5},
			},
			2,
			1,
		},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("CountSafeAreas", countSafeAreas, tt.grid, tt.waterLevel).(int)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("❌ 실패: countSafeAreas(%v, %d) = %v; want %v", tt.grid, tt.waterLevel, got, tt.want)
		} else {
			t.Logf("✅ 성공: got = %v", got)
		}
	}
}
