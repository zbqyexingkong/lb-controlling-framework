/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta1 "git.tencent.com/tke/lb-controlling-framework/pkg/client-go/clientset/versioned/typed/lbcf.tke.cloud.tencent.com/v1beta1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeLbcfV1beta1 struct {
	*testing.Fake
}

func (c *FakeLbcfV1beta1) BackendGroups(namespace string) v1beta1.BackendGroupInterface {
	return &FakeBackendGroups{c, namespace}
}

func (c *FakeLbcfV1beta1) LoadBalancers(namespace string) v1beta1.LoadBalancerInterface {
	return &FakeLoadBalancers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeLbcfV1beta1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
