package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceEnterpriseProfile() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEnterpriseProfileRead,
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
			"bgp_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dhcp_lease_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vnf_management_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"receive_multi_cast_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_multi_cast_list_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_advanced_qos_configuration": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_gateway_management": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allow_trusted_forwarding_class": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"allowed_forwarding_classes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"floating_ips_quota": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enable_application_performance_management": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"encryption_management_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
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

func dataSourceEnterpriseProfileRead(d *schema.ResourceData, m interface{}) error {
	filteredEnterpriseProfiles := vspk.EnterpriseProfilesList{}
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
	filteredEnterpriseProfiles, err = parent.EnterpriseProfiles(fetchFilter)
	if err != nil {
		return err
	}

	EnterpriseProfile := &vspk.EnterpriseProfile{}

	if len(filteredEnterpriseProfiles) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredEnterpriseProfiles) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	EnterpriseProfile = filteredEnterpriseProfiles[0]

	d.Set("bgp_enabled", EnterpriseProfile.BGPEnabled)
	d.Set("dhcp_lease_interval", EnterpriseProfile.DHCPLeaseInterval)
	d.Set("vnf_management_enabled", EnterpriseProfile.VNFManagementEnabled)
	d.Set("name", EnterpriseProfile.Name)
	d.Set("last_updated_by", EnterpriseProfile.LastUpdatedBy)
	d.Set("receive_multi_cast_list_id", EnterpriseProfile.ReceiveMultiCastListID)
	d.Set("send_multi_cast_list_id", EnterpriseProfile.SendMultiCastListID)
	d.Set("description", EnterpriseProfile.Description)
	d.Set("allow_advanced_qos_configuration", EnterpriseProfile.AllowAdvancedQOSConfiguration)
	d.Set("allow_gateway_management", EnterpriseProfile.AllowGatewayManagement)
	d.Set("allow_trusted_forwarding_class", EnterpriseProfile.AllowTrustedForwardingClass)
	d.Set("allowed_forwarding_classes", EnterpriseProfile.AllowedForwardingClasses)
	d.Set("floating_ips_quota", EnterpriseProfile.FloatingIPsQuota)
	d.Set("enable_application_performance_management", EnterpriseProfile.EnableApplicationPerformanceManagement)
	d.Set("encryption_management_mode", EnterpriseProfile.EncryptionManagementMode)
	d.Set("entity_scope", EnterpriseProfile.EntityScope)
	d.Set("external_id", EnterpriseProfile.ExternalID)

	d.Set("id", EnterpriseProfile.Identifier())
	d.Set("parent_id", EnterpriseProfile.ParentID)
	d.Set("parent_type", EnterpriseProfile.ParentType)
	d.Set("owner", EnterpriseProfile.Owner)

	d.SetId(EnterpriseProfile.Identifier())

	return nil
}
