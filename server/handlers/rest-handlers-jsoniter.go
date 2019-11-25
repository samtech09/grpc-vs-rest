package handlers

import (
	"net/http"

	jsoniter "github.com/json-iterator/go"
	"github.com/samtech09/grpc-vs-rest/grpc/model"
	"github.com/samtech09/grpc-vs-rest/server/reposit"
)

var json2 = jsoniter.ConfigCompatibleWithStandardLibrary

func GetDetailRestByPostJsoniter(w http.ResponseWriter, r *http.Request) {
	// create decoder
	decoder := json2.NewDecoder(r.Body)
	fltr := model.Filter{}

	err := decoder.Decode(&fltr)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if fltr.From == "" {
		http.Error(w, "Url Param 'from' is missing", 500)
		return
	}

	if fltr.To == "" {
		http.Error(w, "Url Param 'to' is missing", 500)
		return
	}

	dd, _ := reposit.GetData(fltr.From, fltr.To)
	json2.NewEncoder(w).Encode(dd)
}
