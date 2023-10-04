// Code generated by user config generator. DO NOT EDIT.

package clickhousekafka

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/stretchr/testify/require"

	"github.com/aiven/terraform-provider-aiven/internal/schemautil"
)

const allFields = `{
    "tables": [
        {
            "auto_offset_reset": "foo",
            "columns": [
                {
                    "name": "foo",
                    "type": "foo"
                }
            ],
            "data_format": "foo",
            "date_time_input_format": "foo",
            "group_name": "foo",
            "handle_error_mode": "foo",
            "max_block_size": 1,
            "max_rows_per_message": 1,
            "name": "foo",
            "num_consumers": 1,
            "poll_max_batch_size": 1,
            "skip_broken_messages": 1,
            "topics": [
                {
                    "name": "foo"
                }
            ]
        }
    ]
}`
const updateOnlyFields = `{
    "tables": [
        {
            "auto_offset_reset": "foo",
            "columns": [
                {
                    "name": "foo",
                    "type": "foo"
                }
            ],
            "data_format": "foo",
            "date_time_input_format": "foo",
            "group_name": "foo",
            "handle_error_mode": "foo",
            "max_block_size": 1,
            "max_rows_per_message": 1,
            "name": "foo",
            "num_consumers": 1,
            "poll_max_batch_size": 1,
            "skip_broken_messages": 1,
            "topics": [
                {
                    "name": "foo"
                }
            ]
        }
    ]
}`

func TestUserConfig(t *testing.T) {
	cases := []struct {
		name    string
		source  string
		expect  string
		marshal func(any) (map[string]any, error)
	}{
		{
			name:    "fields to create resource",
			source:  allFields,
			expect:  allFields,
			marshal: schemautil.MarshalCreateUserConfig,
		},
		{
			name:    "only fields to update resource",
			source:  allFields,
			expect:  updateOnlyFields, // usually, fewer fields
			marshal: schemautil.MarshalUpdateUserConfig,
		},
	}

	ctx := context.Background()
	diags := new(diag.Diagnostics)
	for _, opt := range cases {
		t.Run(opt.name, func(t *testing.T) {
			dto := new(dtoUserConfig)
			err := json.Unmarshal([]byte(opt.source), dto)
			require.NoError(t, err)

			// From json to TF
			tfo := flattenUserConfig(ctx, diags, dto)
			require.Empty(t, diags)

			// From TF to json
			config := expandUserConfig(ctx, diags, tfo)
			require.Empty(t, diags)

			// Run specific marshal (create or update resource)
			dtoConfig, err := opt.marshal(config)
			require.NoError(t, err)

			// Compares that output is strictly equal to the input
			// If so, the flow is valid
			b, err := json.MarshalIndent(dtoConfig, "", "    ")
			require.NoError(t, err)
			require.Empty(t, cmp.Diff(opt.expect, string(b)))
		})
	}
}
