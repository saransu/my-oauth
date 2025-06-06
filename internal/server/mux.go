package server

import (
	"encoding/json"
	"my-oauth-server/internal/calculator"
	"my-oauth-server/internal/oauth"
	"my-oauth-server/internal/utils"
	"net/http"

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
	body := map[string]any{
		"jobId": 123,
		"name":  "Saran",
		"stock": "NVDA",
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

	utils.WriteResponseBody(w, utils.ResponseBody{Message: "done"})
}
