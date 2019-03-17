# Google App Engine and Cloud SQL/Storage Demonstration

## What this app does
This all will allow students to comment on and view other people's resumes.

## How we did this
This app is created using NextJS/React for the frontend and Golang for the API. The database of choice will be Google Cloud SQL's MySQL offering. The blob storage provider (for resumes) of choice will be Google Cloud Storage.

## App Engine Setup
1. We will be using Golang as our language of choice. If you would like to get started with installation, visit [this link](https://golang.org/doc/install)
2. Install the gcloud SDK by following the instructions on this [page](https://cloud.google.com/sdk/). You can check your installation by running `gcloud version` and making sure that no errors pop up
3. Check out the `app.yaml` file in this [repo](https://github.com/GoogleCloudPlatform/golang-samples/tree/master/appengine/go11x/helloworld) to get an idea of what is needed for you to get started on App Engine (Yes, it is one line)
4. Create a `main.go` file
	```go
	package main

	import (
		"fmt"
		"log"
		"net/http"
		"os"
	)

	func main() {
		http.HandleFunc("/", indexHandler)

		port := os.Getenv("PORT")
		if port == "" {
			port = "8080"
			log.Printf("Defaulting to port %s", port)
		}

		log.Printf("Listening on port %s", port)
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
	}

	func indexHandler(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		fmt.Fprint(w, "Hello, World!")
	}
	```
5. Run `gcloud app deploy` in the same directory as your `app.yaml` and everything should deploy
6. Check out the results in the [console](https://console.cloud.google.com/appengine)
