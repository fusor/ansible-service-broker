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

package clients

import (
	"sync"

	etcd "github.com/coreos/etcd/client"
	k8s "k8s.io/kubernetes/pkg/client/clientset_generated/clientset"
)

var instances struct {
	Etcd       etcd.Client
	Kubernetes *k8s.Clientset
}

var once struct {
	Etcd       sync.Once
	Kubernetes sync.Once
}
