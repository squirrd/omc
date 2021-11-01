/*
Copyright The Kubernetes Authors.

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

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its types. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_ImageReview = map[string]string{
	"":         "ImageReview checks if the set of images in a pod are allowed.",
	"metadata": "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata",
	"spec":     "Spec holds information about the pod being evaluated",
	"status":   "Status is filled in by the backend and indicates whether the pod should be allowed.",
}

func (ImageReview) SwaggerDoc() map[string]string {
	return map_ImageReview
}

var map_ImageReviewContainerSpec = map[string]string{
	"":      "ImageReviewContainerSpec is a description of a container within the pod creation request.",
	"image": "This can be in the form image:tag or image@SHA:012345679abcdef.",
}

func (ImageReviewContainerSpec) SwaggerDoc() map[string]string {
	return map_ImageReviewContainerSpec
}

var map_ImageReviewSpec = map[string]string{
	"":            "ImageReviewSpec is a description of the pod creation request.",
	"containers":  "Containers is a list of a subset of the information in each container of the Pod being created.",
	"annotations": "Annotations is a list of key-value pairs extracted from the Pod's annotations. It only includes keys which match the pattern `*.image-policy.k8s.io/*`. It is up to each webhook backend to determine how to interpret these annotations, if at all.",
	"namespace":   "Namespace is the namespace the pod is being created in.",
}

func (ImageReviewSpec) SwaggerDoc() map[string]string {
	return map_ImageReviewSpec
}

var map_ImageReviewStatus = map[string]string{
	"":                 "ImageReviewStatus is the result of the review for the pod creation request.",
	"allowed":          "Allowed indicates that all images were allowed to be run.",
	"reason":           "Reason should be empty unless Allowed is false in which case it may contain a short description of what is wrong.  Kubernetes may truncate excessively long errors when displaying to the user.",
	"auditAnnotations": "AuditAnnotations will be added to the attributes object of the admission controller request using 'AddAnnotation'.  The keys should be prefix-less (i.e., the admission controller will add an appropriate prefix).",
}

func (ImageReviewStatus) SwaggerDoc() map[string]string {
	return map_ImageReviewStatus
}

// AUTO-GENERATED FUNCTIONS END HERE