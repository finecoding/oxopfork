// Data types go here as the need arises to define them.

package oxo

// board represents the state of a turn of the game on a 3 x 3 grid
type Grid [9]byte

// a turn has a board and status string which will be one of these strings XWIN,OWIN,DRAW,PLAY
type Turn struct {
	Board  Grid
	Status string
}

// a complete game has a minimum of 5 turns and a maximum of 9 turns
// should I set a capacity for the slice?
type Game []Turn

func (board Grid) InPlay() ([]byte, bool) {
	var spc []byte
	for k, v := range board {
		if v == 32 {
			spc = append(spc, byte(k))
		}
	}
	return spc, len(spc) != 0
}
