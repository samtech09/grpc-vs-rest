package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/samtech09/grpc-vs-rest/grpc/model"
	"github.com/samtech09/grpc-vs-rest/grpc/service"
	"github.com/samtech09/grpc-vs-rest/server/handlers"

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

func TestGetDetailRest(t *testing.T) {
	req, err := http.NewRequest("GET", "/getDetail", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("from", "2016-01-01")
	q.Add("to", "2019-11-11")
	req.URL.RawQuery = q.Encode()
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetDetailRest)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"name":"Test1","rollNo":"Roll-1","age":21,"examCleared":true},{"id":2,"name":"Test2","rollNo":"Roll-2","age":22,"examCleared":true}]` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func BenchmarkGetDetailRest(b *testing.B) {
	req, err := http.NewRequest("GET", "/getDetail", nil)
	if err != nil {
		b.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("from", "2016-01-01")
	q.Add("to", "2019-11-11")
	req.URL.RawQuery = q.Encode()

	for n := 0; n < b.N; n++ {
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handlers.GetDetailRest)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			b.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `[{"id":1,"name":"Test1","rollNo":"Roll-1","age":21,"examCleared":true},{"id":2,"name":"Test2","rollNo":"Roll-2","age":22,"examCleared":true}]` + "\n"
		if rr.Body.String() != expected {
			b.Errorf("handler returned unexpected body: got %v want %v",
				rr.Body.String(), expected)
		}
	}
}

func TestGetDetailRestLive(t *testing.T) {
	// Test as we do in real-world (server/client)
	req, err := http.NewRequest("GET", "http://localhost:8080/getDetail", nil)
	if err != nil {
		t.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("from", "2016-01-01")
	q.Add("to", "2019-11-11")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if status := resp.StatusCode; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":1,"name":"Test1","rollNo":"Roll-1","age":21,"examCleared":true},{"id":2,"name":"Test2","rollNo":"Roll-2","age":22,"examCleared":true}]` + "\n"
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	strbody := string(body)
	if strbody != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			strbody, expected)
	}
}

func BenchmarkGetDetailRestLive(b *testing.B) {
	req, err := http.NewRequest("GET", "http://localhost:8080/getDetail", nil)
	if err != nil {
		b.Fatal(err)
	}
	q := req.URL.Query()
	q.Add("from", "2016-01-01")
	q.Add("to", "2019-11-11")
	req.URL.RawQuery = q.Encode()

	for n := 0; n < b.N; n++ {
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			panic(err)
		}

		if status := resp.StatusCode; status != http.StatusOK {
			b.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}

		expected := `[{"id":1,"name":"Test1","rollNo":"Roll-1","age":21,"examCleared":true},{"id":2,"name":"Test2","rollNo":"Roll-2","age":22,"examCleared":true}]` + "\n"
		defer resp.Body.Close()

		body, _ := ioutil.ReadAll(resp.Body)
		strbody := string(body)
		if strbody != expected {
			b.Errorf("handler returned unexpected body: got %v want %v",
				strbody, expected)
		}
	}
}

func TestGetDetailRestLiveByPost(t *testing.T) {
	filter := model.Filter{}
	filter.From = "2016-01-01"
	filter.To = "2019-11-11"

	payload, err := json.Marshal(filter)
	if err != nil {
		t.Error("Failed marshling payload")
	}

	resp, err := http.Post("http://localhost:8080/getDetailbyost", "application/json", bytes.NewBuffer(payload))
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

func BenchmarkGetDetailRestLiveByPost(b *testing.B) {
	for n := 0; n < b.N; n++ {
		filter := model.Filter{}
		filter.From = "2016-01-01"
		filter.To = "2019-11-11"

		payload, err := json.Marshal(filter)
		if err != nil {
			b.Error("Failed marshling payload")
		}

		resp, err := http.Post("http://localhost:8080/getDetailbyost", "application/json", bytes.NewBuffer(payload))
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
