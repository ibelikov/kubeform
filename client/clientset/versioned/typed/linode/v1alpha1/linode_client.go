/*
Copyright The Kubeform Authors.

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

package v1alpha1

import (
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
	v1alpha1 "kubeform.dev/kubeform/apis/linode/v1alpha1"
	"kubeform.dev/kubeform/client/clientset/versioned/scheme"
)

type LinodeV1alpha1Interface interface {
	RESTClient() rest.Interface
	DomainsGetter
	DomainRecordsGetter
	ImagesGetter
	InstancesGetter
	NodebalancersGetter
	NodebalancerConfigsGetter
	NodebalancerNodesGetter
	RdnsesGetter
	SshkeysGetter
	StackscriptsGetter
	TokensGetter
	VolumesGetter
}

// LinodeV1alpha1Client is used to interact with features provided by the linode.kubeform.com group.
type LinodeV1alpha1Client struct {
	restClient rest.Interface
}

func (c *LinodeV1alpha1Client) Domains(namespace string) DomainInterface {
	return newDomains(c, namespace)
}

func (c *LinodeV1alpha1Client) DomainRecords(namespace string) DomainRecordInterface {
	return newDomainRecords(c, namespace)
}

func (c *LinodeV1alpha1Client) Images(namespace string) ImageInterface {
	return newImages(c, namespace)
}

func (c *LinodeV1alpha1Client) Instances(namespace string) InstanceInterface {
	return newInstances(c, namespace)
}

func (c *LinodeV1alpha1Client) Nodebalancers(namespace string) NodebalancerInterface {
	return newNodebalancers(c, namespace)
}

func (c *LinodeV1alpha1Client) NodebalancerConfigs(namespace string) NodebalancerConfigInterface {
	return newNodebalancerConfigs(c, namespace)
}

func (c *LinodeV1alpha1Client) NodebalancerNodes(namespace string) NodebalancerNodeInterface {
	return newNodebalancerNodes(c, namespace)
}

func (c *LinodeV1alpha1Client) Rdnses(namespace string) RdnsInterface {
	return newRdnses(c, namespace)
}

func (c *LinodeV1alpha1Client) Sshkeys(namespace string) SshkeyInterface {
	return newSshkeys(c, namespace)
}

func (c *LinodeV1alpha1Client) Stackscripts(namespace string) StackscriptInterface {
	return newStackscripts(c, namespace)
}

func (c *LinodeV1alpha1Client) Tokens(namespace string) TokenInterface {
	return newTokens(c, namespace)
}

func (c *LinodeV1alpha1Client) Volumes(namespace string) VolumeInterface {
	return newVolumes(c, namespace)
}

// NewForConfig creates a new LinodeV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*LinodeV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &LinodeV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new LinodeV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *LinodeV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new LinodeV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *LinodeV1alpha1Client {
	return &LinodeV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *LinodeV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
