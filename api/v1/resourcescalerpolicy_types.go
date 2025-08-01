/*
Copyright 2025.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourceScalerPolicySpec defines the desired state of ResourceScalerPolicy
type ResourceScalerPolicySpec struct {
	// TargetDeployment specifies the deployment to monitor and scale
	TargetDeployment TargetDeployment `json:"targetDeployment"`

	// ScalePolicy defines scaling behavior
	//+optional
	ScalePolicy *ScalePolicy `json:"scalePolicy,omitempty"`

	// ResourcePolicy defines resource adjustment behavior
	//+optional
	ResourcePolicy *ResourcePolicy `json:"resourcePolicy,omitempty"`

	// CircuitBreaker defines circuit breaker configuration
	//+optional
	CircuitBreaker *CircuitBreakerConfig `json:"circuitBreaker,omitempty"`

	// CostAwareness defines cost monitoring configuration
	//+optional
	CostAwareness *CostAwarenessConfig `json:"costAwareness,omitempty"`

	// MetricsConfig defines how metrics are collected
	//+optional
	MetricsConfig *MetricsConfig `json:"metricsConfig,omitempty"`

	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	// foo is an example field of ResourceScalerPolicy. Edit resourcescalerpolicy_types.go to remove/update
	// +optional
	Foo *string `json:"foo,omitempty"`
}

// TargetDeployment defines the target deployment to monitor
type TargetDeployment struct {
	// Name of the deployment
	Name string `json:"name"`

	// Namespace of the deployment
	//+optional
	Namespace string `json:"namespace,omitempty"`
}

// ScalePolicy defines scaling behavior
type ScalePolicy struct {
	// MinReplicas is the minimum number of replicas
	//+kubebuilder:validation:Minimum=1
	//+optional
	MinReplicas *int32 `json:"minReplicas,omitempty"`

	// MaxReplicas is the maximum number of replicas
	//+kubebuilder:validation:Minimum=1
	MaxReplicas int32 `json:"maxReplicas"`

	// ScaleUpCooldown is the cooldown period for scale up operations
	//+optional
	ScaleUpCooldown *metav1.Duration `json:"scaleUpCooldown,omitempty"`

	// ScaleDownCooldown is the cooldown period for scale down operations
	//+optional
	ScaleDownCooldown *metav1.Duration `json:"scaleDownCooldown,omitempty"`

	// ScaleUpStep defines how many replicas to add during scale up
	//+optional
	ScaleUpStep *int32 `json:"scaleUpStep,omitempty"`

	// ScaleDownStep defines how many replicas to remove during scale down
	//+optional
	ScaleDownStep *int32 `json:"scaleDownStep,omitempty"`
}

// ResourcePolicy defines resource adjustment behavior
type ResourcePolicy struct {
	// MemoryThreshold is the memory usage percentage that triggers scaling
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=100
	MemoryThreshold int32 `json:"memoryThreshold"`

	// CPUThreshold is the CPU usage percentage that triggers scaling
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=100
	//+optional
	CPUThreshold *int32 `json:"cpuThreshold,omitempty"`

	// AdjustmentMode defines how to handle resource pressure: scale, patch, or both
	//+kubebuilder:validation:Enum=scale;patch;both
	//+optional
	AdjustmentMode string `json:"adjustmentMode,omitempty"`

	// ResourceAdjustmentStep defines the increment for resource adjustments
	//+optional
	ResourceAdjustmentStep *ResourceAdjustmentStep `json:"resourceAdjustmentStep,omitempty"`
}

// ResourceAdjustmentStep defines resource adjustment increments
type ResourceAdjustmentStep struct {
	// Memory adjustment step
	//+optional
	Memory *intstr.IntOrString `json:"memory,omitempty"`

	// CPU adjustment step
	//+optional
	CPU *intstr.IntOrString `json:"cpu,omitempty"`
}

// CircuitBreakerConfig defines circuit breaker behavior
type CircuitBreakerConfig struct {
	// Enabled indicates if circuit breaker is active
	//+optional
	Enabled bool `json:"enabled,omitempty"`

	// FailureThreshold is the number of failures before opening the circuit
	//+kubebuilder:validation:Minimum=1
	//+optional
	FailureThreshold *int32 `json:"failureThreshold,omitempty"`

	// RecoveryTimeout is the time to wait before attempting recovery
	//+optional
	RecoveryTimeout *metav1.Duration `json:"recoveryTimeout,omitempty"`

	// MaxActionsPerWindow limits actions in a time window
	//+optional
	MaxActionsPerWindow *int32 `json:"maxActionsPerWindow,omitempty"`

	// TimeWindow defines the time window for action counting
	//+optional
	TimeWindow *metav1.Duration `json:"timeWindow,omitempty"`
}

// CostAwarenessConfig defines cost monitoring behavior
type CostAwarenessConfig struct {
	// Enabled indicates if cost awareness is active
	//+optional
	Enabled bool `json:"enabled,omitempty"`

	// MaxCostPerHour is the maximum cost per hour
	//+optional
	MaxCostPerHour *float64 `json:"maxCostPerHour,omitempty"`

	// BudgetThreshold is the percentage of the max cost per hour that triggers scaling
	//+kubebuilder:validation:Minimum=1
	//+kubebuilder:validation:Maximum=100
	//+optional
	BudgetThreshold *int32 `json:"budgetThreshold,omitempty"`
	// CostProvider specifies the cloud provider for cost calculation
	//+kubebuilder:validation:Enum=aws;azure;gcp
	//+optional
	CostProvider string `json:"costProvider,omitempty"`
}

// MetricsConfig defines metrics collection behavior
type MetricsConfig struct {
	// CollectionInterval defines how often to collect metrics
	//+optional
	CollectionInterval *metav1.Duration `json:"collectionInterval,omitempty"`

	// RetentionPeriod defines how long to retain metrics
	//+optional
	RetentionPeriod *metav1.Duration `json:"retentionPeriod,omitempty"`

	// AggregationWindow defines the window for metric aggregation
	//+optional
	AggregationWindow *metav1.Duration `json:"aggregationWindow,omitempty"`
}

type ResourceScalerPolicyStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// CurrentReplicas is the current number of replicas
	//+optional
	CurrentReplicas *int32 `json:"currentReplicas,omitempty"`

	// LastScaleTime is the last time scaling occurred
	//+optional
	LastScaleTime *metav1.Time `json:"lastScaleTime,omitempty"`

	// CircuitBreakerState represents the current circuit breaker state
	//+optional
	CircuitBreakerState string `json:"circuitBreakerState,omitempty"`

	// CurrentMemoryUsage is the current memory usage percentage
	//+optional
	CurrentMemoryUsage *int32 `json:"currentMemoryUsage,omitempty"`

	// CurrentCPUUsage is the current CPU usage percentage
	//+optional
	CurrentCPUUsage *int32 `json:"currentCPUUsage,omitempty"`

	// LastCostCalculation is the last calculated cost
	//+optional
	LastCostCalculation *float64 `json:"lastCostCalculation,omitempty"`

	// Conditions represent the latest available observations
	//+optional
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// Events is a list of recent events
	//+optional
	Events []string `json:"events,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status
//+kubebuilder:subresource:scale:specpath=.spec.scalePolicy.maxReplicas,statuspath=.status.currentReplicas
//+kubebuilder:printcolumn:name="Target",type="string",JSONPath=".spec.targetDeployment.name"
//+kubebuilder:printcolumn:name="Current Replicas",type="integer",JSONPath=".status.currentReplicas"
//+kubebuilder:printcolumn:name="Memory Usage",type="integer",JSONPath=".status.currentMemoryUsage"
//+kubebuilder:printcolumn:name="Circuit Breaker",type="string",JSONPath=".status.circuitBreakerState"
//+kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"

// ResourceScalerPolicy is the Schema for the resourcescalerpolicies API
type ResourceScalerPolicy struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty,omitzero"`

	// spec defines the desired state of ResourceScalerPolicy
	// +required
	Spec ResourceScalerPolicySpec `json:"spec"`

	// status defines the observed state of ResourceScalerPolicy
	// +optional
	Status ResourceScalerPolicyStatus `json:"status,omitempty,omitzero"`
}

// +kubebuilder:object:root=true

// ResourceScalerPolicyList contains a list of ResourceScalerPolicy
type ResourceScalerPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceScalerPolicy `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceScalerPolicy{}, &ResourceScalerPolicyList{})
}
