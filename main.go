package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func repoSimulation() error {
	x := true
	var cES []CustomErrors
	if x {
		xr := errors.New("invalid")
		cES = ErrArray(cES, 400, "wrong", "main", xr)
	}
	if y := false; !y {
		xr := errors.New("invalid")
		cES = ErrArray(cES, 400, "wrong", "main", xr)
	}
	fmt.Println(len(cES))
	if len(cES) > 0 {
		return ErrMessageClient(cES)
		
	}
	return nil
}
func ResponseError(rw http.ResponseWriter, err error) {
	rw.Header().Set("Content-Type", "Application/json")
	var ew CustomClientError
	if errors.As(err, &ew) {
		rw.WriteHeader(404)
		for _, val := range ew.Errors {
			fmt.Println(val.Err)
		}
		_ = json.NewEncoder(rw).Encode(ew)
		return
	}
	// handle non CustomErrorWrapper types
	rw.WriteHeader(500)
	log.Println(err.Error())
	_ = json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": err.Error(),
	})
}

func ResponseSuccess(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set("Content-Type", "Application/json")
	body := map[string]interface{}{
		"data": data,
	}
	_ = json.NewEncoder(rw).Encode(body)
}

func handler(rw http.ResponseWriter, r *http.Request) {
	err := repoSimulation()
	if err != nil {
		ResponseError(rw, err)
		return
	}
	ResponseSuccess(rw, "this code should not be reachable")
}

func main() {
	// Routes everything to handler
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	log.Println("server is running on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
