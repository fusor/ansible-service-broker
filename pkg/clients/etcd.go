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
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/coreos/etcd/version"
	"github.com/openshift/ansible-service-broker/pkg/config"

	logging "github.com/op/go-logging"

	etcd "github.com/coreos/etcd/client"
)

// EtcdConfig - Etcd configuration
type EtcdConfig struct {
	EtcdHost string `yaml:"etcd_host"`
	EtcdPort string `yaml:"etcd_port"`
}

// GetEtcdVersion - Connects to ETCD cluster and retrieves server/version info
func GetEtcdVersion(ec EtcdConfig) (string, string, error) {
	// The next etcd release (1.4) will have client.GetVersion()
	// We'll use this to test our etcd connection for now
	etcdURL := fmt.Sprintf("http://%s:%s/version", ec.EtcdHost, ec.EtcdPort)
	resp, err := http.Get(etcdURL)
	if err != nil {
		return "", "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	switch resp.StatusCode {
	case http.StatusOK:
		var vresp version.Versions
		if err := json.Unmarshal(body, &vresp); err != nil {
			return "", "", err
		}
		return vresp.Server, vresp.Cluster, nil
	default:
		var connectErr error
		if err := json.Unmarshal(body, &connectErr); err != nil {
			return "", "", err
		}
		return "", "", connectErr
	}
}

// Etcd - Create a new etcd client if needed, returns reference
func Etcd(config *config.Config, log *logging.Logger) (etcd.Client, error) {
	errMsg := "Something went wrong intializing etcd client!"
	once.Etcd.Do(func() {
		client, err := newEtcd(config, log)
		if err != nil {
			log.Error(errMsg)
			// NOTE: Looking to leverage panic recovery to gracefully handle this
			// with things like retries or better intelligence, but the environment
			// is probably in a unrecoverable state as far as the broker is concerned,
			// and demands the attention of an operator.
			panic(err.Error())
		}
		instances.Etcd = client
	})

	if instances.Etcd == nil {
		return nil, errors.New("Etcd client instance is nil")
	}

	return instances.Etcd, nil
}

func newEtcd(config *config.Config, log *logging.Logger) (etcd.Client, error) {
	host := config.GetString("dao.etcd_host")
	port := config.GetString("dao.etcd_port")
	if host == "" {
		return nil, errors.New("invalid etcd host")
	}
	endpoints := []string{etcdEndpoint(host, port)}

	log.Info("== ETCD CX ==")
	log.Infof("EtcdHost: %s", host)
	log.Infof("EtcdPort: %s", port)
	log.Infof("Endpoints: %v", endpoints)

	etcdClient, err := etcd.New(etcd.Config{
		Endpoints:               endpoints,
		Transport:               etcd.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	})
	if err != nil {
		return nil, err
	}

	return etcdClient, err
}

func etcdEndpoint(host string, port string) string {
	return fmt.Sprintf("http://%s:%s", host, port)
}
