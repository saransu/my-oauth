package calculator

import (
	"math/rand"
	"my-oauth-server/internal/utils"
	"net/http"
	"sync"
)

func RegisterController(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/v1/perf-test", performanceTest)
	mux.HandleFunc("GET /api/v1/perf-test-2", performanceTest2)
}

func heavyDutyCalculation(count int) int {
	result := 0
	for i := range count {
		if i%2 == 0 {
			result += int(rand.Int31())
		} else {
			result -= int(rand.Int31())
		}
	}
	return result
}

func heavyDutyCalculationChan(count int, ch chan<- int, wg *sync.WaitGroup) {
	result := 0
	for i := range count {
		if i%2 == 0 {
			result += int(rand.Int31())
		} else {
			result -= int(rand.Int31())
		}
	}

	ch <- result
	wg.Done()
	return
}

func performanceTest(w http.ResponseWriter, _ *http.Request) {
	results := []int{}
	for range 10 {
		results = append(results, heavyDutyCalculation(100000000))
	}

	w.WriteHeader(http.StatusOK)
	utils.WriteResponseBody(w, utils.ResponseBody{Data: results, Message: "Done ja"})
}

func performanceTest2(w http.ResponseWriter, _ *http.Request) {
	results := []int{}
	resultsCh := make(chan int)
	var wg sync.WaitGroup

	go func() {
		for v := range resultsCh {
			results = append(results, v)
		}
	}()

	wg.Add(10)
	for range 10 {
		go heavyDutyCalculationChan(100000000, resultsCh, &wg)
	}
	wg.Wait()
	close(resultsCh)

	w.WriteHeader(http.StatusOK)
	utils.WriteResponseBody(w, utils.ResponseBody{Data: results, Message: "Done ja"})
}
