package p

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CalculatorFunction(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/add":
		addHandler(w, r)
	case "/subtract":
		subtractHandler(w, r)
	default:
		http.Error(w, "Endpoint not found", http.StatusNotFound)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	//description of my endpoints and how to use them
	fmt.Fprintf(w, "Welcome to the calculator server!\n")
	fmt.Fprintf(w, "Available endpoints:\n")
	fmt.Fprintf(w, "- /add?num1=value1&num2=value2: Performs addition\n")
	fmt.Fprintf(w, "- /subtract?num1=value1&num2=value2: Performs subtraction\n")
}

type operationResult struct {
	//I define my JSON-Structure like so
	Operation string  `json:"operation"`
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Result    float64 `json:"result"`
}

func addHandler(w http.ResponseWriter, r *http.Request) {
	//setting the given path values
	num1, err1 := parseFloatQuery(r, "num1")
	num2, err2 := parseFloatQuery(r, "num2")
	if err1 != nil || err2 != nil {
		http.Error(w, "Invalid query parameters", http.StatusBadRequest)
		return
	}

	//performing the actual operation
	result := num1 + num2

	//Storing all that into my JSON
	jsonResponse := operationResult{
		Operation: "Addition",
		Num1:      num1,
		Num2:      num2,
		Result:    result,
	}

	//setting that the content type should be representated as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonResponse)
}

// similar to my addHandler
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
	json.NewEncoder(w).Encode(jsonResponse)
}

// helper function to get my given path param numbers
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
