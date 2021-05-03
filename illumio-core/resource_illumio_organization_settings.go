package illumiocore

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/illumio/terraform-provider-illumio-core/models"
)

func resourceIllumioOrganizationSettings() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIllumioOrganizationSettingsCreate,
		ReadContext:   resourceIllumioOrganizationSettingsRead,
		UpdateContext: resourceIllumioOrganizationSettingsUpdate,
		DeleteContext: resourceIllumioOrganizationSettingsDelete,

		SchemaVersion: version,
		Description:   "Manages Illumio Organization Settings",

		Schema: map[string]*schema.Schema{
			"audit_event_retention_seconds": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The time in seconds an audit event is stored in the database. The value should be in between 86400 and 17280000",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.IntBetween(86400, 17280000),
				),
			},
			"audit_event_min_severity": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Minimum severity level of audit event messages. Allowed values : \"error\", \"warning\", and \"informational\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"error", "warning", "informational"}, false),
				),
			},
			"format": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The log format (JSON, CEF, LEEF), which applies to all remote syslog destinations. Allowed values : \"JSON\", \"CEF\", and \"LEEF\" ",
				ValidateDiagFunc: validation.ToDiagFunc(
					validation.StringInSlice([]string{"JSON", "CEF", "LEEF"}, false),
				),
			},
		},

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceIllumioOrganizationSettingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Detail:   "Cannot use Create Operation on Organization Settings Resource. Only Read and Update is allowed.",
		Summary:  "Please use terrform import...",
	})

	return diags
}

func resourceIllumioOrganizationSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient
	href := d.Id()

	_, data, err := illumioClient.Get(href, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(href)

	for _, key := range []string{
		"audit_event_retention_seconds",
		"audit_event_min_severity",
		"format",
	} {
		if data.Exists(key) {
			d.Set(key, data.S(key).Data())
		} else {
			d.Set(key, nil)
		}
	}

	return diags
}

func resourceIllumioOrganizationSettingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	pConfig, _ := m.(Config)
	illumioClient := pConfig.IllumioClient

	OrganizationSettings := &models.OrganizationSettings{}

	OrganizationSettings.AuditEventRetentionSeconds = d.Get("audit_event_retention_seconds").(int)

	OrganizationSettings.AuditEventMinSeverity = d.Get("audit_event_min_severity").(string)

	OrganizationSettings.Format = d.Get("format").(string)

	_, err := illumioClient.Update(d.Id(), OrganizationSettings)
	if err != nil {
		return diag.FromErr(err)
	}

	return resourceIllumioOrganizationSettingsRead(ctx, d, m)
}

func resourceIllumioOrganizationSettingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	var diags diag.Diagnostics

	diags = append(diags, diag.Diagnostic{
		Severity: diag.Warning,
		Detail:   "Cannot use Delete Operation on Organization Settings Resource. Only Read and Update is allowed.",
		Summary:  "Setting the ID of Organization Settings to null. Ignoring the Deletion...",
	})

	return diags
}