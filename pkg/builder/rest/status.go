/*
Copyright 2017 The Kubernetes Authors.

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

package rest

import (
	"context"

	"github.com/pwittrock/apiserver-runtime/pkg/builder/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type StatusSubResourceStrategy struct {
	Strategy
}

// PrepareForUpdate calls the PrepareForUpdate function on obj if supported, otherwise does nothing.
func (StatusSubResourceStrategy) PrepareForUpdate(ctx context.Context, obj, old runtime.Object) {
	if v, ok := obj.(resource.StatusGetSetter); ok {
		v.CopySpec(ctx, old)
	}
}
