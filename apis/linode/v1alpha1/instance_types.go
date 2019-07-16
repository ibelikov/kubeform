package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecConfigDevicesSdf struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevicesSdg struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevicesSdh struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevicesSda struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevicesSdb struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevicesSdc struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevicesSdd struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevicesSde struct {
	DiskLabel string `json:"disk_label"`
	DiskId    int    `json:"disk_id"`
	VolumeId  int    `json:"volume_id"`
}

type InstanceSpecConfigDevices struct {
	Sdf []InstanceSpecConfigDevices `json:"sdf"`
	Sdg []InstanceSpecConfigDevices `json:"sdg"`
	Sdh []InstanceSpecConfigDevices `json:"sdh"`
	Sda []InstanceSpecConfigDevices `json:"sda"`
	Sdb []InstanceSpecConfigDevices `json:"sdb"`
	Sdc []InstanceSpecConfigDevices `json:"sdc"`
	Sdd []InstanceSpecConfigDevices `json:"sdd"`
	Sde []InstanceSpecConfigDevices `json:"sde"`
}

type InstanceSpecConfigHelpers struct {
	UpdatedbDisabled  bool `json:"updatedb_disabled"`
	Distro            bool `json:"distro"`
	ModulesDep        bool `json:"modules_dep"`
	Network           bool `json:"network"`
	DevtmpfsAutomount bool `json:"devtmpfs_automount"`
}

type InstanceSpecConfig struct {
	MemoryLimit int                  `json:"memory_limit"`
	Label       string               `json:"label"`
	Devices     []InstanceSpecConfig `json:"devices"`
	Kernel      string               `json:"kernel"`
	VirtMode    string               `json:"virt_mode"`
	Helpers     []InstanceSpecConfig `json:"helpers"`
	RunLevel    string               `json:"run_level"`
	RootDevice  string               `json:"root_device"`
	Comments    string               `json:"comments"`
}

type InstanceSpecAlerts struct {
	Cpu           int `json:"cpu"`
	NetworkIn     int `json:"network_in"`
	NetworkOut    int `json:"network_out"`
	TransferQuota int `json:"transfer_quota"`
	Io            int `json:"io"`
}

type InstanceSpecDisk struct {
	StackscriptId   int               `json:"stackscript_id"`
	StackscriptData map[string]string `json:"stackscript_data"`
	Label           string            `json:"label"`
	Id              int               `json:"id"`
	Filesystem      string            `json:"filesystem"`
	ReadOnly        bool              `json:"read_only"`
	AuthorizedKeys  []string          `json:"authorized_keys"`
	AuthorizedUsers []string          `json:"authorized_users"`
	RootPass        string            `json:"root_pass"`
	Size            int               `json:"size"`
	Image           string            `json:"image"`
}

type InstanceSpecSpecs struct {
	Disk     int `json:"disk"`
	Memory   int `json:"memory"`
	Vcpus    int `json:"vcpus"`
	Transfer int `json:"transfer"`
}

type InstanceSpecBackupsSchedule struct {
	Day    string `json:"day"`
	Window string `json:"window"`
}

type InstanceSpecBackups struct {
	Enabled  bool                  `json:"enabled"`
	Schedule []InstanceSpecBackups `json:"schedule"`
}

type InstanceSpec struct {
	Image            string            `json:"image"`
	Label            string            `json:"label"`
	Tags             []string          `json:"tags"`
	Type             string            `json:"type"`
	Status           string            `json:"status"`
	PrivateIp        bool              `json:"private_ip"`
	Group            string            `json:"group"`
	Ipv6             string            `json:"ipv6"`
	AuthorizedKeys   []string          `json:"authorized_keys"`
	Ipv4             []string          `json:"ipv4"`
	Config           []InstanceSpec    `json:"config"`
	BackupId         int               `json:"backup_id"`
	PrivateIpAddress string            `json:"private_ip_address"`
	AuthorizedUsers  []string          `json:"authorized_users"`
	StackscriptId    int               `json:"stackscript_id"`
	Region           string            `json:"region"`
	RootPass         string            `json:"root_pass"`
	SwapSize         int               `json:"swap_size"`
	BackupsEnabled   bool              `json:"backups_enabled"`
	Alerts           []InstanceSpec    `json:"alerts"`
	BootConfigLabel  string            `json:"boot_config_label"`
	Disk             []InstanceSpec    `json:"disk"`
	IpAddress        string            `json:"ip_address"`
	WatchdogEnabled  bool              `json:"watchdog_enabled"`
	StackscriptData  map[string]string `json:"stackscript_data"`
	Specs            []InstanceSpec    `json:"specs"`
	Backups          []InstanceSpec    `json:"backups"`
}



type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}