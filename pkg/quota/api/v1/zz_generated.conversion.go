// +build !ignore_autogenerated_openshift

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	api "github.com/openshift/origin/pkg/quota/api"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api_v1 "k8s.io/kubernetes/pkg/api/v1"
	unsafe "unsafe"
)

func init() {
	SchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_AppliedClusterResourceQuota_To_api_AppliedClusterResourceQuota,
		Convert_api_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota,
		Convert_v1_AppliedClusterResourceQuotaList_To_api_AppliedClusterResourceQuotaList,
		Convert_api_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList,
		Convert_v1_ClusterResourceQuota_To_api_ClusterResourceQuota,
		Convert_api_ClusterResourceQuota_To_v1_ClusterResourceQuota,
		Convert_v1_ClusterResourceQuotaList_To_api_ClusterResourceQuotaList,
		Convert_api_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList,
		Convert_v1_ClusterResourceQuotaSelector_To_api_ClusterResourceQuotaSelector,
		Convert_api_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector,
		Convert_v1_ClusterResourceQuotaSpec_To_api_ClusterResourceQuotaSpec,
		Convert_api_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec,
		Convert_v1_ClusterResourceQuotaStatus_To_api_ClusterResourceQuotaStatus,
		Convert_api_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus,
	)
}

func autoConvert_v1_AppliedClusterResourceQuota_To_api_AppliedClusterResourceQuota(in *AppliedClusterResourceQuota, out *api.AppliedClusterResourceQuota, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterResourceQuotaSpec_To_api_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterResourceQuotaStatus_To_api_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_AppliedClusterResourceQuota_To_api_AppliedClusterResourceQuota(in *AppliedClusterResourceQuota, out *api.AppliedClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_v1_AppliedClusterResourceQuota_To_api_AppliedClusterResourceQuota(in, out, s)
}

func autoConvert_api_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in *api.AppliedClusterResourceQuota, out *AppliedClusterResourceQuota, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_api_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in *api.AppliedClusterResourceQuota, out *AppliedClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_api_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(in, out, s)
}

func autoConvert_v1_AppliedClusterResourceQuotaList_To_api_AppliedClusterResourceQuotaList(in *AppliedClusterResourceQuotaList, out *api.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_v1_AppliedClusterResourceQuota_To_api_AppliedClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_AppliedClusterResourceQuotaList_To_api_AppliedClusterResourceQuotaList(in *AppliedClusterResourceQuotaList, out *api.AppliedClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_v1_AppliedClusterResourceQuotaList_To_api_AppliedClusterResourceQuotaList(in, out, s)
}

func autoConvert_api_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in *api.AppliedClusterResourceQuotaList, out *AppliedClusterResourceQuotaList, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AppliedClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_api_AppliedClusterResourceQuota_To_v1_AppliedClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in *api.AppliedClusterResourceQuotaList, out *AppliedClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_api_AppliedClusterResourceQuotaList_To_v1_AppliedClusterResourceQuotaList(in, out, s)
}

func autoConvert_v1_ClusterResourceQuota_To_api_ClusterResourceQuota(in *ClusterResourceQuota, out *api.ClusterResourceQuota, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ClusterResourceQuotaSpec_To_api_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ClusterResourceQuotaStatus_To_api_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_ClusterResourceQuota_To_api_ClusterResourceQuota(in *ClusterResourceQuota, out *api.ClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuota_To_api_ClusterResourceQuota(in, out, s)
}

func autoConvert_api_ClusterResourceQuota_To_v1_ClusterResourceQuota(in *api.ClusterResourceQuota, out *ClusterResourceQuota, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_api_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_api_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_ClusterResourceQuota_To_v1_ClusterResourceQuota(in *api.ClusterResourceQuota, out *ClusterResourceQuota, s conversion.Scope) error {
	return autoConvert_api_ClusterResourceQuota_To_v1_ClusterResourceQuota(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaList_To_api_ClusterResourceQuotaList(in *ClusterResourceQuotaList, out *api.ClusterResourceQuotaList, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]api.ClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_v1_ClusterResourceQuota_To_api_ClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_v1_ClusterResourceQuotaList_To_api_ClusterResourceQuotaList(in *ClusterResourceQuotaList, out *api.ClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaList_To_api_ClusterResourceQuotaList(in, out, s)
}

func autoConvert_api_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in *api.ClusterResourceQuotaList, out *ClusterResourceQuotaList, s conversion.Scope) error {
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterResourceQuota, len(*in))
		for i := range *in {
			if err := Convert_api_ClusterResourceQuota_To_v1_ClusterResourceQuota(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func Convert_api_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in *api.ClusterResourceQuotaList, out *ClusterResourceQuotaList, s conversion.Scope) error {
	return autoConvert_api_ClusterResourceQuotaList_To_v1_ClusterResourceQuotaList(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaSelector_To_api_ClusterResourceQuotaSelector(in *ClusterResourceQuotaSelector, out *api.ClusterResourceQuotaSelector, s conversion.Scope) error {
	out.LabelSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.AnnotationSelector = *(*map[string]string)(unsafe.Pointer(&in.AnnotationSelector))
	return nil
}

func Convert_v1_ClusterResourceQuotaSelector_To_api_ClusterResourceQuotaSelector(in *ClusterResourceQuotaSelector, out *api.ClusterResourceQuotaSelector, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaSelector_To_api_ClusterResourceQuotaSelector(in, out, s)
}

func autoConvert_api_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in *api.ClusterResourceQuotaSelector, out *ClusterResourceQuotaSelector, s conversion.Scope) error {
	out.LabelSelector = (*meta_v1.LabelSelector)(unsafe.Pointer(in.LabelSelector))
	out.AnnotationSelector = *(*map[string]string)(unsafe.Pointer(&in.AnnotationSelector))
	return nil
}

func Convert_api_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in *api.ClusterResourceQuotaSelector, out *ClusterResourceQuotaSelector, s conversion.Scope) error {
	return autoConvert_api_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaSpec_To_api_ClusterResourceQuotaSpec(in *ClusterResourceQuotaSpec, out *api.ClusterResourceQuotaSpec, s conversion.Scope) error {
	if err := Convert_v1_ClusterResourceQuotaSelector_To_api_ClusterResourceQuotaSelector(&in.Selector, &out.Selector, s); err != nil {
		return err
	}
	if err := api_v1.Convert_v1_ResourceQuotaSpec_To_api_ResourceQuotaSpec(&in.Quota, &out.Quota, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_ClusterResourceQuotaSpec_To_api_ClusterResourceQuotaSpec(in *ClusterResourceQuotaSpec, out *api.ClusterResourceQuotaSpec, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaSpec_To_api_ClusterResourceQuotaSpec(in, out, s)
}

func autoConvert_api_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in *api.ClusterResourceQuotaSpec, out *ClusterResourceQuotaSpec, s conversion.Scope) error {
	if err := Convert_api_ClusterResourceQuotaSelector_To_v1_ClusterResourceQuotaSelector(&in.Selector, &out.Selector, s); err != nil {
		return err
	}
	if err := api_v1.Convert_api_ResourceQuotaSpec_To_v1_ResourceQuotaSpec(&in.Quota, &out.Quota, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in *api.ClusterResourceQuotaSpec, out *ClusterResourceQuotaSpec, s conversion.Scope) error {
	return autoConvert_api_ClusterResourceQuotaSpec_To_v1_ClusterResourceQuotaSpec(in, out, s)
}

func autoConvert_v1_ClusterResourceQuotaStatus_To_api_ClusterResourceQuotaStatus(in *ClusterResourceQuotaStatus, out *api.ClusterResourceQuotaStatus, s conversion.Scope) error {
	if err := api_v1.Convert_v1_ResourceQuotaStatus_To_api_ResourceQuotaStatus(&in.Total, &out.Total, s); err != nil {
		return err
	}
	if err := Convert_v1_ResourceQuotasStatusByNamespace_To_api_ResourceQuotasStatusByNamespace(&in.Namespaces, &out.Namespaces, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1_ClusterResourceQuotaStatus_To_api_ClusterResourceQuotaStatus(in *ClusterResourceQuotaStatus, out *api.ClusterResourceQuotaStatus, s conversion.Scope) error {
	return autoConvert_v1_ClusterResourceQuotaStatus_To_api_ClusterResourceQuotaStatus(in, out, s)
}

func autoConvert_api_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in *api.ClusterResourceQuotaStatus, out *ClusterResourceQuotaStatus, s conversion.Scope) error {
	if err := api_v1.Convert_api_ResourceQuotaStatus_To_v1_ResourceQuotaStatus(&in.Total, &out.Total, s); err != nil {
		return err
	}
	if err := Convert_api_ResourceQuotasStatusByNamespace_To_v1_ResourceQuotasStatusByNamespace(&in.Namespaces, &out.Namespaces, s); err != nil {
		return err
	}
	return nil
}

func Convert_api_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in *api.ClusterResourceQuotaStatus, out *ClusterResourceQuotaStatus, s conversion.Scope) error {
	return autoConvert_api_ClusterResourceQuotaStatus_To_v1_ClusterResourceQuotaStatus(in, out, s)
}
