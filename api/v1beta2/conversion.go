/*
Copyright 2021 The Kubernetes Authors.

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
	"sigs.k8s.io/cluster-api/api/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (*Cluster) Hub()                {}
func (*ClusterList) Hub()            {}
func (*ClusterClass) Hub()           {}
func (*ClusterClassList) Hub()       {}
func (*Machine) Hub()                {}
func (*MachineList) Hub()            {}
func (*MachineSet) Hub()             {}
func (*MachineSetList) Hub()         {}
func (*MachineDeployment) Hub()      {}
func (*MachineDeploymentList) Hub()  {}
func (*MachineHealthCheck) Hub()     {}
func (*MachineHealthCheckList) Hub() {}

// ConvertTo convertit de v1beta2 vers v1beta1 (Hub)
func (src *Cluster) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1beta1.Cluster)

	// Conversion basique
	dst.ObjectMeta = src.ObjectMeta

	// Conversion du ControlPlaneEndpoint avec proxyURL
	dst.Spec.ControlPlaneEndpoint.Host = src.Spec.ControlPlaneEndpoint.Host
	dst.Spec.ControlPlaneEndpoint.Port = src.Spec.ControlPlaneEndpoint.Port

	return nil
}

// ConvertFrom convertit de v1beta1 (Hub) vers v1beta2
func (dst *Cluster) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.Cluster)

	// Conversion basique
	dst.ObjectMeta = src.ObjectMeta

	// Conversion du ControlPlaneEndpoint avec proxyURL
	dst.Spec.ControlPlaneEndpoint.Host = src.Spec.ControlPlaneEndpoint.Host
	dst.Spec.ControlPlaneEndpoint.Port = src.Spec.ControlPlaneEndpoint.Port
	dst.Spec.ControlPlaneEndpoint.ProxyURL = ""

	return nil
}
