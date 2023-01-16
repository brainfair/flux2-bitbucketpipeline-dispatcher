package main

import (
	"encoding/json"
	"flux2-bitbucketpipeline-dispatcher/pkg/pipeline"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	Username, ok1    = os.LookupEnv("USERNAME")
	Password, ok2    = os.LookupEnv("PASSWORD")
	RepoOwner, ok3   = os.LookupEnv("REPO_OWNER")
	RepoSlug, ok4    = os.LookupEnv("REPO_SLUG")
	PipelineKey, ok5 = os.LookupEnv("PIPELINE_KEY")
)

type Webhook struct {
	InvolvedObject struct {
		APIVersion string `json:"apiVersion"`
		Kind       string `json:"kind"`
		Name       string `json:"name"`
		Namespace  string `json:"namespace"`
		UID        string `json:"uid"`
	} `json:"involvedObject"`
	Metadata struct {
		Revision string `json:"kustomize.toolkit.fluxcd.io/revision"`
	} `json:"metadata"`
	Severity            string `json:"severity"`
	Reason              string `json:"reason"`
	Message             string `json:"message"`
	ReportingController string `json:"reportingController"`
	ReportingInstance   string `json:"reportingInstance"`
	Timestamp           string `json:"timestamp"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var data Webhook
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	fmt.Println("API Version:", data.InvolvedObject.APIVersion)
	fmt.Println("Kind:", data.InvolvedObject.Kind)
	fmt.Println("Name:", data.InvolvedObject.Name)
	fmt.Println("Namespace:", data.InvolvedObject.Namespace)
	fmt.Println("Revision:", data.Metadata.Revision)
	fmt.Println("Severity:", data.Severity)
	fmt.Println("Reason:", data.Reason)
	fmt.Println("Message:", data.Message)
	fmt.Println("ReportingController:", data.ReportingController)
	fmt.Println("ReportingInstance:", data.ReportingInstance)
	fmt.Println("Timestamp:", data.Timestamp)
	fmt.Fprint(w, "JSON received and parsed!")

	if err := pipeline.TriggerPipeline(Username, Password, RepoOwner, RepoSlug, PipelineKey); err != nil {
		log.Println("Error:", err)
	}
}

func main() {
	if !ok1 {
		fmt.Println("USERNAME environment variable not set")
	}
	if !ok2 {
		fmt.Println("PASSWORD environment variable not set")
	}
	if !ok3 {
		fmt.Println("REPO_OWNER environment variable not set")
	}
	if !ok4 {
		fmt.Println("REPO_SLUG environment variable not set")
	}
	if !ok5 {
		fmt.Println("PIPELINE_KEY environment variable not set")
	}
	http.HandleFunc("/webhook", handleWebhook)
	http.ListenAndServe(":8000", nil)
}
