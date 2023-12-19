package main

import (
	s "forum/server"
)

func main() {
	server := s.New() //create new server
	server.Start()    // start the server
}
