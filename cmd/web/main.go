package main

import (
	"fmt"

	"github.com/finecoding/oxopfork/internal/oxoweb"
)

func main() {

	board := oxoweb.Grid([]byte("X........"))
	fmt.Println(board)
}
