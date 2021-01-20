// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Action) DeepCopyInto(out *Action) {
	*out = *in
	out.FillInterval.Seconds = in.FillInterval.Seconds
	out.FillInterval.Nanos = in.FillInterval.Nanos
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Action.
func (in *Action) DeepCopy() *Action {
	if in == nil {
		return nil
	}
	out := new(Action)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionStatus) DeepCopyInto(out *ActionStatus) {
	*out = *in
	out.FillInterval.Seconds = in.FillInterval.Seconds
	out.FillInterval.Nanos = in.FillInterval.Nanos
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionStatus.
func (in *ActionStatus) DeepCopy() *ActionStatus {
	if in == nil {
		return nil
	}
	out := new(ActionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Destinations) DeepCopyInto(out *Destinations) {
	*out = *in
	if in.RecentlyCalled != nil {
		in, out := &in.RecentlyCalled, &out.RecentlyCalled
		*out = new(Timestamp)
		(*in).DeepCopyInto(*out)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Destinations.
func (in *Destinations) DeepCopy() *Destinations {
	if in == nil {
		return nil
	}
	out := new(Destinations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowControlConfSpec) DeepCopyInto(out *FlowControlConfSpec) {
	*out = *in
	if in.RateLimitConf != nil {
		in, out := &in.RateLimitConf, &out.RateLimitConf
		*out = new(RateLimitConfSpec)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowControlConfSpec.
func (in *FlowControlConfSpec) DeepCopy() *FlowControlConfSpec {
	if in == nil {
		return nil
	}
	out := new(FlowControlConfSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FlowControlConfStatus) DeepCopyInto(out *FlowControlConfStatus) {
	*out = *in
	if in.RateLimitConf != nil {
		in, out := &in.RateLimitConf, &out.RateLimitConf
		*out = new(RateLimitConfStatus)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FlowControlConfStatus.
func (in *FlowControlConfStatus) DeepCopy() *FlowControlConfStatus {
	if in == nil {
		return nil
	}
	out := new(FlowControlConfStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HeaderMatcher) DeepCopyInto(out *HeaderMatcher) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HeaderMatcher.
func (in *HeaderMatcher) DeepCopy() *HeaderMatcher {
	if in == nil {
		return nil
	}
	out := new(HeaderMatcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateConfig) DeepCopyInto(out *RateConfig) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateConfig.
func (in *RateConfig) DeepCopy() *RateConfig {
	if in == nil {
		return nil
	}
	out := new(RateConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitConfSpec) DeepCopyInto(out *RateLimitConfSpec) {
	*out = *in
	if in.Descriptors != nil {
		in, out := &in.Descriptors, &out.Descriptors
		*out = make([]*RateLimitDescriptorConfigSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RateLimitDescriptorConfigSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitConfSpec.
func (in *RateLimitConfSpec) DeepCopy() *RateLimitConfSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitConfSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitConfStatus) DeepCopyInto(out *RateLimitConfStatus) {
	*out = *in
	if in.Descriptors != nil {
		in, out := &in.Descriptors, &out.Descriptors
		*out = make([]*RateLimitDescriptorConfigStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RateLimitDescriptorConfigStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitConfStatus.
func (in *RateLimitConfStatus) DeepCopy() *RateLimitConfStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitConfStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitDescriptorConfigSpec) DeepCopyInto(out *RateLimitDescriptorConfigSpec) {
	*out = *in
	if in.Descriptors != nil {
		in, out := &in.Descriptors, &out.Descriptors
		*out = make([]*RateLimitDescriptorConfigSpec, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RateLimitDescriptorConfigSpec)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitDescriptorConfigSpec.
func (in *RateLimitDescriptorConfigSpec) DeepCopy() *RateLimitDescriptorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(RateLimitDescriptorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RateLimitDescriptorConfigStatus) DeepCopyInto(out *RateLimitDescriptorConfigStatus) {
	*out = *in
	if in.RateLimit != nil {
		in, out := &in.RateLimit, &out.RateLimit
		*out = new(RateConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Descriptors != nil {
		in, out := &in.Descriptors, &out.Descriptors
		*out = make([]*RateLimitDescriptorConfigStatus, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(RateLimitDescriptorConfigStatus)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RateLimitDescriptorConfigStatus.
func (in *RateLimitDescriptorConfigStatus) DeepCopy() *RateLimitDescriptorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(RateLimitDescriptorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecyclingStrategy) DeepCopyInto(out *RecyclingStrategy) {
	*out = *in
	if in.Stable != nil {
		in, out := &in.Stable, &out.Stable
		*out = new(RecyclingStrategy_Stable)
		(*in).DeepCopyInto(*out)
	}
	if in.Deadline != nil {
		in, out := &in.Deadline, &out.Deadline
		*out = new(RecyclingStrategy_Deadline)
		(*in).DeepCopyInto(*out)
	}
	if in.Auto != nil {
		in, out := &in.Auto, &out.Auto
		*out = new(RecyclingStrategy_Auto)
		(*in).DeepCopyInto(*out)
	}
	if in.RecentlyCalled != nil {
		in, out := &in.RecentlyCalled, &out.RecentlyCalled
		*out = new(Timestamp)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecyclingStrategy.
func (in *RecyclingStrategy) DeepCopy() *RecyclingStrategy {
	if in == nil {
		return nil
	}
	out := new(RecyclingStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecyclingStrategy_Auto) DeepCopyInto(out *RecyclingStrategy_Auto) {
	*out = *in
	if in.Duration != nil {
		in, out := &in.Duration, &out.Duration
		*out = new(Timestamp)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecyclingStrategy_Auto.
func (in *RecyclingStrategy_Auto) DeepCopy() *RecyclingStrategy_Auto {
	if in == nil {
		return nil
	}
	out := new(RecyclingStrategy_Auto)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecyclingStrategy_Deadline) DeepCopyInto(out *RecyclingStrategy_Deadline) {
	*out = *in
	if in.Expire != nil {
		in, out := &in.Expire, &out.Expire
		*out = new(Timestamp)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecyclingStrategy_Deadline.
func (in *RecyclingStrategy_Deadline) DeepCopy() *RecyclingStrategy_Deadline {
	if in == nil {
		return nil
	}
	out := new(RecyclingStrategy_Deadline)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecyclingStrategy_Stable) DeepCopyInto(out *RecyclingStrategy_Stable) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecyclingStrategy_Stable.
func (in *RecyclingStrategy_Stable) DeepCopy() *RecyclingStrategy_Stable {
	if in == nil {
		return nil
	}
	out := new(RecyclingStrategy_Stable)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFence) DeepCopyInto(out *ServiceFence) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFence.
func (in *ServiceFence) DeepCopy() *ServiceFence {
	if in == nil {
		return nil
	}
	out := new(ServiceFence)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceFence) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFenceList) DeepCopyInto(out *ServiceFenceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ServiceFence, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFenceList.
func (in *ServiceFenceList) DeepCopy() *ServiceFenceList {
	if in == nil {
		return nil
	}
	out := new(ServiceFenceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ServiceFenceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFenceSpec) DeepCopyInto(out *ServiceFenceSpec) {
	*out = *in
	if in.Host != nil {
		in, out := &in.Host, &out.Host
		*out = make(map[string]*RecyclingStrategy, len(*in))
		for key, val := range *in {
			var outVal *RecyclingStrategy
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(RecyclingStrategy)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFenceSpec.
func (in *ServiceFenceSpec) DeepCopy() *ServiceFenceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceFenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceFenceStatus) DeepCopyInto(out *ServiceFenceStatus) {
	*out = *in
	if in.Domains != nil {
		in, out := &in.Domains, &out.Domains
		*out = make(map[string]*Destinations, len(*in))
		for key, val := range *in {
			var outVal *Destinations
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(Destinations)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.Visitor != nil {
		in, out := &in.Visitor, &out.Visitor
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceFenceStatus.
func (in *ServiceFenceStatus) DeepCopy() *ServiceFenceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceFenceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimitDescriptor) DeepCopyInto(out *SmartLimitDescriptor) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(Action)
		(*in).DeepCopyInto(*out)
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]*HeaderMatcher, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HeaderMatcher)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptor.
func (in *SmartLimitDescriptor) DeepCopy() *SmartLimitDescriptor {
	if in == nil {
		return nil
	}
	out := new(SmartLimitDescriptor)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimitDescriptorStatus) DeepCopyInto(out *SmartLimitDescriptorStatus) {
	*out = *in
	if in.Action != nil {
		in, out := &in.Action, &out.Action
		*out = new(ActionStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]*HeaderMatcher, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(HeaderMatcher)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimitDescriptorStatus.
func (in *SmartLimitDescriptorStatus) DeepCopy() *SmartLimitDescriptorStatus {
	if in == nil {
		return nil
	}
	out := new(SmartLimitDescriptorStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimiter) DeepCopyInto(out *SmartLimiter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiter.
func (in *SmartLimiter) DeepCopy() *SmartLimiter {
	if in == nil {
		return nil
	}
	out := new(SmartLimiter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmartLimiter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimiterList) DeepCopyInto(out *SmartLimiterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SmartLimiter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterList.
func (in *SmartLimiterList) DeepCopy() *SmartLimiterList {
	if in == nil {
		return nil
	}
	out := new(SmartLimiterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SmartLimiterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimiterSpec) DeepCopyInto(out *SmartLimiterSpec) {
	*out = *in
	if in.Descriptors != nil {
		in, out := &in.Descriptors, &out.Descriptors
		*out = make([]*SmartLimitDescriptor, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SmartLimitDescriptor)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterSpec.
func (in *SmartLimiterSpec) DeepCopy() *SmartLimiterSpec {
	if in == nil {
		return nil
	}
	out := new(SmartLimiterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimiterStatus) DeepCopyInto(out *SmartLimiterStatus) {
	*out = *in
	if in.RatelimitStatus != nil {
		in, out := &in.RatelimitStatus, &out.RatelimitStatus
		*out = make([]*SmartLimitDescriptor, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SmartLimitDescriptor)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	if in.EndPointStatus != nil {
		in, out := &in.EndPointStatus, &out.EndPointStatus
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ServiceStatus != nil {
		in, out := &in.ServiceStatus, &out.ServiceStatus
		*out = new(SmartLimiterStatus_ServiceStatus)
		(*in).DeepCopyInto(*out)
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterStatus.
func (in *SmartLimiterStatus) DeepCopy() *SmartLimiterStatus {
	if in == nil {
		return nil
	}
	out := new(SmartLimiterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimiterStatus_ServiceStatus) DeepCopyInto(out *SmartLimiterStatus_ServiceStatus) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Listener != nil {
		in, out := &in.Listener, &out.Listener
		*out = make([]*SmartLimiterStatus_ServiceStatus_Listener, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(SmartLimiterStatus_ServiceStatus_Listener)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterStatus_ServiceStatus.
func (in *SmartLimiterStatus_ServiceStatus) DeepCopy() *SmartLimiterStatus_ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(SmartLimiterStatus_ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SmartLimiterStatus_ServiceStatus_Listener) DeepCopyInto(out *SmartLimiterStatus_ServiceStatus_Listener) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SmartLimiterStatus_ServiceStatus_Listener.
func (in *SmartLimiterStatus_ServiceStatus_Listener) DeepCopy() *SmartLimiterStatus_ServiceStatus_Listener {
	if in == nil {
		return nil
	}
	out := new(SmartLimiterStatus_ServiceStatus_Listener)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Timestamp) DeepCopyInto(out *Timestamp) {
	*out = *in
	out.XXX_NoUnkeyedLiteral = in.XXX_NoUnkeyedLiteral
	if in.XXX_unrecognized != nil {
		in, out := &in.XXX_unrecognized, &out.XXX_unrecognized
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Timestamp.
func (in *Timestamp) DeepCopy() *Timestamp {
	if in == nil {
		return nil
	}
	out := new(Timestamp)
	in.DeepCopyInto(out)
	return out
}