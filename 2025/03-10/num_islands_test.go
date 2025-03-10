package _3_10

import (
	"algorithm/util"
	"reflect"
	"testing"
)

// 섬의 개수를 찾는 함수 (직접 구현)
func numIslandsDFS(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	n, m := len(grid), len(grid[0])
	count := 0

	// DFS 함수 정의
	var dfs func(i, j int)
	dfs = func(i, j int) {
		// 범위를 벗어나거나, 이미 방문한 곳(0)이라면 종료
		if i < 0 || j < 0 || i >= n || j >= m || grid[i][j] != 1 {
			return
		}

		// 방문한 육지를 '0'으로 변경하여 중복 탐색 방지
		grid[i][j] = 0

		// 상하좌우 탐색
		dfs(i+1, j) // 아래
		dfs(i-1, j) // 위
		dfs(i, j+1) // 오른쪽
		dfs(i, j-1) // 왼쪽
	}

	// 그리드를 순회하며 섬(1) 찾기
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 { // 새로운 섬 발견
				count++
				dfs(i, j) // DFS 실행
			}
		}
	}

	return count
}

func numIslandsBFS(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	n, m := len(grid), len(grid[0])
	count := 0
	directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 {
				count++
				queue := [][]int{{i, j}}
				grid[i][j] = 0 // 방문 처리

				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:] // Dequeue

					// 상하좌우 탐색
					for _, d := range directions {
						ni, nj := cur[0]+d[0], cur[1]+d[1]
						if ni >= 0 && nj >= 0 && ni < n && nj < m && grid[ni][nj] == 1 {
							grid[ni][nj] = 0                     // 방문 처리
							queue = append(queue, []int{ni, nj}) // Enqueue
						}
					}
				}
			}
		}
	}

	return count
}

// 테스트 코드
func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 1, 0, 0},
				{0, 0, 0, 1, 1},
			},
			3, // 정답: 3개의 섬이 존재
		},
		{
			[][]int{
				{1, 0, 1, 0, 1},
				{0, 0, 0, 1, 0},
				{1, 1, 0, 0, 1},
			},
			6, // 정답: 4개의 섬이 존재
		},
		{
			[][]int{},
			0, // 빈 배열 테스트
		},
		{
			[][]int{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			0, // 전부 물(0)일 경우 섬 없음
		},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("NumIslandsDFS", numIslandsDFS, tt.grid).(int)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("❌ 실패: numIslandsDFS(%v) = %v; want %v", tt.grid, got, tt.want)
		} else {
			t.Logf("✅ 성공: got = %v", got)
		}
	}

	for _, tt := range tests {
		got := util.RunWithTimer("NumIslandsDFS", numIslandsBFS, tt.grid).(int)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("❌ 실패: numIslandsDFS(%v) = %v; want %v", tt.grid, got, tt.want)
		} else {
			t.Logf("✅ 성공: got = %v", got)
		}
	}
}
