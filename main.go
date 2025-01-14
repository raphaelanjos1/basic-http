package main

import "net/http"

func main() {
	// implements the httpHandler interface
	api := &api{addr: ":8080"}

	// initialize mux - mux is a router.
	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:    api.addr,
		Handler: mux,
	}

	// mux.HandleFunc("GET /users", api.getUsersHandler)

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
