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

package broker

import (
	"encoding/json"

	"github.com/openshift/ansible-service-broker/pkg/apb"
	"github.com/openshift/ansible-service-broker/pkg/dao"
	"github.com/openshift/ansible-service-broker/pkg/metrics"
)

// DeprovisionJob - Job to deprovision.
type DeprovisionJob struct {
	serviceInstance  *apb.ServiceInstance
	skipApbExecution bool
	dao              *dao.Dao
}

// DeprovisionMsg - Message returned for a deprovison job.
type DeprovisionMsg struct {
	InstanceUUID string `json:"instance_uuid"`
	PodName      string `json:"podname"`
	JobToken     string `json:"job_token"`
	SpecID       string `json:"spec_id"`
	Error        string `json:"error"`
}

// Render - render the message
func (m DeprovisionMsg) Render() string {
	render, _ := json.Marshal(m)
	return string(render)
}

// NewDeprovisionJob - Create a deprovision job.
func NewDeprovisionJob(serviceInstance *apb.ServiceInstance,
	skipApbExecution bool, dao *dao.Dao,
) *DeprovisionJob {
	return &DeprovisionJob{
		serviceInstance:  serviceInstance,
		skipApbExecution: skipApbExecution,
		dao:              dao,
	}
}

// Run - will run the deprovision job.
func (p *DeprovisionJob) Run(token string, msgBuffer chan<- WorkMsg) {
	metrics.DeprovisionJobStarted()

	if p.skipApbExecution {
		log.Debug("skipping deprovision and sending complete msg to channel")
		msgBuffer <- DeprovisionMsg{InstanceUUID: p.serviceInstance.ID.String(), PodName: "",
			JobToken: token, SpecID: p.serviceInstance.Spec.ID, Error: ""}
		return
	}

	podName, err := apb.Deprovision(p.serviceInstance)
	if err != nil {
		log.Error("broker::Deprovision error occurred.")
		log.Errorf("%s", err.Error())
		msgBuffer <- DeprovisionMsg{InstanceUUID: p.serviceInstance.ID.String(), PodName: podName,
			JobToken: token, SpecID: p.serviceInstance.Spec.ID, Error: err.Error()}
		return
	}

	log.Debug("sending deprovision complete msg to channel")
	msgBuffer <- DeprovisionMsg{InstanceUUID: p.serviceInstance.ID.String(), PodName: podName,
		JobToken: token, SpecID: p.serviceInstance.Spec.ID, Error: ""}
}
