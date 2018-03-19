package nuagenetworks

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nuagenetworks/go-bambou/bambou"
	"github.com/nuagenetworks/vspk-go/vspk"
)

func dataSourceKeyServerMonitorSeed() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceKeyServerMonitorSeedRead,
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_traffic_authentication_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_traffic_encryption_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"seed_traffic_encryption_key_lifetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lifetime": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"creation_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"parent_key_server_monitor": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceKeyServerMonitorSeedRead(d *schema.ResourceData, m interface{}) error {
	filteredKeyServerMonitorSeeds := vspk.KeyServerMonitorSeedsList{}
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
	parent := &vspk.KeyServerMonitor{ID: d.Get("parent_key_server_monitor").(string)}
	filteredKeyServerMonitorSeeds, err = parent.KeyServerMonitorSeeds(fetchFilter)
	if err != nil {
		return err
	}

	KeyServerMonitorSeed := &vspk.KeyServerMonitorSeed{}

	if len(filteredKeyServerMonitorSeeds) < 1 {
		return fmt.Errorf("Your query returned no results. Please change your search criteria and try again.")
	}

	if len(filteredKeyServerMonitorSeeds) > 1 {
		return fmt.Errorf("Your query returned more than one result. Please try a more " +
			"specific search criteria.")
	}

	KeyServerMonitorSeed = filteredKeyServerMonitorSeeds[0]

	d.Set("last_updated_by", KeyServerMonitorSeed.LastUpdatedBy)
	d.Set("seed_traffic_authentication_algorithm", KeyServerMonitorSeed.SeedTrafficAuthenticationAlgorithm)
	d.Set("seed_traffic_encryption_algorithm", KeyServerMonitorSeed.SeedTrafficEncryptionAlgorithm)
	d.Set("seed_traffic_encryption_key_lifetime", KeyServerMonitorSeed.SeedTrafficEncryptionKeyLifetime)
	d.Set("lifetime", KeyServerMonitorSeed.Lifetime)
	d.Set("entity_scope", KeyServerMonitorSeed.EntityScope)
	d.Set("creation_time", KeyServerMonitorSeed.CreationTime)
	d.Set("start_time", KeyServerMonitorSeed.StartTime)
	d.Set("external_id", KeyServerMonitorSeed.ExternalID)

	d.Set("id", KeyServerMonitorSeed.Identifier())
	d.Set("parent_id", KeyServerMonitorSeed.ParentID)
	d.Set("parent_type", KeyServerMonitorSeed.ParentType)
	d.Set("owner", KeyServerMonitorSeed.Owner)

	d.SetId(KeyServerMonitorSeed.Identifier())

	return nil
}
