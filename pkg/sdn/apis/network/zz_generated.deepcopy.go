// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package network

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_ClusterNetwork, InType: reflect.TypeOf(&ClusterNetwork{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_ClusterNetworkList, InType: reflect.TypeOf(&ClusterNetworkList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_EgressNetworkPolicy, InType: reflect.TypeOf(&EgressNetworkPolicy{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_EgressNetworkPolicyList, InType: reflect.TypeOf(&EgressNetworkPolicyList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_EgressNetworkPolicyPeer, InType: reflect.TypeOf(&EgressNetworkPolicyPeer{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_EgressNetworkPolicyRule, InType: reflect.TypeOf(&EgressNetworkPolicyRule{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_EgressNetworkPolicySpec, InType: reflect.TypeOf(&EgressNetworkPolicySpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_HostSubnet, InType: reflect.TypeOf(&HostSubnet{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_HostSubnetList, InType: reflect.TypeOf(&HostSubnetList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_NetNamespace, InType: reflect.TypeOf(&NetNamespace{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_network_NetNamespaceList, InType: reflect.TypeOf(&NetNamespaceList{})},
	)
}

// DeepCopy_network_ClusterNetwork is an autogenerated deepcopy function.
func DeepCopy_network_ClusterNetwork(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterNetwork)
		out := out.(*ClusterNetwork)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

// DeepCopy_network_ClusterNetworkList is an autogenerated deepcopy function.
func DeepCopy_network_ClusterNetworkList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*ClusterNetworkList)
		out := out.(*ClusterNetworkList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]ClusterNetwork, len(*in))
			for i := range *in {
				if err := DeepCopy_network_ClusterNetwork(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_network_EgressNetworkPolicy is an autogenerated deepcopy function.
func DeepCopy_network_EgressNetworkPolicy(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EgressNetworkPolicy)
		out := out.(*EgressNetworkPolicy)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_network_EgressNetworkPolicySpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_network_EgressNetworkPolicyList is an autogenerated deepcopy function.
func DeepCopy_network_EgressNetworkPolicyList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EgressNetworkPolicyList)
		out := out.(*EgressNetworkPolicyList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]EgressNetworkPolicy, len(*in))
			for i := range *in {
				if err := DeepCopy_network_EgressNetworkPolicy(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_network_EgressNetworkPolicyPeer is an autogenerated deepcopy function.
func DeepCopy_network_EgressNetworkPolicyPeer(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EgressNetworkPolicyPeer)
		out := out.(*EgressNetworkPolicyPeer)
		*out = *in
		return nil
	}
}

// DeepCopy_network_EgressNetworkPolicyRule is an autogenerated deepcopy function.
func DeepCopy_network_EgressNetworkPolicyRule(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EgressNetworkPolicyRule)
		out := out.(*EgressNetworkPolicyRule)
		*out = *in
		return nil
	}
}

// DeepCopy_network_EgressNetworkPolicySpec is an autogenerated deepcopy function.
func DeepCopy_network_EgressNetworkPolicySpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*EgressNetworkPolicySpec)
		out := out.(*EgressNetworkPolicySpec)
		*out = *in
		if in.Egress != nil {
			in, out := &in.Egress, &out.Egress
			*out = make([]EgressNetworkPolicyRule, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_network_HostSubnet is an autogenerated deepcopy function.
func DeepCopy_network_HostSubnet(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HostSubnet)
		out := out.(*HostSubnet)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

// DeepCopy_network_HostSubnetList is an autogenerated deepcopy function.
func DeepCopy_network_HostSubnetList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*HostSubnetList)
		out := out.(*HostSubnetList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]HostSubnet, len(*in))
			for i := range *in {
				if err := DeepCopy_network_HostSubnet(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_network_NetNamespace is an autogenerated deepcopy function.
func DeepCopy_network_NetNamespace(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetNamespace)
		out := out.(*NetNamespace)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		return nil
	}
}

// DeepCopy_network_NetNamespaceList is an autogenerated deepcopy function.
func DeepCopy_network_NetNamespaceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*NetNamespaceList)
		out := out.(*NetNamespaceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]NetNamespace, len(*in))
			for i := range *in {
				if err := DeepCopy_network_NetNamespace(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
