package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/samtech09/grpc-vs-rest/grpc/model"
	"github.com/samtech09/grpc-vs-rest/grpc/service"
	"github.com/vmihailenco/msgpack"

	"google.golang.org/grpc"
)

var testclient service.ReportClient

func init() {
	go startGrpcServer()
	go startRestServer()
	initGrpcClient()
}

func initGrpcClient() {
	serverAddress := "localhost:8888"
	conn, e := grpc.Dial(serverAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatal(e)
	}
	testclient = service.NewReportClient(conn)
}

// GRPC

func TestGetDetailGrpc(t *testing.T) {
	filter := model.Filter{From: "2016-01-01", To: "2019-10-10"}

	_, err := testclient.GetDetail(context.Background(), &filter)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkGetDetailGrpc(b *testing.B) {
	filter := model.Filter{From: "2016-01-01", To: "2019-10-10"}

	for n := 0; n < b.N; n++ {
		_, err := testclient.GetDetail(context.Background(), &filter)
		if err != nil {
			b.Error(err)
		}
	}
}

// REST

func TestGetDetailRestLiveByPost(t *testing.T) {
	filter := model.Filter{}
	filter.From = "2016-01-01"
	filter.To = "2019-11-11"

	payload, err := json.Marshal(filter)
	if err != nil {
		t.Error("Failed marshling payload")
	}

	resp, err := http.Post("http://localhost:8080/getDetailbyPost", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		t.Error("Error making request. ", err)
	}

	defer resp.Body.Close()

	// body, _ := ioutil.ReadAll(resp.Body)
	// strbody := string(body)
	// fmt.Println(strbody)

	dest := []model.StudentDetails{}
	err = json.NewDecoder(resp.Body).Decode(&dest)
	if err != nil {
		t.Error("Error decoding response: ", err.Error())
	}

	expected := 2
	if len(dest) != expected {
		t.Errorf("handler returned unexpected number of records: got %d want %d", len(dest), expected)
	}
}

func TestGetDetailRestLiveByPostMsgpack(t *testing.T) {
	filter := model.Filter{}
	filter.From = "2016-01-01"
	filter.To = "2019-11-11"

	payload, err := msgpack.Marshal(filter)
	if err != nil {
		t.Error("Failed marshling payload")
	}

	resp, err := http.Post("http://localhost:8080/getDetailbyPostMsgpack", "application/x-msgpack", bytes.NewBuffer(payload))
	if err != nil {
		t.Error("Error making request. ", err)
	}

	defer resp.Body.Close()

	dest := []model.StudentDetails{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("Error reading response: ", err.Error())
	}

	err = msgpack.Unmarshal(body, &dest)
	if err != nil {
		t.Error("Error decoding response: ", err.Error())
	}

	expected := 2
	if len(dest) != expected {
		t.Errorf("handler returned unexpected number of records: got %d want %d", len(dest), expected)
	}
}

func BenchmarkGetDetailRestLiveByPostJsoniter(b *testing.B) {
	json2 := jsoniter.ConfigCompatibleWithStandardLibrary

	for n := 0; n < b.N; n++ {
		filter := model.Filter{}
		filter.From = "2016-01-01"
		filter.To = "2019-11-11"

		payload, err := json2.Marshal(filter)
		if err != nil {
			b.Error("Failed marshling payload")
		}

		resp, err := http.Post("http://localhost:8080/getDetailbyPostJsoniter", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			b.Error("Error making request. ", err)
		}

		defer resp.Body.Close()
		dest := []model.StudentDetails{}
		err = json2.NewDecoder(resp.Body).Decode(&dest)
		if err != nil {
			b.Error("Error decoding response: ", err)
		}

		expected := 2
		if len(dest) != expected {
			b.Errorf("handler returned unexpected number of records: got %d want %d", len(dest), expected)
		}
	}
}

func BenchmarkGetDetailRestLiveByPostMsgpack(b *testing.B) {
	for n := 0; n < b.N; n++ {
		filter := model.Filter{}
		filter.From = "2016-01-01"
		filter.To = "2019-11-11"

		payload, err := msgpack.Marshal(filter)
		if err != nil {
			b.Error("Failed marshling payload")
		}

		resp, err := http.Post("http://localhost:8080/getDetailbyPostMsgpack", "application/x-msgpack", bytes.NewBuffer(payload))
		if err != nil {
			b.Error("Error making request. ", err)
		}

		defer resp.Body.Close()
		dest := []model.StudentDetails{}
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			b.Error("Error reading response: ", err.Error())
		}

		err = msgpack.Unmarshal(body, &dest)
		if err != nil {
			b.Error("Error decoding response: ", err.Error())
		}

		expected := 2
		if len(dest) != expected {
			b.Errorf("handler returned unexpected number of records: got %d want %d", len(dest), expected)
		}
	}
}

func BenchmarkGetDetailRestLiveByPost(b *testing.B) {
	for n := 0; n < b.N; n++ {
		filter := model.Filter{}
		filter.From = "2016-01-01"
		filter.To = "2019-11-11"

		payload, err := json.Marshal(filter)
		if err != nil {
			b.Error("Failed marshling payload")
		}

		resp, err := http.Post("http://localhost:8080/getDetailbyPost", "application/json", bytes.NewBuffer(payload))
		if err != nil {
			b.Error("Error making request. ", err)
		}

		defer resp.Body.Close()
		dest := []model.StudentDetails{}
		err = json.NewDecoder(resp.Body).Decode(&dest)
		if err != nil {
			b.Error("Error decoding response: ", err)
		}

		expected := 2
		if len(dest) != expected {
			b.Errorf("handler returned unexpected number of records: got %d want %d", len(dest), expected)
		}
	}
}
