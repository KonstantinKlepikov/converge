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

package apt

import (
	"os/exec"
	"strings"

	"github.com/asteris-llc/converge/resource"
	"github.com/pkg/errors"
)

// Package is an API for package state
type Package struct {
	resource.TaskStatus

	Name  string
	State string
}

// Check if the package has to be 'installed', or 'absent'
func (p *Package) Check(resource.Renderer) (resource.TaskStatus, error) {
	status := resource.NewStatus()

	currentPkgStatusRaw, err := exec.Command("sh", "-c", "apt-cache policy "+p.Name+" | grep Installed | awk '{print $2}'").Output()
	if err != nil {
		return status, errors.Wrapf(err, "checking package %s", p.Name)
	}
	currentPkgStatus := strings.TrimSpace(string(currentPkgStatusRaw))

	if p.State == "installed" {
		if string(currentPkgStatus) == "(none)" {
			status.AddDifference(p.Name, string(currentPkgStatus), "installed", "")
		} else {
			status.AddMessage("Package is installed")
		}
	} else if p.State == "absent" {
		if string(currentPkgStatus) != "(none)" {
			status.AddDifference(p.Name, string(currentPkgStatus), "uninstalled", "")
		} else {
			status.AddMessage("Package is absent")
		}
	}

	p.TaskStatus = status
	return p, nil
}

// Apply desired package state
func (p *Package) Apply() (resource.TaskStatus, error) {
	status := resource.NewStatus()

	if p.State == "installed" {
		aptStatus, err := exec.Command("sh", "-c", "apt-get install -y "+p.Name).Output()
		if err != nil {
			return status, errors.Wrapf(err, "installing package %s, what happened: %s", p.Name, aptStatus)
		}

		status.AddDifference(p.Name, "absent", "installed", "")

	} else if p.State == "absent" {
		aptStatus, err := exec.Command("sh", "-c", "apt-get remove -y "+p.Name).Output()
		if err != nil {
			return status, errors.Wrapf(err, "removing package %s, what happened: %s", p.Name, aptStatus)
		}

		status.AddDifference(p.Name, "installed", "absent", "")
	}

	p.TaskStatus = status

	return p, nil
}
