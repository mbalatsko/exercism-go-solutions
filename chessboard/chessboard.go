package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) (cnt int) {
	foundFile := cb[file]

	for _, v := range foundFile {
		if v {
			cnt++
		}
	}

	return
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) (cnt int) {
	if rank < 1 || rank > 8 {
		return
	}

	for i := range cb {
		if cb[i][rank-1] {
			cnt++
		}
	}
	return
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) (cnt int) {
	for _, v := range cb {
		cnt += len(v)
	}
	return
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) (cnt int) {
	for i := 1; i < 9; i++ {
		cnt += CountInRank(cb, i)
	}
	return
}
