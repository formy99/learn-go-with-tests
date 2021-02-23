package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestLeague(t *testing.T) {
	wantedLeague := []Player{
		{"Cleo", 32},
		{"Chris", 20},
		{"Tiest", 14},
	}
	store := StubPlayStore{nil, nil, wantedLeague}
	server := NewPlayerServer(&store)
	t.Run("it returns 200 on /league", func(t *testing.T) {
		request := NewLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		got := GetLeagueFromResponse(t, response)

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContextType(t, response, "application/json")

	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want status %d", got, want)
	}
}

func assertLeague(tb testing.TB, got, want []Player) {
	tb.Helper()
	if !reflect.DeepEqual(got, want) {
		tb.Errorf("want %v , got %v", want, got)
	}
}

func assertContextType(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("context-type") != want {
		t.Errorf("reponse did not have context-type of application/json, got '%v' want 'context-type' ", response.HeaderMap)
	}
}

func NewLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func GetLeagueFromResponse(t testing.TB, response *httptest.ResponseRecorder) (league []Player) {
	t.Helper()
	err := json.NewDecoder(response.Body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
	}
	return
}
