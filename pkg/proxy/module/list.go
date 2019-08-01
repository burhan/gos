/*
 * Copyright 2019 storyicon@foxmail.com
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

package module

import (
    "strings"
    "time"
)

// List is the standard return value, which corresponds to the go list -json command.
type List struct {
    Path     string    `json:"Path"`
    Version  string    `json:"Version"`
    Versions Versions  `json:",omitempty"`
    Time     time.Time `json:"time"`
    Dir      string    `json:"Dir"`
    GoMod    string    `json:"GoMod"`
}

// GetVersions is used to get all valid versions of package
func (s *List) GetVersions() Versions {
    versions := s.Versions
    if len(versions) == 0 {
        versions = append(versions, s.Version)
    }
    return versions
}

// Versions is the slice of version
type Versions []string

func (v Versions) String() string {
    return strings.Join(v, "\r\n")
}
