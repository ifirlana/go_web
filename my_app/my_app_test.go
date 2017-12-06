package my_app_test

import (
	"testing"
	"net/http/httptest"
	"net/http"
)

func TestHomePageHandler(t *testing.T) {
	//t.Error("I failed on purpose")
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET","/", nil)

	//my_app.HomePagehandler()
}