// Copyright (c) 2018 Cisco and/or its affiliates.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/contiv/vpp/plugins/crd/pkg/apis/telemetry/v1"
	scheme "github.com/contiv/vpp/plugins/crd/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TelemetryReportsGetter has a method to return a TelemetryReportInterface.
// A group's client should implement this interface.
type TelemetryReportsGetter interface {
	TelemetryReports(namespace string) TelemetryReportInterface
}

// TelemetryReportInterface has methods to work with TelemetryReport resources.
type TelemetryReportInterface interface {
	Create(*v1.TelemetryReport) (*v1.TelemetryReport, error)
	Update(*v1.TelemetryReport) (*v1.TelemetryReport, error)
	UpdateStatus(*v1.TelemetryReport) (*v1.TelemetryReport, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error
	Get(name string, options metav1.GetOptions) (*v1.TelemetryReport, error)
	List(opts metav1.ListOptions) (*v1.TelemetryReportList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TelemetryReport, err error)
	TelemetryReportExpansion
}

// telemetryReports implements TelemetryReportInterface
type telemetryReports struct {
	client rest.Interface
	ns     string
}

// newTelemetryReports returns a TelemetryReports
func newTelemetryReports(c *TelemetryV1Client, namespace string) *telemetryReports {
	return &telemetryReports{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the telemetryReport, and returns the corresponding telemetryReport object, and an error if there is any.
func (c *telemetryReports) Get(name string, options metav1.GetOptions) (result *v1.TelemetryReport, err error) {
	result = &v1.TelemetryReport{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("telemetryreports").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TelemetryReports that match those selectors.
func (c *telemetryReports) List(opts metav1.ListOptions) (result *v1.TelemetryReportList, err error) {
	result = &v1.TelemetryReportList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("telemetryreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested telemetryReports.
func (c *telemetryReports) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("telemetryreports").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a telemetryReport and creates it.  Returns the server's representation of the telemetryReport, and an error, if there is any.
func (c *telemetryReports) Create(telemetryReport *v1.TelemetryReport) (result *v1.TelemetryReport, err error) {
	result = &v1.TelemetryReport{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("telemetryreports").
		Body(telemetryReport).
		Do().
		Into(result)
	return
}

// Update takes the representation of a telemetryReport and updates it. Returns the server's representation of the telemetryReport, and an error, if there is any.
func (c *telemetryReports) Update(telemetryReport *v1.TelemetryReport) (result *v1.TelemetryReport, err error) {
	result = &v1.TelemetryReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("telemetryreports").
		Name(telemetryReport.Name).
		Body(telemetryReport).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *telemetryReports) UpdateStatus(telemetryReport *v1.TelemetryReport) (result *v1.TelemetryReport, err error) {
	result = &v1.TelemetryReport{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("telemetryreports").
		Name(telemetryReport.Name).
		SubResource("status").
		Body(telemetryReport).
		Do().
		Into(result)
	return
}

// Delete takes name of the telemetryReport and deletes it. Returns an error if one occurs.
func (c *telemetryReports) Delete(name string, options *metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("telemetryreports").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *telemetryReports) DeleteCollection(options *metav1.DeleteOptions, listOptions metav1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("telemetryreports").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched telemetryReport.
func (c *telemetryReports) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.TelemetryReport, err error) {
	result = &v1.TelemetryReport{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("telemetryreports").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
