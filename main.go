package main

import "loadbalancer/loadbalancer"

// First uncomment the servers.RunServers to start 5 server instances
// Then uncomment the loadbalancer.MakeLoadBalancer and comment out the
// servers.RunServers and run the file in a new terminal to start the
// loadbalancer then visit the loadbalancer path localhost:8090/loadbalancer

func main() {
	// servers.RunServers(5)
	loadbalancer.MakeLoadBalancer(5)
}
