// Code generated by user config generator. DO NOT EDIT.

package m3aggregator

import (
	"context"
	"encoding/json"

	setvalidator "github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	attr "github.com/hashicorp/terraform-plugin-framework/attr"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	resource "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"

	schemautil "github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

// NewResourceSchema returns resource schema
func NewResourceSchema() resource.SetNestedBlock {
	return resource.SetNestedBlock{
		Description: "M3aggregator user configurable settings",
		NestedObject: resource.NestedBlockObject{
			Attributes: map[string]resource.Attribute{
				"custom_domain": resource.StringAttribute{
					Computed:    true,
					Description: "Serve the web frontend using a custom CNAME pointing to the Aiven DNS name.",
					Optional:    true,
				},
				"m3_version": resource.StringAttribute{
					Computed:    true,
					Description: "M3 major version (deprecated, use m3aggregator_version).",
					Optional:    true,
				},
				"m3aggregator_version": resource.StringAttribute{
					Computed:    true,
					Description: "M3 major version (the minimum compatible version).",
					Optional:    true,
				},
				"static_ips": resource.BoolAttribute{
					Computed:    true,
					Description: "Use static public IP addresses.",
					Optional:    true,
				},
			},
			Blocks: map[string]resource.Block{"ip_filter": resource.SetNestedBlock{
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'",
				NestedObject: resource.NestedBlockObject{Attributes: map[string]resource.Attribute{
					"description": resource.StringAttribute{
						Computed:    true,
						Description: "Description for IP filter list entry.",
						Optional:    true,
					},
					"network": resource.StringAttribute{
						Description: "CIDR address block.",
						Required:    true,
					},
				}},
				Validators: []validator.Set{setvalidator.SizeAtMost(1024)},
			}},
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(1)},
	}
}

// NewDataSourceSchema returns datasource schema
func NewDataSourceSchema() datasource.SetNestedBlock {
	return datasource.SetNestedBlock{
		Description: "M3aggregator user configurable settings",
		NestedObject: datasource.NestedBlockObject{
			Attributes: map[string]datasource.Attribute{
				"custom_domain": datasource.StringAttribute{
					Computed:    true,
					Description: "Serve the web frontend using a custom CNAME pointing to the Aiven DNS name.",
				},
				"m3_version": datasource.StringAttribute{
					Computed:    true,
					Description: "M3 major version (deprecated, use m3aggregator_version).",
				},
				"m3aggregator_version": datasource.StringAttribute{
					Computed:    true,
					Description: "M3 major version (the minimum compatible version).",
				},
				"static_ips": datasource.BoolAttribute{
					Computed:    true,
					Description: "Use static public IP addresses.",
				},
			},
			Blocks: map[string]datasource.Block{"ip_filter": datasource.SetNestedBlock{
				Description: "Allow incoming connections from CIDR address block, e.g. '10.20.0.0/16'",
				NestedObject: datasource.NestedBlockObject{Attributes: map[string]datasource.Attribute{
					"description": datasource.StringAttribute{
						Computed:    true,
						Description: "Description for IP filter list entry.",
					},
					"network": datasource.StringAttribute{
						Computed:    true,
						Description: "CIDR address block.",
					},
				}},
				Validators: []validator.Set{setvalidator.SizeAtMost(1024)},
			}},
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(1)},
	}
}

// tfoUserConfig M3aggregator user configurable settings
type tfoUserConfig struct {
	CustomDomain        types.String `tfsdk:"custom_domain"`
	IpFilter            types.Set    `tfsdk:"ip_filter"`
	M3Version           types.String `tfsdk:"m3_version"`
	M3aggregatorVersion types.String `tfsdk:"m3aggregator_version"`
	StaticIps           types.Bool   `tfsdk:"static_ips"`
}

// dtoUserConfig request/response object
type dtoUserConfig struct {
	CustomDomain        *string        `groups:"create,update" json:"custom_domain,omitempty"`
	IpFilter            []*dtoIpFilter `groups:"create,update" json:"ip_filter,omitempty"`
	M3Version           *string        `groups:"create,update" json:"m3_version,omitempty"`
	M3aggregatorVersion *string        `groups:"create,update" json:"m3aggregator_version,omitempty"`
	StaticIps           *bool          `groups:"create,update" json:"static_ips,omitempty"`
}

// expandUserConfig expands tf object into dto object
func expandUserConfig(ctx context.Context, diags diag.Diagnostics, o *tfoUserConfig) *dtoUserConfig {
	ipFilterVar := schemautil.ExpandSetNested(ctx, diags, expandIpFilter, o.IpFilter)
	if diags.HasError() {
		return nil
	}
	return &dtoUserConfig{
		CustomDomain:        schemautil.ValueStringPointer(o.CustomDomain),
		IpFilter:            ipFilterVar,
		M3Version:           schemautil.ValueStringPointer(o.M3Version),
		M3aggregatorVersion: schemautil.ValueStringPointer(o.M3aggregatorVersion),
		StaticIps:           schemautil.ValueBoolPointer(o.StaticIps),
	}
}

// flattenUserConfig flattens dto object into tf object
func flattenUserConfig(ctx context.Context, diags diag.Diagnostics, o *dtoUserConfig) *tfoUserConfig {
	ipFilterVar := schemautil.FlattenSetNested(ctx, diags, flattenIpFilter, o.IpFilter, ipFilterAttrs)
	if diags.HasError() {
		return nil
	}
	return &tfoUserConfig{
		CustomDomain:        types.StringPointerValue(o.CustomDomain),
		IpFilter:            ipFilterVar,
		M3Version:           types.StringPointerValue(o.M3Version),
		M3aggregatorVersion: types.StringPointerValue(o.M3aggregatorVersion),
		StaticIps:           types.BoolPointerValue(o.StaticIps),
	}
}

var userConfigAttrs = map[string]attr.Type{
	"custom_domain":        types.StringType,
	"ip_filter":            types.SetType{ElemType: types.ObjectType{AttrTypes: ipFilterAttrs}},
	"m3_version":           types.StringType,
	"m3aggregator_version": types.StringType,
	"static_ips":           types.BoolType,
}

// tfoIpFilter CIDR address block, either as a string, or in a dict with an optional description field
type tfoIpFilter struct {
	Description types.String `tfsdk:"description"`
	Network     types.String `tfsdk:"network"`
}

// dtoIpFilter request/response object
type dtoIpFilter struct {
	Description *string `groups:"create,update" json:"description,omitempty"`
	Network     string  `groups:"create,update" json:"network"`
}

func (d *dtoIpFilter) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)
	if err == nil {
		d.Network = s
		return nil
	}

	type obj dtoIpFilter
	o := &struct {
		Description *string `groups:"create,update" json:"description,omitempty"`
		Network     string  `groups:"create,update" json:"network"`
	}{}
	err = json.Unmarshal(data, o)
	if err != nil {
		return err
	}

	d.Description = o.Description
	d.Network = o.Network
	return nil
}

// expandIpFilter expands tf object into dto object
func expandIpFilter(ctx context.Context, diags diag.Diagnostics, o *tfoIpFilter) *dtoIpFilter {
	return &dtoIpFilter{
		Description: schemautil.ValueStringPointer(o.Description),
		Network:     o.Network.ValueString(),
	}
}

// flattenIpFilter flattens dto object into tf object
func flattenIpFilter(ctx context.Context, diags diag.Diagnostics, o *dtoIpFilter) *tfoIpFilter {
	return &tfoIpFilter{
		Description: types.StringPointerValue(o.Description),
		Network:     types.StringValue(o.Network),
	}
}

var ipFilterAttrs = map[string]attr.Type{
	"description": types.StringType,
	"network":     types.StringType,
}

// Expand public function that converts tf object into dto
func Expand(ctx context.Context, diags diag.Diagnostics, set types.Set) *dtoUserConfig {
	return schemautil.ExpandSetBlockNested[tfoUserConfig, dtoUserConfig](ctx, diags, expandUserConfig, set)
}

// Flatten public function that converts dto into tf object
func Flatten(ctx context.Context, diags diag.Diagnostics, m map[string]any) types.Set {
	o := new(dtoUserConfig)
	err := schemautil.MapToDTO(m, o)
	if err != nil {
		diags.AddError("Failed to marshal map user config to dto", err.Error())
		return types.SetNull(types.ObjectType{AttrTypes: userConfigAttrs})
	}
	return schemautil.FlattenSetBlockNested[dtoUserConfig, tfoUserConfig](ctx, diags, flattenUserConfig, o, userConfigAttrs)
}
