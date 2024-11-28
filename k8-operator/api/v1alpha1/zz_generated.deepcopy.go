//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSIamAuthDetails) DeepCopyInto(out *AWSIamAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSIamAuthDetails.
func (in *AWSIamAuthDetails) DeepCopy() *AWSIamAuthDetails {
	if in == nil {
		return nil
	}
	out := new(AWSIamAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	out.ServiceAccount = in.ServiceAccount
	out.ServiceToken = in.ServiceToken
	out.UniversalAuth = in.UniversalAuth
	out.KubernetesAuth = in.KubernetesAuth
	out.AwsIamAuth = in.AwsIamAuth
	out.AzureAuth = in.AzureAuth
	out.GcpIdTokenAuth = in.GcpIdTokenAuth
	out.GcpIamAuth = in.GcpIamAuth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAuthDetails) DeepCopyInto(out *AzureAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAuthDetails.
func (in *AzureAuthDetails) DeepCopy() *AzureAuthDetails {
	if in == nil {
		return nil
	}
	out := new(AzureAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaReference) DeepCopyInto(out *CaReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaReference.
func (in *CaReference) DeepCopy() *CaReference {
	if in == nil {
		return nil
	}
	out := new(CaReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GCPIdTokenAuthDetails) DeepCopyInto(out *GCPIdTokenAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GCPIdTokenAuthDetails.
func (in *GCPIdTokenAuthDetails) DeepCopy() *GCPIdTokenAuthDetails {
	if in == nil {
		return nil
	}
	out := new(GCPIdTokenAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GcpIamAuthDetails) DeepCopyInto(out *GcpIamAuthDetails) {
	*out = *in
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GcpIamAuthDetails.
func (in *GcpIamAuthDetails) DeepCopy() *GcpIamAuthDetails {
	if in == nil {
		return nil
	}
	out := new(GcpIamAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecret) DeepCopyInto(out *InfisicalPushSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecret.
func (in *InfisicalPushSecret) DeepCopy() *InfisicalPushSecret {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalPushSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretDestination) DeepCopyInto(out *InfisicalPushSecretDestination) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretDestination.
func (in *InfisicalPushSecretDestination) DeepCopy() *InfisicalPushSecretDestination {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretDestination)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretList) DeepCopyInto(out *InfisicalPushSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfisicalPushSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretList.
func (in *InfisicalPushSecretList) DeepCopy() *InfisicalPushSecretList {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalPushSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretSpec) DeepCopyInto(out *InfisicalPushSecretSpec) {
	*out = *in
	out.Destination = in.Destination
	out.Authentication = in.Authentication
	out.Push = in.Push
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretSpec.
func (in *InfisicalPushSecretSpec) DeepCopy() *InfisicalPushSecretSpec {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalPushSecretStatus) DeepCopyInto(out *InfisicalPushSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ManagedSecrets != nil {
		in, out := &in.ManagedSecrets, &out.ManagedSecrets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalPushSecretStatus.
func (in *InfisicalPushSecretStatus) DeepCopy() *InfisicalPushSecretStatus {
	if in == nil {
		return nil
	}
	out := new(InfisicalPushSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecret) DeepCopyInto(out *InfisicalSecret) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecret.
func (in *InfisicalSecret) DeepCopy() *InfisicalSecret {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecret)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalSecret) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretList) DeepCopyInto(out *InfisicalSecretList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]InfisicalSecret, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretList.
func (in *InfisicalSecretList) DeepCopy() *InfisicalSecretList {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InfisicalSecretList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretSpec) DeepCopyInto(out *InfisicalSecretSpec) {
	*out = *in
	out.TokenSecretReference = in.TokenSecretReference
	out.Authentication = in.Authentication
	out.ManagedSecretReference = in.ManagedSecretReference
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretSpec.
func (in *InfisicalSecretSpec) DeepCopy() *InfisicalSecretSpec {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InfisicalSecretStatus) DeepCopyInto(out *InfisicalSecretStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InfisicalSecretStatus.
func (in *InfisicalSecretStatus) DeepCopy() *InfisicalSecretStatus {
	if in == nil {
		return nil
	}
	out := new(InfisicalSecretStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeSecretReference) DeepCopyInto(out *KubeSecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeSecretReference.
func (in *KubeSecretReference) DeepCopy() *KubeSecretReference {
	if in == nil {
		return nil
	}
	out := new(KubeSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesAuthDetails) DeepCopyInto(out *KubernetesAuthDetails) {
	*out = *in
	out.ServiceAccountRef = in.ServiceAccountRef
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesAuthDetails.
func (in *KubernetesAuthDetails) DeepCopy() *KubernetesAuthDetails {
	if in == nil {
		return nil
	}
	out := new(KubernetesAuthDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubernetesServiceAccountRef) DeepCopyInto(out *KubernetesServiceAccountRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubernetesServiceAccountRef.
func (in *KubernetesServiceAccountRef) DeepCopy() *KubernetesServiceAccountRef {
	if in == nil {
		return nil
	}
	out := new(KubernetesServiceAccountRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MachineIdentityScopeInWorkspace) DeepCopyInto(out *MachineIdentityScopeInWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MachineIdentityScopeInWorkspace.
func (in *MachineIdentityScopeInWorkspace) DeepCopy() *MachineIdentityScopeInWorkspace {
	if in == nil {
		return nil
	}
	out := new(MachineIdentityScopeInWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MangedKubeSecretConfig) DeepCopyInto(out *MangedKubeSecretConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MangedKubeSecretConfig.
func (in *MangedKubeSecretConfig) DeepCopy() *MangedKubeSecretConfig {
	if in == nil {
		return nil
	}
	out := new(MangedKubeSecretConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretAuthentication) DeepCopyInto(out *PushSecretAuthentication) {
	*out = *in
	out.UniversalAuth = in.UniversalAuth
	out.KubernetesAuth = in.KubernetesAuth
	out.AwsIamAuth = in.AwsIamAuth
	out.AzureAuth = in.AzureAuth
	out.GcpIdTokenAuth = in.GcpIdTokenAuth
	out.GcpIamAuth = in.GcpIamAuth
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretAuthentication.
func (in *PushSecretAuthentication) DeepCopy() *PushSecretAuthentication {
	if in == nil {
		return nil
	}
	out := new(PushSecretAuthentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretAwsIamAuth) DeepCopyInto(out *PushSecretAwsIamAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretAwsIamAuth.
func (in *PushSecretAwsIamAuth) DeepCopy() *PushSecretAwsIamAuth {
	if in == nil {
		return nil
	}
	out := new(PushSecretAwsIamAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretAzureAuth) DeepCopyInto(out *PushSecretAzureAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretAzureAuth.
func (in *PushSecretAzureAuth) DeepCopy() *PushSecretAzureAuth {
	if in == nil {
		return nil
	}
	out := new(PushSecretAzureAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretGcpIamAuth) DeepCopyInto(out *PushSecretGcpIamAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretGcpIamAuth.
func (in *PushSecretGcpIamAuth) DeepCopy() *PushSecretGcpIamAuth {
	if in == nil {
		return nil
	}
	out := new(PushSecretGcpIamAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretGcpIdTokenAuth) DeepCopyInto(out *PushSecretGcpIdTokenAuth) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretGcpIdTokenAuth.
func (in *PushSecretGcpIdTokenAuth) DeepCopy() *PushSecretGcpIdTokenAuth {
	if in == nil {
		return nil
	}
	out := new(PushSecretGcpIdTokenAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretKubernetesAuth) DeepCopyInto(out *PushSecretKubernetesAuth) {
	*out = *in
	out.ServiceAccountRef = in.ServiceAccountRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretKubernetesAuth.
func (in *PushSecretKubernetesAuth) DeepCopy() *PushSecretKubernetesAuth {
	if in == nil {
		return nil
	}
	out := new(PushSecretKubernetesAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretTlsConfig) DeepCopyInto(out *PushSecretTlsConfig) {
	*out = *in
	out.CaRef = in.CaRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretTlsConfig.
func (in *PushSecretTlsConfig) DeepCopy() *PushSecretTlsConfig {
	if in == nil {
		return nil
	}
	out := new(PushSecretTlsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSecretUniversalAuth) DeepCopyInto(out *PushSecretUniversalAuth) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSecretUniversalAuth.
func (in *PushSecretUniversalAuth) DeepCopy() *PushSecretUniversalAuth {
	if in == nil {
		return nil
	}
	out := new(PushSecretUniversalAuth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretPush) DeepCopyInto(out *SecretPush) {
	*out = *in
	out.Secret = in.Secret
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretPush.
func (in *SecretPush) DeepCopy() *SecretPush {
	if in == nil {
		return nil
	}
	out := new(SecretPush)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretScopeInWorkspace) DeepCopyInto(out *SecretScopeInWorkspace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretScopeInWorkspace.
func (in *SecretScopeInWorkspace) DeepCopy() *SecretScopeInWorkspace {
	if in == nil {
		return nil
	}
	out := new(SecretScopeInWorkspace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceAccountDetails) DeepCopyInto(out *ServiceAccountDetails) {
	*out = *in
	out.ServiceAccountSecretReference = in.ServiceAccountSecretReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceAccountDetails.
func (in *ServiceAccountDetails) DeepCopy() *ServiceAccountDetails {
	if in == nil {
		return nil
	}
	out := new(ServiceAccountDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceTokenDetails) DeepCopyInto(out *ServiceTokenDetails) {
	*out = *in
	out.ServiceTokenSecretReference = in.ServiceTokenSecretReference
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceTokenDetails.
func (in *ServiceTokenDetails) DeepCopy() *ServiceTokenDetails {
	if in == nil {
		return nil
	}
	out := new(ServiceTokenDetails)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSConfig) DeepCopyInto(out *TLSConfig) {
	*out = *in
	out.CaRef = in.CaRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSConfig.
func (in *TLSConfig) DeepCopy() *TLSConfig {
	if in == nil {
		return nil
	}
	out := new(TLSConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UniversalAuthDetails) DeepCopyInto(out *UniversalAuthDetails) {
	*out = *in
	out.CredentialsRef = in.CredentialsRef
	out.SecretsScope = in.SecretsScope
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UniversalAuthDetails.
func (in *UniversalAuthDetails) DeepCopy() *UniversalAuthDetails {
	if in == nil {
		return nil
	}
	out := new(UniversalAuthDetails)
	in.DeepCopyInto(out)
	return out
}
