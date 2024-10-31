package race

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedHttpServer(10 * time.Millisecond)
		fastServer := makeDelayedHttpServer(0 * time.Millisecond)

		// make resource clean up near where the resource were created for readability
		// `defer` it will now call that function at the end of the containing function.
		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		expected := fastUrl
		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Error("did not expect an error but got one %v", err)
		}

		if got != expected {
			t.Errorf("got %q, want %q", got, expected)
		}
	})

	t.Run("returns an error if a server doesn't respond within 10s", func(t *testing.T) {
		serverA := makeDelayedHttpServer(20 * time.Millisecond)
		serverB := makeDelayedHttpServer(20 * time.Millisecond)

		defer serverA.Close()
		defer serverB.Close()

		_, err := ConfigurableRacer(serverA.URL, serverB.URL, 10*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedHttpServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
