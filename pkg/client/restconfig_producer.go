/*
SPDX-License-Identifier: Apache-2.0

Copyright Contributors to the Submariner project.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package client

import (
	"github.com/pkg/errors"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func NewProducerFromRestConfig(config *rest.Config) (Producer, error) {
	var err error
	p := &DefaultProducer{}

	p.KubeClient, err = kubernetes.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "error creating kube client")
	}

	p.DynamicClient, err = dynamic.NewForConfig(config)
	if err != nil {
		return nil, errors.Wrap(err, "error creating dynamic client")
	}

	p.GeneralClient, err = client.New(config, client.Options{})
	if err != nil {
		return nil, errors.Wrap(err, "error creating controller client")
	}

	return p, nil
}
