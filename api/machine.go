// SPDX-FileCopyrightText: 2023 SAP SE or an SAP affiliate company and IronCore contributors
// SPDX-License-Identifier: Apache-2.0

package api

import (
	"time"

	apiutils "github.com/ironcore-dev/provider-utils/apiutils/api"
)

type Machine struct {
	apiutils.Metadata `json:"metadata,omitempty"`

	Spec   MachineSpec   `json:"spec"`
	Status MachineStatus `json:"status"`
}

type MachineSpec struct {
	ApiSocketPath *string `json:"api"`

	Power PowerState `json:"power"`

	Cpu         int64 `json:"cpuMillis"`
	MemoryBytes int64 `json:"memoryBytes"`

	Image    *string `json:"image"`
	Ignition []byte  `json:"ignition"`

	Volumes           []*VolumeSpec           `json:"volumes"`
	NetworkInterfaces []*NetworkInterfaceSpec `json:"networkInterfaces"`

	ShutdownAt time.Time `json:"shutdownAt,omitempty"`
}

type MachineStatus struct {
	VolumeStatus           []VolumeStatus           `json:"volumeStatus"`
	NetworkInterfaceStatus []NetworkInterfaceStatus `json:"networkInterfaceStatus"`
	State                  MachineState             `json:"state"`
	ImageRef               string                   `json:"imageRef"`
}

type MachineState string

const (
	MachineStatePending     MachineState = "Pending"
	MachineStateRunning     MachineState = "Running"
	MachineStateSuspended   MachineState = "Suspended"
	MachineStateTerminating MachineState = "Terminating"
	MachineStateTerminated  MachineState = "Terminated"
)

type PowerState int32

const (
	PowerStatePowerOn  PowerState = 0
	PowerStatePowerOff PowerState = 1
)

type VolumeSpec struct {
	Name       string            `json:"name"`
	Device     string            `json:"device"`
	EmptyDisk  *EmptyDiskSpec    `json:"emptyDisk,omitempty"`
	Connection *VolumeConnection `json:"cephDisk,omitempty"`
	DeletedAt  *time.Time        `json:"deletedAt,omitempty"`
}

type VolumeStatus struct {
	Name   string      `json:"name,omitempty"`
	Type   VolumeType  `json:"type,omitempty"`
	Path   string      `json:"path,omitempty"`
	Handle string      `json:"handle,omitempty"`
	State  VolumeState `json:"state,omitempty"`
	Size   int64       `json:"size,omitempty"`
}

type EmptyDiskSpec struct {
	Size int64 `json:"size"`
}

type VolumeConnection struct {
	Driver         string            ` json:"driver,omitempty"`
	Handle         string            ` json:"handle,omitempty"`
	Attributes     map[string]string ` json:"attributes,omitempty"`
	SecretData     map[string][]byte ` json:"secret_data,omitempty"`
	EncryptionData map[string][]byte ` json:"encryption_data,omitempty"`
}

type VolumeState string

const (
	VolumeStatePending  VolumeState = "Pending"
	VolumeStatePrepared VolumeState = "Prepared"
	VolumeStateAttached VolumeState = "Attached"
)

type VolumeType string

const (
	VolumeSocketType VolumeType = "socket"
	VolumeFileType   VolumeType = "file"
)

type NetworkInterfaceSpec struct {
	Name       string            `json:"name"`
	NetworkId  string            `json:"networkId"`
	Ips        []string          `json:"ips"`
	Attributes map[string]string `json:"attributes"`
	DeletedAt  *time.Time        `json:"deletedAt,omitempty"`
}

type NetworkInterfaceStatus struct {
	Name   string                `json:"name"`
	Handle string                `json:"handle"`
	State  NetworkInterfaceState `json:"state"`
	Type   NetworkInterfaceType  `json:"type,omitempty"`
	Path   string                `json:"path,omitempty"`
}

type NetworkInterfaceState string

const (
	NetworkInterfaceStatePending  NetworkInterfaceState = "Pending"
	NetworkInterfaceStatePrepared NetworkInterfaceState = "Prepared"
	NetworkInterfaceStateAttached NetworkInterfaceState = "Attached"
)

type NetworkInterfaceType string

const (
	NetworkInterfacePCIType NetworkInterfaceType = "pci"
	NetworkInterfaceTAPType NetworkInterfaceType = "tap"
)
