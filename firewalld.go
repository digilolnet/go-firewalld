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

type FirewalldClient struct {
	client *dbus.Conn
}

func NewFirewalldClient() (*FirewalldClient, error) {
	conn, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return &FirewalldClient{client: conn}, nil
}
