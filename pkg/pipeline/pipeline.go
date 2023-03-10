package pipeline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func TriggerPipeline(accessToken, repoOwner, repoSlug, pipelineRef string, pipeVariables []map[string]string) error {
	client := &http.Client{}
	data := map[string]interface{}{
		"target": map[string]interface{}{
			"ref_type": "branch",
			"type":     "pipeline_ref_target",
			"ref_name": pipelineRef,
			"selector": map[string]string{
				"type":    "custom",
				"pattern": "pr-promotion",
			},
		},
		"variables": pipeVariables,
	}
	body, err := json.Marshal(data)
	if err != nil {
		log.Println("Error while marshalling json:", err)
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s/pipelines/", repoOwner, repoSlug), bytes.NewReader(body))
	if err != nil {
		log.Println("Error while creating request:", err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error while triggering pipeline:", err)
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		log.Println("Error while triggering pipeline, status code:", resp.StatusCode)
		return fmt.Errorf("received non-201 status: %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading response body:", err)
		return err
	}
	fmt.Println(string(respBody))
	return nil
}
