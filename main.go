package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/pykmi/api-seed/datastore"
)

// type ServerOptions struct {
// 	HttpHost string
// 	HttpPort string
// }

var (
	StoreOpt = datastore.StoreOptions{
		DBPath:        "couchbase://localhost",
		Namespace:     "default",
		RetryAttempts: 5,
	}
)

func main() {
	// set default commandline flags and parse them
	httphost := flag.String("host", "localhost", "HTTP hostname")
	httpport := flag.String("port", "80", "HTTP port number")

	flag.Parse()

	server := net.JoinHostPort(*httphost, *httpport)

	// Create new datastore
	store, err := datastore.New(StoreOpt)
	if err != nil {
		log.Println(err)
	}

	// start the server
	if err := startServer(server, store); err != nil {
		log.Printf("%#v", err)
		return
	}
}

/**
 * Starts the HTTP server.
 */
func startServer(server string, store *datastore.Store) error {
	log.Println("Server started on at: ", server)

	// create http routes & database middleware
	APIrouter := setupRoutes(store)

	// start listening for the client connections
	err := http.ListenAndServe(server, APIrouter)
	if err != nil {

		fmt.Println(err)
		return err
	}

	return nil
}
