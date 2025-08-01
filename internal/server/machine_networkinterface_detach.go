// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"
	"fmt"
	"time"

	"github.com/ironcore-dev/cloud-hypervisor-provider/api"
	iri "github.com/ironcore-dev/ironcore/iri/apis/machine/v1alpha1"
	"k8s.io/utils/ptr"
)

func (s *Server) DetachNetworkInterface(
	ctx context.Context,
	req *iri.DetachNetworkInterfaceRequest,
) (*iri.DetachNetworkInterfaceResponse, error) {
	log := s.loggerFrom(ctx)
	log.V(1).Info("Detaching nic from machine")

	if req == nil {
		return nil, fmt.Errorf("DetachNetworkInterface is nil")
	}

	apiMachine, err := s.machineStore.Get(ctx, req.MachineId)
	if err != nil {
		return nil, fmt.Errorf("failed to get machine: %w", err)
	}

	var updatedNICS []*api.NetworkInterfaceSpec
	found := false
	for _, nic := range apiMachine.Spec.NetworkInterfaces {
		if nic.Name == req.Name {
			nic.DeletedAt = ptr.To(time.Now())
			found = true
		}
		updatedNICS = append(updatedNICS, nic)
	}

	if !found {
		return nil, fmt.Errorf("nic '%s' not found in machine '%s'", req.Name, req.MachineId)
	}

	apiMachine.Spec.NetworkInterfaces = updatedNICS

	if _, err := s.machineStore.Update(ctx, apiMachine); err != nil {
		return nil, fmt.Errorf("failed to update machine: %w", err)
	}

	return &iri.DetachNetworkInterfaceResponse{}, nil
}
