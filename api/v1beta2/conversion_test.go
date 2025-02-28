/*
Copyright 2025 The Kubernetes Authors.

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

package v1beta2

import (
	"testing"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/cluster-api/api/v1beta1"
)

func TestConversion(t *testing.T) {
    g := NewWithT(t)

    v1beta2Cluster := &Cluster{
        Spec: ClusterSpec{
            ControlPlaneEndpoint: APIEndpoint{
                Host: "test-host",
                Port: 6443,
                ProxyURL: "http://proxy:3128",
            },
        },
    }

    v1beta1Cluster := &v1beta1.Cluster{}
    
    g.Expect(v1beta2Cluster.ConvertTo(v1beta1Cluster)).To(Succeed())
	g.Expect(v1beta1Cluster.Spec.ControlPlaneEndpoint.Host).To(Equal("test-host"))
    g.Expect(v1beta1Cluster.Spec.ControlPlaneEndpoint.Port).To(Equal(int32(6443)))

    newV1beta2Cluster := &Cluster{}
    g.Expect(newV1beta2Cluster.ConvertFrom(v1beta1Cluster)).To(Succeed())
	g.Expect(newV1beta2Cluster.Spec.ControlPlaneEndpoint.Host).To(Equal("test-host"))
    g.Expect(newV1beta2Cluster.Spec.ControlPlaneEndpoint.Port).To(Equal(int32(6443)))
    g.Expect(newV1beta2Cluster.Spec.ControlPlaneEndpoint.ProxyURL).To(Equal(""))
}