/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2015 Intel Corporation
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

package pkgname

import (
	"github.com/intelsdi-x/snap/control/plugin"
	"github.com/intelsdi-x/snap/control/plugin/cpolicy"
	"os"
	"github.com/intelsdi-x/snap-plugin-utilities/config"
	"strconv"
	"time"
	"math/rand"
)

const (
	name = "many-mts"
	version = 1
	plgtype = plugin.CollectorPluginType
	vendor = "intel"
)

func NewCollector() *collector {
	h, err := os.Hostname()
	if err != nil {
		h = "localhost"
	}
	return &collector{hostname: h}
}

func (p *collector) GetMetricTypes(cfg plugin.PluginConfigType) ([]plugin.PluginMetricType, error) {
	var numberOfMetrics int
	mts := []plugin.PluginMetricType{}
	item, err := config.GetConfigItem(cfg, "number-of-metrics")
	if err != nil {
		numberOfMetrics = 1000
	} else {
		numberOfMetrics = item.(int)
	}
	for i := 1; i < numberOfMetrics + 1; i++ {
		mts = append(mts, plugin.PluginMetricType{Namespace_: []string{vendor, name, strconv.Itoa(i)}})
	}
	// Gather available metrics here
	return mts, nil
}

func (p *collector) CollectMetrics(metricTypes []plugin.PluginMetricType) ([]plugin.PluginMetricType, error) {
	metrics := []plugin.PluginMetricType{}

	for _, metricType := range metricTypes {
		metrics = append(metrics, plugin.PluginMetricType{
			Source_: p.hostname,
			Timestamp_: time.Now(),
			Namespace_: metricType.Namespace(),
			Data_: rand.Int(),
		})
	}

	return metrics, nil
}

// Commenting exported items is very important
func (p *collector) GetConfigPolicy() (*cpolicy.ConfigPolicy, error) {
	c := cpolicy.New()
	return c, nil
}

func Meta() *plugin.PluginMeta {
	return plugin.NewPluginMeta(
		name,
		version,
		plgtype,
		[]string{plugin.SnapGOBContentType},
		[]string{plugin.SnapGOBContentType},
	)
}

type collector struct {
	hostname string
}