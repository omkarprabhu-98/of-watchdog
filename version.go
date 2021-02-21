// Copyright (c) OpenFaaS Author(s) 2021. All rights reserved.
// Licensed under the MIT license. See LICENSE file in the project root for full license information.

package main

var (
	// Version release version of the watchdog
	Version string
	// GitCommit SHA of the last git commit
	GitCommit string
	// DevVersion string for the development version
	DevVersion = "dev"
)

// BuildVersion returns current version of watchdog
func BuildVersion() string {
	if len(Version) == 0 {
		return DevVersion
	}
	return Version
}
