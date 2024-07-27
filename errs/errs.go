package errs

import (
	"encoding/json"
	"net/http"
)


type Error struct{
    StatusCode  int `json:"statusCode"`
    Err         string  `json:"err"`
}

func JSONError(w http.ResponseWriter)  {
    err := Error{ StatusCode: 400, Err: "invalid json" }
    json.NewEncoder(w).Encode(err)
}

func ConnError(w http.ResponseWriter) {
    err := Error{ StatusCode: 400, Err: "failed to connect" } 
    json.NewEncoder(w).Encode(err)
}
