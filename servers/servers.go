package servers

import (
	"fmt"
	"log"
)

type ServerList struct {
	Ports []int
}

func (s *ServerList) Populate(amount int) {
	if amount > 10 {
		log.Fatalf("Amount of ports cant exceed 10")
	}

	for i := 0; i < amount; i++ {
		s.Ports = append(s.Ports, i)
	}
}

func (s *ServerList) Pop() int {
	port := s.Ports[0]
	s.Ports = s.Ports[1:]
	return port
}

func Test() {
	var myServerList ServerList

	myServerList.Populate(5)
	limit := len(myServerList.Ports)

	for i := 0; i < limit; i++ {
		fmt.Println(myServerList.Pop())
	}
}
