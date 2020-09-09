/*
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

package clients

import (
	"github.com/ellistarn/karpenter/pkg/apis/horizontalautoscaler/v1alpha1"
	"github.com/ellistarn/karpenter/pkg/metrics"
	"github.com/prometheus/client_golang/api"
)

// PrometheusMetricsClient is a metrics client for Prometheus
type PrometheusMetricsClient struct {
	Client api.Client
}

// GetCurrentValue for the metric
func (c *PrometheusMetricsClient) GetCurrentValue(v1alpha1.Metric) (metrics.Metric, error) {
	// TODO, implement this
	return metrics.Metric{}, nil
}
