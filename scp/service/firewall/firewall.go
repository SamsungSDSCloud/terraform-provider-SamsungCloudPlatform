package firewall

import (
	"context"
	"github.com/ScpDevTerra/trf-provider/scp/client"
	"github.com/ScpDevTerra/trf-provider/scp/common"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func DatasourceFirewall() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceFirewallRead,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"filter": common.DatasourceFilter(),
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Firewall id",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of firewall",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Firewall status",
			},
			"vpc_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "VPC id",
			},
			"target_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: "Target firewall resource id",
			},
			"target_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target firewall resource type",
			},
		},
		Description: "Provides firewall details",
	}
}

func datasourceFirewallRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	vpcId := rd.Get("vpc_id").(string)
	targetId := rd.Get("target_id").(string)

	firewalls, _, err := inst.Client.Firewall.GetFirewallList(ctx, vpcId, targetId, "")
	if err != nil {
		return diag.FromErr(err)
	}

	setFirewalls, _ := convertFirewallListToHclSet(firewalls.Contents)

	if f, ok := rd.GetOk("filter"); ok {
		setFirewalls = common.ApplyFilter(DatasourceFirewall().Schema, f.(*schema.Set), setFirewalls)
	}

	if len(setFirewalls) == 0 {
		return diag.Errorf("no matching firewall found")
	}

	for k, v := range setFirewalls[0] {
		if k == "id" {
			rd.SetId(v.(string))
			continue
		}
		rd.Set(k, v)
	}

	return nil
}

func ResourceFirewall() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFirewallCreate,
		ReadContext:   resourceFirewallRead,
		UpdateContext: resourceFirewallUpdate,
		DeleteContext: resourceFirewallDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"vpc_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "VPC id",
			},
			"target_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Target firewall resource id",
			},
			"target_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Target firewall resource type",
			},
			"enabled": {
				Type:        schema.TypeBool,
				Required:    true,
				Description: "Firewall enabled status",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Name of firewall",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Firewall status",
			},
		},
		Description: "Provides a Firewall resource.",
	}
}

func resourceFirewallCreate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	vpcId := rd.Get("vpc_id").(string)
	targetId := rd.Get("target_id").(string)
	isEnabled := rd.Get("enabled").(bool)

	firewalls, _, err := inst.Client.Firewall.GetFirewallList(ctx, vpcId, targetId, "")
	if err != nil {
		return diag.FromErr(err)
	}
	if len(firewalls.Contents) == 0 {
		return diag.Errorf("firewall not found")
	}

	firewallInfo, _, err := inst.Client.Firewall.GetFirewall(ctx, firewalls.Contents[0].FirewallId)
	if err != nil {
		return diag.FromErr(err)
	}
	if firewallInfo.IsEnabled != isEnabled {
		inst.Client.Firewall.UpdateFirewallEnabled(ctx, firewalls.Contents[0].FirewallId, isEnabled)
		targetState := common.ActiveState
		if !isEnabled {
			targetState = common.InActiveState
		}
		err = WaitForFirewallStatus(ctx, inst.Client, firewalls.Contents[0].FirewallId, FirewallPendingStates(), []string{targetState}, true)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	rd.SetId(firewalls.Contents[0].FirewallId)

	return resourceFirewallRead(ctx, rd, meta)
}

func resourceFirewallRead(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	firewallInfo, _, err := inst.Client.Firewall.GetFirewall(ctx, rd.Id())
	if err != nil {
		rd.SetId("")
		return diag.FromErr(err)
	}

	rd.Set("vpc_id", firewallInfo.VpcId)
	rd.Set("target_id", firewallInfo.ObjectId)
	rd.Set("target_type", firewallInfo.ObjectType)
	rd.Set("name", firewallInfo.FirewallName)
	rd.Set("state", firewallInfo.FirewallState)
	rd.Set("enabled", firewallInfo.IsEnabled)

	return nil
}

func resourceFirewallUpdate(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	inst := meta.(*client.Instance)

	if rd.HasChanges("enabled") {

		isEnabled := rd.Get("enabled").(bool)

		inst.Client.Firewall.UpdateFirewallEnabled(ctx, rd.Id(), isEnabled)
		targetState := common.ActiveState
		if !isEnabled {
			targetState = common.InActiveState
		}
		err := WaitForFirewallStatus(ctx, inst.Client, rd.Id(), FirewallPendingStates(), []string{targetState}, false)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceFirewallRead(ctx, rd, meta)
}

func resourceFirewallDelete(ctx context.Context, rd *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// Must be deleted by parent's resource
	return nil
}

func FirewallPendingStates() []string {
	return []string{common.DeployingState}
}

func WaitForFirewallStatus(ctx context.Context, scpClient *client.SCPClient, firewallId string, pendingStates []string, targetStates []string, errorOnNotFound bool) error {
	return client.WaitForStatus(ctx, scpClient, pendingStates, targetStates, func() (interface{}, string, error) {
		info, c, err := scpClient.Firewall.GetFirewall(ctx, firewallId)
		if err != nil {
			if c == 404 && !errorOnNotFound {
				return "", common.DeletedState, nil
			}
			if c == 403 && !errorOnNotFound {
				return "", common.DeletedState, nil
			}
			return nil, "", err
		}
		return info, info.FirewallState, nil
	})
}
