// Code generated by user config generator. DO NOT EDIT.

package datadog

import (
	"context"

	setvalidator "github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	attr "github.com/hashicorp/terraform-plugin-framework/attr"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	resource "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	booldefault "github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"

	schemautil "github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

// NewResourceSchema returns resource schema
func NewResourceSchema() resource.SetNestedBlock {
	return resource.SetNestedBlock{
		Description: "Datadog user configurable settings",
		NestedObject: resource.NestedBlockObject{
			Attributes: map[string]resource.Attribute{
				"datadog_dbm_enabled": resource.BoolAttribute{
					Computed:    true,
					Description: "Enable Datadog Database Monitoring.",
					Optional:    true,
				},
				"exclude_consumer_groups": resource.SetAttribute{
					Computed:    true,
					Description: "List of custom metrics.",
					ElementType: types.StringType,
					Optional:    true,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"exclude_topics": resource.SetAttribute{
					Computed:    true,
					Description: "List of topics to exclude.",
					ElementType: types.StringType,
					Optional:    true,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"include_consumer_groups": resource.SetAttribute{
					Computed:    true,
					Description: "List of custom metrics.",
					ElementType: types.StringType,
					Optional:    true,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"include_topics": resource.SetAttribute{
					Computed:    true,
					Description: "List of topics to include.",
					ElementType: types.StringType,
					Optional:    true,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"kafka_custom_metrics": resource.SetAttribute{
					Computed:    true,
					Description: "List of custom metrics.",
					ElementType: types.StringType,
					Optional:    true,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"max_jmx_metrics": resource.Int64Attribute{
					Computed:    true,
					Description: "Maximum number of JMX metrics to send.",
					Optional:    true,
				},
			},
			Blocks: map[string]resource.Block{
				"datadog_tags": resource.SetNestedBlock{
					Description: "Custom tags provided by user",
					NestedObject: resource.NestedBlockObject{Attributes: map[string]resource.Attribute{
						"comment": resource.StringAttribute{
							Computed:    true,
							Description: "Optional tag explanation.",
							Optional:    true,
						},
						"tag": resource.StringAttribute{
							Description: "Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix 'aiven-' are reserved for Aiven.",
							Required:    true,
						},
					}},
					Validators: []validator.Set{setvalidator.SizeAtMost(32)},
				},
				"opensearch": resource.SetNestedBlock{
					Description: "Datadog Opensearch Options",
					NestedObject: resource.NestedBlockObject{Attributes: map[string]resource.Attribute{
						"index_stats_enabled": resource.BoolAttribute{
							Computed:    true,
							Description: "Enable Datadog Opensearch Index Monitoring.",
							Optional:    true,
						},
						"pending_task_stats_enabled": resource.BoolAttribute{
							Computed:    true,
							Description: "Enable Datadog Opensearch Pending Task Monitoring.",
							Optional:    true,
						},
						"pshard_stats_enabled": resource.BoolAttribute{
							Computed:    true,
							Description: "Enable Datadog Opensearch Primary Shard Monitoring.",
							Optional:    true,
						},
					}},
				},
				"redis": resource.SetNestedBlock{
					Description: "Datadog Redis Options",
					NestedObject: resource.NestedBlockObject{Attributes: map[string]resource.Attribute{"command_stats_enabled": resource.BoolAttribute{
						Computed:    true,
						Default:     booldefault.StaticBool(false),
						Description: "Enable command_stats option in the agent's configuration. The default value is `false`.",
						Optional:    true,
					}}},
				},
			},
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(1)},
	}
}

// NewDataSourceSchema returns datasource schema
func NewDataSourceSchema() datasource.SetNestedBlock {
	return datasource.SetNestedBlock{
		Description: "Datadog user configurable settings",
		NestedObject: datasource.NestedBlockObject{
			Attributes: map[string]datasource.Attribute{
				"datadog_dbm_enabled": datasource.BoolAttribute{
					Computed:    true,
					Description: "Enable Datadog Database Monitoring.",
				},
				"exclude_consumer_groups": datasource.SetAttribute{
					Computed:    true,
					Description: "List of custom metrics.",
					ElementType: types.StringType,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"exclude_topics": datasource.SetAttribute{
					Computed:    true,
					Description: "List of topics to exclude.",
					ElementType: types.StringType,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"include_consumer_groups": datasource.SetAttribute{
					Computed:    true,
					Description: "List of custom metrics.",
					ElementType: types.StringType,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"include_topics": datasource.SetAttribute{
					Computed:    true,
					Description: "List of topics to include.",
					ElementType: types.StringType,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"kafka_custom_metrics": datasource.SetAttribute{
					Computed:    true,
					Description: "List of custom metrics.",
					ElementType: types.StringType,
					Validators:  []validator.Set{setvalidator.SizeAtMost(1024)},
				},
				"max_jmx_metrics": datasource.Int64Attribute{
					Computed:    true,
					Description: "Maximum number of JMX metrics to send.",
				},
			},
			Blocks: map[string]datasource.Block{
				"datadog_tags": datasource.SetNestedBlock{
					Description: "Custom tags provided by user",
					NestedObject: datasource.NestedBlockObject{Attributes: map[string]datasource.Attribute{
						"comment": datasource.StringAttribute{
							Computed:    true,
							Description: "Optional tag explanation.",
						},
						"tag": datasource.StringAttribute{
							Computed:    true,
							Description: "Tag format and usage are described here: https://docs.datadoghq.com/getting_started/tagging. Tags with prefix 'aiven-' are reserved for Aiven.",
						},
					}},
					Validators: []validator.Set{setvalidator.SizeAtMost(32)},
				},
				"opensearch": datasource.SetNestedBlock{
					Description: "Datadog Opensearch Options",
					NestedObject: datasource.NestedBlockObject{Attributes: map[string]datasource.Attribute{
						"index_stats_enabled": datasource.BoolAttribute{
							Computed:    true,
							Description: "Enable Datadog Opensearch Index Monitoring.",
						},
						"pending_task_stats_enabled": datasource.BoolAttribute{
							Computed:    true,
							Description: "Enable Datadog Opensearch Pending Task Monitoring.",
						},
						"pshard_stats_enabled": datasource.BoolAttribute{
							Computed:    true,
							Description: "Enable Datadog Opensearch Primary Shard Monitoring.",
						},
					}},
				},
				"redis": datasource.SetNestedBlock{
					Description: "Datadog Redis Options",
					NestedObject: datasource.NestedBlockObject{Attributes: map[string]datasource.Attribute{"command_stats_enabled": datasource.BoolAttribute{
						Computed:    true,
						Description: "Enable command_stats option in the agent's configuration. The default value is `false`.",
					}}},
				},
			},
		},
		Validators: []validator.Set{setvalidator.SizeAtMost(1)},
	}
}

// tfoUserConfig Datadog user configurable settings
type tfoUserConfig struct {
	DatadogDbmEnabled     types.Bool  `tfsdk:"datadog_dbm_enabled"`
	DatadogTags           types.Set   `tfsdk:"datadog_tags"`
	ExcludeConsumerGroups types.Set   `tfsdk:"exclude_consumer_groups"`
	ExcludeTopics         types.Set   `tfsdk:"exclude_topics"`
	IncludeConsumerGroups types.Set   `tfsdk:"include_consumer_groups"`
	IncludeTopics         types.Set   `tfsdk:"include_topics"`
	KafkaCustomMetrics    types.Set   `tfsdk:"kafka_custom_metrics"`
	MaxJmxMetrics         types.Int64 `tfsdk:"max_jmx_metrics"`
	Opensearch            types.Set   `tfsdk:"opensearch"`
	Redis                 types.Set   `tfsdk:"redis"`
}

// dtoUserConfig request/response object
type dtoUserConfig struct {
	DatadogDbmEnabled     *bool             `groups:"create,update" json:"datadog_dbm_enabled,omitempty"`
	DatadogTags           []*dtoDatadogTags `groups:"create,update" json:"datadog_tags,omitempty"`
	ExcludeConsumerGroups []string          `groups:"create,update" json:"exclude_consumer_groups,omitempty"`
	ExcludeTopics         []string          `groups:"create,update" json:"exclude_topics,omitempty"`
	IncludeConsumerGroups []string          `groups:"create,update" json:"include_consumer_groups,omitempty"`
	IncludeTopics         []string          `groups:"create,update" json:"include_topics,omitempty"`
	KafkaCustomMetrics    []string          `groups:"create,update" json:"kafka_custom_metrics,omitempty"`
	MaxJmxMetrics         *int64            `groups:"create,update" json:"max_jmx_metrics,omitempty"`
	Opensearch            *dtoOpensearch    `groups:"create,update" json:"opensearch,omitempty"`
	Redis                 *dtoRedis         `groups:"create,update" json:"redis,omitempty"`
}

// expandUserConfig expands tf object into dto object
func expandUserConfig(ctx context.Context, diags diag.Diagnostics, o *tfoUserConfig) *dtoUserConfig {
	datadogTagsVar := schemautil.ExpandSetNested(ctx, diags, expandDatadogTags, o.DatadogTags)
	if diags.HasError() {
		return nil
	}
	excludeConsumerGroupsVar := schemautil.ExpandSet[string](ctx, diags, o.ExcludeConsumerGroups)
	if diags.HasError() {
		return nil
	}
	excludeTopicsVar := schemautil.ExpandSet[string](ctx, diags, o.ExcludeTopics)
	if diags.HasError() {
		return nil
	}
	includeConsumerGroupsVar := schemautil.ExpandSet[string](ctx, diags, o.IncludeConsumerGroups)
	if diags.HasError() {
		return nil
	}
	includeTopicsVar := schemautil.ExpandSet[string](ctx, diags, o.IncludeTopics)
	if diags.HasError() {
		return nil
	}
	kafkaCustomMetricsVar := schemautil.ExpandSet[string](ctx, diags, o.KafkaCustomMetrics)
	if diags.HasError() {
		return nil
	}
	opensearchVar := schemautil.ExpandSetBlockNested(ctx, diags, expandOpensearch, o.Opensearch)
	if diags.HasError() {
		return nil
	}
	redisVar := schemautil.ExpandSetBlockNested(ctx, diags, expandRedis, o.Redis)
	if diags.HasError() {
		return nil
	}
	return &dtoUserConfig{
		DatadogDbmEnabled:     schemautil.ValueBoolPointer(o.DatadogDbmEnabled),
		DatadogTags:           datadogTagsVar,
		ExcludeConsumerGroups: excludeConsumerGroupsVar,
		ExcludeTopics:         excludeTopicsVar,
		IncludeConsumerGroups: includeConsumerGroupsVar,
		IncludeTopics:         includeTopicsVar,
		KafkaCustomMetrics:    kafkaCustomMetricsVar,
		MaxJmxMetrics:         schemautil.ValueInt64Pointer(o.MaxJmxMetrics),
		Opensearch:            opensearchVar,
		Redis:                 redisVar,
	}
}

// flattenUserConfig flattens dto object into tf object
func flattenUserConfig(ctx context.Context, diags diag.Diagnostics, o *dtoUserConfig) *tfoUserConfig {
	datadogTagsVar := schemautil.FlattenSetNested(ctx, diags, flattenDatadogTags, o.DatadogTags, datadogTagsAttrs)
	if diags.HasError() {
		return nil
	}
	excludeConsumerGroupsVar, d := types.SetValueFrom(ctx, types.StringType, o.ExcludeConsumerGroups)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	excludeTopicsVar, d := types.SetValueFrom(ctx, types.StringType, o.ExcludeTopics)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	includeConsumerGroupsVar, d := types.SetValueFrom(ctx, types.StringType, o.IncludeConsumerGroups)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	includeTopicsVar, d := types.SetValueFrom(ctx, types.StringType, o.IncludeTopics)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	kafkaCustomMetricsVar, d := types.SetValueFrom(ctx, types.StringType, o.KafkaCustomMetrics)
	diags.Append(d...)
	if diags.HasError() {
		return nil
	}
	opensearchVar := schemautil.FlattenSetBlockNested(ctx, diags, flattenOpensearch, o.Opensearch, opensearchAttrs)
	if diags.HasError() {
		return nil
	}
	redisVar := schemautil.FlattenSetBlockNested(ctx, diags, flattenRedis, o.Redis, redisAttrs)
	if diags.HasError() {
		return nil
	}
	return &tfoUserConfig{
		DatadogDbmEnabled:     types.BoolPointerValue(o.DatadogDbmEnabled),
		DatadogTags:           datadogTagsVar,
		ExcludeConsumerGroups: excludeConsumerGroupsVar,
		ExcludeTopics:         excludeTopicsVar,
		IncludeConsumerGroups: includeConsumerGroupsVar,
		IncludeTopics:         includeTopicsVar,
		KafkaCustomMetrics:    kafkaCustomMetricsVar,
		MaxJmxMetrics:         types.Int64PointerValue(o.MaxJmxMetrics),
		Opensearch:            opensearchVar,
		Redis:                 redisVar,
	}
}

var userConfigAttrs = map[string]attr.Type{
	"datadog_dbm_enabled":     types.BoolType,
	"datadog_tags":            types.SetType{ElemType: types.ObjectType{AttrTypes: datadogTagsAttrs}},
	"exclude_consumer_groups": types.SetType{ElemType: types.StringType},
	"exclude_topics":          types.SetType{ElemType: types.StringType},
	"include_consumer_groups": types.SetType{ElemType: types.StringType},
	"include_topics":          types.SetType{ElemType: types.StringType},
	"kafka_custom_metrics":    types.SetType{ElemType: types.StringType},
	"max_jmx_metrics":         types.Int64Type,
	"opensearch":              types.SetType{ElemType: types.ObjectType{AttrTypes: opensearchAttrs}},
	"redis":                   types.SetType{ElemType: types.ObjectType{AttrTypes: redisAttrs}},
}

// tfoDatadogTags Datadog tag defined by user
type tfoDatadogTags struct {
	Comment types.String `tfsdk:"comment"`
	Tag     types.String `tfsdk:"tag"`
}

// dtoDatadogTags request/response object
type dtoDatadogTags struct {
	Comment *string `groups:"create,update" json:"comment,omitempty"`
	Tag     string  `groups:"create,update" json:"tag"`
}

// expandDatadogTags expands tf object into dto object
func expandDatadogTags(ctx context.Context, diags diag.Diagnostics, o *tfoDatadogTags) *dtoDatadogTags {
	return &dtoDatadogTags{
		Comment: schemautil.ValueStringPointer(o.Comment),
		Tag:     o.Tag.ValueString(),
	}
}

// flattenDatadogTags flattens dto object into tf object
func flattenDatadogTags(ctx context.Context, diags diag.Diagnostics, o *dtoDatadogTags) *tfoDatadogTags {
	return &tfoDatadogTags{
		Comment: types.StringPointerValue(o.Comment),
		Tag:     types.StringValue(o.Tag),
	}
}

var datadogTagsAttrs = map[string]attr.Type{
	"comment": types.StringType,
	"tag":     types.StringType,
}

// tfoOpensearch Datadog Opensearch Options
type tfoOpensearch struct {
	IndexStatsEnabled       types.Bool `tfsdk:"index_stats_enabled"`
	PendingTaskStatsEnabled types.Bool `tfsdk:"pending_task_stats_enabled"`
	PshardStatsEnabled      types.Bool `tfsdk:"pshard_stats_enabled"`
}

// dtoOpensearch request/response object
type dtoOpensearch struct {
	IndexStatsEnabled       *bool `groups:"create,update" json:"index_stats_enabled,omitempty"`
	PendingTaskStatsEnabled *bool `groups:"create,update" json:"pending_task_stats_enabled,omitempty"`
	PshardStatsEnabled      *bool `groups:"create,update" json:"pshard_stats_enabled,omitempty"`
}

// expandOpensearch expands tf object into dto object
func expandOpensearch(ctx context.Context, diags diag.Diagnostics, o *tfoOpensearch) *dtoOpensearch {
	return &dtoOpensearch{
		IndexStatsEnabled:       schemautil.ValueBoolPointer(o.IndexStatsEnabled),
		PendingTaskStatsEnabled: schemautil.ValueBoolPointer(o.PendingTaskStatsEnabled),
		PshardStatsEnabled:      schemautil.ValueBoolPointer(o.PshardStatsEnabled),
	}
}

// flattenOpensearch flattens dto object into tf object
func flattenOpensearch(ctx context.Context, diags diag.Diagnostics, o *dtoOpensearch) *tfoOpensearch {
	return &tfoOpensearch{
		IndexStatsEnabled:       types.BoolPointerValue(o.IndexStatsEnabled),
		PendingTaskStatsEnabled: types.BoolPointerValue(o.PendingTaskStatsEnabled),
		PshardStatsEnabled:      types.BoolPointerValue(o.PshardStatsEnabled),
	}
}

var opensearchAttrs = map[string]attr.Type{
	"index_stats_enabled":        types.BoolType,
	"pending_task_stats_enabled": types.BoolType,
	"pshard_stats_enabled":       types.BoolType,
}

// tfoRedis Datadog Redis Options
type tfoRedis struct {
	CommandStatsEnabled types.Bool `tfsdk:"command_stats_enabled"`
}

// dtoRedis request/response object
type dtoRedis struct {
	CommandStatsEnabled *bool `groups:"create,update" json:"command_stats_enabled,omitempty"`
}

// expandRedis expands tf object into dto object
func expandRedis(ctx context.Context, diags diag.Diagnostics, o *tfoRedis) *dtoRedis {
	return &dtoRedis{CommandStatsEnabled: schemautil.ValueBoolPointer(o.CommandStatsEnabled)}
}

// flattenRedis flattens dto object into tf object
func flattenRedis(ctx context.Context, diags diag.Diagnostics, o *dtoRedis) *tfoRedis {
	return &tfoRedis{CommandStatsEnabled: types.BoolPointerValue(o.CommandStatsEnabled)}
}

var redisAttrs = map[string]attr.Type{"command_stats_enabled": types.BoolType}

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
