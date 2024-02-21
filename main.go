package main

import (
	"fmt"

	"github.com/Chanawee/cinema/movie"
	"github.com/Chanawee/cinema/ticket"
)

func init() {
	fmt.Println("init: main")
}

func main(){
	movieName := movie.FindMovieName("tt4154796")
	fmt.Println(movieName)
	movie.RevieMovie(movieName, 8.4)
	ticket.BuyTicket(movieName)
}