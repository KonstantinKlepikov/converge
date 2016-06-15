// Copyright © 2016 Asteris, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package resource

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Template is a task defined by content and a destination
type Template struct {
	Name           string
	RawContent     string   `hcl:"content"`
	RawDestination string   `hcl:"destination"`
	Dependencies   []string `hcl:"depends"`

	renderer    *Renderer
	content     string
	destination string
}

// String returns the name of this template
func (t *Template) String() string {
	return "template." + t.Name
}

// SetDepends overwrites the Dependencies of this resource
func (t *Template) SetDepends(deps []string) {
	t.Dependencies = deps
}

// Depends list dependencies for this task
func (t *Template) Depends() []string {
	return t.Dependencies
}

// Check satisfies the Monitor interface
func (t *Template) Check() (string, bool, error) {
	stat, err := os.Stat(t.destination)
	if os.IsNotExist(err) {
		return "", true, nil
	} else if err != nil {
		return "", false, err
	} else if stat.IsDir() {
		return "", true, fmt.Errorf("cannot template %q, is a directory", t.destination)
	}

	actual, err := ioutil.ReadFile(t.destination)
	if err != nil {
		return "", false, err
	}

	return string(actual), t.content != string(actual), nil
}

// Apply (plus Check) satisfies the Task interface
func (t *Template) Apply() (string, bool, error) {
	err := ioutil.WriteFile(t.destination, []byte(t.content), 0600)
	if err != nil {
		return "", false, err
	}

	actual, err := ioutil.ReadFile(t.destination)
	if err != nil {
		return "", false, err
	}

	return string(actual), t.content == string(actual), err
}

// Prepare this module for use
func (t *Template) Prepare(parent *Module) (err error) {
	t.renderer, err = NewRenderer(parent)
	if err != nil {
		return err
	}

	// check the rendered input is good
	t.content, err = t.renderer.Render(t.String()+".content", t.RawContent)
	if err != nil {
		return ValidationError{Location: t.String() + ".content", Err: err}
	}

	// check the rendered destination
	t.destination, err = t.renderer.Render(t.String()+".destination", t.RawDestination)
	if err != nil {
		return ValidationError{Location: t.String() + ".destination", Err: err}
	}

	// get param dependencies
	t.Dependencies, err = t.renderer.Dependencies(
		t.String()+".dependencies",
		t.Dependencies,
		t.RawContent,
		t.RawDestination,
	)
	if err != nil {
		return ValidationError{Location: t.String() + ".dependencies", Err: err}
	}

	return nil
}
