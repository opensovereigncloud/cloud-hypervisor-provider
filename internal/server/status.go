// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package server

import (
	"context"

	iri "github.com/ironcore-dev/ironcore/iri/apis/machine/v1alpha1"
)

func (s *Server) Status(ctx context.Context, _ *iri.StatusRequest) (*iri.StatusResponse, error) {
	log := s.loggerFrom(ctx)

	var classes []*iri.MachineClassStatus
	for _, class := range s.machineClassRegistry.List() {
		classes = append(classes, &iri.MachineClassStatus{
			MachineClass: &iri.MachineClass{
				Name: class.Name,
				Capabilities: &iri.MachineClassCapabilities{
					Resources: map[string]int64{
						"cpu":    class.Cpu,
						"memory": class.MemoryBytes,
					},
				},
			},
			//TODO will be deprecated soon
			Quantity: 1000,
		})
	}

	log.V(1).Info("Returning machine classes")
	return &iri.StatusResponse{
		MachineClassStatus: classes,
	}, nil
}
