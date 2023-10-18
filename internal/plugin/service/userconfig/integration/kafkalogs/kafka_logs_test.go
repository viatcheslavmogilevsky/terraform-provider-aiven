// Code generated by user config generator. DO NOT EDIT.

package kafkalogs

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
    "kafka_topic": "foo",
    "selected_log_fields": [
        "foo"
    ]
}`
const updateOnlyFields = `{
    "kafka_topic": "foo",
    "selected_log_fields": [
        "foo"
    ]
}`

func TestUserConfig(t *testing.T) {
	cases := []struct {
		name   string
		source string
		expect string
		create bool
	}{
		{
			name:   "fields to create resource",
			source: allFields,
			expect: allFields,
			create: true,
		},
		{
			name:   "only fields to update resource",
			source: allFields,
			expect: updateOnlyFields, // usually, fewer fields
			create: false,
		},
	}

	ctx := context.Background()
	diags := make(diag.Diagnostics, 0)
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
			dtoConfig, err := schemautil.MarshalUserConfig(config, opt.create)
			require.NoError(t, err)

			// Compares that output is strictly equal to the input
			// If so, the flow is valid
			b, err := json.MarshalIndent(dtoConfig, "", "    ")
			require.NoError(t, err)
			require.Empty(t, cmp.Diff(opt.expect, string(b)))
		})
	}
}
