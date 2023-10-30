package oxo

// board represents the state of a turn of the game on a 3 x 3 grid
type Grid [9]byte

// a turn has a board and status string which will be one of these strings XWIN,OWIN,DRAW,PLAY
type Turn struct {
	board  Grid
	status string
}
type Game []Turn
