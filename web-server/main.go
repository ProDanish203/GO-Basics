package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// const myUrl = "https://web-veritas.com/about?mode=admin"

func main() {
	// response, err := http.Get(url)
	// showNilError(err)

	// fmt.Println("Status code: ", response)

	// defer response.Body.Close()
	// dataBytes, err := ioutil.ReadAll(response.Body)
	// showNilError(err)
	// fmt.Println("Data: ", string(dataBytes))

	// // Handling URL in Go
	// result, err := url.Parse(myUrl)
	// showNilError(err)
	// fmt.Println("Scheme: ", result.Scheme)
	// fmt.Println("Host: ", result.Host)
	// fmt.Println("Path: ", result.Path)
	// fmt.Println("Query: ", result.RawQuery)
	// fmt.Println("Fragment: ", result.Fragment)
	// fmt.Println("RequestURI: ", result.RequestURI())
	// fmt.Println("Opaque: ", result.Opaque)

	// qparams := result.Query()
	// fmt.Println("Query Params (Mode): ", qparams["mode"])

	// PerformGetRequest()

	// EncodeJSON()
	DecodeJSON()
}

func showNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func PerformGetRequest() {
	const myUrl = "http://localhost:8000"

	response, err := http.Get(myUrl)
	showNilError(err)

	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	fmt.Println("Content Type: ", response.Header.Get("Content-Type"))

	var responseString strings.Builder
	dataBytes, err := io.ReadAll(response.Body)
	byteCount, err := responseString.Write(dataBytes)
	showNilError(err)

	fmt.Println("Byte Count: ", byteCount)
	fmt.Println("Data: ", responseString.String())

	// fmt.Println("Data: ", string(dataBytes))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/formdata"

	formData := url.Values{}
	formData.Add("name", "John Doe")
	formData.Add("age", "25")

	response, err := http.PostForm(myUrl, formData)
	showNilError(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	dataBytes, err := io.ReadAll(response.Body)
	showNilError(err)
	fmt.Println("Data: ", string(dataBytes))

}

func PerformPostRequest() {
	const myUrl = "http://localhost:8000"

	reqBody := strings.NewReader(`
		{
			"name": "John Doe",
			"age": 25
		}
	`)

	response, err := http.Post(myUrl, "application/json", reqBody)
	showNilError(err)

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)

	dataBytes, err := io.ReadAll(response.Body)
	showNilError(err)
	fmt.Println("Data: ", string(dataBytes))
}

func EncodeJSON() {
	courses := []Course{
		{1, "Course 1", 100, []string{"tag1", "tag2"}},
		{2, "Course 2", 200, []string{"tag3", "tag4"}},
	}

	// Package this data into JSON
	finalJson, err := json.MarshalIndent(courses, "", "\t")
	showNilError(err)
	fmt.Println("JSON: ", string(finalJson))

}

type Course struct {
	ID    int      `json:"course_id"`
	Name  string   `json:"course_name"`
	Price int      `json:"course_price"`
	Tags  []string `json:"course_tags,omitempty"`
	// Password string `json:"-"`
}

func DecodeJSON() {
	jsonData := `
		[
			{
				"course_id": 1,
				"course_name": "Course 1",
				"course_price": 100,
				"course_tags": ["tag1", "tag2"]
			},
			{
				"course_id": 2,
				"course_name": "Course 2",
				"course_price": 200,
				"course_tags": ["tag3", "tag4"]
			}
		]
	`

	var courses []Course

	// Check validity of JSON
	isValid := json.Valid([]byte(jsonData))
	fmt.Println("Is JSON Valid: ", isValid)

	err := json.Unmarshal([]byte(jsonData), &courses)
	showNilError(err)

	fmt.Printf("Courses: %#v\n", courses)
}
