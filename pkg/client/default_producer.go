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
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type DefaultProducer struct {
	KubeClient    kubernetes.Interface
	DynamicClient dynamic.Interface
	GeneralClient client.Client
}

func (p *DefaultProducer) ForKubernetes() kubernetes.Interface {
	return p.KubeClient
}

func (p *DefaultProducer) ForDynamic() dynamic.Interface {
	return p.DynamicClient
}

func (p *DefaultProducer) ForGeneral() client.Client {
	return p.GeneralClient
}
