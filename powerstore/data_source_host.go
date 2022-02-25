package powerstore

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceHost() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceHostRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func dataSourceHostRead(d *schema.ResourceData, meta interface{}) error {

	c := meta.(*Client)

	hostName := d.Get("name").(string)

	host, err := c.GetHost(c, c.HostURL, hostName)
	if err != nil {
		return err
	}

	d.Set("id", host.ID)
	d.SetId(host.ID)
	return nil

}
