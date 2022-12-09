package main

type coord struct{ x, y int }

// dist returns the x distance and y distance from this coord to the provided coord
func (c coord) dist(other coord) (int, int) {
	return c.x - other.x, c.y - other.y
}

// touching takes two coords and returns whether they are touching
//
// points are considered touching if they are vertically or horizontally adjacent or one cell diagonal.
// practically, this means their x delta and y delta must both be between 1 and -1
func (c coord) touching(other coord) bool {
	xDist, yDist := c.dist(other)
	return -1 <= yDist && yDist <= 1 && -1 <= xDist && xDist <= 1
}

// direction will return the sequence of directions this coord needs to move by 1 space in order to be
// touching the other coord
//
// if the coords are in the same column or row, this will return 1 direction, otherwise this will provide
// the two directions required to move in the appropriate diagonal
func (c coord) direction(other coord) []direction {
	xDist, yDist := c.dist(other)
	if xDist == 0 {
		// same column
		if yDist > 1 {
			return []direction{down}
		}
		return []direction{up}
	}
	if yDist == 0 {
		// same row
		if xDist > 1 {
			return []direction{left}
		}
		return []direction{right}
	}
	if xDist >= 1 {
		// right one column
		if yDist >= 1 {
			// up one column
			return []direction{left, down}
		}
		return []direction{left, up}
	}
	if xDist <= -1 {
		// left one column
		if yDist >= 1 {
			// up one column
			return []direction{right, down}
		}
		return []direction{right, up}
	}
	return nil
}
