package ecs

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

// Disk is a nested struct in ecs response
type Disk struct {
	BdfId                         string                        `json:"BdfId" xml:"BdfId"`
	ImageId                       string                        `json:"ImageId" xml:"ImageId"`
	DeleteAutoSnapshot            bool                          `json:"DeleteAutoSnapshot" xml:"DeleteAutoSnapshot"`
	AutoSnapshotPolicyId          string                        `json:"AutoSnapshotPolicyId" xml:"AutoSnapshotPolicyId"`
	ProvisionedIops               int64                         `json:"ProvisionedIops" xml:"ProvisionedIops"`
	DiskId                        string                        `json:"DiskId" xml:"DiskId"`
	IOPS                          int                           `json:"IOPS" xml:"IOPS"`
	MountInstanceNum              int                           `json:"MountInstanceNum" xml:"MountInstanceNum"`
	StorageSetId                  string                        `json:"StorageSetId" xml:"StorageSetId"`
	Type                          string                        `json:"Type" xml:"Type"`
	MultiAttach                   string                        `json:"MultiAttach" xml:"MultiAttach"`
	BurstingEnabled               bool                          `json:"BurstingEnabled" xml:"BurstingEnabled"`
	ProductCode                   string                        `json:"ProductCode" xml:"ProductCode"`
	Encrypted                     bool                          `json:"Encrypted" xml:"Encrypted"`
	DetachedTime                  string                        `json:"DetachedTime" xml:"DetachedTime"`
	DeleteWithInstance            bool                          `json:"DeleteWithInstance" xml:"DeleteWithInstance"`
	ZoneId                        string                        `json:"ZoneId" xml:"ZoneId"`
	PerformanceLevel              string                        `json:"PerformanceLevel" xml:"PerformanceLevel"`
	ThroughputRead                int                           `json:"ThroughputRead" xml:"ThroughputRead"`
	Status                        string                        `json:"Status" xml:"Status"`
	IOPSWrite                     int                           `json:"IOPSWrite" xml:"IOPSWrite"`
	AttachedTime                  string                        `json:"AttachedTime" xml:"AttachedTime"`
	Category                      string                        `json:"Category" xml:"Category"`
	StorageClusterId              string                        `json:"StorageClusterId" xml:"StorageClusterId"`
	EnableAutomatedSnapshotPolicy bool                          `json:"EnableAutomatedSnapshotPolicy" xml:"EnableAutomatedSnapshotPolicy"`
	SerialNumber                  string                        `json:"SerialNumber" xml:"SerialNumber"`
	Throughput                    int                           `json:"Throughput" xml:"Throughput"`
	Size                          int                           `json:"Size" xml:"Size"`
	RegionId                      string                        `json:"RegionId" xml:"RegionId"`
	ResourceGroupId               string                        `json:"ResourceGroupId" xml:"ResourceGroupId"`
	InstanceId                    string                        `json:"InstanceId" xml:"InstanceId"`
	Description                   string                        `json:"Description" xml:"Description"`
	ExpiredTime                   string                        `json:"ExpiredTime" xml:"ExpiredTime"`
	Device                        string                        `json:"Device" xml:"Device"`
	CreationTime                  string                        `json:"CreationTime" xml:"CreationTime"`
	IOPSRead                      int                           `json:"IOPSRead" xml:"IOPSRead"`
	SourceSnapshotId              string                        `json:"SourceSnapshotId" xml:"SourceSnapshotId"`
	StorageSetPartitionNumber     int                           `json:"StorageSetPartitionNumber" xml:"StorageSetPartitionNumber"`
	Portable                      bool                          `json:"Portable" xml:"Portable"`
	KMSKeyId                      string                        `json:"KMSKeyId" xml:"KMSKeyId"`
	EnableAutoSnapshot            bool                          `json:"EnableAutoSnapshot" xml:"EnableAutoSnapshot"`
	DiskChargeType                string                        `json:"DiskChargeType" xml:"DiskChargeType"`
	ThroughputWrite               int                           `json:"ThroughputWrite" xml:"ThroughputWrite"`
	DiskName                      string                        `json:"DiskName" xml:"DiskName"`
	Placement                     Placement                     `json:"Placement" xml:"Placement"`
	Tags                          TagsInDescribeDisks           `json:"Tags" xml:"Tags"`
	OperationLocks                OperationLocksInDescribeDisks `json:"OperationLocks" xml:"OperationLocks"`
	Attachments                   Attachments                   `json:"Attachments" xml:"Attachments"`
	MountInstances                MountInstances                `json:"MountInstances" xml:"MountInstances"`
}
