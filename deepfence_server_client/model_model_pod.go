/*
Deepfence ThreatMapper

Deepfence Runtime API provides programmatic control over Deepfence microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: 2.0.0
Contact: community@deepfence.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deepfence_server_client

import (
	"encoding/json"
)

// checks if the ModelPod type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelPod{}

// ModelPod struct for ModelPod
type ModelPod struct {
	Containers []ModelContainer `json:"containers"`
	Image string `json:"image"`
	Metadata map[string]interface{} `json:"metadata"`
	Metrics ModelComputeMetrics `json:"metrics"`
	Name string `json:"name"`
	NodeId string `json:"node_id"`
	Processes []ModelProcess `json:"processes"`
}

// NewModelPod instantiates a new ModelPod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelPod(containers []ModelContainer, image string, metadata map[string]interface{}, metrics ModelComputeMetrics, name string, nodeId string, processes []ModelProcess) *ModelPod {
	this := ModelPod{}
	this.Containers = containers
	this.Image = image
	this.Metadata = metadata
	this.Metrics = metrics
	this.Name = name
	this.NodeId = nodeId
	this.Processes = processes
	return &this
}

// NewModelPodWithDefaults instantiates a new ModelPod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelPodWithDefaults() *ModelPod {
	this := ModelPod{}
	return &this
}

// GetContainers returns the Containers field value
// If the value is explicit nil, the zero value for []ModelContainer will be returned
func (o *ModelPod) GetContainers() []ModelContainer {
	if o == nil {
		var ret []ModelContainer
		return ret
	}

	return o.Containers
}

// GetContainersOk returns a tuple with the Containers field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPod) GetContainersOk() ([]ModelContainer, bool) {
	if o == nil || isNil(o.Containers) {
		return nil, false
	}
	return o.Containers, true
}

// SetContainers sets field value
func (o *ModelPod) SetContainers(v []ModelContainer) {
	o.Containers = v
}

// GetImage returns the Image field value
func (o *ModelPod) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Image, true
}

// SetImage sets field value
func (o *ModelPod) SetImage(v string) {
	o.Image = v
}

// GetMetadata returns the Metadata field value
func (o *ModelPod) GetMetadata() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetMetadataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Metadata, true
}

// SetMetadata sets field value
func (o *ModelPod) SetMetadata(v map[string]interface{}) {
	o.Metadata = v
}

// GetMetrics returns the Metrics field value
func (o *ModelPod) GetMetrics() ModelComputeMetrics {
	if o == nil {
		var ret ModelComputeMetrics
		return ret
	}

	return o.Metrics
}

// GetMetricsOk returns a tuple with the Metrics field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetMetricsOk() (*ModelComputeMetrics, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metrics, true
}

// SetMetrics sets field value
func (o *ModelPod) SetMetrics(v ModelComputeMetrics) {
	o.Metrics = v
}

// GetName returns the Name field value
func (o *ModelPod) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ModelPod) SetName(v string) {
	o.Name = v
}

// GetNodeId returns the NodeId field value
func (o *ModelPod) GetNodeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value
// and a boolean to check if the value has been set.
func (o *ModelPod) GetNodeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeId, true
}

// SetNodeId sets field value
func (o *ModelPod) SetNodeId(v string) {
	o.NodeId = v
}

// GetProcesses returns the Processes field value
// If the value is explicit nil, the zero value for []ModelProcess will be returned
func (o *ModelPod) GetProcesses() []ModelProcess {
	if o == nil {
		var ret []ModelProcess
		return ret
	}

	return o.Processes
}

// GetProcessesOk returns a tuple with the Processes field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelPod) GetProcessesOk() ([]ModelProcess, bool) {
	if o == nil || isNil(o.Processes) {
		return nil, false
	}
	return o.Processes, true
}

// SetProcesses sets field value
func (o *ModelPod) SetProcesses(v []ModelProcess) {
	o.Processes = v
}

func (o ModelPod) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelPod) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	toSerialize["image"] = o.Image
	toSerialize["metadata"] = o.Metadata
	toSerialize["metrics"] = o.Metrics
	toSerialize["name"] = o.Name
	toSerialize["node_id"] = o.NodeId
	if o.Processes != nil {
		toSerialize["processes"] = o.Processes
	}
	return toSerialize, nil
}

type NullableModelPod struct {
	value *ModelPod
	isSet bool
}

func (v NullableModelPod) Get() *ModelPod {
	return v.value
}

func (v *NullableModelPod) Set(val *ModelPod) {
	v.value = val
	v.isSet = true
}

func (v NullableModelPod) IsSet() bool {
	return v.isSet
}

func (v *NullableModelPod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelPod(val *ModelPod) *NullableModelPod {
	return &NullableModelPod{value: val, isSet: true}
}

func (v NullableModelPod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelPod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


