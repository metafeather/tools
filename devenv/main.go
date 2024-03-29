/*
Copyright © 2024 Liam Clancy (metafeather) github@metafeather.com

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
package main

import (
	"embed"
	"time"

	"github.com/carlmjohnson/versioninfo"
	"github.com/metafeather/tools/devenv/cmd"
)

// Submodules cannot be embedded directly so zip them up
//
//go:generate zip -r embed/stdlib/global/publish.zip embed/stdlib/global/publish
//go:embed all:embed/stdlib
var stdlib embed.FS

func main() {
	// Share embedded stdlib filesystem with all commands
	cmd.StdlibFS = stdlib

	cmd.SetVersionInfo(versioninfo.Version, versioninfo.Revision, versioninfo.LastCommit.Format(time.RFC3339))
	cmd.Execute()
}
