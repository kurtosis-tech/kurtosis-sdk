/*
 *    Copyright 2021 Kurtosis Technologies Inc.
 *
 *    Licensed under the Apache License, Version 2.0 (the "License");
 *    you may not use this file except in compliance with the License.
 *    You may obtain a copy of the License at
 *
 *        http://www.apache.org/licenses/LICENSE-2.0
 *
 *    Unless required by applicable law or agreed to in writing, software
 *    distributed under the License is distributed on an "AS IS" BASIS,
 *    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *    See the License for the specific language governing permissions and
 *    limitations under the License.
 *
 */

package services

import (
	"context"
	"github.com/kurtosis-tech/kurtosis-sdk/api/golang/core/kurtosis_core_rpc_api_bindings"
	"github.com/kurtosis-tech/kurtosis-sdk/api/golang/core/lib/binding_constructors"
	"github.com/kurtosis-tech/stacktrace"
)

// Docs available at https://docs.kurtosistech.com/kurtosis-core/lib-documentation
type ServiceContext struct {
	client    kurtosis_core_rpc_api_bindings.ApiContainerServiceClient
	serviceId ServiceID

	// Network location inside the enclave
	privateIpAddr string
	privatePorts  map[string]*PortSpec

	// Network location outside the enclave
	publicIpAddr string
	publicPorts  map[string]*PortSpec
}

func NewServiceContext(client kurtosis_core_rpc_api_bindings.ApiContainerServiceClient, serviceId ServiceID, privateIpAddr string, privatePorts map[string]*PortSpec, publicIpAddr string, publicPorts map[string]*PortSpec) *ServiceContext {
	return &ServiceContext{client: client, serviceId: serviceId, privateIpAddr: privateIpAddr, privatePorts: privatePorts, publicIpAddr: publicIpAddr, publicPorts: publicPorts}
}

// Docs available at https://docs.kurtosistech.com/kurtosis-core/lib-documentation
func (self *ServiceContext) GetServiceID() ServiceID {
	return self.serviceId
}

// Docs available at https://docs.kurtosistech.com/kurtosis-core/lib-documentation
func (self *ServiceContext) GetPrivateIPAddress() string {
	return self.privateIpAddr
}

// Docs available at https://docs.kurtosistech.com/kurtosis-core/lib-documentation
func (self *ServiceContext) GetPrivatePorts() map[string]*PortSpec {
	return self.privatePorts
}

// Docs available at https://docs.kurtosistech.com/kurtosis-core/lib-documentation
func (self *ServiceContext) GetMaybePublicIPAddress() string {
	return self.publicIpAddr
}

// Docs available at https://docs.kurtosistech.com/kurtosis-core/lib-documentation
func (self *ServiceContext) GetPublicPorts() map[string]*PortSpec {
	return self.publicPorts
}

// Docs available at https://docs.kurtosistech.com/kurtosis-core/lib-documentation
func (self *ServiceContext) ExecCommand(command []string) (int32, string, error) {
	serviceId := self.serviceId
	args := binding_constructors.NewExecCommandArgs(string(serviceId), command)
	resp, err := self.client.ExecCommand(context.Background(), args)
	if err != nil {
		return 0, "", stacktrace.Propagate(
			err,
			"An error occurred executing command '%v' on service '%v'",
			command,
			serviceId)
	}
	return resp.ExitCode, resp.LogOutput, nil
}
