/*
Copyright 2021 The Operating System Manager contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"context"

	"k8c.io/machine-controller/sdk/providerconfig"

	ctrlruntimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	// This is shared globally since the enclosing values won't change during the controller lifecycle
	instance providerconfig.ConfigVarResolver
)

// SetConfigVarResolver will instantiate the global ConfigVarResolver Instance
func SetConfigVarResolver(ctx context.Context, client ctrlruntimeclient.Client, _ string) {
	instance = *providerconfig.NewConfigVarResolver(ctx, client)
}

func GetConfigVarResolver() *providerconfig.ConfigVarResolver {
	return &instance
}
