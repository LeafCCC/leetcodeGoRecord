package leetcode

func squareIsWhite(coordinates string) bool {
	return (int(coordinates[0]-'a')+int(coordinates[1]-'1'))%2 == 1
}
