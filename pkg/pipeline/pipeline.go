package pipeline

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func TriggerPipeline(username, password, repoOwner, repoSlug, pipelineKey string) error {
	client := &http.Client{}
	data := map[string]interface{}{
		"target": map[string]string{
			"ref_type": "branch",
			"type":     "pipeline_ref_target",
			"ref_name": "master",
		},
	}
	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://api.bitbucket.org/2.0/repositories/%s/%s/pipelines/", repoOwner, repoSlug), bytes.NewReader(body))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)
	if resp.StatusCode != 201 {
		return fmt.Errorf("received non-201 status: %s", resp.Status)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBody))
	return nil
}
