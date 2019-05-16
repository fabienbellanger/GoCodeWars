package kata

import "strings"

// Points takes such collection and counts the points of our team in the championship
func Points(games []string) int {
	points := 0

	for _, game := range games {
		score := strings.Split(game, ":")

		if len(score) == 2 {
			if score[0] > score[1] {
				points += 3
			} else if score[0] == score[1] {
				points++
			}
		}
	}

	return points
}
