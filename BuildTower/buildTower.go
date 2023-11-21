package BuildTower

import (
	"strings"
)

func TowerBuilder(nFloors int) []string {

	tower := make([]string, 0)

	for i := 0; i < nFloors; i++ {
		spaces := strings.Repeat(" ", nFloors-i-1)
		star := strings.Repeat("*", 2*i+1)
		twr := spaces + star + spaces
		tower = append(tower, twr)
	}

	return tower
}
