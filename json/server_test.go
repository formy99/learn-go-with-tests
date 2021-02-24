package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestLeague(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	t.Run("get league", func(t *testing.T) {
		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)

		got := getLeagueFromResponse(t, response)
		want := []Player{
			{"Pepper", 3},
		}

		assertLeague(t, got, want)
		assertContextType(t, response, "application/json")

	})

	t.Run("get score", func(t *testing.T) {
		request := newGetScoreRequest(player)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response, "3")

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

func assertResponseBody(t testing.TB, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Body.String() != want {
		t.Errorf("reponse did not have context-type of application/json, got '%v' want 'context-type' ", response.Body.String())
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func getLeagueFromResponse(t testing.TB, response *httptest.ResponseRecorder) (league []Player) {
	t.Helper()
	fmt.Println(response.Body)
	err := json.NewDecoder(response.Body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
	}
	return
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
