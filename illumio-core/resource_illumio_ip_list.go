package illumiocore

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioIPList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioIPListCreate,
		ReadContext:   resourceIllumioIPListRead,
		UpdateContext: resourceIllumioIPListUpdate,
		DeleteContext: resourceIllumioIPListDelete,
		SchemaVersion: version,
		Description:   "Manages Illumio IP List",

		Schema: map[string]*schema.Schema{
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URI of this IPList",
			},
			"name": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "Name of the IPList",
				ValidateDiagFunc: nameValidation,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the IPList",
			},
			"ip_ranges": {
				Type:         schema.TypeSet,
				Optional:     true,
				AtLeastOneOf: []string{"ip_ranges", "fqdns"},
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Desciption of IP Range",
						},
						"from_ip": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "IP address or a low end of IP range. Might be specified with CIDR notation",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.Any(validation.IsIPAddress, validation.IsCIDR),
							),
						},
						"to_ip": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "High end of an IP range. Might be specified with CIDR notation",
							ValidateDiagFunc: validation.ToDiagFunc(
								validation.IsIPAddress,
							),
						},
						"exclusion": {
							Type:        schema.TypeBool,
							Optional:    true,
							Description: "Whether this IP address is an exclusion. Exclusions must be a strict subset of inclusive IP addresses.",
							Default:     false,
						},
					},
				},
				Description: "IP addresses or ranges",
			},
			"fqdns": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fqdn": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Full Qualified Domain Name for IP List. Allowed formats: hostname, IP, or URI",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Desciption of FQDN",
						},
					},
				},
				Description: "Collection of Fully Qualified Domain Names",
			},
			"external_data_set": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "The data source from which a resource originates",
			},
			"external_data_reference": {
				Type:             schema.TypeString,
				Optional:         true,
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringIsNotEmpty),
				Description:      "A unique identifier within the external data source",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IPList was first created",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IPList was last updated",
			},
			"deleted_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp when this IPList was last deleted",
			},
			"created_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who originally created this IPList",
			},
			"updated_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last updated this IPList",
			},
			"deleted_by": {
				Type:     schema.TypeMap,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Description: "User who last deleted this IPList",
			},
		},
	}
}

func resourceIllumioIPListCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	orgID := pConfig.OrgID

	ipList := &models.IPList{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
	}

	ipList.IPRanges = expandIllumioIPListIPRanges(d.Get("ip_ranges").(*schema.Set).List())
	ipList.FQDNs = expandIllumioIPListFQDNs(d.Get("fqdns").(*schema.Set).List())

	_, data, err := illumioClient.Create(fmt.Sprintf("/orgs/%d/sec_policy/draft/ip_lists", orgID), ipList)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "ip_lists", data.S("href").Data().(string))
	d.SetId(data.S("href").Data().(string))
	return resourceIllumioIPListRead(ctx, d, m)
}

func resourceIllumioIPListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(data.S("href").Data().(string))
	for _, key := range []string{
		"href",
		"name",
		"description",
		"ip_ranges",
		"fqdns",
		"external_data_set",
		"external_data_reference",
		"created_at",
		"updated_at",
		"created_by",
		"updated_by",
		"deleted_by",
		"deleted_at",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}
	return diagnostics
}

func resourceIllumioIPListUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	ipList := &models.IPList{
		Name:                  d.Get("name").(string),
		Description:           d.Get("description").(string),
		ExternalDataSet:       d.Get("external_data_set").(string),
		ExternalDataReference: d.Get("external_data_reference").(string),
		IPRanges:              expandIllumioIPListIPRanges(d.Get("ip_ranges").(*schema.Set).List()),
		FQDNs:                 expandIllumioIPListFQDNs(d.Get("fqdns").(*schema.Set).List()),
	}

	_, err := illumioClient.Update(d.Id(), ipList)
	if err != nil {
		return diag.FromErr(err)
	}
	pConfig.StoreHref(pConfig.OrgID, "ip_lists", d.Id())

	return resourceIllumioIPListRead(ctx, d, m)
}

func resourceIllumioIPListDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diagnostics diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	href := d.Id()
	_, err := illumioClient.Delete(href)
	if err != nil {
		return diag.FromErr(err)
	}

	pConfig.StoreHref(pConfig.OrgID, "ip_lists", href)

	d.SetId("")
	return diagnostics
}

func expandIllumioIPListIPRanges(arr []interface{}) []models.IPRange {
	var ipranges []models.IPRange
	for _, elem := range arr {
		ipranges = append(ipranges, models.IPRange{
			Description: elem.(map[string]interface{})["description"].(string),
			FromIP:      elem.(map[string]interface{})["from_ip"].(string),
			ToIP:        elem.(map[string]interface{})["to_ip"].(string),
			Exclusion:   elem.(map[string]interface{})["exclusion"].(bool),
		})
	}
	return ipranges
}

func expandIllumioIPListFQDNs(arr []interface{}) []models.FQDN {
	var fqdns []models.FQDN
	for _, elem := range arr {
		fqdns = append(fqdns, models.FQDN{
			FQDN:        elem.(map[string]interface{})["fqdn"].(string),
			Description: elem.(map[string]interface{})["description"].(string),
		})
	}
	return fqdns
}