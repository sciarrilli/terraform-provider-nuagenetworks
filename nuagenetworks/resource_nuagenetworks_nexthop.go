package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceNextHop() *schema.Resource {
	return &schema.Resource{
		Create: resourceNextHopCreate,
		Read:   resourceNextHopRead,
		Update: resourceNextHopUpdate,
		Delete: resourceNextHopDelete,
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
			"last_updated_by": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entity_scope": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"route_distinguisher": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"parent_link": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceNextHopCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize NextHop object
	o := &vspk.NextHop{}
	if attr, ok := d.GetOk("route_distinguisher"); ok {
		o.RouteDistinguisher = attr.(string)
	}
	if attr, ok := d.GetOk("ip"); ok {
		o.Ip = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}
	parent := &vspk.Link{ID: d.Get("parent_link").(string)}
	err := parent.CreateNextHop(o)
	if err != nil {
		return err
	}

	d.SetId(o.Identifier())
	return resourceNextHopRead(d, m)
}

func resourceNextHopRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NextHop{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("last_updated_by", o.LastUpdatedBy)
	d.Set("entity_scope", o.EntityScope)
	d.Set("route_distinguisher", o.RouteDistinguisher)
	d.Set("ip", o.Ip)
	d.Set("external_id", o.ExternalID)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceNextHopUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NextHop{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("route_distinguisher"); ok {
		o.RouteDistinguisher = attr.(string)
	}
	if attr, ok := d.GetOk("ip"); ok {
		o.Ip = attr.(string)
	}
	if attr, ok := d.GetOk("external_id"); ok {
		o.ExternalID = attr.(string)
	}

	o.Save()

	return nil
}

func resourceNextHopDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.NextHop{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
