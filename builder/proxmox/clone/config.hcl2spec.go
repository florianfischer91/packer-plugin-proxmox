// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package proxmoxclone

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	proxmox "github.com/hashicorp/packer-plugin-proxmox/builder/proxmox/common"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName           *string                            `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType         *string                            `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion         *string                            `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug               *bool                              `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce               *bool                              `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError             *string                            `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars            map[string]string                  `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars       []string                           `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	HTTPDir                   *string                            `mapstructure:"http_directory" cty:"http_directory" hcl:"http_directory"`
	HTTPContent               map[string]string                  `mapstructure:"http_content" cty:"http_content" hcl:"http_content"`
	HTTPPortMin               *int                               `mapstructure:"http_port_min" cty:"http_port_min" hcl:"http_port_min"`
	HTTPPortMax               *int                               `mapstructure:"http_port_max" cty:"http_port_max" hcl:"http_port_max"`
	HTTPAddress               *string                            `mapstructure:"http_bind_address" cty:"http_bind_address" hcl:"http_bind_address"`
	HTTPInterface             *string                            `mapstructure:"http_interface" undocumented:"true" cty:"http_interface" hcl:"http_interface"`
	BootGroupInterval         *string                            `mapstructure:"boot_keygroup_interval" cty:"boot_keygroup_interval" hcl:"boot_keygroup_interval"`
	BootWait                  *string                            `mapstructure:"boot_wait" cty:"boot_wait" hcl:"boot_wait"`
	BootCommand               []string                           `mapstructure:"boot_command" cty:"boot_command" hcl:"boot_command"`
	BootKeyInterval           *string                            `mapstructure:"boot_key_interval" cty:"boot_key_interval" hcl:"boot_key_interval"`
	Type                      *string                            `mapstructure:"communicator" cty:"communicator" hcl:"communicator"`
	PauseBeforeConnect        *string                            `mapstructure:"pause_before_connecting" cty:"pause_before_connecting" hcl:"pause_before_connecting"`
	SSHHost                   *string                            `mapstructure:"ssh_host" cty:"ssh_host" hcl:"ssh_host"`
	SSHPort                   *int                               `mapstructure:"ssh_port" cty:"ssh_port" hcl:"ssh_port"`
	SSHUsername               *string                            `mapstructure:"ssh_username" cty:"ssh_username" hcl:"ssh_username"`
	SSHPassword               *string                            `mapstructure:"ssh_password" cty:"ssh_password" hcl:"ssh_password"`
	SSHKeyPairName            *string                            `mapstructure:"ssh_keypair_name" undocumented:"true" cty:"ssh_keypair_name" hcl:"ssh_keypair_name"`
	SSHTemporaryKeyPairName   *string                            `mapstructure:"temporary_key_pair_name" undocumented:"true" cty:"temporary_key_pair_name" hcl:"temporary_key_pair_name"`
	SSHTemporaryKeyPairType   *string                            `mapstructure:"temporary_key_pair_type" cty:"temporary_key_pair_type" hcl:"temporary_key_pair_type"`
	SSHTemporaryKeyPairBits   *int                               `mapstructure:"temporary_key_pair_bits" cty:"temporary_key_pair_bits" hcl:"temporary_key_pair_bits"`
	SSHCiphers                []string                           `mapstructure:"ssh_ciphers" cty:"ssh_ciphers" hcl:"ssh_ciphers"`
	SSHClearAuthorizedKeys    *bool                              `mapstructure:"ssh_clear_authorized_keys" cty:"ssh_clear_authorized_keys" hcl:"ssh_clear_authorized_keys"`
	SSHKEXAlgos               []string                           `mapstructure:"ssh_key_exchange_algorithms" cty:"ssh_key_exchange_algorithms" hcl:"ssh_key_exchange_algorithms"`
	SSHPrivateKeyFile         *string                            `mapstructure:"ssh_private_key_file" undocumented:"true" cty:"ssh_private_key_file" hcl:"ssh_private_key_file"`
	SSHCertificateFile        *string                            `mapstructure:"ssh_certificate_file" cty:"ssh_certificate_file" hcl:"ssh_certificate_file"`
	SSHPty                    *bool                              `mapstructure:"ssh_pty" cty:"ssh_pty" hcl:"ssh_pty"`
	SSHTimeout                *string                            `mapstructure:"ssh_timeout" cty:"ssh_timeout" hcl:"ssh_timeout"`
	SSHWaitTimeout            *string                            `mapstructure:"ssh_wait_timeout" undocumented:"true" cty:"ssh_wait_timeout" hcl:"ssh_wait_timeout"`
	SSHAgentAuth              *bool                              `mapstructure:"ssh_agent_auth" undocumented:"true" cty:"ssh_agent_auth" hcl:"ssh_agent_auth"`
	SSHDisableAgentForwarding *bool                              `mapstructure:"ssh_disable_agent_forwarding" cty:"ssh_disable_agent_forwarding" hcl:"ssh_disable_agent_forwarding"`
	SSHHandshakeAttempts      *int                               `mapstructure:"ssh_handshake_attempts" cty:"ssh_handshake_attempts" hcl:"ssh_handshake_attempts"`
	SSHBastionHost            *string                            `mapstructure:"ssh_bastion_host" cty:"ssh_bastion_host" hcl:"ssh_bastion_host"`
	SSHBastionPort            *int                               `mapstructure:"ssh_bastion_port" cty:"ssh_bastion_port" hcl:"ssh_bastion_port"`
	SSHBastionAgentAuth       *bool                              `mapstructure:"ssh_bastion_agent_auth" cty:"ssh_bastion_agent_auth" hcl:"ssh_bastion_agent_auth"`
	SSHBastionUsername        *string                            `mapstructure:"ssh_bastion_username" cty:"ssh_bastion_username" hcl:"ssh_bastion_username"`
	SSHBastionPassword        *string                            `mapstructure:"ssh_bastion_password" cty:"ssh_bastion_password" hcl:"ssh_bastion_password"`
	SSHBastionInteractive     *bool                              `mapstructure:"ssh_bastion_interactive" cty:"ssh_bastion_interactive" hcl:"ssh_bastion_interactive"`
	SSHBastionPrivateKeyFile  *string                            `mapstructure:"ssh_bastion_private_key_file" cty:"ssh_bastion_private_key_file" hcl:"ssh_bastion_private_key_file"`
	SSHBastionCertificateFile *string                            `mapstructure:"ssh_bastion_certificate_file" cty:"ssh_bastion_certificate_file" hcl:"ssh_bastion_certificate_file"`
	SSHFileTransferMethod     *string                            `mapstructure:"ssh_file_transfer_method" cty:"ssh_file_transfer_method" hcl:"ssh_file_transfer_method"`
	SSHProxyHost              *string                            `mapstructure:"ssh_proxy_host" cty:"ssh_proxy_host" hcl:"ssh_proxy_host"`
	SSHProxyPort              *int                               `mapstructure:"ssh_proxy_port" cty:"ssh_proxy_port" hcl:"ssh_proxy_port"`
	SSHProxyUsername          *string                            `mapstructure:"ssh_proxy_username" cty:"ssh_proxy_username" hcl:"ssh_proxy_username"`
	SSHProxyPassword          *string                            `mapstructure:"ssh_proxy_password" cty:"ssh_proxy_password" hcl:"ssh_proxy_password"`
	SSHKeepAliveInterval      *string                            `mapstructure:"ssh_keep_alive_interval" cty:"ssh_keep_alive_interval" hcl:"ssh_keep_alive_interval"`
	SSHReadWriteTimeout       *string                            `mapstructure:"ssh_read_write_timeout" cty:"ssh_read_write_timeout" hcl:"ssh_read_write_timeout"`
	SSHRemoteTunnels          []string                           `mapstructure:"ssh_remote_tunnels" cty:"ssh_remote_tunnels" hcl:"ssh_remote_tunnels"`
	SSHLocalTunnels           []string                           `mapstructure:"ssh_local_tunnels" cty:"ssh_local_tunnels" hcl:"ssh_local_tunnels"`
	SSHPublicKey              []byte                             `mapstructure:"ssh_public_key" undocumented:"true" cty:"ssh_public_key" hcl:"ssh_public_key"`
	SSHPrivateKey             []byte                             `mapstructure:"ssh_private_key" undocumented:"true" cty:"ssh_private_key" hcl:"ssh_private_key"`
	WinRMUser                 *string                            `mapstructure:"winrm_username" cty:"winrm_username" hcl:"winrm_username"`
	WinRMPassword             *string                            `mapstructure:"winrm_password" cty:"winrm_password" hcl:"winrm_password"`
	WinRMHost                 *string                            `mapstructure:"winrm_host" cty:"winrm_host" hcl:"winrm_host"`
	WinRMNoProxy              *bool                              `mapstructure:"winrm_no_proxy" cty:"winrm_no_proxy" hcl:"winrm_no_proxy"`
	WinRMPort                 *int                               `mapstructure:"winrm_port" cty:"winrm_port" hcl:"winrm_port"`
	WinRMTimeout              *string                            `mapstructure:"winrm_timeout" cty:"winrm_timeout" hcl:"winrm_timeout"`
	WinRMUseSSL               *bool                              `mapstructure:"winrm_use_ssl" cty:"winrm_use_ssl" hcl:"winrm_use_ssl"`
	WinRMInsecure             *bool                              `mapstructure:"winrm_insecure" cty:"winrm_insecure" hcl:"winrm_insecure"`
	WinRMUseNTLM              *bool                              `mapstructure:"winrm_use_ntlm" cty:"winrm_use_ntlm" hcl:"winrm_use_ntlm"`
	ProxmoxURLRaw             *string                            `mapstructure:"proxmox_url" cty:"proxmox_url" hcl:"proxmox_url"`
	SkipCertValidation        *bool                              `mapstructure:"insecure_skip_tls_verify" cty:"insecure_skip_tls_verify" hcl:"insecure_skip_tls_verify"`
	Username                  *string                            `mapstructure:"username" cty:"username" hcl:"username"`
	Password                  *string                            `mapstructure:"password" cty:"password" hcl:"password"`
	Token                     *string                            `mapstructure:"token" cty:"token" hcl:"token"`
	Node                      *string                            `mapstructure:"node" cty:"node" hcl:"node"`
	Pool                      *string                            `mapstructure:"pool" cty:"pool" hcl:"pool"`
	TaskTimeout               *string                            `mapstructure:"task_timeout" cty:"task_timeout" hcl:"task_timeout"`
	VMName                    *string                            `mapstructure:"vm_name" cty:"vm_name" hcl:"vm_name"`
	VMID                      *int                               `mapstructure:"vm_id" cty:"vm_id" hcl:"vm_id"`
	Boot                      *string                            `mapstructure:"boot" cty:"boot" hcl:"boot"`
	Memory                    *int                               `mapstructure:"memory" cty:"memory" hcl:"memory"`
	BalloonMinimum            *int                               `mapstructure:"ballooning_minimum" cty:"ballooning_minimum" hcl:"ballooning_minimum"`
	Cores                     *int                               `mapstructure:"cores" cty:"cores" hcl:"cores"`
	CPUType                   *string                            `mapstructure:"cpu_type" cty:"cpu_type" hcl:"cpu_type"`
	Sockets                   *int                               `mapstructure:"sockets" cty:"sockets" hcl:"sockets"`
	Numa                      *bool                              `mapstructure:"numa" cty:"numa" hcl:"numa"`
	OS                        *string                            `mapstructure:"os" cty:"os" hcl:"os"`
	BIOS                      *string                            `mapstructure:"bios" cty:"bios" hcl:"bios"`
	EFIConfig                 *proxmox.FlatefiConfig             `mapstructure:"efi_config" cty:"efi_config" hcl:"efi_config"`
	EFIDisk                   *string                            `mapstructure:"efidisk" cty:"efidisk" hcl:"efidisk"`
	Machine                   *string                            `mapstructure:"machine" cty:"machine" hcl:"machine"`
	VGA                       *proxmox.FlatvgaConfig             `mapstructure:"vga" cty:"vga" hcl:"vga"`
	NICs                      []proxmox.FlatNICConfig            `mapstructure:"network_adapters" cty:"network_adapters" hcl:"network_adapters"`
	Disks                     []proxmox.FlatdiskConfig           `mapstructure:"disks" cty:"disks" hcl:"disks"`
	Serials                   []string                           `mapstructure:"serials" cty:"serials" hcl:"serials"`
	Agent                     *bool                              `mapstructure:"qemu_agent" cty:"qemu_agent" hcl:"qemu_agent"`
	SCSIController            *string                            `mapstructure:"scsi_controller" cty:"scsi_controller" hcl:"scsi_controller"`
	Onboot                    *bool                              `mapstructure:"onboot" cty:"onboot" hcl:"onboot"`
	DisableKVM                *bool                              `mapstructure:"disable_kvm" cty:"disable_kvm" hcl:"disable_kvm"`
	ProxmoxDefaultBehavior    *bool                              `mapstructure:"proxmox_default_behavior" cty:"proxmox_default_behavior" hcl:"proxmox_default_behavior"`
	TemplateName              *string                            `mapstructure:"template_name" cty:"template_name" hcl:"template_name"`
	TemplateDescription       *string                            `mapstructure:"template_description" cty:"template_description" hcl:"template_description"`
	SkipConvertToTemplate     *bool                              `mapstructure:"skip_convert_to_template" cty:"skip_convert_to_template" hcl:"skip_convert_to_template"`
	CloudInit                 *bool                              `mapstructure:"cloud_init" cty:"cloud_init" hcl:"cloud_init"`
	CloudInitStoragePool      *string                            `mapstructure:"cloud_init_storage_pool" cty:"cloud_init_storage_pool" hcl:"cloud_init_storage_pool"`
	AdditionalISOFiles        []proxmox.FlatadditionalISOsConfig `mapstructure:"additional_iso_files" cty:"additional_iso_files" hcl:"additional_iso_files"`
	VMInterface               *string                            `mapstructure:"vm_interface" cty:"vm_interface" hcl:"vm_interface"`
	CloneVM                   *string                            `mapstructure:"clone_vm" required:"true" cty:"clone_vm" hcl:"clone_vm"`
	CloneVMID                 *int                               `mapstructure:"clone_vm_id" required:"true" cty:"clone_vm_id" hcl:"clone_vm_id"`
	FullClone                 *bool                              `mapstructure:"full_clone" required:"false" cty:"full_clone" hcl:"full_clone"`
	Nameserver                *string                            `mapstructure:"nameserver" required:"false" cty:"nameserver" hcl:"nameserver"`
	Searchdomain              *string                            `mapstructure:"searchdomain" required:"false" cty:"searchdomain" hcl:"searchdomain"`
	Ipconfigs                 []FlatcloudInitIpconfig            `mapstructure:"ipconfig" required:"false" cty:"ipconfig" hcl:"ipconfig"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":            &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":          &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":          &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":                 &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":                 &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":              &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":        &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables":   &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"http_directory":               &hcldec.AttrSpec{Name: "http_directory", Type: cty.String, Required: false},
		"http_content":                 &hcldec.AttrSpec{Name: "http_content", Type: cty.Map(cty.String), Required: false},
		"http_port_min":                &hcldec.AttrSpec{Name: "http_port_min", Type: cty.Number, Required: false},
		"http_port_max":                &hcldec.AttrSpec{Name: "http_port_max", Type: cty.Number, Required: false},
		"http_bind_address":            &hcldec.AttrSpec{Name: "http_bind_address", Type: cty.String, Required: false},
		"http_interface":               &hcldec.AttrSpec{Name: "http_interface", Type: cty.String, Required: false},
		"boot_keygroup_interval":       &hcldec.AttrSpec{Name: "boot_keygroup_interval", Type: cty.String, Required: false},
		"boot_wait":                    &hcldec.AttrSpec{Name: "boot_wait", Type: cty.String, Required: false},
		"boot_command":                 &hcldec.AttrSpec{Name: "boot_command", Type: cty.List(cty.String), Required: false},
		"boot_key_interval":            &hcldec.AttrSpec{Name: "boot_key_interval", Type: cty.String, Required: false},
		"communicator":                 &hcldec.AttrSpec{Name: "communicator", Type: cty.String, Required: false},
		"pause_before_connecting":      &hcldec.AttrSpec{Name: "pause_before_connecting", Type: cty.String, Required: false},
		"ssh_host":                     &hcldec.AttrSpec{Name: "ssh_host", Type: cty.String, Required: false},
		"ssh_port":                     &hcldec.AttrSpec{Name: "ssh_port", Type: cty.Number, Required: false},
		"ssh_username":                 &hcldec.AttrSpec{Name: "ssh_username", Type: cty.String, Required: false},
		"ssh_password":                 &hcldec.AttrSpec{Name: "ssh_password", Type: cty.String, Required: false},
		"ssh_keypair_name":             &hcldec.AttrSpec{Name: "ssh_keypair_name", Type: cty.String, Required: false},
		"temporary_key_pair_name":      &hcldec.AttrSpec{Name: "temporary_key_pair_name", Type: cty.String, Required: false},
		"temporary_key_pair_type":      &hcldec.AttrSpec{Name: "temporary_key_pair_type", Type: cty.String, Required: false},
		"temporary_key_pair_bits":      &hcldec.AttrSpec{Name: "temporary_key_pair_bits", Type: cty.Number, Required: false},
		"ssh_ciphers":                  &hcldec.AttrSpec{Name: "ssh_ciphers", Type: cty.List(cty.String), Required: false},
		"ssh_clear_authorized_keys":    &hcldec.AttrSpec{Name: "ssh_clear_authorized_keys", Type: cty.Bool, Required: false},
		"ssh_key_exchange_algorithms":  &hcldec.AttrSpec{Name: "ssh_key_exchange_algorithms", Type: cty.List(cty.String), Required: false},
		"ssh_private_key_file":         &hcldec.AttrSpec{Name: "ssh_private_key_file", Type: cty.String, Required: false},
		"ssh_certificate_file":         &hcldec.AttrSpec{Name: "ssh_certificate_file", Type: cty.String, Required: false},
		"ssh_pty":                      &hcldec.AttrSpec{Name: "ssh_pty", Type: cty.Bool, Required: false},
		"ssh_timeout":                  &hcldec.AttrSpec{Name: "ssh_timeout", Type: cty.String, Required: false},
		"ssh_wait_timeout":             &hcldec.AttrSpec{Name: "ssh_wait_timeout", Type: cty.String, Required: false},
		"ssh_agent_auth":               &hcldec.AttrSpec{Name: "ssh_agent_auth", Type: cty.Bool, Required: false},
		"ssh_disable_agent_forwarding": &hcldec.AttrSpec{Name: "ssh_disable_agent_forwarding", Type: cty.Bool, Required: false},
		"ssh_handshake_attempts":       &hcldec.AttrSpec{Name: "ssh_handshake_attempts", Type: cty.Number, Required: false},
		"ssh_bastion_host":             &hcldec.AttrSpec{Name: "ssh_bastion_host", Type: cty.String, Required: false},
		"ssh_bastion_port":             &hcldec.AttrSpec{Name: "ssh_bastion_port", Type: cty.Number, Required: false},
		"ssh_bastion_agent_auth":       &hcldec.AttrSpec{Name: "ssh_bastion_agent_auth", Type: cty.Bool, Required: false},
		"ssh_bastion_username":         &hcldec.AttrSpec{Name: "ssh_bastion_username", Type: cty.String, Required: false},
		"ssh_bastion_password":         &hcldec.AttrSpec{Name: "ssh_bastion_password", Type: cty.String, Required: false},
		"ssh_bastion_interactive":      &hcldec.AttrSpec{Name: "ssh_bastion_interactive", Type: cty.Bool, Required: false},
		"ssh_bastion_private_key_file": &hcldec.AttrSpec{Name: "ssh_bastion_private_key_file", Type: cty.String, Required: false},
		"ssh_bastion_certificate_file": &hcldec.AttrSpec{Name: "ssh_bastion_certificate_file", Type: cty.String, Required: false},
		"ssh_file_transfer_method":     &hcldec.AttrSpec{Name: "ssh_file_transfer_method", Type: cty.String, Required: false},
		"ssh_proxy_host":               &hcldec.AttrSpec{Name: "ssh_proxy_host", Type: cty.String, Required: false},
		"ssh_proxy_port":               &hcldec.AttrSpec{Name: "ssh_proxy_port", Type: cty.Number, Required: false},
		"ssh_proxy_username":           &hcldec.AttrSpec{Name: "ssh_proxy_username", Type: cty.String, Required: false},
		"ssh_proxy_password":           &hcldec.AttrSpec{Name: "ssh_proxy_password", Type: cty.String, Required: false},
		"ssh_keep_alive_interval":      &hcldec.AttrSpec{Name: "ssh_keep_alive_interval", Type: cty.String, Required: false},
		"ssh_read_write_timeout":       &hcldec.AttrSpec{Name: "ssh_read_write_timeout", Type: cty.String, Required: false},
		"ssh_remote_tunnels":           &hcldec.AttrSpec{Name: "ssh_remote_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_local_tunnels":            &hcldec.AttrSpec{Name: "ssh_local_tunnels", Type: cty.List(cty.String), Required: false},
		"ssh_public_key":               &hcldec.AttrSpec{Name: "ssh_public_key", Type: cty.List(cty.Number), Required: false},
		"ssh_private_key":              &hcldec.AttrSpec{Name: "ssh_private_key", Type: cty.List(cty.Number), Required: false},
		"winrm_username":               &hcldec.AttrSpec{Name: "winrm_username", Type: cty.String, Required: false},
		"winrm_password":               &hcldec.AttrSpec{Name: "winrm_password", Type: cty.String, Required: false},
		"winrm_host":                   &hcldec.AttrSpec{Name: "winrm_host", Type: cty.String, Required: false},
		"winrm_no_proxy":               &hcldec.AttrSpec{Name: "winrm_no_proxy", Type: cty.Bool, Required: false},
		"winrm_port":                   &hcldec.AttrSpec{Name: "winrm_port", Type: cty.Number, Required: false},
		"winrm_timeout":                &hcldec.AttrSpec{Name: "winrm_timeout", Type: cty.String, Required: false},
		"winrm_use_ssl":                &hcldec.AttrSpec{Name: "winrm_use_ssl", Type: cty.Bool, Required: false},
		"winrm_insecure":               &hcldec.AttrSpec{Name: "winrm_insecure", Type: cty.Bool, Required: false},
		"winrm_use_ntlm":               &hcldec.AttrSpec{Name: "winrm_use_ntlm", Type: cty.Bool, Required: false},
		"proxmox_url":                  &hcldec.AttrSpec{Name: "proxmox_url", Type: cty.String, Required: false},
		"insecure_skip_tls_verify":     &hcldec.AttrSpec{Name: "insecure_skip_tls_verify", Type: cty.Bool, Required: false},
		"username":                     &hcldec.AttrSpec{Name: "username", Type: cty.String, Required: false},
		"password":                     &hcldec.AttrSpec{Name: "password", Type: cty.String, Required: false},
		"token":                        &hcldec.AttrSpec{Name: "token", Type: cty.String, Required: false},
		"node":                         &hcldec.AttrSpec{Name: "node", Type: cty.String, Required: false},
		"pool":                         &hcldec.AttrSpec{Name: "pool", Type: cty.String, Required: false},
		"task_timeout":                 &hcldec.AttrSpec{Name: "task_timeout", Type: cty.String, Required: false},
		"vm_name":                      &hcldec.AttrSpec{Name: "vm_name", Type: cty.String, Required: false},
		"vm_id":                        &hcldec.AttrSpec{Name: "vm_id", Type: cty.Number, Required: false},
		"boot":                         &hcldec.AttrSpec{Name: "boot", Type: cty.String, Required: false},
		"memory":                       &hcldec.AttrSpec{Name: "memory", Type: cty.Number, Required: false},
		"ballooning_minimum":           &hcldec.AttrSpec{Name: "ballooning_minimum", Type: cty.Number, Required: false},
		"cores":                        &hcldec.AttrSpec{Name: "cores", Type: cty.Number, Required: false},
		"cpu_type":                     &hcldec.AttrSpec{Name: "cpu_type", Type: cty.String, Required: false},
		"sockets":                      &hcldec.AttrSpec{Name: "sockets", Type: cty.Number, Required: false},
		"numa":                         &hcldec.AttrSpec{Name: "numa", Type: cty.Bool, Required: false},
		"os":                           &hcldec.AttrSpec{Name: "os", Type: cty.String, Required: false},
		"bios":                         &hcldec.AttrSpec{Name: "bios", Type: cty.String, Required: false},
		"efi_config":                   &hcldec.BlockSpec{TypeName: "efi_config", Nested: hcldec.ObjectSpec((*proxmox.FlatefiConfig)(nil).HCL2Spec())},
		"efidisk":                      &hcldec.AttrSpec{Name: "efidisk", Type: cty.String, Required: false},
		"machine":                      &hcldec.AttrSpec{Name: "machine", Type: cty.String, Required: false},
		"vga":                          &hcldec.BlockSpec{TypeName: "vga", Nested: hcldec.ObjectSpec((*proxmox.FlatvgaConfig)(nil).HCL2Spec())},
		"network_adapters":             &hcldec.BlockListSpec{TypeName: "network_adapters", Nested: hcldec.ObjectSpec((*proxmox.FlatNICConfig)(nil).HCL2Spec())},
		"disks":                        &hcldec.BlockListSpec{TypeName: "disks", Nested: hcldec.ObjectSpec((*proxmox.FlatdiskConfig)(nil).HCL2Spec())},
		"serials":                      &hcldec.AttrSpec{Name: "serials", Type: cty.List(cty.String), Required: false},
		"qemu_agent":                   &hcldec.AttrSpec{Name: "qemu_agent", Type: cty.Bool, Required: false},
		"scsi_controller":              &hcldec.AttrSpec{Name: "scsi_controller", Type: cty.String, Required: false},
		"onboot":                       &hcldec.AttrSpec{Name: "onboot", Type: cty.Bool, Required: false},
		"disable_kvm":                  &hcldec.AttrSpec{Name: "disable_kvm", Type: cty.Bool, Required: false},
		"proxmox_default_behavior":     &hcldec.AttrSpec{Name: "proxmox_default_behavior", Type: cty.Bool, Required: false},
		"template_name":                &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"template_description":         &hcldec.AttrSpec{Name: "template_description", Type: cty.String, Required: false},
		"skip_convert_to_template":     &hcldec.AttrSpec{Name: "skip_convert_to_template", Type: cty.Bool, Required: false},
		"cloud_init":                   &hcldec.AttrSpec{Name: "cloud_init", Type: cty.Bool, Required: false},
		"cloud_init_storage_pool":      &hcldec.AttrSpec{Name: "cloud_init_storage_pool", Type: cty.String, Required: false},
		"additional_iso_files":         &hcldec.BlockListSpec{TypeName: "additional_iso_files", Nested: hcldec.ObjectSpec((*proxmox.FlatadditionalISOsConfig)(nil).HCL2Spec())},
		"vm_interface":                 &hcldec.AttrSpec{Name: "vm_interface", Type: cty.String, Required: false},
		"clone_vm":                     &hcldec.AttrSpec{Name: "clone_vm", Type: cty.String, Required: false},
		"clone_vm_id":                  &hcldec.AttrSpec{Name: "clone_vm_id", Type: cty.Number, Required: false},
		"full_clone":                   &hcldec.AttrSpec{Name: "full_clone", Type: cty.Bool, Required: false},
		"nameserver":                   &hcldec.AttrSpec{Name: "nameserver", Type: cty.String, Required: false},
		"searchdomain":                 &hcldec.AttrSpec{Name: "searchdomain", Type: cty.String, Required: false},
		"ipconfig":                     &hcldec.BlockListSpec{TypeName: "ipconfig", Nested: hcldec.ObjectSpec((*FlatcloudInitIpconfig)(nil).HCL2Spec())},
	}
	return s
}

// FlatcloudInitIpconfig is an auto-generated flat version of cloudInitIpconfig.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatcloudInitIpconfig struct {
	Ip       *string `mapstructure:"ip" required:"false" cty:"ip" hcl:"ip"`
	Gateway  *string `mapstructure:"gateway" required:"false" cty:"gateway" hcl:"gateway"`
	Ip6      *string `mapstructure:"ip6" required:"false" cty:"ip6" hcl:"ip6"`
	Gateway6 *string `mapstructure:"gateway6" required:"false" cty:"gateway6" hcl:"gateway6"`
}

// FlatMapstructure returns a new FlatcloudInitIpconfig.
// FlatcloudInitIpconfig is an auto-generated flat version of cloudInitIpconfig.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*cloudInitIpconfig) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatcloudInitIpconfig)
}

// HCL2Spec returns the hcl spec of a cloudInitIpconfig.
// This spec is used by HCL to read the fields of cloudInitIpconfig.
// The decoded values from this spec will then be applied to a FlatcloudInitIpconfig.
func (*FlatcloudInitIpconfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ip":       &hcldec.AttrSpec{Name: "ip", Type: cty.String, Required: false},
		"gateway":  &hcldec.AttrSpec{Name: "gateway", Type: cty.String, Required: false},
		"ip6":      &hcldec.AttrSpec{Name: "ip6", Type: cty.String, Required: false},
		"gateway6": &hcldec.AttrSpec{Name: "gateway6", Type: cty.String, Required: false},
	}
	return s
}
