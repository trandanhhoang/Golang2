package v1

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()
	fmt.Println("req", request)
	fmt.Println("res", response)

	PlayerServer(response, request)

	t.Run("returns Pepper's score", func(t *testing.T) {
		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

}
