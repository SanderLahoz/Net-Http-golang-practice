package servers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
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

func RunServers(amount int) {
	// ServerList Object - creation
	var myServerList ServerList
	myServerList.Populate(amount)

	// Wait group - initialization
	var wg sync.WaitGroup
	wg.Add(amount)
	defer wg.Wait()

	for i := 0; i < amount; i++ {
		go initializeServer(&myServerList, &wg)
	}
}

func initializeServer(sl *ServerList, wg *sync.WaitGroup) {
	defer wg.Done()
	r := http.NewServeMux()

	port := sl.Pop()
	server := http.Server{
		Addr:    fmt.Sprintf(":808%d", port),
		Handler: r,
	}

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := fmt.Fprintf(w, "Server %d", port)
		if err != nil {
			log.Fatal(err)
		}
	})

	r.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		_, err := w.Write([]byte("503 - Server shutdown"))
		if err != nil {
			return
		}
		err2 := server.Shutdown(context.Background())
		if err2 != nil {
			return
		}
	})

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
