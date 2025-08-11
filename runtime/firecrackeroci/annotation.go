// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package firecrackeroci

import (
	"context"
    "strconv"

	"github.com/containerd/containerd/containers"
	"github.com/containerd/containerd/oci"
)

const (
	// VMIDAnnotationKey is the key specified in an OCI-runtime config annotation section
	// specifying the ID of the VM in which the container should be spun up.
	VMIDAnnotationKey = "aws.firecracker.vm.id"

    // VMMemoryMiBAnnotationKey is the key specified in an OCI-runtime config annotation section
    // specifying the memory size (in MiB) of the Firecracker microVM to run the container in.
    VMMemoryMiBAnnotationKey = "aws.firecracker.vm.mem_mib"
)

// WithVMID annotates a containerd client's container object with a given firecracker VMID.
func WithVMID(vmID string) oci.SpecOpts {
	return func(_ context.Context, _ oci.Client, _ *containers.Container, s *oci.Spec) error {
		if s.Annotations == nil {
			s.Annotations = make(map[string]string)
		}

		s.Annotations[VMIDAnnotationKey] = vmID
		return nil
	}
}

// WithVMMemoryMiB annotates a containerd client's container object with the desired Firecracker VM memory size in MiB.
func WithVMMemoryMiB(memoryMiB uint32) oci.SpecOpts {
    return func(_ context.Context, _ oci.Client, _ *containers.Container, s *oci.Spec) error {
        if s.Annotations == nil {
            s.Annotations = make(map[string]string)
        }

        s.Annotations[VMMemoryMiBAnnotationKey] = strconv.FormatUint(uint64(memoryMiB), 10)
        return nil
    }
}
