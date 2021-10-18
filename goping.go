package main

import (
    "log"
    "net/http"
    "fmt"
    "os"
)

func apiResponse(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  switch r.Method {
    case "HEAD":
      w.WriteHeader(http.StatusOK)
    case "GET":
      w.WriteHeader(http.StatusOK)
      w.Write([]byte(`{"message": "OK: GET method requested"}`))
    default:
      w.WriteHeader(http.StatusNotFound)
      w.Write([]byte(`{"message": "ERROR: Can't find method requested"}`))
    }
}

func livenessResponse(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "application/json")

  switch r.Method {
    case "HEAD":
      w.WriteHeader(http.StatusOK)
    case "GET":
      hostname, err := os.Hostname()
	    if err != nil {
		    log.Println(err)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write([]byte(`{"message": "ERROR: Can't get Hostname"}`))
	    }
      w.WriteHeader(http.StatusOK)
      json := fmt.Sprintf("{\"message\": \"OK: PONG from %s\"}", hostname)
      w.Write([]byte(json))
    default:
      w.WriteHeader(http.StatusNotFound)
      w.Write([]byte(`{"message": "ERROR: Can't find method requested"}`))
    }
}

func main() {
  log.Println("Goping started")
  //http.HandleFunc("/",apiResponse)
  http.HandleFunc("/ping",livenessResponse)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

