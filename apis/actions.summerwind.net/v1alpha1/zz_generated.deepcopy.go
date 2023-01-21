//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2020 The actions-runner-controller authors.

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
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CacheEntry) DeepCopyInto(out *CacheEntry) {
	*out = *in
	in.ExpirationTime.DeepCopyInto(&out.ExpirationTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CacheEntry.
func (in *CacheEntry) DeepCopy() *CacheEntry {
	if in == nil {
		return nil
	}
	out := new(CacheEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservation) DeepCopyInto(out *CapacityReservation) {
	*out = *in
	in.ExpirationTime.DeepCopyInto(&out.ExpirationTime)
	in.EffectiveTime.DeepCopyInto(&out.EffectiveTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservation.
func (in *CapacityReservation) DeepCopy() *CapacityReservation {
	if in == nil {
		return nil
	}
	out := new(CapacityReservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CheckRunSpec) DeepCopyInto(out *CheckRunSpec) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Names != nil {
		in, out := &in.Names, &out.Names
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CheckRunSpec.
func (in *CheckRunSpec) DeepCopy() *CheckRunSpec {
	if in == nil {
		return nil
	}
	out := new(CheckRunSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubAPICredentialsFrom) DeepCopyInto(out *GitHubAPICredentialsFrom) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubAPICredentialsFrom.
func (in *GitHubAPICredentialsFrom) DeepCopy() *GitHubAPICredentialsFrom {
	if in == nil {
		return nil
	}
	out := new(GitHubAPICredentialsFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubEventScaleUpTriggerSpec) DeepCopyInto(out *GitHubEventScaleUpTriggerSpec) {
	*out = *in
	if in.CheckRun != nil {
		in, out := &in.CheckRun, &out.CheckRun
		*out = new(CheckRunSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(PullRequestSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Push != nil {
		in, out := &in.Push, &out.Push
		*out = new(PushSpec)
		**out = **in
	}
	if in.WorkflowJob != nil {
		in, out := &in.WorkflowJob, &out.WorkflowJob
		*out = new(WorkflowJobSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubEventScaleUpTriggerSpec.
func (in *GitHubEventScaleUpTriggerSpec) DeepCopy() *GitHubEventScaleUpTriggerSpec {
	if in == nil {
		return nil
	}
	out := new(GitHubEventScaleUpTriggerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalRunnerAutoscaler) DeepCopyInto(out *HorizontalRunnerAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalRunnerAutoscaler.
func (in *HorizontalRunnerAutoscaler) DeepCopy() *HorizontalRunnerAutoscaler {
	if in == nil {
		return nil
	}
	out := new(HorizontalRunnerAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorizontalRunnerAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalRunnerAutoscalerList) DeepCopyInto(out *HorizontalRunnerAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HorizontalRunnerAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalRunnerAutoscalerList.
func (in *HorizontalRunnerAutoscalerList) DeepCopy() *HorizontalRunnerAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(HorizontalRunnerAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HorizontalRunnerAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalRunnerAutoscalerSpec) DeepCopyInto(out *HorizontalRunnerAutoscalerSpec) {
	*out = *in
	out.ScaleTargetRef = in.ScaleTargetRef
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int)
		**out = **in
	}
	if in.MaxReplicas != nil {
		in, out := &in.MaxReplicas, &out.MaxReplicas
		*out = new(int)
		**out = **in
	}
	if in.ScaleDownDelaySecondsAfterScaleUp != nil {
		in, out := &in.ScaleDownDelaySecondsAfterScaleUp, &out.ScaleDownDelaySecondsAfterScaleUp
		*out = new(int)
		**out = **in
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]MetricSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScaleUpTriggers != nil {
		in, out := &in.ScaleUpTriggers, &out.ScaleUpTriggers
		*out = make([]ScaleUpTrigger, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.CapacityReservations != nil {
		in, out := &in.CapacityReservations, &out.CapacityReservations
		*out = make([]CapacityReservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScheduledOverrides != nil {
		in, out := &in.ScheduledOverrides, &out.ScheduledOverrides
		*out = make([]ScheduledOverride, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.GitHubAPICredentialsFrom != nil {
		in, out := &in.GitHubAPICredentialsFrom, &out.GitHubAPICredentialsFrom
		*out = new(GitHubAPICredentialsFrom)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalRunnerAutoscalerSpec.
func (in *HorizontalRunnerAutoscalerSpec) DeepCopy() *HorizontalRunnerAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(HorizontalRunnerAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HorizontalRunnerAutoscalerStatus) DeepCopyInto(out *HorizontalRunnerAutoscalerStatus) {
	*out = *in
	if in.DesiredReplicas != nil {
		in, out := &in.DesiredReplicas, &out.DesiredReplicas
		*out = new(int)
		**out = **in
	}
	if in.LastSuccessfulScaleOutTime != nil {
		in, out := &in.LastSuccessfulScaleOutTime, &out.LastSuccessfulScaleOutTime
		*out = (*in).DeepCopy()
	}
	if in.CacheEntries != nil {
		in, out := &in.CacheEntries, &out.CacheEntries
		*out = make([]CacheEntry, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ScheduledOverridesSummary != nil {
		in, out := &in.ScheduledOverridesSummary, &out.ScheduledOverridesSummary
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HorizontalRunnerAutoscalerStatus.
func (in *HorizontalRunnerAutoscalerStatus) DeepCopy() *HorizontalRunnerAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(HorizontalRunnerAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricSpec) DeepCopyInto(out *MetricSpec) {
	*out = *in
	if in.RepositoryNames != nil {
		in, out := &in.RepositoryNames, &out.RepositoryNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricSpec.
func (in *MetricSpec) DeepCopy() *MetricSpec {
	if in == nil {
		return nil
	}
	out := new(MetricSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestSpec) DeepCopyInto(out *PullRequestSpec) {
	*out = *in
	if in.Types != nil {
		in, out := &in.Types, &out.Types
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Branches != nil {
		in, out := &in.Branches, &out.Branches
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestSpec.
func (in *PullRequestSpec) DeepCopy() *PullRequestSpec {
	if in == nil {
		return nil
	}
	out := new(PullRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PushSpec) DeepCopyInto(out *PushSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PushSpec.
func (in *PushSpec) DeepCopy() *PushSpec {
	if in == nil {
		return nil
	}
	out := new(PushSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecurrenceRule) DeepCopyInto(out *RecurrenceRule) {
	*out = *in
	in.UntilTime.DeepCopyInto(&out.UntilTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecurrenceRule.
func (in *RecurrenceRule) DeepCopy() *RecurrenceRule {
	if in == nil {
		return nil
	}
	out := new(RecurrenceRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Runner) DeepCopyInto(out *Runner) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Runner.
func (in *Runner) DeepCopy() *Runner {
	if in == nil {
		return nil
	}
	out := new(Runner)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Runner) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerConfig) DeepCopyInto(out *RunnerConfig) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Ephemeral != nil {
		in, out := &in.Ephemeral, &out.Ephemeral
		*out = new(bool)
		**out = **in
	}
	if in.DockerdWithinRunnerContainer != nil {
		in, out := &in.DockerdWithinRunnerContainer, &out.DockerdWithinRunnerContainer
		*out = new(bool)
		**out = **in
	}
	if in.DockerEnabled != nil {
		in, out := &in.DockerEnabled, &out.DockerEnabled
		*out = new(bool)
		**out = **in
	}
	if in.DockerMTU != nil {
		in, out := &in.DockerMTU, &out.DockerMTU
		*out = new(int64)
		**out = **in
	}
	if in.DockerRegistryMirror != nil {
		in, out := &in.DockerRegistryMirror, &out.DockerRegistryMirror
		*out = new(string)
		**out = **in
	}
	if in.VolumeSizeLimit != nil {
		in, out := &in.VolumeSizeLimit, &out.VolumeSizeLimit
		x := (*in).DeepCopy()
		*out = &x
	}
	if in.VolumeStorageMedium != nil {
		in, out := &in.VolumeStorageMedium, &out.VolumeStorageMedium
		*out = new(string)
		**out = **in
	}
	if in.GitHubAPICredentialsFrom != nil {
		in, out := &in.GitHubAPICredentialsFrom, &out.GitHubAPICredentialsFrom
		*out = new(GitHubAPICredentialsFrom)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerConfig.
func (in *RunnerConfig) DeepCopy() *RunnerConfig {
	if in == nil {
		return nil
	}
	out := new(RunnerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerDeployment) DeepCopyInto(out *RunnerDeployment) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerDeployment.
func (in *RunnerDeployment) DeepCopy() *RunnerDeployment {
	if in == nil {
		return nil
	}
	out := new(RunnerDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerDeployment) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerDeploymentList) DeepCopyInto(out *RunnerDeploymentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RunnerDeployment, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerDeploymentList.
func (in *RunnerDeploymentList) DeepCopy() *RunnerDeploymentList {
	if in == nil {
		return nil
	}
	out := new(RunnerDeploymentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerDeploymentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerDeploymentSpec) DeepCopyInto(out *RunnerDeploymentSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.EffectiveTime != nil {
		in, out := &in.EffectiveTime, &out.EffectiveTime
		*out = (*in).DeepCopy()
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerDeploymentSpec.
func (in *RunnerDeploymentSpec) DeepCopy() *RunnerDeploymentSpec {
	if in == nil {
		return nil
	}
	out := new(RunnerDeploymentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerDeploymentStatus) DeepCopyInto(out *RunnerDeploymentStatus) {
	*out = *in
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int)
		**out = **in
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int)
		**out = **in
	}
	if in.UpdatedReplicas != nil {
		in, out := &in.UpdatedReplicas, &out.UpdatedReplicas
		*out = new(int)
		**out = **in
	}
	if in.DesiredReplicas != nil {
		in, out := &in.DesiredReplicas, &out.DesiredReplicas
		*out = new(int)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerDeploymentStatus.
func (in *RunnerDeploymentStatus) DeepCopy() *RunnerDeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(RunnerDeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerList) DeepCopyInto(out *RunnerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Runner, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerList.
func (in *RunnerList) DeepCopy() *RunnerList {
	if in == nil {
		return nil
	}
	out := new(RunnerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerPodSpec) DeepCopyInto(out *RunnerPodSpec) {
	*out = *in
	in.DockerdContainerResources.DeepCopyInto(&out.DockerdContainerResources)
	if in.DockerVolumeMounts != nil {
		in, out := &in.DockerVolumeMounts, &out.DockerVolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DockerEnv != nil {
		in, out := &in.DockerEnv, &out.DockerEnv
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Containers != nil {
		in, out := &in.Containers, &out.Containers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnvFrom != nil {
		in, out := &in.EnvFrom, &out.EnvFrom
		*out = make([]v1.EnvFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Resources.DeepCopyInto(&out.Resources)
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EnableServiceLinks != nil {
		in, out := &in.EnableServiceLinks, &out.EnableServiceLinks
		*out = new(bool)
		**out = **in
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AutomountServiceAccountToken != nil {
		in, out := &in.AutomountServiceAccountToken, &out.AutomountServiceAccountToken
		*out = new(bool)
		**out = **in
	}
	if in.SidecarContainers != nil {
		in, out := &in.SidecarContainers, &out.SidecarContainers
		*out = make([]v1.Container, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecrets != nil {
		in, out := &in.ImagePullSecrets, &out.ImagePullSecrets
		*out = make([]v1.LocalObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TerminationGracePeriodSeconds != nil {
		in, out := &in.TerminationGracePeriodSeconds, &out.TerminationGracePeriodSeconds
		*out = new(int64)
		**out = **in
	}
	if in.EphemeralContainers != nil {
		in, out := &in.EphemeralContainers, &out.EphemeralContainers
		*out = make([]v1.EphemeralContainer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.HostAliases != nil {
		in, out := &in.HostAliases, &out.HostAliases
		*out = make([]v1.HostAlias, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.TopologySpreadConstraints != nil {
		in, out := &in.TopologySpreadConstraints, &out.TopologySpreadConstraints
		*out = make([]v1.TopologySpreadConstraint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RuntimeClassName != nil {
		in, out := &in.RuntimeClassName, &out.RuntimeClassName
		*out = new(string)
		**out = **in
	}
	if in.DnsConfig != nil {
		in, out := &in.DnsConfig, &out.DnsConfig
		*out = new(v1.PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkVolumeClaimTemplate != nil {
		in, out := &in.WorkVolumeClaimTemplate, &out.WorkVolumeClaimTemplate
		*out = new(WorkVolumeClaimTemplate)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerPodSpec.
func (in *RunnerPodSpec) DeepCopy() *RunnerPodSpec {
	if in == nil {
		return nil
	}
	out := new(RunnerPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerReplicaSet) DeepCopyInto(out *RunnerReplicaSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerReplicaSet.
func (in *RunnerReplicaSet) DeepCopy() *RunnerReplicaSet {
	if in == nil {
		return nil
	}
	out := new(RunnerReplicaSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerReplicaSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerReplicaSetList) DeepCopyInto(out *RunnerReplicaSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RunnerReplicaSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerReplicaSetList.
func (in *RunnerReplicaSetList) DeepCopy() *RunnerReplicaSetList {
	if in == nil {
		return nil
	}
	out := new(RunnerReplicaSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerReplicaSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerReplicaSetSpec) DeepCopyInto(out *RunnerReplicaSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.EffectiveTime != nil {
		in, out := &in.EffectiveTime, &out.EffectiveTime
		*out = (*in).DeepCopy()
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	in.Template.DeepCopyInto(&out.Template)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerReplicaSetSpec.
func (in *RunnerReplicaSetSpec) DeepCopy() *RunnerReplicaSetSpec {
	if in == nil {
		return nil
	}
	out := new(RunnerReplicaSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerReplicaSetStatus) DeepCopyInto(out *RunnerReplicaSetStatus) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int)
		**out = **in
	}
	if in.AvailableReplicas != nil {
		in, out := &in.AvailableReplicas, &out.AvailableReplicas
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerReplicaSetStatus.
func (in *RunnerReplicaSetStatus) DeepCopy() *RunnerReplicaSetStatus {
	if in == nil {
		return nil
	}
	out := new(RunnerReplicaSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerSet) DeepCopyInto(out *RunnerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerSet.
func (in *RunnerSet) DeepCopy() *RunnerSet {
	if in == nil {
		return nil
	}
	out := new(RunnerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerSetList) DeepCopyInto(out *RunnerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RunnerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerSetList.
func (in *RunnerSetList) DeepCopy() *RunnerSetList {
	if in == nil {
		return nil
	}
	out := new(RunnerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RunnerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerSetSpec) DeepCopyInto(out *RunnerSetSpec) {
	*out = *in
	in.RunnerConfig.DeepCopyInto(&out.RunnerConfig)
	if in.EffectiveTime != nil {
		in, out := &in.EffectiveTime, &out.EffectiveTime
		*out = (*in).DeepCopy()
	}
	if in.WorkVolumeClaimTemplate != nil {
		in, out := &in.WorkVolumeClaimTemplate, &out.WorkVolumeClaimTemplate
		*out = new(WorkVolumeClaimTemplate)
		(*in).DeepCopyInto(*out)
	}
	in.StatefulSetSpec.DeepCopyInto(&out.StatefulSetSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerSetSpec.
func (in *RunnerSetSpec) DeepCopy() *RunnerSetSpec {
	if in == nil {
		return nil
	}
	out := new(RunnerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerSetStatus) DeepCopyInto(out *RunnerSetStatus) {
	*out = *in
	if in.CurrentReplicas != nil {
		in, out := &in.CurrentReplicas, &out.CurrentReplicas
		*out = new(int)
		**out = **in
	}
	if in.ReadyReplicas != nil {
		in, out := &in.ReadyReplicas, &out.ReadyReplicas
		*out = new(int)
		**out = **in
	}
	if in.UpdatedReplicas != nil {
		in, out := &in.UpdatedReplicas, &out.UpdatedReplicas
		*out = new(int)
		**out = **in
	}
	if in.DesiredReplicas != nil {
		in, out := &in.DesiredReplicas, &out.DesiredReplicas
		*out = new(int)
		**out = **in
	}
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerSetStatus.
func (in *RunnerSetStatus) DeepCopy() *RunnerSetStatus {
	if in == nil {
		return nil
	}
	out := new(RunnerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerSpec) DeepCopyInto(out *RunnerSpec) {
	*out = *in
	in.RunnerConfig.DeepCopyInto(&out.RunnerConfig)
	in.RunnerPodSpec.DeepCopyInto(&out.RunnerPodSpec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerSpec.
func (in *RunnerSpec) DeepCopy() *RunnerSpec {
	if in == nil {
		return nil
	}
	out := new(RunnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerStatus) DeepCopyInto(out *RunnerStatus) {
	*out = *in
	in.Registration.DeepCopyInto(&out.Registration)
	if in.WorkflowStatus != nil {
		in, out := &in.WorkflowStatus, &out.WorkflowStatus
		*out = new(WorkflowStatus)
		**out = **in
	}
	if in.LastRegistrationCheckTime != nil {
		in, out := &in.LastRegistrationCheckTime, &out.LastRegistrationCheckTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerStatus.
func (in *RunnerStatus) DeepCopy() *RunnerStatus {
	if in == nil {
		return nil
	}
	out := new(RunnerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerStatusRegistration) DeepCopyInto(out *RunnerStatusRegistration) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.ExpiresAt.DeepCopyInto(&out.ExpiresAt)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerStatusRegistration.
func (in *RunnerStatusRegistration) DeepCopy() *RunnerStatusRegistration {
	if in == nil {
		return nil
	}
	out := new(RunnerStatusRegistration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunnerTemplate) DeepCopyInto(out *RunnerTemplate) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunnerTemplate.
func (in *RunnerTemplate) DeepCopy() *RunnerTemplate {
	if in == nil {
		return nil
	}
	out := new(RunnerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleTargetRef) DeepCopyInto(out *ScaleTargetRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleTargetRef.
func (in *ScaleTargetRef) DeepCopy() *ScaleTargetRef {
	if in == nil {
		return nil
	}
	out := new(ScaleTargetRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScaleUpTrigger) DeepCopyInto(out *ScaleUpTrigger) {
	*out = *in
	if in.GitHubEvent != nil {
		in, out := &in.GitHubEvent, &out.GitHubEvent
		*out = new(GitHubEventScaleUpTriggerSpec)
		(*in).DeepCopyInto(*out)
	}
	out.Duration = in.Duration
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScaleUpTrigger.
func (in *ScaleUpTrigger) DeepCopy() *ScaleUpTrigger {
	if in == nil {
		return nil
	}
	out := new(ScaleUpTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduledOverride) DeepCopyInto(out *ScheduledOverride) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.MinReplicas != nil {
		in, out := &in.MinReplicas, &out.MinReplicas
		*out = new(int)
		**out = **in
	}
	in.RecurrenceRule.DeepCopyInto(&out.RecurrenceRule)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduledOverride.
func (in *ScheduledOverride) DeepCopy() *ScheduledOverride {
	if in == nil {
		return nil
	}
	out := new(ScheduledOverride)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretReference) DeepCopyInto(out *SecretReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretReference.
func (in *SecretReference) DeepCopy() *SecretReference {
	if in == nil {
		return nil
	}
	out := new(SecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkVolumeClaimTemplate) DeepCopyInto(out *WorkVolumeClaimTemplate) {
	*out = *in
	if in.AccessModes != nil {
		in, out := &in.AccessModes, &out.AccessModes
		*out = make([]v1.PersistentVolumeAccessMode, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkVolumeClaimTemplate.
func (in *WorkVolumeClaimTemplate) DeepCopy() *WorkVolumeClaimTemplate {
	if in == nil {
		return nil
	}
	out := new(WorkVolumeClaimTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowJobSpec) DeepCopyInto(out *WorkflowJobSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowJobSpec.
func (in *WorkflowJobSpec) DeepCopy() *WorkflowJobSpec {
	if in == nil {
		return nil
	}
	out := new(WorkflowJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkflowStatus) DeepCopyInto(out *WorkflowStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkflowStatus.
func (in *WorkflowStatus) DeepCopy() *WorkflowStatus {
	if in == nil {
		return nil
	}
	out := new(WorkflowStatus)
	in.DeepCopyInto(out)
	return out
}