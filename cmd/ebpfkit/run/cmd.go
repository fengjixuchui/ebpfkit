/*
Copyright © 2020 GUILLAUME FOURNIER

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
package run

import (
	"github.com/spf13/cobra"
)

// EBPFKit represents the base command of ebpfKit
var EBPFKit = &cobra.Command{
	Use:  "ebpfkit",
	RunE: ebpfKitCmd,
}

var options CLIOptions

func init() {
	EBPFKit.Flags().VarP(
		NewLogLevelSanitizer(&options.LogLevel),
		"log-level",
		"l",
		`log level, options: panic, fatal, error, warn, info, debug or trace`)
	EBPFKit.Flags().IntVarP(
		&options.EBPFKit.TargetHTTPServerPort,
		"target-http-server-port",
		"p",
		8000,
		"Target HTTP server port used for C&C")
}