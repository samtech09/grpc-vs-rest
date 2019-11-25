package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/samtech09/grpc-vs-rest/server/reposit"
)

func GetDetailRest(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["from"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Url Param 'from' is missing", 500)
		return
	}
	from := keys[0]

	key, ok := r.URL.Query()["to"]
	if !ok || len(key[0]) < 1 {
		http.Error(w, "Url Param 'to' is missing", 500)
		return
	}
	to := key[0]

	dd, _ := reposit.GetData(from, to)
	json.NewEncoder(w).Encode(dd)
}
