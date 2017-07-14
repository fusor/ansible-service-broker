package apb

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	logging "github.com/op/go-logging"
)

var timeoutFreq = 6    // Seconds
var totalTimeout = 900 // 15min

// ExtractCredentials - Extract credentials from pod in a certain namespace.
func ExtractCredentials(
	podname string, namespace string, log *logging.Logger,
) (*ExtractedCredentials, error) {
	log.Debug("Calling monitorOutput on " + podname)
	bindOutput, err := monitorOutput(namespace, podname, log)
	if bindOutput == nil {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return buildExtractedCredentials(bindOutput)
}

func monitorOutput(namespace string, podname string, log *logging.Logger) ([]byte, error) {
	// TODO: Error handling here
	// It would also be nice to gather the script output that exec runs
	// instead of only getting the credentials

	for r := 1; r <= credentialExtRetries; r++ {
		output, err := RunCommand("oc", "exec", podname, gatherCredentialsCMD, "--namespace="+namespace)
		if err != nil {
			// Since we combine stderr and stdout in RunCommand, log
			// output of RunCommand as Info
			log.Info(err.Error())
		}

		stillWaiting := strings.Contains(string(output), "ContainerCreating") ||
			strings.Contains(string(output), "NotFound") ||
			strings.Contains(string(output), "container not found")
		podCompleted := strings.Contains(string(output), "current phase is Succeeded") ||
			strings.Contains(string(output), "cannot exec into a container in a completed pod")

		// TODO: Replace the string parsing by passing around the pod
		// object and checking its status
		if stillWaiting {
			log.Info("[%s] Retry attempt %d: Waiting for container to start", podname, r)
		} else if podCompleted {
			log.Notice("[%s] APB completed", podname)
			return nil, nil
		} else if strings.Contains(string(output), "BIND_CREDENTIALS") {
			log.Notice("[%s] Bind credentials found", podname)
			return output, nil
		}

		log.Warning("[%s] Retry attempt %d: exec into %s failed", podname, r, podname)
		time.Sleep(time.Duration(credentialExtInterval) * time.Second)
	}
	timeout := fmt.Sprintf("[%s] ExecTimeout: Failed to gather bind credentials after %d retries", podname, credentialExtRetries)
	return nil, errors.New(timeout)
}

func buildExtractedCredentials(output []byte) (*ExtractedCredentials, error) {
	result, err := decodeOutput(output)
	if err != nil {
		return nil, err
	}

	creds := make(map[string]interface{})
	for k, v := range result {
		creds[k] = v
	}

	return &ExtractedCredentials{Credentials: creds}, nil
}

func decodeOutput(output []byte) (map[string]string, error) {
	// Possible return states
	// 1) nil, nil -> No credentials found, no errors occurred. Valid.
	// 2) creds, nil -> Credentials found, no errors occurred. Valid.
	// 3) nil, err -> Credentials found, no errors occurred. Error state.
	str := string(output)

	startIdx := strings.Index(str, "<BIND_CREDENTIALS>")
	startOffset := startIdx + len("<BIND_CREDENTIALS>")
	endIdx := strings.Index(str, "</BIND_CREDENTIALS>")

	if startIdx < 0 || endIdx < 0 {
		startIdx = strings.Index(str, "<BIND_ERROR>")
		startOffset := startIdx + len("<BIND_ERROR>")
		endIdx := strings.Index(str, "</BIND_ERROR>")
		if startIdx > -1 && endIdx > -1 {
			// Case 3, error reported
			return nil, errors.New(str[startOffset:endIdx])
		}

		if strings.Contains(str, "image can't be pulled") {
			return nil, errors.New("image can't be pulled")
		} else if strings.Contains(str, "FAILED! =>") {
			return nil, errors.New("provision failed, INSERT MESSAGE HERE")
		} else {
			// Case 1, no creds found, no errors occurred
			return nil, nil
		}
	}

	decodedjson, err := base64.StdEncoding.DecodeString(str[startOffset:endIdx])
	if err != nil {
		return nil, err
	}

	decoded := make(map[string]string)
	json.Unmarshal(decodedjson, &decoded)
	// Case 2, creds successfully found and decoded
	return decoded, nil
}
