// Copyright 2016 Palantir Technologies, Inc.
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

package disterfactory

import (
	"github.com/palantir/distgo/dister"
	"github.com/palantir/distgo/dister/bin"
	binconfig "github.com/palantir/distgo/dister/bin/config"
	"github.com/palantir/distgo/dister/manual"
	manualconfig "github.com/palantir/distgo/dister/manual/config"
	"github.com/palantir/distgo/dister/osarchbin"
	osarchbinconfig "github.com/palantir/distgo/dister/osarchbin/config"
	"github.com/palantir/distgo/distgo"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type creatorWithUpgrader struct {
	creator  dister.CreatorFunction
	upgrader distgo.ConfigUpgrader
}

func builtinDisters() map[string]creatorWithUpgrader {
	return map[string]creatorWithUpgrader{
		bin.TypeName: {
			creator: func(cfgYML []byte) (distgo.Dister, error) {
				return bin.New(), nil
			},
			upgrader: distgo.NewConfigUpgrader(bin.TypeName, binconfig.UpgradeConfig),
		},
		osarchbin.TypeName: {
			creator: func(cfgYML []byte) (distgo.Dister, error) {
				var cfg osarchbinconfig.OSArchBin
				if err := yaml.UnmarshalStrict(cfgYML, &cfg); err != nil {
					return nil, errors.Wrapf(err, "failed to unmarshal YAML")
				}
				return cfg.ToDister(), nil
			},
			upgrader: distgo.NewConfigUpgrader(osarchbin.TypeName, osarchbinconfig.UpgradeConfig),
		},
		manual.TypeName: {
			creator: func(cfgYML []byte) (distgo.Dister, error) {
				var cfg manualconfig.Manual
				if err := yaml.UnmarshalStrict(cfgYML, &cfg); err != nil {
					return nil, errors.Wrapf(err, "failed to unmarshal YAML")
				}
				return cfg.ToDister(), nil
			},
			upgrader: distgo.NewConfigUpgrader(manual.TypeName, manualconfig.UpgradeConfig),
		},
	}
}
