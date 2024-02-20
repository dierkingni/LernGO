package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/add", addHandler)
	mux.HandleFunc("/subtract", subtractHandler)

	http.ListenAndServe(":8080", mux)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the calculator server!\n")
	fmt.Fprintf(w, "Available endpoints:\n")
	fmt.Fprintf(w, "- /add?num1=value1&num2=value2: Performs addition\n")
	fmt.Fprintf(w, "- /subtract?num1=value1&num2=value2: Performs subtraction\n")
}

type operationResult struct {
	Operation string  `json:"operation"`
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Result    float64 `json:"result"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	num1, err1 := parseFloatQuery(r, "num1")
	num2, err2 := parseFloatQuery(r, "num2")
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	result := num1 + num2

	jsonResponse := operationResult{
		Operation: "Addition",
		Num1:      num1,
		Num2:      num2,
		Result:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	num1, err1 := parseFloatQuery(r, "num1")
	num2, err2 := parseFloatQuery(r, "num2")
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	result := num1 - num2

	jsonResponse := operationResult{
		Operation: "Subtraction",
		Num1:      num1,
		Num2:      num2,
		Result:    result,
	}

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(jsonResponse)
	if err != nil {
		http.Error(w, "Error encoding JSON response", http.StatusInternalServerError)
		return
	}
}

func parseFloatQuery(r *http.Request, param string) (float64, error) {
	value := r.URL.Query().Get(param)
	if value == "" {
		return 0, fmt.Errorf("missing query parameter %s", param)
	}

	num, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid value for %s: %w", param, err)
	}

	return num, nil
}
