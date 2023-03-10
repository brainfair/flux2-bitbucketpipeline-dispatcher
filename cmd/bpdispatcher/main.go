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
	AccessToken, ok1 = os.LookupEnv("ACCESS_TOKEN")
	RepoOwner, ok2   = os.LookupEnv("REPO_OWNER")
	RepoSlug, ok3    = os.LookupEnv("REPO_SLUG")
	PipelineRef, ok4 = os.LookupEnv("PIPELINE_REF")
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
		Revision string `json:"revision"`
		Summary  string `json:"summary"`
	} `json:"metadata"`
	Severity            string `json:"severity"`
	Reason              string `json:"reason"`
	Message             string `json:"message"`
	ReportingController string `json:"reportingController"`
	ReportingInstance   string `json:"reportingInstance"`
	Timestamp           string `json:"timestamp"`
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	var pipeVariables []map[string]string
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
	fmt.Println("Summary:", data.Metadata.Summary)
	fmt.Println("Severity:", data.Severity)
	fmt.Println("Reason:", data.Reason)
	fmt.Println("Message:", data.Message)
	fmt.Println("ReportingController:", data.ReportingController)
	fmt.Println("ReportingInstance:", data.ReportingInstance)
	fmt.Println("Timestamp:", data.Timestamp)
	fmt.Fprint(w, "JSON received and parsed!")

	pipeVariables = append(pipeVariables, map[string]string{"key": "KIND", "value": data.InvolvedObject.Kind})
	pipeVariables = append(pipeVariables, map[string]string{"key": "NAME", "value": data.InvolvedObject.Name})
	pipeVariables = append(pipeVariables, map[string]string{"key": "NAMESPACE", "value": data.InvolvedObject.Namespace})
	pipeVariables = append(pipeVariables, map[string]string{"key": "REVISION", "value": data.Metadata.Revision})
	pipeVariables = append(pipeVariables, map[string]string{"key": "SUMMARY", "value": data.Metadata.Summary})
	pipeVariables = append(pipeVariables, map[string]string{"key": "SEVERITY", "value": data.Severity})
	pipeVariables = append(pipeVariables, map[string]string{"key": "REASON", "value": data.Reason})
	pipeVariables = append(pipeVariables, map[string]string{"key": "MESSAGE", "value": data.Message})

	if err := pipeline.TriggerPipeline(AccessToken, RepoOwner, RepoSlug, PipelineRef, pipeVariables); err != nil {
		log.Println("Error:", err)
	}
}

func livenessProbe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func readinessProbe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func checkEnv(envVar string, ok bool) {
	if !ok {
		fmt.Printf("%s environment variable not set\n", envVar)
		os.Exit(1)
	}
}

func main() {
	checkEnv("ACCESS_TOKEN", ok1)
	checkEnv("REPO_OWNER", ok2)
	checkEnv("REPO_SLUG", ok3)
	checkEnv("PIPELINE_REF", ok4)
	http.HandleFunc("/webhook", handleWebhook)
	http.HandleFunc("/healthz", livenessProbe)
	http.HandleFunc("/ready", readinessProbe)

	http.ListenAndServe(":8000", nil)
}
