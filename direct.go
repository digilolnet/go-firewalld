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

import "strings"

// ipv is either "ipv4" or "ipv6", for table refer to iptables docs, chain is the name of the new chain
func (fw *FirewalldClient) DirectAddChainPermanent(ipv, table, chain string) error {
	obj := fw.conn.Object("org.fedoraproject.FirewallD1", "/org/fedoraproject/FirewallD1/config")
	call := obj.Call("org.fedoraproject.FirewallD1.config.direct.addChain", 0, ipv, table, chain)
	return call.Err
}

func (fw *FirewalldClient) DirectAddRulePermanent(ipv, table, chain string, priority int, rules string) error {
	obj := fw.conn.Object("org.fedoraproject.FirewallD1", "/org/fedoraproject/FirewallD1/config")
	call := obj.Call("org.fedoraproject.FirewallD1.config.direct.addRule", 0, ipv, table, chain, priority, strings.Split(rules, " "))
	return call.Err
}

func (fw *FirewalldClient) DirectRemoveRulesPermanent(ipv, table, chain string) error {
	obj := fw.conn.Object("org.fedoraproject.FirewallD1", "/org/fedoraproject/FirewallD1/config")
	call := obj.Call("org.fedoraproject.FirewallD1.config.direct.removeRules", 0, ipv, table, chain)
	return call.Err
}

// Check whether if the specified chain exists
func (fw *FirewalldClient) DirectQueryChainPermanent(ipv, table, chain string) (bool, error) {
        var exists bool
        obj := fw.conn.Object("org.fedoraproject.FirewallD1", "/org/fedoraproject/FirewallD1/config")
        err := obj.Call("org.fedoraproject.FirewallD1.config.direct.queryChain", 0, ipv, table, chain).Store(&exists)
        return exists, err
}
