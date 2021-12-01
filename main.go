package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

//function returns error
func repo() error {
	x := true
	var customErrorArray []CustomError
	if x {
		xr := errors.New("invalid")
		customErrorArray = ReturnErrorArray(customErrorArray, 400, "wrong", "main", xr)
	}
	if y := false; !y {
		xr := errors.New("invalid")
		customErrorArray = ReturnErrorArray(customErrorArray, 400, "wrong", "main", xr)
	}
	fmt.Println(len(customErrorArray))
	if len(customErrorArray) > 0 {
		return ErrMessageClient(customErrorArray)
		
	}
	return nil
}

//handles error responses
func ResponseError(rw http.ResponseWriter, err error) {
	rw.Header().Set("Content-Type", "Application/json")
	var ew CustomClientErrorBody
	if errors.As(err, &ew) {
		rw.WriteHeader(404)
		for _, val := range ew.Errors {
			fmt.Println(val.Err)
		}
		_ = json.NewEncoder(rw).Encode(ew)
		return
	}

	// handles non CustomErrorWrapper types
	//structure in progress
	//this works for now
	rw.WriteHeader(500)
	log.Println(err.Error())
	_ = json.NewEncoder(rw).Encode(map[string]interface{}{
		"message": err.Error(),
	})
}

//handles success responses
func ResponseSuccess(rw http.ResponseWriter, data interface{}) {
	rw.Header().Set("Content-Type", "Application/json")
	body := map[string]interface{}{
		"data": data,
	}
	_ = json.NewEncoder(rw).Encode(body)
}

//http handler
func handler(rw http.ResponseWriter, r *http.Request) {
	err := repo()
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
