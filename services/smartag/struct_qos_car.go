package smartag

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// QosCar is a nested struct in smartag response
type QosCar struct {
	QosCarId            string `json:"QosCarId" xml:"QosCarId"`
	QosId               string `json:"QosId" xml:"QosId"`
	Description         string `json:"Description" xml:"Description"`
	Priority            int    `json:"Priority" xml:"Priority"`
	LimitType           string `json:"LimitType" xml:"LimitType"`
	MinBandwidthAbs     int    `json:"MinBandwidthAbs" xml:"MinBandwidthAbs"`
	MaxBandwidthAbs     int    `json:"MaxBandwidthAbs" xml:"MaxBandwidthAbs"`
	MinBandwidthPercent int    `json:"MinBandwidthPercent" xml:"MinBandwidthPercent"`
	MaxBandwidthPercent int    `json:"MaxBandwidthPercent" xml:"MaxBandwidthPercent"`
	PercentSourceType   string `json:"PercentSourceType" xml:"PercentSourceType"`
	Name                string `json:"Name" xml:"Name"`
}
