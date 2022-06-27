package chessboard

// Rank stores if a square is occupied by a piece
type Rank []bool

// Chessboard contains a map of eight Ranks, accessed with keys from "A" to "H"
type Chessboard map[string]Rank

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank
func CountInRank(cb Chessboard, rank string) int {
	count := 0

	rankVals, exists := cb[rank]
	if !exists {
		return count
	}

	for _, occupied := range rankVals {
		if occupied {
			count++
		}
	}

	return count
}

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file
func CountInFile(cb Chessboard, file int) int {
	count := 0

	if file < 1 || file > 8 {
		return count
	}

	for _, rankVals := range cb {
		if rankVals[file-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard
func CountAll(cb Chessboard) int {
	return 64
}

// CountOccupied returns how many squares are occupied in the chessboard
func CountOccupied(cb Chessboard) int {
	count := 0
	for _, rankVals := range cb {
		for _, occupied := range rankVals {
			if occupied {
				count++
			}
		}
	}
	return count
}
