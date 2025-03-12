package _3_12

import (
	"algorithm/util"
	"reflect"
	"testing"
)

/*
문제 4: 벽을 부수고 이동하기 (Break the Wall)
난이도: 중급~상급
문제 유형: #DFS #BFS #최단거리 탐색
권장 풀이 시간: ⏳ 40분
출처: 직접 만든 문제

📝 문제 설명
n x m 크기의 2D 격자가 주어진다.
0 → 이동할 수 있는 공간
1 → 벽(장애물)
(0,0)에서 (n-1, m-1)까지 최단 거리로 이동해야 한다.
단, 벽(1)은 최대 1번까지만 부술 수 있다.
최단 경로의 길이를 반환하라.
이동 방향: 상하좌우 (4방향 이동 가능, 대각선 이동 불가능)
📥 입력
grid: n x m 크기의 2D 정수 배열 (0 또는 1)
📤 출력
int 값 하나를 반환 (최단 경로의 길이)
도착할 수 없는 경우 -1 반환
*/

// 벽을 부수고 최단 경로 찾는 함수 (직접 구현)
func shortestPathWithWall(grid [][]int) int {
	if len(grid) == 0 {
		return -1
	}

	n, m := len(grid), len(grid[0])
	visited := make([][][]bool, n)
	for i := range visited {
		visited[i] = make([][]bool, m)
		for j := range visited[i] {
			visited[i][j] = make([]bool, 2)
		}
	}

	var dfs func(i, j, broken, distance int) int
	dfs = func(i, j, broken, distance int) int {
		// 1. 범위 체크
		if i < 0 || i >= n || j < 0 || j >= m {
			return -1
		}
		// 2. 이미 방문했으면 종료
		if visited[i][j][broken] {
			return -1
		}
		// 3. 도착지점에 도달하면 거리 반환
		if i == n-1 && j == m-1 {
			return distance
		}

		// 4. 방문처리
		visited[i][j][broken] = true

		// 5. 벽(1)을 만난경우 처리
		if grid[i][j] == 1 {
			if broken == 1 {
				return -1
			}
			broken = 1
		}

		// 6. 상하좌우 이동
		minPath := -1
		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
		for _, d := range directions {
			ni, nj := i+d[0], j+d[1]
			res := dfs(ni, nj, broken, distance+1)
			if res != -1 && (minPath == -1 || res < minPath) {
				minPath = res
			}
		}

		//7. 방문해제 (백트래킹)
		visited[i][j][broken] = false

		return minPath
	}
	return dfs(0, 0, 0, 1)
}

// 테스트 코드
func TestShortestPathWithWall(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{
			[][]int{
				{0, 0, 0, 0, 1},
				{1, 1, 1, 0, 1},
				{0, 0, 0, 0, 0},
			},
			7,
		},
		{
			[][]int{
				{0, 1, 1, 1, 1},
				{0, 0, 1, 1, 1},
				{1, 0, 0, 0, 0},
			},
			7,
		},
		{
			[][]int{
				{0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 1, 1, 1, 0},
			},
			-1,
		},
	}

	for _, tt := range tests {
		got := util.RunWithTimer("ShortestPathWithWall", shortestPathWithWall, tt.grid).(int)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("❌ 실패: shortestPathWithWall(%v) = %v; want %v", tt.grid, got, tt.want)
		} else {
			t.Logf("✅ 성공: got = %v", got)
		}
	}
}
