package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourcePATIPEntry() *schema.Resource {
	return &schema.Resource{
		Create: resourcePATIPEntryCreate,
		Read:   resourcePATIPEntryRead,
		Update: resourcePATIPEntryUpdate,
		Delete: resourcePATIPEntryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_type": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"owner": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pat_centralized": {
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ip_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "IPV4",
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"associated_domain_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hypervisor_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"parent_shared_network_resource": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourcePATIPEntryCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize PATIPEntry object
	o := &vspk.PATIPEntry{}
	if attr, ok := d.GetOk("pat_centralized"); ok {
		PATCentralized := attr.(bool)
		o.PATCentralized = &PATCentralized
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_domain_id"); ok {
		o.AssociatedDomainID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("hypervisor_id"); ok {
		o.HypervisorID = attr.(string)
	}
	parent := &vspk.SharedNetworkResource{ID: d.Get("parent_shared_network_resource").(string)}
	err := parent.CreatePATIPEntry(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourcePATIPEntryRead(d, m)
}

func resourcePATIPEntryRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PATIPEntry{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("pat_centralized", o.PATCentralized)
	d.Set("ip_address", o.IPAddress)
	d.Set("ip_type", o.IPType)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("entity_scope", o.EntityScope)
	d.Set("associated_domain_id", o.AssociatedDomainID)
	d.Set("external_id", o.ExternalID)
	d.Set("hypervisor_id", o.HypervisorID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourcePATIPEntryUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PATIPEntry{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("pat_centralized"); ok {
		PATCentralized := attr.(bool)
		o.PATCentralized = &PATCentralized
	}
	if attr, ok := d.GetOk("ip_address"); ok {
		o.IPAddress = attr.(string)
	}
	if attr, ok := d.GetOk("ip_type"); ok {
		o.IPType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_domain_id"); ok {
		o.AssociatedDomainID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("hypervisor_id"); ok {
		o.HypervisorID = attr.(string)
	}

	o.Save()

	return nil
}

func resourcePATIPEntryDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.PATIPEntry{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
