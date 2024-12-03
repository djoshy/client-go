// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/console/v1"
)

// ConsolePluginCSPApplyConfiguration represents a declarative configuration of the ConsolePluginCSP type for use
// with apply.
type ConsolePluginCSPApplyConfiguration struct {
	Directive *v1.DirectiveType      `json:"directive,omitempty"`
	Values    []v1.CSPDirectiveValue `json:"values,omitempty"`
}

// ConsolePluginCSPApplyConfiguration constructs a declarative configuration of the ConsolePluginCSP type for use with
// apply.
func ConsolePluginCSP() *ConsolePluginCSPApplyConfiguration {
	return &ConsolePluginCSPApplyConfiguration{}
}

// WithDirective sets the Directive field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Directive field is set to the value of the last call.
func (b *ConsolePluginCSPApplyConfiguration) WithDirective(value v1.DirectiveType) *ConsolePluginCSPApplyConfiguration {
	b.Directive = &value
	return b
}

// WithValues adds the given value to the Values field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Values field.
func (b *ConsolePluginCSPApplyConfiguration) WithValues(values ...v1.CSPDirectiveValue) *ConsolePluginCSPApplyConfiguration {
	for i := range values {
		b.Values = append(b.Values, values[i])
	}
	return b
}
