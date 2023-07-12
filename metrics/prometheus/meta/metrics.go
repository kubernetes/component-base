/*
Copyright 2023 The Kubernetes Authors.

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

package meta

import (
	k8smetrics "k8s.io/component-base/metrics"
)

var (
	RegisteredMetrics = k8smetrics.NewCounterVec(
		&k8smetrics.CounterOpts{
			Name:           "registered_metric_total",
			Help:           "The count of registered metrics broken by stability level and deprecation version.",
			StabilityLevel: k8smetrics.BETA,
		},
		[]string{"stability_level", "deprecated_version"},
	)

	DisabledMetricsTotal = k8smetrics.NewCounter(
		&k8smetrics.CounterOpts{
			Name:           "disabled_metric_total",
			Help:           "The count of disabled metrics.",
			StabilityLevel: k8smetrics.BETA,
		},
	)

	HiddenMetricsTotal = k8smetrics.NewCounter(
		&k8smetrics.CounterOpts{
			Name:           "hidden_metric_total",
			Help:           "The count of hidden metrics.",
			StabilityLevel: k8smetrics.BETA,
		},
	)
)
