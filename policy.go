/* Copyright 2024 İrem Kuyucu <irem@digilol.net>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package firewalld

import "github.com/godbus/dbus/v5"

// PolicyParams represents the parameters for a firewall policy
type PolicyParams struct {
	IngressZones []string
	EgressZones  []string
	Priority     int
	Target       string
	RichRules    []string
}

func (fw *FirewalldClient) AddPolicyPermanent(policyName string, params PolicyParams) error {
	obj := fw.conn.Object("org.fedoraproject.FirewallD1", "/org/fedoraproject/FirewallD1/config")

	// Convert struct to a map
	policyParams := map[string]interface{}{
		"ingress_zones": params.IngressZones,
		"egress_zones":  params.EgressZones,
		"priority":      params.Priority,
		"target":        params.Target,
		"rich_rules":    params.RichRules,
	}

	call := obj.Call("org.fedoraproject.FirewallD1.config.addPolicy", 0, policyName, policyParams)
	return call.Err
}

func (fw *FirewalldClient) GetPolicyPathPermanent(policyName string) (string, error) {
	obj := fw.conn.Object("org.fedoraproject.FirewallD1", "/org/fedoraproject/FirewallD1/config")

	var policyPath string
	err := obj.Call("org.fedoraproject.FirewallD1.config.getPolicyByName", 0, policyName).Store(&policyPath)
	if err != nil {
		return "", err
	}

	return policyPath, nil
}

func (fw *FirewalldClient) GetPolicySettingsPermanent(policyPath string) (map[string]interface{}, error) {
	obj := fw.conn.Object("org.fedoraproject.FirewallD1", dbus.ObjectPath(policyPath))

	var settings map[string]interface{}
	err := obj.Call("org.fedoraproject.FirewallD1.config.policy.getSettings", 0).Store(&settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (fw *FirewalldClient) UpdatePolicyPermanent(policyPath string, params PolicyParams) error {
	obj := fw.conn.Object("org.fedoraproject.FirewallD1", dbus.ObjectPath(policyPath))

	// Convert struct to a map
	policyParams := make(map[string]interface{})

	if len(params.IngressZones) != 0 {
		policyParams["ingress_zones"] = params.IngressZones
	}
	if len(params.EgressZones) != 0 {
		policyParams["egress_zones"] = params.EgressZones
	}
	if params.Priority != 0 {
		policyParams["priority"] = params.Priority
	}
	if params.Target != "" {
		policyParams["target"] = params.Target
	}
	if len(params.RichRules) != 0 {
		policyParams["rich_rules"] = params.RichRules
	}

	call := obj.Call("org.fedoraproject.FirewallD1.config.policy.update", 0, policyParams)
	return call.Err
}
