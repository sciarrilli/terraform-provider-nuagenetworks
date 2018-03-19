package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceDomainFIPAclTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceDomainFIPAclTemplateCreate,
		Read:   resourceDomainFIPAclTemplateRead,
		Update: resourceDomainFIPAclTemplateUpdate,
		Delete: resourceDomainFIPAclTemplateDelete,
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
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"active": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_allow_ip": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"default_allow_non_ip": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"policy_state": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"priority_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"associated_live_entity_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_domain": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain_template"},
			},
			"parent_domain_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_domain"},
			},
		},
	}
}

func resourceDomainFIPAclTemplateCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize DomainFIPAclTemplate object
	o := &vspk.DomainFIPAclTemplate{}
	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("active"); ok {
		o.Active = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_ip"); ok {
		o.DefaultAllowIP = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_non_ip"); ok {
		o.DefaultAllowNonIP = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("entries"); ok {
		o.Entries = attr.([]interface{})
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("priority_type"); ok {
		o.PriorityType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	if attr, ok := d.GetOk("parent_domain"); ok {
		parent := &vspk.Domain{ID: attr.(string)}
		err := parent.CreateDomainFIPAclTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_me"); ok {
		parent := &vspk.Me{ID: attr.(string)}
		err := parent.CreateDomainFIPAclTemplate(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_domain_template"); ok {
		parent := &vspk.DomainTemplate{ID: attr.(string)}
		err := parent.CreateDomainFIPAclTemplate(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceDomainFIPAclTemplateRead(d, m)
}

func resourceDomainFIPAclTemplateRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DomainFIPAclTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("name", o.Name)
	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("active", o.Active)
	d.Set("default_allow_ip", o.DefaultAllowIP)
	d.Set("default_allow_non_ip", o.DefaultAllowNonIP)
	d.Set("description", o.Description)
	d.Set("entity_scope", o.EntityScope)
	d.Set("entries", o.Entries)
	d.Set("policy_state", o.PolicyState)
	d.Set("priority", o.Priority)
	d.Set("priority_type", o.PriorityType)
	d.Set("associated_live_entity_id", o.AssociatedLiveEntityID)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceDomainFIPAclTemplateUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DomainFIPAclTemplate{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("name"); ok {
		o.Name = attr.(string)
	}
	if attr, ok := d.GetOk("active"); ok {
		o.Active = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_ip"); ok {
		o.DefaultAllowIP = attr.(bool)
	}
	if attr, ok := d.GetOk("default_allow_non_ip"); ok {
		o.DefaultAllowNonIP = attr.(bool)
	}
	if attr, ok := d.GetOk("description"); ok {
		o.Description = attr.(string)
	}
	if attr, ok := d.GetOk("entries"); ok {
		o.Entries = attr.([]interface{})
	}
	if attr, ok := d.GetOk("policy_state"); ok {
		o.PolicyState = attr.(string)
	}
	if attr, ok := d.GetOk("priority"); ok {
		o.Priority = attr.(int)
	}
	if attr, ok := d.GetOk("priority_type"); ok {
		o.PriorityType = attr.(string)
	}
	if attr, ok := d.GetOk("associated_live_entity_id"); ok {
		o.AssociatedLiveEntityID = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceDomainFIPAclTemplateDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.DomainFIPAclTemplate{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
