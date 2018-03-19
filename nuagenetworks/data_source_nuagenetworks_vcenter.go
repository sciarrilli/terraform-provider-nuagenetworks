package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceVCenter() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVCenterRead,
		Schema: map[string]*schema.Schema{
			"filter": dataSourceFiltersSchema(),
			"parent_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_configuration_time_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"v_require_nuage_metadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_dns1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_dns2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datapath_sync_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_data_uplink_dhcp_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secondary_data_uplink_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"secondary_data_uplink_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_data_uplink_primary_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_secondary_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"secondary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"secondary_nuage_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_size_in_gb": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_syslog_server_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_syslog_server_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"remote_syslog_server_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"generic_split_activation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"separate_data_network": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"personality": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"destination_mirror_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_server_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_server_listen_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"metadata_server_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"metadata_service_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"network_uplink_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_uplink_interface_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_uplink_interface_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_uplink_interface_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"revertive_controller_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"revertive_timer": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"nfs_log_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nfs_mount_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_dns1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_dns2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhcp_relay_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mirror_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_gro_on_datapath": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disable_lro_on_datapath": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"disable_network_discovery": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"old_agency_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_data_dhcp": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_mgmt_dhcp": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"flow_eviction_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vm_network_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_vrs_resource_reservation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configured_metrics_push_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"connection_status": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"portgroup_metadata": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"host_level_management": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"nova_client_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"nova_identity_url_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_auth_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_endpoint": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_tenant": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_service_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_metadata_shared_secret": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_os_keystone_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_project_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_project_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_region_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nova_user_domain_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_package_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_package_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_package_username": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"upgrade_script_time_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cpu_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_data_uplink_underlay_id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"primary_nuage_controller": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_config_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vrs_user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_route": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_route_gateway": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_route_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server1": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ntp_server2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"https_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"mtu": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"multi_vmssupport": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"multicast_receive_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_receive_interface_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_receive_interface_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_receive_range": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_send_interface": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_send_interface_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_send_interface_netmask": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"multicast_source_portgroup": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"customized_script_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"auto_resolve_frequency": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ovf_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"avrs_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"avrs_profile": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceVCenterRead(d *schema.ResourceData, m interface{}) error {
	filteredVCenters := vspk.VCentersList{}
	err := &bambou.Error{}
	fetchFilter := &bambou.FetchingInfo{}

	filters, filtersOk := d.GetOk("filter")
	if filtersOk {
		fetchFilter = bambou.NewFetchingInfo()
		for _, v := range filters.(*schema.Set).List() {
			m := v.(map[string]interface{})
			if fetchFilter.Filter != "" {
				fetchFilter.Filter = fmt.Sprintf("%s AND %s %s '%s'", fetchFilter.Filter, m["key"].(string), m["operator"].(string), m["value"].(string))
			} else {
				fetchFilter.Filter = fmt.Sprintf("%s %s '%s'", m["key"].(string), m["operator"].(string), m["value"].(string))
			}

		}
	}
	parent := m.(*vspk.Me)
	filteredVCenters, err = parent.VCenters(fetchFilter)
	if err != nil {
		return err
	}

	VCenter := &vspk.VCenter{}

	if len(filteredVCenters) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVCenters) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VCenter = filteredVCenters[0]

	d.Set("vrs_configuration_time_limit", VCenter.VRSConfigurationTimeLimit)
	d.Set("v_require_nuage_metadata", VCenter.VRequireNuageMetadata)
	d.Set("name", VCenter.Name)
	d.Set("password", VCenter.Password)
	d.Set("last_updated_by", VCenter.LastUpdatedBy)
	d.Set("data_dns1", VCenter.DataDNS1)
	d.Set("data_dns2", VCenter.DataDNS2)
	d.Set("data_gateway", VCenter.DataGateway)
	d.Set("data_network_portgroup", VCenter.DataNetworkPortgroup)
	d.Set("datapath_sync_timeout", VCenter.DatapathSyncTimeout)
	d.Set("secondary_data_uplink_dhcp_enabled", VCenter.SecondaryDataUplinkDHCPEnabled)
	d.Set("secondary_data_uplink_enabled", VCenter.SecondaryDataUplinkEnabled)
	d.Set("secondary_data_uplink_interface", VCenter.SecondaryDataUplinkInterface)
	d.Set("secondary_data_uplink_mtu", VCenter.SecondaryDataUplinkMTU)
	d.Set("secondary_data_uplink_primary_controller", VCenter.SecondaryDataUplinkPrimaryController)
	d.Set("secondary_data_uplink_secondary_controller", VCenter.SecondaryDataUplinkSecondaryController)
	d.Set("secondary_data_uplink_underlay_id", VCenter.SecondaryDataUplinkUnderlayID)
	d.Set("secondary_nuage_controller", VCenter.SecondaryNuageController)
	d.Set("memory_size_in_gb", VCenter.MemorySizeInGB)
	d.Set("remote_syslog_server_ip", VCenter.RemoteSyslogServerIP)
	d.Set("remote_syslog_server_port", VCenter.RemoteSyslogServerPort)
	d.Set("remote_syslog_server_type", VCenter.RemoteSyslogServerType)
	d.Set("generic_split_activation", VCenter.GenericSplitActivation)
	d.Set("separate_data_network", VCenter.SeparateDataNetwork)
	d.Set("personality", VCenter.Personality)
	d.Set("description", VCenter.Description)
	d.Set("destination_mirror_port", VCenter.DestinationMirrorPort)
	d.Set("metadata_server_ip", VCenter.MetadataServerIP)
	d.Set("metadata_server_listen_port", VCenter.MetadataServerListenPort)
	d.Set("metadata_server_port", VCenter.MetadataServerPort)
	d.Set("metadata_service_enabled", VCenter.MetadataServiceEnabled)
	d.Set("network_uplink_interface", VCenter.NetworkUplinkInterface)
	d.Set("network_uplink_interface_gateway", VCenter.NetworkUplinkInterfaceGateway)
	d.Set("network_uplink_interface_ip", VCenter.NetworkUplinkInterfaceIp)
	d.Set("network_uplink_interface_netmask", VCenter.NetworkUplinkInterfaceNetmask)
	d.Set("revertive_controller_enabled", VCenter.RevertiveControllerEnabled)
	d.Set("revertive_timer", VCenter.RevertiveTimer)
	d.Set("nfs_log_server", VCenter.NfsLogServer)
	d.Set("nfs_mount_path", VCenter.NfsMountPath)
	d.Set("mgmt_dns1", VCenter.MgmtDNS1)
	d.Set("mgmt_dns2", VCenter.MgmtDNS2)
	d.Set("mgmt_gateway", VCenter.MgmtGateway)
	d.Set("mgmt_network_portgroup", VCenter.MgmtNetworkPortgroup)
	d.Set("dhcp_relay_server", VCenter.DhcpRelayServer)
	d.Set("mirror_network_portgroup", VCenter.MirrorNetworkPortgroup)
	d.Set("disable_gro_on_datapath", VCenter.DisableGROOnDatapath)
	d.Set("disable_lro_on_datapath", VCenter.DisableLROOnDatapath)
	d.Set("disable_network_discovery", VCenter.DisableNetworkDiscovery)
	d.Set("site_id", VCenter.SiteId)
	d.Set("old_agency_name", VCenter.OldAgencyName)
	d.Set("allow_data_dhcp", VCenter.AllowDataDHCP)
	d.Set("allow_mgmt_dhcp", VCenter.AllowMgmtDHCP)
	d.Set("flow_eviction_threshold", VCenter.FlowEvictionThreshold)
	d.Set("vm_network_portgroup", VCenter.VmNetworkPortgroup)
	d.Set("enable_vrs_resource_reservation", VCenter.EnableVRSResourceReservation)
	d.Set("entity_scope", VCenter.EntityScope)
	d.Set("configured_metrics_push_interval", VCenter.ConfiguredMetricsPushInterval)
	d.Set("connection_status", VCenter.ConnectionStatus)
	d.Set("portgroup_metadata", VCenter.PortgroupMetadata)
	d.Set("host_level_management", VCenter.HostLevelManagement)
	d.Set("nova_client_version", VCenter.NovaClientVersion)
	d.Set("nova_identity_url_version", VCenter.NovaIdentityURLVersion)
	d.Set("nova_metadata_service_auth_url", VCenter.NovaMetadataServiceAuthUrl)
	d.Set("nova_metadata_service_endpoint", VCenter.NovaMetadataServiceEndpoint)
	d.Set("nova_metadata_service_password", VCenter.NovaMetadataServicePassword)
	d.Set("nova_metadata_service_tenant", VCenter.NovaMetadataServiceTenant)
	d.Set("nova_metadata_service_username", VCenter.NovaMetadataServiceUsername)
	d.Set("nova_metadata_shared_secret", VCenter.NovaMetadataSharedSecret)
	d.Set("nova_os_keystone_username", VCenter.NovaOSKeystoneUsername)
	d.Set("nova_project_domain_name", VCenter.NovaProjectDomainName)
	d.Set("nova_project_name", VCenter.NovaProjectName)
	d.Set("nova_region_name", VCenter.NovaRegionName)
	d.Set("nova_user_domain_name", VCenter.NovaUserDomainName)
	d.Set("ip_address", VCenter.IpAddress)
	d.Set("upgrade_package_password", VCenter.UpgradePackagePassword)
	d.Set("upgrade_package_url", VCenter.UpgradePackageURL)
	d.Set("upgrade_package_username", VCenter.UpgradePackageUsername)
	d.Set("upgrade_script_time_limit", VCenter.UpgradeScriptTimeLimit)
	d.Set("cpu_count", VCenter.CpuCount)
	d.Set("primary_data_uplink_underlay_id", VCenter.PrimaryDataUplinkUnderlayID)
	d.Set("primary_nuage_controller", VCenter.PrimaryNuageController)
	d.Set("vrs_config_id", VCenter.VrsConfigID)
	d.Set("vrs_password", VCenter.VrsPassword)
	d.Set("vrs_user_name", VCenter.VrsUserName)
	d.Set("user_name", VCenter.UserName)
	d.Set("static_route", VCenter.StaticRoute)
	d.Set("static_route_gateway", VCenter.StaticRouteGateway)
	d.Set("static_route_netmask", VCenter.StaticRouteNetmask)
	d.Set("ntp_server1", VCenter.NtpServer1)
	d.Set("ntp_server2", VCenter.NtpServer2)
	d.Set("http_port", VCenter.HttpPort)
	d.Set("https_port", VCenter.HttpsPort)
	d.Set("mtu", VCenter.Mtu)
	d.Set("multi_vmssupport", VCenter.MultiVMSsupport)
	d.Set("multicast_receive_interface", VCenter.MulticastReceiveInterface)
	d.Set("multicast_receive_interface_ip", VCenter.MulticastReceiveInterfaceIP)
	d.Set("multicast_receive_interface_netmask", VCenter.MulticastReceiveInterfaceNetmask)
	d.Set("multicast_receive_range", VCenter.MulticastReceiveRange)
	d.Set("multicast_send_interface", VCenter.MulticastSendInterface)
	d.Set("multicast_send_interface_ip", VCenter.MulticastSendInterfaceIP)
	d.Set("multicast_send_interface_netmask", VCenter.MulticastSendInterfaceNetmask)
	d.Set("multicast_source_portgroup", VCenter.MulticastSourcePortgroup)
	d.Set("customized_script_url", VCenter.CustomizedScriptURL)
	d.Set("auto_resolve_frequency", VCenter.AutoResolveFrequency)
	d.Set("ovf_url", VCenter.OvfURL)
	d.Set("avrs_enabled", VCenter.AvrsEnabled)
	d.Set("avrs_profile", VCenter.AvrsProfile)
	d.Set("external_id", VCenter.ExternalID)

	d.Set("id", VCenter.Identifier())
	d.Set("parent_id", VCenter.ParentID)
	d.Set("parent_type", VCenter.ParentType)
	d.Set("owner", VCenter.Owner)

	d.SetId(VCenter.Identifier())

	return nil
}
