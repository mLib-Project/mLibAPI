package tests

import (
	fmt "fmt"
	http "net/http"
	httptest "net/http/httptest"
	testing "testing"

	Config "mLibAPI/src/config"
	App "mLibAPI/src/router"
)

func TestBooksRoute(t *testing.T) {
	Config.Init()

	ts := httptest.NewServer(App.CreateServer())
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api", ts.URL))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}

	val, ok := resp.Header["Content-Type"]
	if !ok {
		t.Fatalf("Expected Content-Type header to be set")
	}
	if val[0] != "application/json; charset=utf-8" {
		t.Fatalf("Expected \"application/json; charset=utf-8\", got %s", val[0])
	}
}
