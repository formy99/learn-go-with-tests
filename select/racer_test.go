package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowDelayTime := 20 * time.Millisecond
		fastDelyaTime := 10 * time.Millisecond

		slowDelayServer := makeDelayedServer(slowDelayTime)
		fastDelyaServer := makeDelayedServer(fastDelyaTime)

		defer slowDelayServer.Close()
		defer fastDelyaServer.Close()

		slowURL := slowDelayServer.URL
		fastURL := fastDelyaServer.URL
		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Errorf("didn't expect err but got one %v", err)
		}

		if want != got {
			t.Errorf("want '%s', got '%s'", want, got)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 2*time.Second)

		if err == nil {
			t.Errorf("expected an error but didn't get one")
		}
	})

}
