package nuagenetworks

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/rvichery/vspk-go/vspk"
)

func resourceUplinkConnection() *schema.Resource {
	return &schema.Resource{
		Create: resourceUplinkConnectionCreate,
		Read:   resourceUplinkConnectionRead,
		Update: resourceUplinkConnectionUpdate,
		Delete: resourceUplinkConnectionDelete,
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
			"pat_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"dns_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gateway": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"advertisement_criteria": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"secondary_address": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"netmask": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"underlay_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"installer_managed": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"interface_connection_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AUTOMATIC",
			},
			"mode": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Dynamic",
			},
			"role": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "PRIMARY",
			},
			"role_order": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"download_rate_limit": {
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  8,
			},
			"uplink_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"assoc_underlay_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_bgp_neighbor_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"associated_underlay_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auxiliary_link": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"parent_vlan": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vlan_template"},
			},
			"parent_vlan_template": {
				Type:          schema.TypeString,
				Optional:      true,
				ConflictsWith: []string{"parent_vlan"},
			},
		},
	}
}

func resourceUplinkConnectionCreate(d *schema.ResourceData, m interface{}) error {

	// Initialize UplinkConnection object
	o := &vspk.UplinkConnection{}
	if attr, ok := d.GetOk("pat_enabled"); ok {
		PATEnabled := attr.(bool)
		o.PATEnabled = &PATEnabled
	}
	if attr, ok := d.GetOk("dns_address"); ok {
		o.DNSAddress = attr.(string)
	}
	if attr, ok := d.GetOk("password"); ok {
		o.Password = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("advertisement_criteria"); ok {
		o.AdvertisementCriteria = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_address"); ok {
		o.SecondaryAddress = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("underlay_enabled"); ok {
		UnderlayEnabled := attr.(bool)
		o.UnderlayEnabled = &UnderlayEnabled
	}
	if attr, ok := d.GetOk("installer_managed"); ok {
		InstallerManaged := attr.(bool)
		o.InstallerManaged = &InstallerManaged
	}
	if attr, ok := d.GetOk("interface_connection_type"); ok {
		o.InterfaceConnectionType = attr.(string)
	}
	if attr, ok := d.GetOk("mode"); ok {
		o.Mode = attr.(string)
	}
	if attr, ok := d.GetOk("role"); ok {
		o.Role = attr.(string)
	}
	if attr, ok := d.GetOk("download_rate_limit"); ok {
		o.DownloadRateLimit = attr.(float64)
	}
	if attr, ok := d.GetOk("uplink_id"); ok {
		UplinkID := attr.(int)
		o.UplinkID = &UplinkID
	}
	if attr, ok := d.GetOk("username"); ok {
		o.Username = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_underlay_id"); ok {
		o.AssocUnderlayID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_bgp_neighbor_id"); ok {
		o.AssociatedBGPNeighborID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_underlay_name"); ok {
		o.AssociatedUnderlayName = attr.(string)
	}
	if attr, ok := d.GetOk("auxiliary_link"); ok {
		AuxiliaryLink := attr.(bool)
		o.AuxiliaryLink = &AuxiliaryLink
	}
	if attr, ok := d.GetOk("parent_vlan"); ok {
		parent := &vspk.VLAN{ID: attr.(string)}
		err := parent.CreateUplinkConnection(o)
		if err != nil {
			return err
		}
	}
	if attr, ok := d.GetOk("parent_vlan_template"); ok {
		parent := &vspk.VLANTemplate{ID: attr.(string)}
		err := parent.CreateUplinkConnection(o)
		if err != nil {
			return err
		}
	}

	d.SetId(o.Identifier())
	return resourceUplinkConnectionRead(d, m)
}

func resourceUplinkConnectionRead(d *schema.ResourceData, m interface{}) error {
	o := &vspk.UplinkConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		d.SetId("")
		return nil
	}

	d.Set("pat_enabled", o.PATEnabled)
	d.Set("dns_address", o.DNSAddress)
	d.Set("password", o.Password)
	d.Set("gateway", o.Gateway)
	d.Set("address", o.Address)
	d.Set("advertisement_criteria", o.AdvertisementCriteria)
	d.Set("secondary_address", o.SecondaryAddress)
	d.Set("netmask", o.Netmask)
	d.Set("vlan_id", o.VlanId)
	d.Set("underlay_enabled", o.UnderlayEnabled)
	d.Set("installer_managed", o.InstallerManaged)
	d.Set("interface_connection_type", o.InterfaceConnectionType)
	d.Set("mode", o.Mode)
	d.Set("role", o.Role)
	d.Set("role_order", o.RoleOrder)
	d.Set("port_name", o.PortName)
	d.Set("download_rate_limit", o.DownloadRateLimit)
	d.Set("uplink_id", o.UplinkID)
	d.Set("username", o.Username)
	d.Set("assoc_underlay_id", o.AssocUnderlayID)
	d.Set("associated_bgp_neighbor_id", o.AssociatedBGPNeighborID)
	d.Set("associated_underlay_name", o.AssociatedUnderlayName)
	d.Set("auxiliary_link", o.AuxiliaryLink)

	d.Set("id", o.Identifier())
	d.Set("parent_id", o.ParentID)
	d.Set("parent_type", o.ParentType)
	d.Set("owner", o.Owner)

	return nil
}

func resourceUplinkConnectionUpdate(d *schema.ResourceData, m interface{}) error {
	o := &vspk.UplinkConnection{
		ID: d.Id(),
	}

	err := o.Fetch()
	if err != nil {
		return err
	}

	if attr, ok := d.GetOk("pat_enabled"); ok {
		PATEnabled := attr.(bool)
		o.PATEnabled = &PATEnabled
	}
	if attr, ok := d.GetOk("dns_address"); ok {
		o.DNSAddress = attr.(string)
	}
	if attr, ok := d.GetOk("password"); ok {
		o.Password = attr.(string)
	}
	if attr, ok := d.GetOk("gateway"); ok {
		o.Gateway = attr.(string)
	}
	if attr, ok := d.GetOk("address"); ok {
		o.Address = attr.(string)
	}
	if attr, ok := d.GetOk("advertisement_criteria"); ok {
		o.AdvertisementCriteria = attr.(string)
	}
	if attr, ok := d.GetOk("secondary_address"); ok {
		o.SecondaryAddress = attr.(string)
	}
	if attr, ok := d.GetOk("netmask"); ok {
		o.Netmask = attr.(string)
	}
	if attr, ok := d.GetOk("underlay_enabled"); ok {
		UnderlayEnabled := attr.(bool)
		o.UnderlayEnabled = &UnderlayEnabled
	}
	if attr, ok := d.GetOk("installer_managed"); ok {
		InstallerManaged := attr.(bool)
		o.InstallerManaged = &InstallerManaged
	}
	if attr, ok := d.GetOk("interface_connection_type"); ok {
		o.InterfaceConnectionType = attr.(string)
	}
	if attr, ok := d.GetOk("mode"); ok {
		o.Mode = attr.(string)
	}
	if attr, ok := d.GetOk("role"); ok {
		o.Role = attr.(string)
	}
	if attr, ok := d.GetOk("download_rate_limit"); ok {
		o.DownloadRateLimit = attr.(float64)
	}
	if attr, ok := d.GetOk("uplink_id"); ok {
		UplinkID := attr.(int)
		o.UplinkID = &UplinkID
	}
	if attr, ok := d.GetOk("username"); ok {
		o.Username = attr.(string)
	}
	if attr, ok := d.GetOk("assoc_underlay_id"); ok {
		o.AssocUnderlayID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_bgp_neighbor_id"); ok {
		o.AssociatedBGPNeighborID = attr.(string)
	}
	if attr, ok := d.GetOk("associated_underlay_name"); ok {
		o.AssociatedUnderlayName = attr.(string)
	}
	if attr, ok := d.GetOk("auxiliary_link"); ok {
		AuxiliaryLink := attr.(bool)
		o.AuxiliaryLink = &AuxiliaryLink
	}

	o.Save()

	return nil
}

func resourceUplinkConnectionDelete(d *schema.ResourceData, m interface{}) error {
	o := &vspk.UplinkConnection{
		ID: d.Id(),
	}

	err := o.Delete()
	if err != nil {
		return err
	}

	return nil
}
