package fortios

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFirewallSecurityPolicySeq() *schema.Resource {
	return &schema.Resource{
		Create: resourceFirewallSecurityPolicySeqCreateUpdate,
		Read:   resourceFirewallSecurityPolicySeqRead,
		Update: resourceFirewallSecurityPolicySeqCreateUpdate,
		Delete: resourceFirewallSecurityPolicySeqDel,

		Schema: map[string]*schema.Schema{
			"policy_src_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"policy_dst_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"alter_position": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceFirewallSecurityPolicySeqCreateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	c.Retries = 1

	//Get Params from d
	srcIdPatch := d.Get("policy_src_id").(int)
	dstIdPatch := d.Get("policy_dst_id").(int)
	alterPos := d.Get("alter_position").(string)

	srcId := strconv.Itoa(srcIdPatch)
	dstId := strconv.Itoa(dstIdPatch)

	if alterPos != "before" && alterPos != "after" {
		return fmt.Errorf("<alter_position> param should be only 'after' or 'before'")
	}

	err := c.CreateUpdateFirewallSecurityPolicySeq(srcId, dstId, alterPos)
	if err != nil {
		return fmt.Errorf("Error Altering Firewall Security Policy Sequence: %s", err)
	}

	d.SetId(srcId)

	return nil
}

// Not suitable operation
func resourceFirewallSecurityPolicySeqRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	return c.ReadFirewallSecurityPolicySeq()
}

// Not suitable operation
func resourceFirewallSecurityPolicySeqDel(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if c == nil {
		return fmt.Errorf("FortiOS connection did not initialize successfully!")
	}

	return c.DelFirewallSecurityPolicySeq()
}
