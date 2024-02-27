package p

/*
get the auth bearer token: gcloud auth print-identity-token
*/

/*
login: gcloud auth application-default login
*/

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	_ "example.com/cloudfunction/docs"

	httpSwagger "github.com/swaggo/http-swagger/v2"

	"cloud.google.com/go/storage"
)

// @title	Simple Calculator RestAPI
func CalculatorFunction(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/swagger/index.html" {
		httpSwagger.WrapHandler(w, r)
		return
	}
	if r.URL.Path == "/swagger/swagger-ui.css" {
		httpSwagger.WrapHandler(w, r)
		return
	}
	if r.URL.Path == "/swagger/swagger-ui-bundle.js" {
		httpSwagger.WrapHandler(w, r)
		return
	}
	if r.URL.Path == "/swagger/swagger-ui-standalone-preset.js" {
		httpSwagger.WrapHandler(w, r)
		return
	}
	if r.URL.Path == "/swagger/doc.json" {
		httpSwagger.WrapHandler(w, r)
		return
	}

	switch r.URL.Path {
	case "/add":
		addHandler(w, r)
	case "/subtract":
		subtractHandler(w, r)
	default:
		http.Error(w, "Endpoint not found", http.StatusNotFound)
	}
}

type operationResult struct {
	//I define my JSON-Structure like so
	Operation string  `json:"operation"`
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Result    float64 `json:"result"`
}

// Add godoc
//
//	@Summary		add
//	@Description	plus
//	@Accept			json
//	@Produce		plain
//	@Param			num1	query		int		true	"used for calc"
//	@Param			num2	query		int		true	"used for calc"
//	@Success		200		{integer}	string	"answer"
//
// @Router			/add [get]
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
	BucketManipulation(jsonResponse)
}

// similar to my addHandler

// Subtract godoc
//
//	@Summary		subtract
//	@Description	plus
//	@Accept			json
//	@Produce		plain
//	@Param			num1	query		int		true	"used for calc"
//	@Param			num2	query		int		true	"used for calc"
//	@Success		200		{integer}	string	"answer"
//
// @Router			/subtract [get]
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
	BucketManipulation(jsonResponse)
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

func BucketManipulation(operationResult operationResult) {
	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	projectID := "mms-clp-playground2402-a-i2ar"

	// Creates a client.
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Sets the name for the new bucket.
	bucketName := "calculator-bucket"

	// Holt den Bucket.
	bucket := client.Bucket(bucketName)
	if err := bucket.Create(ctx, projectID, nil); err != nil {
		//log.Fatalf("Failed to create bucket: %v", err)
	}

	// Create a new object based on the time
	timestamp := time.Now().Format("2006-01-02T15-04-05")

	obj := bucket.Object(timestamp + ".json") // Add ".json" extension

	// Convert the operationResult struct to a JSON byte slice
	jsonData, err := json.Marshal(operationResult)
	if err != nil {
		log.Fatal(err)
	}

	// Create a writer for the object
	writer := obj.NewWriter(ctx)
	_, err = writer.Write(jsonData)
	if err != nil {
		log.Fatal(err)
	}

	// Close the writer
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}
}
