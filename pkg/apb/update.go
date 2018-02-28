//
// Copyright (c) 2018 Red Hat, Inc.
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

package apb

import (
	"github.com/openshift/ansible-service-broker/pkg/runtime"
)

// Update - will run the abp with the provision action.
func (e *executor) Update(instance *ServiceInstance) <-chan StatusMessage {
	log.Notice("============================================================")
	log.Notice("                       UPDATING                             ")
	log.Notice("============================================================")
	log.Noticef("Spec.ID: %s", instance.Spec.ID)
	log.Noticef("Spec.Name: %s", instance.Spec.FQName)
	log.Noticef("Spec.Image: %s", instance.Spec.Image)
	log.Noticef("Spec.Description: %s", instance.Spec.Description)
	log.Notice("============================================================")

	go func() {
		e.actionStarted()
		err := e.provisionOrUpdate(executionMethodUpdate, instance)
		if err != nil {
			log.Errorf("Update APB error: %v", err)
			e.actionFinishedWithError(err)
			return
		}
		if e.extractedCredentials != nil {
			labels := map[string]string{"apbAction": string(executionMethodUpdate), "apbName": instance.Spec.FQName}
			err := runtime.Provider.UpdateExtractedCredential(instance.ID.String(), clusterConfig.Namespace, e.extractedCredentials.Credentials, labels)
			if err != nil {
				log.Errorf("apb::%v error occurred - %v", executionMethodUpdate, err)
				e.actionFinishedWithError(err)
				return
			}
		}
		e.actionFinishedWithSuccess()
	}()

	return e.statusChan
}
