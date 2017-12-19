//
// Copyright (c) 2017 Red Hat, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Red Hat trademarks are not licensed under Apache License, Version 2.
// No permission is granted to use or replicate Red Hat trademarks that
// are incorporated in this software or its documentation.
//

package broker

import (
	"encoding/json"

	"github.com/openshift/ansible-service-broker/pkg/apb"
	"github.com/openshift/ansible-service-broker/pkg/metrics"
)

// BindingJob - Job to provision
type BindingJob struct {
	serviceInstance *apb.ServiceInstance
}

// BindingMsg - Message to be returned from the binding job
type BindingMsg struct {
	InstanceUUID string `json:"instance_uuid"`
	JobToken     string `json:"job_token"`
	SpecID       string `json:"spec_id"`
	PodName      string `json:"podname"`
	Msg          string `json:"msg"`
	Error        string `json:"error"`
}

// Render - Display the binding message.
func (m BindingMsg) Render() string {
	render, _ := json.Marshal(m)
	return string(render)
}

// NewBindingJob - Create a new binding job.
func NewBindingJob(serviceInstance *apb.ServiceInstance) *BindingJob {
	return &BindingJob{
		serviceInstance: serviceInstance,
	}
}

// Run - run the binding job.
func (p *BindingJob) Run(token string, msgBuffer chan<- WorkMsg) {
	metrics.BindingJobStarted()
	var podName, extCreds string
	var err error

	// podName, extCreds, err := apb.Provision(p.serviceInstance)

	if err != nil {
		log.Error("broker::Binding error occurred.")
		log.Errorf("%s", err.Error())

		// send error message
		// can't have an error type in a struct you want marshalled
		// https://github.com/golang/go/issues/5161
		msgBuffer <- BindingMsg{InstanceUUID: p.serviceInstance.ID.String(),
			JobToken: token, SpecID: p.serviceInstance.Spec.ID, PodName: "", Msg: "", Error: err.Error()}
		return
	}

	// send creds
	jsonmsg, err := json.Marshal(extCreds)
	if err != nil {
		msgBuffer <- BindingMsg{InstanceUUID: p.serviceInstance.ID.String(),
			JobToken: token, SpecID: p.serviceInstance.Spec.ID, PodName: "", Msg: "", Error: err.Error()}
		return
	}

	msgBuffer <- BindingMsg{InstanceUUID: p.serviceInstance.ID.String(),
		JobToken: token, SpecID: p.serviceInstance.Spec.ID, PodName: podName, Msg: string(jsonmsg), Error: ""}
}
