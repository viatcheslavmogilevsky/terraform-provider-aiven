// Code generated by user config generator. DO NOT EDIT.

package kafkamirrormaker

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
			Attributes: map[string]resource.Attribute{"cluster_alias": resource.StringAttribute{
				Computed:    true,
				Description: "The alias under which the Kafka cluster is known to MirrorMaker. Can contain the following symbols: ASCII alphanumerics, '.', '_', and '-'.",
				Optional:    true,
			}},
			Blocks: map[string]resource.Block{"kafka_mirrormaker": resource.ListNestedBlock{
				Description: "Kafka MirrorMaker configuration values",
				NestedObject: resource.NestedBlockObject{Attributes: map[string]resource.Attribute{
					"consumer_fetch_min_bytes": resource.Int64Attribute{
						Computed:    true,
						Description: "The minimum amount of data the server should return for a fetch request.",
						Optional:    true,
					},
					"producer_batch_size": resource.Int64Attribute{
						Computed:    true,
						Description: "The batch size in bytes producer will attempt to collect before publishing to broker.",
						Optional:    true,
					},
					"producer_buffer_memory": resource.Int64Attribute{
						Computed:    true,
						Description: "The amount of bytes producer can use for buffering data before publishing to broker.",
						Optional:    true,
					},
					"producer_compression_type": resource.StringAttribute{
						Computed:    true,
						Description: "Specify the default compression type for producers. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'none' which is the default and equivalent to no compression.",
						Optional:    true,
					},
					"producer_linger_ms": resource.Int64Attribute{
						Computed:    true,
						Description: "The linger time (ms) for waiting new data to arrive for publishing.",
						Optional:    true,
					},
					"producer_max_request_size": resource.Int64Attribute{
						Computed:    true,
						Description: "The maximum request size in bytes.",
						Optional:    true,
					},
				}},
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
			Attributes: map[string]datasource.Attribute{"cluster_alias": datasource.StringAttribute{
				Computed:    true,
				Description: "The alias under which the Kafka cluster is known to MirrorMaker. Can contain the following symbols: ASCII alphanumerics, '.', '_', and '-'.",
			}},
			Blocks: map[string]datasource.Block{"kafka_mirrormaker": datasource.ListNestedBlock{
				Description: "Kafka MirrorMaker configuration values",
				NestedObject: datasource.NestedBlockObject{Attributes: map[string]datasource.Attribute{
					"consumer_fetch_min_bytes": datasource.Int64Attribute{
						Computed:    true,
						Description: "The minimum amount of data the server should return for a fetch request.",
					},
					"producer_batch_size": datasource.Int64Attribute{
						Computed:    true,
						Description: "The batch size in bytes producer will attempt to collect before publishing to broker.",
					},
					"producer_buffer_memory": datasource.Int64Attribute{
						Computed:    true,
						Description: "The amount of bytes producer can use for buffering data before publishing to broker.",
					},
					"producer_compression_type": datasource.StringAttribute{
						Computed:    true,
						Description: "Specify the default compression type for producers. This configuration accepts the standard compression codecs ('gzip', 'snappy', 'lz4', 'zstd'). It additionally accepts 'none' which is the default and equivalent to no compression.",
					},
					"producer_linger_ms": datasource.Int64Attribute{
						Computed:    true,
						Description: "The linger time (ms) for waiting new data to arrive for publishing.",
					},
					"producer_max_request_size": datasource.Int64Attribute{
						Computed:    true,
						Description: "The maximum request size in bytes.",
					},
				}},
			}},
		},
		Validators: []validator.List{listvalidator.SizeAtMost(1)},
	}
}

// tfoUserConfig Integration user config
type tfoUserConfig struct {
	ClusterAlias     types.String `tfsdk:"cluster_alias"`
	KafkaMirrormaker types.List   `tfsdk:"kafka_mirrormaker"`
}

// dtoUserConfig request/response object
type dtoUserConfig struct {
	ClusterAlias     *string              `groups:"create,update" json:"cluster_alias,omitempty"`
	KafkaMirrormaker *dtoKafkaMirrormaker `groups:"create,update" json:"kafka_mirrormaker,omitempty"`
}

// expandUserConfig expands tf object into dto object
func expandUserConfig(ctx context.Context, diags *diag.Diagnostics, o *tfoUserConfig) *dtoUserConfig {
	kafkaMirrormakerVar := schemautil.ExpandListBlockNested[tfoKafkaMirrormaker, dtoKafkaMirrormaker](ctx, diags, expandKafkaMirrormaker, o.KafkaMirrormaker)
	if diags.HasError() {
		return nil
	}
	return &dtoUserConfig{
		ClusterAlias:     schemautil.ValueStringPointer(o.ClusterAlias),
		KafkaMirrormaker: kafkaMirrormakerVar,
	}
}

// flattenUserConfig flattens dto object into tf object
func flattenUserConfig(ctx context.Context, diags *diag.Diagnostics, o *dtoUserConfig) *tfoUserConfig {
	kafkaMirrormakerVar := schemautil.FlattenListBlockNested[dtoKafkaMirrormaker, tfoKafkaMirrormaker](ctx, diags, flattenKafkaMirrormaker, kafkaMirrormakerAttrs, o.KafkaMirrormaker)
	if diags.HasError() {
		return nil
	}
	return &tfoUserConfig{
		ClusterAlias:     types.StringPointerValue(o.ClusterAlias),
		KafkaMirrormaker: kafkaMirrormakerVar,
	}
}

var userConfigAttrs = map[string]attr.Type{
	"cluster_alias":     types.StringType,
	"kafka_mirrormaker": types.ListType{ElemType: types.ObjectType{AttrTypes: kafkaMirrormakerAttrs}},
}

// tfoKafkaMirrormaker Kafka MirrorMaker configuration values
type tfoKafkaMirrormaker struct {
	ConsumerFetchMinBytes   types.Int64  `tfsdk:"consumer_fetch_min_bytes"`
	ProducerBatchSize       types.Int64  `tfsdk:"producer_batch_size"`
	ProducerBufferMemory    types.Int64  `tfsdk:"producer_buffer_memory"`
	ProducerCompressionType types.String `tfsdk:"producer_compression_type"`
	ProducerLingerMs        types.Int64  `tfsdk:"producer_linger_ms"`
	ProducerMaxRequestSize  types.Int64  `tfsdk:"producer_max_request_size"`
}

// dtoKafkaMirrormaker request/response object
type dtoKafkaMirrormaker struct {
	ConsumerFetchMinBytes   *int64  `groups:"create,update" json:"consumer_fetch_min_bytes,omitempty"`
	ProducerBatchSize       *int64  `groups:"create,update" json:"producer_batch_size,omitempty"`
	ProducerBufferMemory    *int64  `groups:"create,update" json:"producer_buffer_memory,omitempty"`
	ProducerCompressionType *string `groups:"create,update" json:"producer_compression_type,omitempty"`
	ProducerLingerMs        *int64  `groups:"create,update" json:"producer_linger_ms,omitempty"`
	ProducerMaxRequestSize  *int64  `groups:"create,update" json:"producer_max_request_size,omitempty"`
}

// expandKafkaMirrormaker expands tf object into dto object
func expandKafkaMirrormaker(ctx context.Context, diags *diag.Diagnostics, o *tfoKafkaMirrormaker) *dtoKafkaMirrormaker {
	return &dtoKafkaMirrormaker{
		ConsumerFetchMinBytes:   schemautil.ValueInt64Pointer(o.ConsumerFetchMinBytes),
		ProducerBatchSize:       schemautil.ValueInt64Pointer(o.ProducerBatchSize),
		ProducerBufferMemory:    schemautil.ValueInt64Pointer(o.ProducerBufferMemory),
		ProducerCompressionType: schemautil.ValueStringPointer(o.ProducerCompressionType),
		ProducerLingerMs:        schemautil.ValueInt64Pointer(o.ProducerLingerMs),
		ProducerMaxRequestSize:  schemautil.ValueInt64Pointer(o.ProducerMaxRequestSize),
	}
}

// flattenKafkaMirrormaker flattens dto object into tf object
func flattenKafkaMirrormaker(ctx context.Context, diags *diag.Diagnostics, o *dtoKafkaMirrormaker) *tfoKafkaMirrormaker {
	return &tfoKafkaMirrormaker{
		ConsumerFetchMinBytes:   types.Int64PointerValue(o.ConsumerFetchMinBytes),
		ProducerBatchSize:       types.Int64PointerValue(o.ProducerBatchSize),
		ProducerBufferMemory:    types.Int64PointerValue(o.ProducerBufferMemory),
		ProducerCompressionType: types.StringPointerValue(o.ProducerCompressionType),
		ProducerLingerMs:        types.Int64PointerValue(o.ProducerLingerMs),
		ProducerMaxRequestSize:  types.Int64PointerValue(o.ProducerMaxRequestSize),
	}
}

var kafkaMirrormakerAttrs = map[string]attr.Type{
	"consumer_fetch_min_bytes":  types.Int64Type,
	"producer_batch_size":       types.Int64Type,
	"producer_buffer_memory":    types.Int64Type,
	"producer_compression_type": types.StringType,
	"producer_linger_ms":        types.Int64Type,
	"producer_max_request_size": types.Int64Type,
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
