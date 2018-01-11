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
	"github.com/openshift/ansible-service-broker/pkg/dao"
	"github.com/openshift/ansible-service-broker/pkg/metrics"
)

// UnbindingWorkSubscriber - Listen for binding messages
type UnbindingWorkSubscriber struct {
	dao       *dao.Dao
	msgBuffer <-chan JobMsg
}

// NewUnbindingWorkSubscriber - Creates a new work subscriber
func NewUnbindingWorkSubscriber(dao *dao.Dao) *UnbindingWorkSubscriber {
	return &UnbindingWorkSubscriber{dao: dao}
}

// Subscribe - will start a work subscriber listening for bind job messages
func (b *UnbindingWorkSubscriber) Subscribe(msgBuffer <-chan JobMsg) {
	b.msgBuffer = msgBuffer

	go func() {
		log.Info("Listening for binding messages")
		for {
			msg := <-msgBuffer
			var bmsg *JobMsg
			var extCreds *apb.ExtractedCredentials
			metrics.UnbindingJobFinished()

			log.Debug("Processed binding message from buffer")

			json.Unmarshal([]byte(msg.Render()), &bmsg)

			if bmsg.Error != "" {
				log.Errorf("Unbinding job reporting error: %s", bmsg.Error)
				b.dao.SetState(bmsg.InstanceUUID, apb.JobState{
					Token:   bmsg.JobToken,
					State:   apb.StateFailed,
					Podname: bmsg.PodName,
					Method:  apb.JobMethodBind,
				})
			} else if bmsg.Msg == "" {
				b.dao.SetState(bmsg.InstanceUUID, apb.JobState{
					Token:   bmsg.JobToken,
					State:   apb.StateInProgress,
					Podname: bmsg.PodName,
					Method:  apb.JobMethodBind,
				})
			} else {
				json.Unmarshal([]byte(bmsg.Msg), &extCreds)
				b.dao.SetState(bmsg.InstanceUUID, apb.JobState{
					Token:   bmsg.JobToken,
					State:   apb.StateSucceeded,
					Podname: bmsg.PodName,
					Method:  apb.JobMethodBind,
				})
				b.dao.SetExtractedCredentials(bmsg.BindingUUID, extCreds)
			}
		}
	}()
}
