/*
 * Copyright 2020 The Knative Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package framework

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// ResourceContext allows the manipulation (crud) of resources within
// the current context (cluster and namespace_
type ResourceContext interface {
	// Create a resource from the given object (or fail)
	CreateOrFail(obj runtime.Object)

	// CreateFromYAMLOrFail creates a resource from the given YAML specification (or fail)
	CreateFromYAMLOrFail(yaml string)
}
