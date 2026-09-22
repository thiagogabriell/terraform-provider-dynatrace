//go:build unit

/*
 * @license
 * Copyright 2026 Dynatrace LLC
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package provider_test

import (
	"testing"

	"github.com/dynatrace-oss/terraform-provider-dynatrace/provider"
	"github.com/dynatrace-oss/terraform-provider-dynatrace/provider/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The export command reads the provider configuration through the schema's DefaultFunc rather than
// from Terraform, so the environment variables are the only way to configure it there.
func TestWIFAudienceIsReadFromTheEnvironment(t *testing.T) {
	for _, envVarName := range []string{"DYNATRACE_WIF_AUDIENCE", "DT_WIF_AUDIENCE"} {
		t.Run(envVarName, func(t *testing.T) {
			t.Setenv("DYNATRACE_WIF_AUDIENCE", "")
			t.Setenv("DT_WIF_AUDIENCE", "")
			t.Setenv(envVarName, "dynatrace")

			credentials := createCredentials(&config.ConfigGetter{Provider: provider.Provider()})

			require.NotNil(t, credentials)
			assert.Equal(t, "dynatrace", credentials.Platform.WorkloadIdentityFederationAudience)
		})
	}
}
