package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/rvichery/vspk-go/vspk"
)

func dataSourceVNFDescriptor() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVNFDescriptorRead,
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
			"cpu_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"memory_mb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"metadata_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"visible": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"associated_vnf_threshold_policy_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"storage_gb": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_vnf_catalog": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVNFDescriptorRead(d *schema.ResourceData, m interface{}) error {
	filteredVNFDescriptors := vspk.VNFDescriptorsList{}
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
	parent := &vspk.VNFCatalog{ID: d.Get("parent_vnf_catalog").(string)}
	filteredVNFDescriptors, err = parent.VNFDescriptors(fetchFilter)
	if err != nil {
		return err
	}

	VNFDescriptor := &vspk.VNFDescriptor{}

	if len(filteredVNFDescriptors) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredVNFDescriptors) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	VNFDescriptor = filteredVNFDescriptors[0]

	d.Set("cpu_count", VNFDescriptor.CPUCount)
	d.Set("name", VNFDescriptor.Name)
	d.Set("memory_mb", VNFDescriptor.MemoryMB)
	d.Set("vendor", VNFDescriptor.Vendor)
	d.Set("description", VNFDescriptor.Description)
	d.Set("metadata_id", VNFDescriptor.MetadataID)
	d.Set("visible", VNFDescriptor.Visible)
	d.Set("associated_vnf_threshold_policy_id", VNFDescriptor.AssociatedVNFThresholdPolicyID)
	d.Set("storage_gb", VNFDescriptor.StorageGB)
	d.Set("type", VNFDescriptor.Type)

	d.Set("id", VNFDescriptor.Identifier())
	d.Set("parent_id", VNFDescriptor.ParentID)
	d.Set("parent_type", VNFDescriptor.ParentType)
	d.Set("owner", VNFDescriptor.Owner)

	d.SetId(VNFDescriptor.Identifier())

	return nil
}
