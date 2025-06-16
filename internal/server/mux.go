package server

import (
	"encoding/json"
	"math/rand"
	"my-oauth-server/internal/calculator"
	"my-oauth-server/internal/oauth"
	"my-oauth-server/internal/utils"
	"net/http"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func newMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", healthCheck)

	mux.HandleFunc("/test", test)
	mux.HandleFunc("/test2", test2)
	mux.HandleFunc("/test-queue", testQueue)

	oauth.RegisterController(mux)
	calculator.RegisterController(mux)

	return mux
}

func test(w http.ResponseWriter, _ *http.Request) {
	u := utils.User{ID: 1, Name: "Saran", Age: 24}

	token, err := utils.GenerateToken(u)
	if err != nil {
		utils.WriteResponseBody(w, utils.ResponseBody{Error: err.Error()})
		return
	}

	utils.WriteResponseBody(w, utils.ResponseBody{Data: token, Message: "done"})
}

func test2(w http.ResponseWriter, req *http.Request) {
	token := req.Header.Get("authorization")

	u, err := utils.DecryptToken(token)
	if err != nil {
		utils.WriteResponseBody(w, utils.ResponseBody{Error: err.Error()})
		return
	}

	utils.WriteResponseBody(w, utils.ResponseBody{Data: u, Message: "done"})
}

func testQueue(w http.ResponseWriter, req *http.Request) {
	go infinitePublish(w)
	utils.WriteResponseBody(w, utils.ResponseBody{Message: "done"})
}

func infinitePublish(w http.ResponseWriter) {
	ticker := time.NewTicker(time.Millisecond * 100)
	for {
		body := map[string]any{
			"id":    rand.Int31()%100 + 1,
			"name":  "Saran",
			"stock": "NVDA",
			"price": rand.Int31(),
		}

		jsonBody, err := json.Marshal(body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			utils.WriteResponseBody(w, utils.ResponseBody{Error: err.Error()})
			return
		}

		MainChannel.Publish(
			"",
			MainQueue.Name,
			false,
			false,
			amqp.Publishing{
				ContentType: "application/json",
				Body:        jsonBody,
			},
		)
		<-ticker.C
	}
}
