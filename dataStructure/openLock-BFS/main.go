package main

import (
	"fmt"
)

func openLock(deadends []string, target string) int {
	dead := make(map[string]bool)
	for _, d := range deadends {
		dead[d] = true
	}

	if dead["0000"] {
		return -1
	}

	if target == "0000" {
		return 0
	}

	queue := []string{"0000"}
	visited := make(map[string]bool)
	visited["0000"] = true

	step := 0

	for len(queue) > 0 {
		step++
		size := len(queue)
		for i := 0; i<size; i++ {
			curr := queue[0]
			queue = queue[1:]

			for j := 0; j<4; j++ {
				for k:=-1; k<=1; k+=2 {
					next := []byte(curr)
					next[j] = (next[j]-'0'+byte(k)+10)%10 + '0'

					nextState := string(next)

					if nextState == target {
						return step
					}

					if !dead[nextState] && !visited[nextState] {
						queue = append(queue, nextState)
						visited[nextState] = true
					}
				}
			} 
		}

	}

	return -1 // there is no path from root to target
}

func main() {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
    target := "0202"
    fmt.Println(openLock(deadends, target)) // Output: 6
}