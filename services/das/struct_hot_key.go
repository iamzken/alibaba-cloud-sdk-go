package das

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

// HotKey is a nested struct in das response
type HotKey struct {
	Db      int    `json:"Db" xml:"Db"`
	KeyType string `json:"KeyType" xml:"KeyType"`
	Size    int64  `json:"Size" xml:"Size"`
	Lru     int    `json:"Lru" xml:"Lru"`
	Hot     string `json:"Hot" xml:"Hot"`
	Key     string `json:"Key" xml:"Key"`
}
