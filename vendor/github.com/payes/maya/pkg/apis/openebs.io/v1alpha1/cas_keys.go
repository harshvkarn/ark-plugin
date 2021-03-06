/*
Copyright 2018 The OpenEBS Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// CASVolumeType represents a valid cas volume
type CASVolumeType string

const (
	// JivaVolume represents a volume based on jiva
	JivaVolume CASVolumeType = "jiva"

	// CstorVolume represents a volume based on cstor
	CstorVolume CASVolumeType = "cstor"
)

// CASKey represents the key used either in resource annotation or label
type CASKey string

const (
	// CreatePoolCASTemplateKey is the cas template annotation key whose value is
	// the name of cas template that will be used to provision a storagepool
	// If this field is present on SPC then CAS template pointed by this field is used to provision
	// the storage pool else default CAS template is used.
	CreatePoolCASTemplateKey CASKey = "cas.openebs.io/create-pool-template"

	// DeletePoolCASTemplateKey is the cas template annotation key whose value is
	// the name of cas template that will be used to delete a storagepool
	DeletePoolCASTemplateKey CASKey = "cas.openebs.io/delete-pool-template"

	// OpenEBSVersionKey is the label key which provides the installed version of
	// OpenEBS
	OpenEBSVersionKey CASKey = "openebs.io/version"

	// CASConfigKey is the key to fetch configurations w.r.t a CAS entity
	CASConfigKey CASKey = "cas.openebs.io/config"

	// NamespaceKey is the key to fetch cas entity's namespace
	NamespaceKey CASKey = "openebs.io/namespace"

	// PersistentVolumeClaimKey is the key to fetch name of PersistentVolumeClaim
	PersistentVolumeClaimKey CASKey = "openebs.io/persistentvolumeclaim"

	// StorageClassKey is the key to fetch name of StorageClass
	StorageClassKey CASKey = "openebs.io/storageclass"

	// CASTypeKey is the key to fetch storage engine for the volume
	CASTypeKey CASKey = "openebs.io/cas-type"

	// StorageClassHeaderKey is the key to fetch name of StorageClass
	// This key is present only in get request headers
	StorageClassHeaderKey CASKey = "storageclass"
)

// CASPlainKey represents a openebs key used either in resource annotation
// or label
//
// NOTE:
//  PlainKey (i.e. without 'openebs.io/' ) helps to parse key via
// go templating
type CASPlainKey string

const (
	// OpenEBSVersionPlainKey is the label key which provides the installed
	// version of OpenEBS
	OpenEBSVersionPlainKey CASPlainKey = "version"

	// CASTNamePlainKey is the key to fetch name of CAS template
	CASTNamePlainKey CASPlainKey = "castName"
)

// KubePlainKey represents a kubernetes key used either in resource annotation
// or label
//
// NOTE:
//  PlainKey (i.e. without 'kubernetes.io/' ) helps to parse key via
// go templating
type KubePlainKey string

const (
	// KubeServerVersionPlainKey is the key to fetch Kubernetes server version
	KubeServerVersionPlainKey KubePlainKey = "kubeVersion"
)

// DeprecatedKey is a typed string to represent deprecated annotations' or
// labels' key
type DeprecatedKey string

const (
	// CapacityDeprecatedKey is a label key used to set volume capacity
	//
	// NOTE:
	//  Deprecated in favour of CASVolume.Spec.Capacity
	CapacityDeprecatedKey DeprecatedKey = "volumeprovisioner.mapi.openebs.io/storage-size"

	// NamespaceDeprecatedKey is the key to fetch volume's namespace
	//
	// NOTE:
	//  Deprecated in favour of NamespaceKey
	NamespaceDeprecatedKey DeprecatedKey = "k8s.io/namespace"

	// PersistentVolumeClaimDeprecatedKey is the key to fetch volume's PVC
	//
	// NOTE:
	//  Deprecated in favour of PersistentVolumeClaimKey
	PersistentVolumeClaimDeprecatedKey DeprecatedKey = "k8s.io/pvc"

	// StorageClassDeprecatedKey is the key to fetch name of StorageClass
	//
	// NOTE:
	//  Deprecated in favour of StorageClassKey
	StorageClassDeprecatedKey DeprecatedKey = "k8s.io/storage-class"
)
