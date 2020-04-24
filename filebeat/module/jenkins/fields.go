// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package jenkins

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("filebeat", "jenkins", asset.ModuleFieldsPri, AssetJenkins); err != nil {
		panic(err)
	}
}

// AssetJenkins returns asset data.
// This is the base64 encoded gzipped contents of module/jenkins.
func AssetJenkins() string {
	return "eJyMlF9P2zAUxd/7Ka54Am3E+UNLyaRJ7I9gAyS0VdreKhPfNqaJ7dk3o92nnwwhSzoH0Yc+XB/f8zu+do5hg7sc7lFtpHITAJJUYQ4HX58qBxMAga6w0pDUKof3EwCAdhVutGgqnACsJFbC5Y+Lx6B4jf2m/kc7gzmsrW5MWwn0HTbqN+NFgc515b3dtxVyh8CF6Nd7atzy2vhgPWVb66n+Zwwh/YOyWGvCpeyr98i+GO9l0TnQK6ASoagkKhrs6OiSOMrSKJ1GSZwMFE9kA6NnCJI1jvt/4oRe4d0t/mrQjVinJ+zcWJbGaZwn89z/n8GbOI3jAIjghAGUGqnUYhzmcrG4DYk6iIvPi4DbBncP2oqAoeFUjtvdciqfQ+OII/MtGGlm+Bqje2cC/oRbCpiXRGb5G60bXrX9zEQGQqoOwZ8KS6LQwMeTW3RGK4fLQosXxv944t9aLXgtWKTGqrHzyAYX79ntbkfoxl0++GVwqARI1aGFDaazkyx0tRXhGu1LSV81bK8VKOBB7ok7AD82lzNW79oPVKQqxo2JE3av79jr5984tHy9/5YHUJ0ESuQCLaysrl9+hzf6j6wqzqZRDIc/k+QdXEvVbGE7ny1nJ0dwbkyFP/DuShKbZqdRNoPDq8vFzfVbqOQG4QKLjT6Cj6XVNbJ5HMVRdjY/jZIshe98xa1st40F/RsAAP//WRmo3w=="
}
