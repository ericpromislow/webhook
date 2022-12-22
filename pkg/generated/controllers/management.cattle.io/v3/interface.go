/*
Copyright 2022 Rancher Labs, Inc.

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

// Code generated by main. DO NOT EDIT.

package v3

import (
	"github.com/rancher/lasso/pkg/controller"
	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/schemes"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v3.AddToScheme)
}

type Interface interface {
	Cluster() ClusterController
	ClusterRoleTemplateBinding() ClusterRoleTemplateBindingController
	GlobalRole() GlobalRoleController
	PodSecurityAdmissionConfigurationTemplate() PodSecurityAdmissionConfigurationTemplateController
	ProjectRoleTemplateBinding() ProjectRoleTemplateBindingController
	RoleTemplate() RoleTemplateController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (c *version) Cluster() ClusterController {
	return NewClusterController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "Cluster"}, "clusters", false, c.controllerFactory)
}
func (c *version) ClusterRoleTemplateBinding() ClusterRoleTemplateBindingController {
	return NewClusterRoleTemplateBindingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ClusterRoleTemplateBinding"}, "clusterroletemplatebindings", true, c.controllerFactory)
}
func (c *version) GlobalRole() GlobalRoleController {
	return NewGlobalRoleController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "GlobalRole"}, "globalroles", false, c.controllerFactory)
}
func (c *version) PodSecurityAdmissionConfigurationTemplate() PodSecurityAdmissionConfigurationTemplateController {
	return NewPodSecurityAdmissionConfigurationTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "PodSecurityAdmissionConfigurationTemplate"}, "podsecurityadmissionconfigurationtemplates", false, c.controllerFactory)
}
func (c *version) ProjectRoleTemplateBinding() ProjectRoleTemplateBindingController {
	return NewProjectRoleTemplateBindingController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "ProjectRoleTemplateBinding"}, "projectroletemplatebindings", true, c.controllerFactory)
}
func (c *version) RoleTemplate() RoleTemplateController {
	return NewRoleTemplateController(schema.GroupVersionKind{Group: "management.cattle.io", Version: "v3", Kind: "RoleTemplate"}, "roletemplates", false, c.controllerFactory)
}
