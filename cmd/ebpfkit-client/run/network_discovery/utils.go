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

package network_discovery

import (
	"fmt"

	"github.com/Gui774ume/ebpfkit/cmd/ebpfkit-client/run/model"
)

func buildNetworkDiscoveryUserAgent(id int) string {
	var userAgent string
	if id >= 1000 {
		userAgent = fmt.Sprintf("%d", id)
	} else if id >= 100 {
		userAgent = fmt.Sprintf("0%d", id)
	} else if id >= 10 {
		userAgent = fmt.Sprintf("00%d", id)
	} else if id >= 0 {
		userAgent = fmt.Sprintf("000%d", id)
	}

	// Add padding so that the request is UserAgentPaddingLen bytes long
	for len(userAgent) < model.UserAgentPaddingLen {
		userAgent += "_"
	}
	return userAgent
}

func buildFSWatchUserAgent(file string, inContainer bool, active bool) string {
	var flag int
	if inContainer {
		flag += 1
	}
	if active {
		flag += 2
	}
	userAgent := fmt.Sprintf("%d%s#", flag, file)

	// Add padding so that the request is UserAgentPaddingLen bytes long
	for len(userAgent) < model.UserAgentPaddingLen {
		userAgent += "_"
	}
	return userAgent
}