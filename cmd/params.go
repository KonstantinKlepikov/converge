// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"encoding/json"

	"github.com/asteris-llc/converge/resource"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func addParamsArguments(flags *pflag.FlagSet) {
	flags.StringP("paramsJSON", "p", "{}", "parameters for the top-level module, in JSON format")
}

func getParamsFromFlags() (resource.Values, error) {
	params := resource.Values{}
	err := json.Unmarshal([]byte(viper.GetString("paramsJson")), &params)

	return params, err
}