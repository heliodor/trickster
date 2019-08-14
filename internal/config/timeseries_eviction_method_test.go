/**
* Copyright 2018 Comcast Cable Communications Management, LLC
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
* http://www.apache.org/licenses/LICENSE-2.0
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package config

import "testing"

func TestTEMString(t *testing.T) {

	t1 := EvictionMethodLRU
	t2 := EvictionMethodOldest

	if t1.String() != "lru" {
		t.Errorf("expected %s got %s", "lru", t1.String())
	}

	if t2.String() != "oldest" {
		t.Errorf("expected %s got %s", "oldest", t2.String())
	}

}
