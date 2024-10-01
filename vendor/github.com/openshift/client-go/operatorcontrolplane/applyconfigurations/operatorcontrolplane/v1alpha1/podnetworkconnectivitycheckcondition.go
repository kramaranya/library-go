// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/api/operatorcontrolplane/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PodNetworkConnectivityCheckConditionApplyConfiguration represents a declarative configuration of the PodNetworkConnectivityCheckCondition type for use
// with apply.
type PodNetworkConnectivityCheckConditionApplyConfiguration struct {
	Type               *v1alpha1.PodNetworkConnectivityCheckConditionType `json:"type,omitempty"`
	Status             *v1.ConditionStatus                                `json:"status,omitempty"`
	Reason             *string                                            `json:"reason,omitempty"`
	Message            *string                                            `json:"message,omitempty"`
	LastTransitionTime *v1.Time                                           `json:"lastTransitionTime,omitempty"`
}

// PodNetworkConnectivityCheckConditionApplyConfiguration constructs a declarative configuration of the PodNetworkConnectivityCheckCondition type for use with
// apply.
func PodNetworkConnectivityCheckCondition() *PodNetworkConnectivityCheckConditionApplyConfiguration {
	return &PodNetworkConnectivityCheckConditionApplyConfiguration{}
}

// WithType sets the Type field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Type field is set to the value of the last call.
func (b *PodNetworkConnectivityCheckConditionApplyConfiguration) WithType(value v1alpha1.PodNetworkConnectivityCheckConditionType) *PodNetworkConnectivityCheckConditionApplyConfiguration {
	b.Type = &value
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *PodNetworkConnectivityCheckConditionApplyConfiguration) WithStatus(value v1.ConditionStatus) *PodNetworkConnectivityCheckConditionApplyConfiguration {
	b.Status = &value
	return b
}

// WithReason sets the Reason field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Reason field is set to the value of the last call.
func (b *PodNetworkConnectivityCheckConditionApplyConfiguration) WithReason(value string) *PodNetworkConnectivityCheckConditionApplyConfiguration {
	b.Reason = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *PodNetworkConnectivityCheckConditionApplyConfiguration) WithMessage(value string) *PodNetworkConnectivityCheckConditionApplyConfiguration {
	b.Message = &value
	return b
}

// WithLastTransitionTime sets the LastTransitionTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the LastTransitionTime field is set to the value of the last call.
func (b *PodNetworkConnectivityCheckConditionApplyConfiguration) WithLastTransitionTime(value v1.Time) *PodNetworkConnectivityCheckConditionApplyConfiguration {
	b.LastTransitionTime = &value
	return b
}
