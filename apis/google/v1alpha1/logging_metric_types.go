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

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type LoggingMetric struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LoggingMetricSpec   `json:"spec,omitempty"`
	Status            LoggingMetricStatus `json:"status,omitempty"`
}

type LoggingMetricSpecBucketOptionsExplicitBuckets struct {
	// +optional
	Bounds []float64 `json:"bounds,omitempty" tf:"bounds,omitempty"`
}

type LoggingMetricSpecBucketOptionsExponentialBuckets struct {
	// +optional
	GrowthFactor int64 `json:"growthFactor,omitempty" tf:"growth_factor,omitempty"`
	// +optional
	NumFiniteBuckets int64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`
	// +optional
	Scale float64 `json:"scale,omitempty" tf:"scale,omitempty"`
}

type LoggingMetricSpecBucketOptionsLinearBuckets struct {
	// +optional
	NumFiniteBuckets int64 `json:"numFiniteBuckets,omitempty" tf:"num_finite_buckets,omitempty"`
	// +optional
	Offset float64 `json:"offset,omitempty" tf:"offset,omitempty"`
	// +optional
	Width int64 `json:"width,omitempty" tf:"width,omitempty"`
}

type LoggingMetricSpecBucketOptions struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExplicitBuckets []LoggingMetricSpecBucketOptionsExplicitBuckets `json:"explicitBuckets,omitempty" tf:"explicit_buckets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ExponentialBuckets []LoggingMetricSpecBucketOptionsExponentialBuckets `json:"exponentialBuckets,omitempty" tf:"exponential_buckets,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	LinearBuckets []LoggingMetricSpecBucketOptionsLinearBuckets `json:"linearBuckets,omitempty" tf:"linear_buckets,omitempty"`
}

type LoggingMetricSpecMetricDescriptorLabels struct {
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Key         string `json:"key" tf:"key"`
	// +optional
	ValueType string `json:"valueType,omitempty" tf:"value_type,omitempty"`
}

type LoggingMetricSpecMetricDescriptor struct {
	// +optional
	Labels     []LoggingMetricSpecMetricDescriptorLabels `json:"labels,omitempty" tf:"labels,omitempty"`
	MetricKind string                                    `json:"metricKind" tf:"metric_kind"`
	// +optional
	Unit      string `json:"unit,omitempty" tf:"unit,omitempty"`
	ValueType string `json:"valueType" tf:"value_type"`
}

type LoggingMetricSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	BucketOptions []LoggingMetricSpecBucketOptions `json:"bucketOptions,omitempty" tf:"bucket_options,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Filter      string `json:"filter" tf:"filter"`
	// +optional
	LabelExtractors map[string]string `json:"labelExtractors,omitempty" tf:"label_extractors,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	MetricDescriptor []LoggingMetricSpecMetricDescriptor `json:"metricDescriptor" tf:"metric_descriptor"`
	Name             string                              `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	ValueExtractor string `json:"valueExtractor,omitempty" tf:"value_extractor,omitempty"`
}

type LoggingMetricStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LoggingMetricSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LoggingMetricList is a list of LoggingMetrics
type LoggingMetricList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LoggingMetric CRD objects
	Items []LoggingMetric `json:"items,omitempty"`
}