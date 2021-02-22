package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	store := StubPlayStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
	}

	server := &PlayerServer{&store}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertRequestBody(t, response.Body.String(), "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusOK)
		assertRequestBody(t, response.Body.String(), "10")
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusNotFound)

	})

}

func TestStoreWins(t *testing.T) {
	store := StubPlayStore{
		map[string]int{},
		nil,
	}
	server := &PlayerServer{&store}
	t.Run("it records wins when POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Fatalf("got %d ,want %d", len(store.winCalls), 1)
		}
		if store.winCalls[0] != player {
			t.Errorf("got %s, want %s", store.winCalls[0], player)
		}

	})

}

func TestRecordingWinRetrievingThem(t *testing.T) {
	store := *NewInMemoryPlayerStore()
	server := PlayerServer{&store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertResponseStatus(t, response.Code, http.StatusOK)

	assertRequestBody(t, response.Body.String(), "3")

}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status %d , want status %d", got, want)
	}
}

func assertRequestBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("want %s, got %s", want, got)

	}
}
