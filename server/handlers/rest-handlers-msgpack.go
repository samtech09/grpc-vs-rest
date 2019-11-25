package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/samtech09/grpc-vs-rest/grpc/model"
	"github.com/samtech09/grpc-vs-rest/server/reposit"
	"github.com/vmihailenco/msgpack"
)

func GetDetailRestByPostMsgpack(w http.ResponseWriter, r *http.Request) {
	// decode passed msgpack
	fltr := model.Filter{}
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = msgpack.Unmarshal(b, &fltr)
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
	b, _ = msgpack.Marshal(dd)
	w.Header().Set("Content-Type", "application/x-msgpack")
	w.Write(b)
}
