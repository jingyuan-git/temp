
/*
Copyright 2017 The Kubernetes Authors.
 
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



package demoadmission

import (
	"context"
	aggregatedadmission "temp/kubebuild/plugin/admission"
	aggregatedinformerfactory "temp/kubebuild/pkg/client/informers_generated/externalversions"
	aggregatedclientset "temp/kubebuild/pkg/client/clientset_generated/clientset"
	genericadmissioninitializer "k8s.io/apiserver/pkg/admission/initializer"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/apiserver/pkg/admission"
)

var _ admission.Interface 											= &demoPlugin{}
var _ admission.MutationInterface 									= &demoPlugin{}
var _ admission.ValidationInterface 								= &demoPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeInformerFactory 	= &demoPlugin{}
var _ genericadmissioninitializer.WantsExternalKubeClientSet 		= &demoPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceInformerFactory 	= &demoPlugin{}
var _ aggregatedadmission.WantsAggregatedResourceClientSet 			= &demoPlugin{}

func NewDemoPlugin() *demoPlugin {
	return &demoPlugin{
		Handler: admission.NewHandler(admission.Create, admission.Update),
	}
}

type demoPlugin struct {
	*admission.Handler
}

func (p *demoPlugin) ValidateInitialization() error {
	return nil
}

func (p *demoPlugin) Admit(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *demoPlugin) Validate(ctx context.Context, a admission.Attributes, o admission.ObjectInterfaces) error {
	return nil
}

func (p *demoPlugin) SetAggregatedResourceInformerFactory(aggregatedinformerfactory.SharedInformerFactory) {}

func (p *demoPlugin) SetAggregatedResourceClientSet(aggregatedclientset.Interface) {}

func (p *demoPlugin) SetExternalKubeInformerFactory(informers.SharedInformerFactory) {}

func (p *demoPlugin) SetExternalKubeClientSet(kubernetes.Interface) {}
