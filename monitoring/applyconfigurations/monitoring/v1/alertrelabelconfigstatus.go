// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// AlertRelabelConfigStatusApplyConfiguration represents a declarative configuration of the AlertRelabelConfigStatus type for use
// with apply.
type AlertRelabelConfigStatusApplyConfiguration struct {
	Conditions []metav1.ConditionApplyConfiguration `json:"conditions,omitempty"`
}

// AlertRelabelConfigStatusApplyConfiguration constructs a declarative configuration of the AlertRelabelConfigStatus type for use with
// apply.
func AlertRelabelConfigStatus() *AlertRelabelConfigStatusApplyConfiguration {
	return &AlertRelabelConfigStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *AlertRelabelConfigStatusApplyConfiguration) WithConditions(values ...*metav1.ConditionApplyConfiguration) *AlertRelabelConfigStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
