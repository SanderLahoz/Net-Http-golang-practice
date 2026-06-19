package loadbalancer

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

var (
	baseUrl = "http://localhost:808"
)

type LoadBalancer struct {
	RevProxy httputil.ReverseProxy
}

type Endpoints struct {
	List []*url.URL
}

// Shuffle : Shuffles through the list like a cartwheel
func (e *Endpoints) Shuffle() {
	temp := e.List[0]
	e.List = e.List[1:]
	e.List = append(e.List, temp)
}

func MakeLoadBalancer(amount int) {
	var lb LoadBalancer
	var ep Endpoints

	// Server + Router
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8090",
		Handler: router,
	}

	// Creating the endpoints
	for i := 0; i < amount; i++ {
		ep.List = append(ep.List, createEndpoint(baseUrl, i))
	}

	// Handler functions
	router.HandleFunc("/loadbalancer", makeRequest(&lb, &ep))

	// Listen and server
	log.Fatal(server.ListenAndServe())
}

func makeRequest(lb *LoadBalancer, ep *Endpoints) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for !isHealthy(ep.List[0].String()) {
			ep.Shuffle()
		}

		lb.RevProxy = *httputil.NewSingleHostReverseProxy(ep.List[0])
		lb.RevProxy.ServeHTTP(w, r)
		ep.Shuffle()
	}
}

func createEndpoint(endpoint string, idx int) *url.URL {
	link := endpoint + strconv.Itoa(idx)
	parsedUrl, err := url.Parse(link)
	if err != nil {
		log.Fatal(err)
	}
	return parsedUrl
}

func isHealthy(endpoint string) bool {
	resp, err := http.Get(endpoint)
	if err != nil {
		return false
	}
	if resp.StatusCode != http.StatusOK {
		return false
	}
	return true
}
