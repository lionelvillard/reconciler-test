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

// Suite represents a collection of test cases.
// There is only one suite per Go package and created in TestMain
type Suite interface {
	// ParseFlags parses defined flags in def
	ParseFlags(def interface{}) Suite

	// Require makes sure the cluster-scoped component
	// is deployed and properly configured (ie. ready to be used)
	//
	// When the --build flag is set, Require first builds and installs
	// the required component.
	Require(component Component) Suite

	// Run runs the tests
	Run()
}
