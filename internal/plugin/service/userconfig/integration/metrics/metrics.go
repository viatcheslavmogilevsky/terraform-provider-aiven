// Code generated by user config generator. DO NOT EDIT.

package metrics

import (
	"context"

	listvalidator "github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	attr "github.com/hashicorp/terraform-plugin-framework/attr"
	datasource "github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	diag "github.com/hashicorp/terraform-plugin-framework/diag"
	resource "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	validator "github.com/hashicorp/terraform-plugin-framework/schema/validator"
	types "github.com/hashicorp/terraform-plugin-framework/types"

	schemautil "github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

// NewResourceSchema returns resource schema
func NewResourceSchema() resource.ListNestedBlock {
	return resource.ListNestedBlock{
		Description: "Integration user config",
		NestedObject: resource.NestedBlockObject{
			Attributes: map[string]resource.Attribute{
				"database": resource.StringAttribute{
					Computed:    true,
					Description: "Name of the database where to store metric datapoints. Only affects PostgreSQL destinations. Defaults to 'metrics'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
					Optional:    true,
				},
				"retention_days": resource.Int64Attribute{
					Computed:    true,
					Description: "Number of days to keep old metrics. Only affects PostgreSQL destinations. Set to 0 for no automatic cleanup. Defaults to 30 days.",
					Optional:    true,
				},
				"ro_username": resource.StringAttribute{
					Computed:    true,
					Description: "Name of a user that can be used to read metrics. This will be used for Grafana integration (if enabled) to prevent Grafana users from making undesired changes. Only affects PostgreSQL destinations. Defaults to 'metrics_reader'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
					Optional:    true,
				},
				"username": resource.StringAttribute{
					Computed:    true,
					Description: "Name of the user used to write metrics. Only affects PostgreSQL destinations. Defaults to 'metrics_writer'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
					Optional:    true,
				},
			},
			Blocks: map[string]resource.Block{"source_mysql": resource.ListNestedBlock{
				Description: "Configuration options for metrics where source service is MySQL",
				NestedObject: resource.NestedBlockObject{Blocks: map[string]resource.Block{"telegraf": resource.ListNestedBlock{
					Description: "Configuration options for Telegraf MySQL input plugin",
					NestedObject: resource.NestedBlockObject{Attributes: map[string]resource.Attribute{
						"gather_event_waits": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.EVENT_WAITS.",
							Optional:    true,
						},
						"gather_file_events_stats": resource.BoolAttribute{
							Computed:    true,
							Description: "gather metrics from PERFORMANCE_SCHEMA.FILE_SUMMARY_BY_EVENT_NAME.",
							Optional:    true,
						},
						"gather_index_io_waits": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_INDEX_USAGE.",
							Optional:    true,
						},
						"gather_info_schema_auto_inc": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather auto_increment columns and max values from information schema.",
							Optional:    true,
						},
						"gather_innodb_metrics": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from INFORMATION_SCHEMA.INNODB_METRICS.",
							Optional:    true,
						},
						"gather_perf_events_statements": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.EVENTS_STATEMENTS_SUMMARY_BY_DIGEST.",
							Optional:    true,
						},
						"gather_process_list": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather thread state counts from INFORMATION_SCHEMA.PROCESSLIST.",
							Optional:    true,
						},
						"gather_slave_status": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from SHOW SLAVE STATUS command output.",
							Optional:    true,
						},
						"gather_table_io_waits": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_TABLE.",
							Optional:    true,
						},
						"gather_table_lock_waits": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_LOCK_WAITS.",
							Optional:    true,
						},
						"gather_table_schema": resource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from INFORMATION_SCHEMA.TABLES.",
							Optional:    true,
						},
						"perf_events_statements_digest_text_limit": resource.Int64Attribute{
							Computed:    true,
							Description: "Truncates digest text from perf_events_statements into this many characters.",
							Optional:    true,
						},
						"perf_events_statements_limit": resource.Int64Attribute{
							Computed:    true,
							Description: "Limits metrics from perf_events_statements.",
							Optional:    true,
						},
						"perf_events_statements_time_limit": resource.Int64Attribute{
							Computed:    true,
							Description: "Only include perf_events_statements whose last seen is less than this many seconds.",
							Optional:    true,
						},
					}},
				}}},
			}},
		},
		Validators: []validator.List{listvalidator.SizeAtMost(1)},
	}
}

// NewDataSourceSchema returns datasource schema
func NewDataSourceSchema() datasource.ListNestedBlock {
	return datasource.ListNestedBlock{
		Description: "Integration user config",
		NestedObject: datasource.NestedBlockObject{
			Attributes: map[string]datasource.Attribute{
				"database": datasource.StringAttribute{
					Computed:    true,
					Description: "Name of the database where to store metric datapoints. Only affects PostgreSQL destinations. Defaults to 'metrics'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
				},
				"retention_days": datasource.Int64Attribute{
					Computed:    true,
					Description: "Number of days to keep old metrics. Only affects PostgreSQL destinations. Set to 0 for no automatic cleanup. Defaults to 30 days.",
				},
				"ro_username": datasource.StringAttribute{
					Computed:    true,
					Description: "Name of a user that can be used to read metrics. This will be used for Grafana integration (if enabled) to prevent Grafana users from making undesired changes. Only affects PostgreSQL destinations. Defaults to 'metrics_reader'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
				},
				"username": datasource.StringAttribute{
					Computed:    true,
					Description: "Name of the user used to write metrics. Only affects PostgreSQL destinations. Defaults to 'metrics_writer'. Note that this must be the same for all metrics integrations that write data to the same PostgreSQL service.",
				},
			},
			Blocks: map[string]datasource.Block{"source_mysql": datasource.ListNestedBlock{
				Description: "Configuration options for metrics where source service is MySQL",
				NestedObject: datasource.NestedBlockObject{Blocks: map[string]datasource.Block{"telegraf": datasource.ListNestedBlock{
					Description: "Configuration options for Telegraf MySQL input plugin",
					NestedObject: datasource.NestedBlockObject{Attributes: map[string]datasource.Attribute{
						"gather_event_waits": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.EVENT_WAITS.",
						},
						"gather_file_events_stats": datasource.BoolAttribute{
							Computed:    true,
							Description: "gather metrics from PERFORMANCE_SCHEMA.FILE_SUMMARY_BY_EVENT_NAME.",
						},
						"gather_index_io_waits": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_INDEX_USAGE.",
						},
						"gather_info_schema_auto_inc": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather auto_increment columns and max values from information schema.",
						},
						"gather_innodb_metrics": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from INFORMATION_SCHEMA.INNODB_METRICS.",
						},
						"gather_perf_events_statements": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.EVENTS_STATEMENTS_SUMMARY_BY_DIGEST.",
						},
						"gather_process_list": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather thread state counts from INFORMATION_SCHEMA.PROCESSLIST.",
						},
						"gather_slave_status": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from SHOW SLAVE STATUS command output.",
						},
						"gather_table_io_waits": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_IO_WAITS_SUMMARY_BY_TABLE.",
						},
						"gather_table_lock_waits": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from PERFORMANCE_SCHEMA.TABLE_LOCK_WAITS.",
						},
						"gather_table_schema": datasource.BoolAttribute{
							Computed:    true,
							Description: "Gather metrics from INFORMATION_SCHEMA.TABLES.",
						},
						"perf_events_statements_digest_text_limit": datasource.Int64Attribute{
							Computed:    true,
							Description: "Truncates digest text from perf_events_statements into this many characters.",
						},
						"perf_events_statements_limit": datasource.Int64Attribute{
							Computed:    true,
							Description: "Limits metrics from perf_events_statements.",
						},
						"perf_events_statements_time_limit": datasource.Int64Attribute{
							Computed:    true,
							Description: "Only include perf_events_statements whose last seen is less than this many seconds.",
						},
					}},
				}}},
			}},
		},
		Validators: []validator.List{listvalidator.SizeAtMost(1)},
	}
}

// tfoUserConfig Integration user config
type tfoUserConfig struct {
	Database      types.String `tfsdk:"database"`
	RetentionDays types.Int64  `tfsdk:"retention_days"`
	RoUsername    types.String `tfsdk:"ro_username"`
	SourceMysql   types.List   `tfsdk:"source_mysql"`
	Username      types.String `tfsdk:"username"`
}

// dtoUserConfig request/response object
type dtoUserConfig struct {
	Database      *string         `groups:"create,update" json:"database,omitempty"`
	RetentionDays *int64          `groups:"create,update" json:"retention_days,omitempty"`
	RoUsername    *string         `groups:"create,update" json:"ro_username,omitempty"`
	SourceMysql   *dtoSourceMysql `groups:"create,update" json:"source_mysql,omitempty"`
	Username      *string         `groups:"create,update" json:"username,omitempty"`
}

// expandUserConfig expands tf object into dto object
func expandUserConfig(ctx context.Context, diags *diag.Diagnostics, o *tfoUserConfig) *dtoUserConfig {
	sourceMysqlVar := schemautil.ExpandListBlockNested[tfoSourceMysql, dtoSourceMysql](ctx, diags, expandSourceMysql, o.SourceMysql)
	if diags.HasError() {
		return nil
	}
	return &dtoUserConfig{
		Database:      schemautil.ValueStringPointer(o.Database),
		RetentionDays: schemautil.ValueInt64Pointer(o.RetentionDays),
		RoUsername:    schemautil.ValueStringPointer(o.RoUsername),
		SourceMysql:   sourceMysqlVar,
		Username:      schemautil.ValueStringPointer(o.Username),
	}
}

// flattenUserConfig flattens dto object into tf object
func flattenUserConfig(ctx context.Context, diags *diag.Diagnostics, o *dtoUserConfig) *tfoUserConfig {
	sourceMysqlVar := schemautil.FlattenListBlockNested[dtoSourceMysql, tfoSourceMysql](ctx, diags, flattenSourceMysql, sourceMysqlAttrs, o.SourceMysql)
	if diags.HasError() {
		return nil
	}
	return &tfoUserConfig{
		Database:      types.StringPointerValue(o.Database),
		RetentionDays: types.Int64PointerValue(o.RetentionDays),
		RoUsername:    types.StringPointerValue(o.RoUsername),
		SourceMysql:   sourceMysqlVar,
		Username:      types.StringPointerValue(o.Username),
	}
}

var userConfigAttrs = map[string]attr.Type{
	"database":       types.StringType,
	"retention_days": types.Int64Type,
	"ro_username":    types.StringType,
	"source_mysql":   types.ListType{ElemType: types.ObjectType{AttrTypes: sourceMysqlAttrs}},
	"username":       types.StringType,
}

// tfoSourceMysql Configuration options for metrics where source service is MySQL
type tfoSourceMysql struct {
	Telegraf types.List `tfsdk:"telegraf"`
}

// dtoSourceMysql request/response object
type dtoSourceMysql struct {
	Telegraf *dtoTelegraf `groups:"create,update" json:"telegraf,omitempty"`
}

// expandSourceMysql expands tf object into dto object
func expandSourceMysql(ctx context.Context, diags *diag.Diagnostics, o *tfoSourceMysql) *dtoSourceMysql {
	telegrafVar := schemautil.ExpandListBlockNested[tfoTelegraf, dtoTelegraf](ctx, diags, expandTelegraf, o.Telegraf)
	if diags.HasError() {
		return nil
	}
	return &dtoSourceMysql{Telegraf: telegrafVar}
}

// flattenSourceMysql flattens dto object into tf object
func flattenSourceMysql(ctx context.Context, diags *diag.Diagnostics, o *dtoSourceMysql) *tfoSourceMysql {
	telegrafVar := schemautil.FlattenListBlockNested[dtoTelegraf, tfoTelegraf](ctx, diags, flattenTelegraf, telegrafAttrs, o.Telegraf)
	if diags.HasError() {
		return nil
	}
	return &tfoSourceMysql{Telegraf: telegrafVar}
}

var sourceMysqlAttrs = map[string]attr.Type{"telegraf": types.ListType{ElemType: types.ObjectType{AttrTypes: telegrafAttrs}}}

// tfoTelegraf Configuration options for Telegraf MySQL input plugin
type tfoTelegraf struct {
	GatherEventWaits                    types.Bool  `tfsdk:"gather_event_waits"`
	GatherFileEventsStats               types.Bool  `tfsdk:"gather_file_events_stats"`
	GatherIndexIoWaits                  types.Bool  `tfsdk:"gather_index_io_waits"`
	GatherInfoSchemaAutoInc             types.Bool  `tfsdk:"gather_info_schema_auto_inc"`
	GatherInnodbMetrics                 types.Bool  `tfsdk:"gather_innodb_metrics"`
	GatherPerfEventsStatements          types.Bool  `tfsdk:"gather_perf_events_statements"`
	GatherProcessList                   types.Bool  `tfsdk:"gather_process_list"`
	GatherSlaveStatus                   types.Bool  `tfsdk:"gather_slave_status"`
	GatherTableIoWaits                  types.Bool  `tfsdk:"gather_table_io_waits"`
	GatherTableLockWaits                types.Bool  `tfsdk:"gather_table_lock_waits"`
	GatherTableSchema                   types.Bool  `tfsdk:"gather_table_schema"`
	PerfEventsStatementsDigestTextLimit types.Int64 `tfsdk:"perf_events_statements_digest_text_limit"`
	PerfEventsStatementsLimit           types.Int64 `tfsdk:"perf_events_statements_limit"`
	PerfEventsStatementsTimeLimit       types.Int64 `tfsdk:"perf_events_statements_time_limit"`
}

// dtoTelegraf request/response object
type dtoTelegraf struct {
	GatherEventWaits                    *bool  `groups:"create,update" json:"gather_event_waits,omitempty"`
	GatherFileEventsStats               *bool  `groups:"create,update" json:"gather_file_events_stats,omitempty"`
	GatherIndexIoWaits                  *bool  `groups:"create,update" json:"gather_index_io_waits,omitempty"`
	GatherInfoSchemaAutoInc             *bool  `groups:"create,update" json:"gather_info_schema_auto_inc,omitempty"`
	GatherInnodbMetrics                 *bool  `groups:"create,update" json:"gather_innodb_metrics,omitempty"`
	GatherPerfEventsStatements          *bool  `groups:"create,update" json:"gather_perf_events_statements,omitempty"`
	GatherProcessList                   *bool  `groups:"create,update" json:"gather_process_list,omitempty"`
	GatherSlaveStatus                   *bool  `groups:"create,update" json:"gather_slave_status,omitempty"`
	GatherTableIoWaits                  *bool  `groups:"create,update" json:"gather_table_io_waits,omitempty"`
	GatherTableLockWaits                *bool  `groups:"create,update" json:"gather_table_lock_waits,omitempty"`
	GatherTableSchema                   *bool  `groups:"create,update" json:"gather_table_schema,omitempty"`
	PerfEventsStatementsDigestTextLimit *int64 `groups:"create,update" json:"perf_events_statements_digest_text_limit,omitempty"`
	PerfEventsStatementsLimit           *int64 `groups:"create,update" json:"perf_events_statements_limit,omitempty"`
	PerfEventsStatementsTimeLimit       *int64 `groups:"create,update" json:"perf_events_statements_time_limit,omitempty"`
}

// expandTelegraf expands tf object into dto object
func expandTelegraf(ctx context.Context, diags *diag.Diagnostics, o *tfoTelegraf) *dtoTelegraf {
	return &dtoTelegraf{
		GatherEventWaits:                    schemautil.ValueBoolPointer(o.GatherEventWaits),
		GatherFileEventsStats:               schemautil.ValueBoolPointer(o.GatherFileEventsStats),
		GatherIndexIoWaits:                  schemautil.ValueBoolPointer(o.GatherIndexIoWaits),
		GatherInfoSchemaAutoInc:             schemautil.ValueBoolPointer(o.GatherInfoSchemaAutoInc),
		GatherInnodbMetrics:                 schemautil.ValueBoolPointer(o.GatherInnodbMetrics),
		GatherPerfEventsStatements:          schemautil.ValueBoolPointer(o.GatherPerfEventsStatements),
		GatherProcessList:                   schemautil.ValueBoolPointer(o.GatherProcessList),
		GatherSlaveStatus:                   schemautil.ValueBoolPointer(o.GatherSlaveStatus),
		GatherTableIoWaits:                  schemautil.ValueBoolPointer(o.GatherTableIoWaits),
		GatherTableLockWaits:                schemautil.ValueBoolPointer(o.GatherTableLockWaits),
		GatherTableSchema:                   schemautil.ValueBoolPointer(o.GatherTableSchema),
		PerfEventsStatementsDigestTextLimit: schemautil.ValueInt64Pointer(o.PerfEventsStatementsDigestTextLimit),
		PerfEventsStatementsLimit:           schemautil.ValueInt64Pointer(o.PerfEventsStatementsLimit),
		PerfEventsStatementsTimeLimit:       schemautil.ValueInt64Pointer(o.PerfEventsStatementsTimeLimit),
	}
}

// flattenTelegraf flattens dto object into tf object
func flattenTelegraf(ctx context.Context, diags *diag.Diagnostics, o *dtoTelegraf) *tfoTelegraf {
	return &tfoTelegraf{
		GatherEventWaits:                    types.BoolPointerValue(o.GatherEventWaits),
		GatherFileEventsStats:               types.BoolPointerValue(o.GatherFileEventsStats),
		GatherIndexIoWaits:                  types.BoolPointerValue(o.GatherIndexIoWaits),
		GatherInfoSchemaAutoInc:             types.BoolPointerValue(o.GatherInfoSchemaAutoInc),
		GatherInnodbMetrics:                 types.BoolPointerValue(o.GatherInnodbMetrics),
		GatherPerfEventsStatements:          types.BoolPointerValue(o.GatherPerfEventsStatements),
		GatherProcessList:                   types.BoolPointerValue(o.GatherProcessList),
		GatherSlaveStatus:                   types.BoolPointerValue(o.GatherSlaveStatus),
		GatherTableIoWaits:                  types.BoolPointerValue(o.GatherTableIoWaits),
		GatherTableLockWaits:                types.BoolPointerValue(o.GatherTableLockWaits),
		GatherTableSchema:                   types.BoolPointerValue(o.GatherTableSchema),
		PerfEventsStatementsDigestTextLimit: types.Int64PointerValue(o.PerfEventsStatementsDigestTextLimit),
		PerfEventsStatementsLimit:           types.Int64PointerValue(o.PerfEventsStatementsLimit),
		PerfEventsStatementsTimeLimit:       types.Int64PointerValue(o.PerfEventsStatementsTimeLimit),
	}
}

var telegrafAttrs = map[string]attr.Type{
	"gather_event_waits":                       types.BoolType,
	"gather_file_events_stats":                 types.BoolType,
	"gather_index_io_waits":                    types.BoolType,
	"gather_info_schema_auto_inc":              types.BoolType,
	"gather_innodb_metrics":                    types.BoolType,
	"gather_perf_events_statements":            types.BoolType,
	"gather_process_list":                      types.BoolType,
	"gather_slave_status":                      types.BoolType,
	"gather_table_io_waits":                    types.BoolType,
	"gather_table_lock_waits":                  types.BoolType,
	"gather_table_schema":                      types.BoolType,
	"perf_events_statements_digest_text_limit": types.Int64Type,
	"perf_events_statements_limit":             types.Int64Type,
	"perf_events_statements_time_limit":        types.Int64Type,
}

// Expand public function that converts tf object into dto
func Expand(ctx context.Context, diags *diag.Diagnostics, list types.List) *dtoUserConfig {
	return schemautil.ExpandListBlockNested[tfoUserConfig, dtoUserConfig](ctx, diags, expandUserConfig, list)
}

// Flatten public function that converts dto into tf object
func Flatten(ctx context.Context, diags *diag.Diagnostics, m map[string]any) types.List {
	o := new(dtoUserConfig)
	err := schemautil.MapToDTO(m, o)
	if err != nil {
		diags.AddError("failed to marshal map user config to dto", err.Error())
		return types.ListNull(types.ObjectType{AttrTypes: userConfigAttrs})
	}
	return schemautil.FlattenListBlockNested[dtoUserConfig, tfoUserConfig](ctx, diags, flattenUserConfig, userConfigAttrs, o)
}
