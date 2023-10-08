/*
 * Copyright 2023 Caio Matheus Marcatti Calimério
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
package cors

import (
	"sync"
)

var corsStore *CorsStore
var corsStoreOnce sync.Once

// GetCorsStore returns the global instance of CorsStore.
func GetCorsStore() *CorsStore {
	corsStoreOnce.Do(func() {
		corsStore = &CorsStore{
			config: &CORSConfig{},
		}
	})
	return corsStore
}

// Add updates the CORS configuration.
func (rs *CorsStore) Add(config *CORSConfig) {
	rs.mu.Lock()
	defer rs.mu.Unlock()
	rs.config = config
}

// Get retrieves the current CORS configuration.
func (rs *CorsStore) Get() *CORSConfig {
	rs.mu.RLock()
	defer rs.mu.RUnlock()

	return rs.config
}
