package chessboard

//import "fmt"
// Declare a type named Rank which stores if a square is occupied by a piece - this will be a slice of bools
type Rank []bool

// Declare a type named Chessboard contains a map of eight Ranks, accessed with values from "A" to "H"
type Chessboard = map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	var count int = 0
	for _, occupeid := range cb[rank] {
		if occupeid {
			count += 1
		}
	}
	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	var count int = 0
	if file >= 1 && file <= 8 {
		for _, cells := range cb {
			if cells[file-1] {
				count += 1
			}
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	var count int = 0
	for _, squares := range cb {
		count += len(squares)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	var count int = 0
	for _, squares := range cb {
		for _, occupied := range squares {
			if occupied {
				count += 1
			}
		}
	}
	return count
}
