/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	"fmt"

	"kubedb.dev/apimachinery/apis"
	"kubedb.dev/apimachinery/apis/autoscaling"
	"kubedb.dev/apimachinery/crds"

	"kmodules.xyz/client-go/apiextensions"
)

func (s SinglestoreAutoscaler) CustomResourceDefinition() *apiextensions.CustomResourceDefinition {
	return crds.MustCustomResourceDefinition(SchemeGroupVersion.WithResource(ResourcePluralSinglestoreAutoscaler))
}

var _ apis.ResourceInfo = &SinglestoreAutoscaler{}

func (s SinglestoreAutoscaler) ResourceFQN() string {
	return fmt.Sprintf("%s.%s", ResourcePluralSinglestoreAutoscaler, autoscaling.GroupName)
}

func (s SinglestoreAutoscaler) ResourceShortCode() string {
	return ResourceCodeSinglestoreAutoscaler
}

func (s SinglestoreAutoscaler) ResourceKind() string {
	return ResourceKindSinglestoreAutoscaler
}

func (s SinglestoreAutoscaler) ResourceSingular() string {
	return ResourceSingularSinglestoreAutoscaler
}

func (s SinglestoreAutoscaler) ResourcePlural() string {
	return ResourcePluralSinglestoreAutoscaler
}

func (s SinglestoreAutoscaler) ValidateSpecs() error {
	return nil
}

var _ StatusAccessor = &SinglestoreAutoscaler{}

func (s *SinglestoreAutoscaler) GetStatus() AutoscalerStatus {
	return s.Status
}

func (s *SinglestoreAutoscaler) SetStatus(m AutoscalerStatus) {
	s.Status = m
}
