// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package args

import (
	"fmt"
	"github.com/ByteBam/thirftbam/pkg/plugin"
	"runtime"
	"strings"
	"time"
)

// StringSlice implements the flag.Value interface_info on string slices
// to allow a flag to be set multiple times.
type StringSlice []string

func (ss *StringSlice) String() string {
	return fmt.Sprintf("%v", *ss)
}

// Set implements the flag.Value interface_info.
func (ss *StringSlice) Set(value string) error {
	*ss = append(*ss, value)
	return nil
}

// Arguments contains command line arguments for thriftgo.
type Arguments struct {
	AskVersion      bool
	Recursive       bool
	Verbose         bool
	Quiet           bool
	CheckKeyword    bool
	OutputPath      string
	Includes        StringSlice
	Plugins         StringSlice
	Langs           StringSlice
	IDL             string
	PluginTimeLimit time.Duration
}

const WINDOWS_REPLACER = "#$$#"

// UsedPlugins returns a list of plugin.Desc for plugins.
func (a *Arguments) UsedPlugins() (descs []*plugin.Desc, err error) {
	for _, str := range a.Plugins {
		if runtime.GOOS == "windows" {
			// windows should replace :\ because thriftgo will separates args by ":"
			str = strings.ReplaceAll(str, ":\\", WINDOWS_REPLACER)
		}
		desc, err := plugin.ParseCompactArguments(str)
		if err != nil {
			return nil, err
		}
		if runtime.GOOS == "windows" {
			desc.Name = strings.ReplaceAll(desc.Name, WINDOWS_REPLACER, ":\\")
			for i, o := range desc.Options {
				desc.Options[i].Name = strings.ReplaceAll(o.Name, WINDOWS_REPLACER, ":\\")
				desc.Options[i].Desc = strings.ReplaceAll(o.Desc, WINDOWS_REPLACER, ":\\")
			}
		}
		descs = append(descs, desc)
	}
	return
}
