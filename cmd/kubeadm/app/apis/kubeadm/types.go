/*
Copyright 2016 The Kubernetes Authors.

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

package kubeadm

import (
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type MasterConfiguration struct {
	metav1.TypeMeta

	API                API
	Etcd               Etcd
	Networking         Networking
	KubernetesVersion  string
	CloudProvider      string
	NodeName           string
	AuthorizationModes []string

	Token    string
	TokenTTL time.Duration

	APIServerExtraArgs         map[string]string
	ControllerManagerExtraArgs map[string]string
	SchedulerExtraArgs         map[string]string

	// APIServerCertSANs sets extra Subject Alternative Names for the API Server signing cert
	APIServerCertSANs []string
	// CertificatesDir specifies where to store or look for all required certificates
	CertificatesDir string

	// ImageRepository what container registry to pull control plane images from
	ImageRepository string
	// UnifiedControlPlaneImage specifies if a specific container image should be used for all control plane components
	UnifiedControlPlaneImage string

	// FeatureFlags enabled by the user
	FeatureFlags map[string]bool
}

type API struct {
	// AdvertiseAddress sets the address for the API server to advertise.
	AdvertiseAddress string
	// BindPort sets the secure port for the API Server to bind to
	BindPort int32
}

type TokenDiscovery struct {
	ID        string
	Secret    string
	Addresses []string
}

type Networking struct {
	ServiceSubnet string
	PodSubnet     string
	DNSDomain     string
}

type Etcd struct {
	Endpoints []string
	CAFile    string
	CertFile  string
	KeyFile   string
	DataDir   string
	ExtraArgs map[string]string
	// Image specifies which container image to use for running etcd. If empty, automatically populated by kubeadm using the image repository and default etcd version
	Image string
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type NodeConfiguration struct {
	metav1.TypeMeta

	CACertPath     string
	DiscoveryFile  string
	DiscoveryToken string
	// Currently we only pay attention to one api server but hope to support >1 in the future
	DiscoveryTokenAPIServers []string
	NodeName                 string
	TLSBootstrapToken        string
	Token                    string
}

func (cfg *MasterConfiguration) GetMasterEndpoint() string {
	return fmt.Sprintf("https://%s:%d", cfg.API.AdvertiseAddress, cfg.API.BindPort)
}
