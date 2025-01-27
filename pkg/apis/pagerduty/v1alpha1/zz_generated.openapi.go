//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/openshift/pagerduty-operator/pkg/apis/pagerduty/v1alpha1.PagerDutyIntegration":       schema_pkg_apis_pagerduty_v1alpha1_PagerDutyIntegration(ref),
		"github.com/openshift/pagerduty-operator/pkg/apis/pagerduty/v1alpha1.PagerDutyIntegrationSpec":   schema_pkg_apis_pagerduty_v1alpha1_PagerDutyIntegrationSpec(ref),
		"github.com/openshift/pagerduty-operator/pkg/apis/pagerduty/v1alpha1.PagerDutyIntegrationStatus": schema_pkg_apis_pagerduty_v1alpha1_PagerDutyIntegrationStatus(ref),
	}
}

func schema_pkg_apis_pagerduty_v1alpha1_PagerDutyIntegration(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PagerDutyIntegration is the Schema for the pagerdutyintegrations API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/pagerduty-operator/pkg/apis/pagerduty/v1alpha1.PagerDutyIntegrationSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/openshift/pagerduty-operator/pkg/apis/pagerduty/v1alpha1.PagerDutyIntegrationStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/openshift/pagerduty-operator/pkg/apis/pagerduty/v1alpha1.PagerDutyIntegrationSpec", "github.com/openshift/pagerduty-operator/pkg/apis/pagerduty/v1alpha1.PagerDutyIntegrationStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_pagerduty_v1alpha1_PagerDutyIntegrationSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PagerDutyIntegrationSpec defines the desired state of PagerDutyIntegration",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"acknowledgeTimeout": {
						SchemaProps: spec.SchemaProps{
							Description: "Time in seconds that an incident changes to the Triggered State after being Acknowledged. Value must not be negative. Omitting or setting this field to 0 will disable the feature.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"escalationPolicy": {
						SchemaProps: spec.SchemaProps{
							Description: "ID of an existing Escalation Policy in PagerDuty.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resolveTimeout": {
						SchemaProps: spec.SchemaProps{
							Description: "Time in seconds that an incident is automatically resolved if left open for that long. Value must not be negative. Omitting or setting this field to 0 will disable the feature.",
							Type:        []string{"integer"},
							Format:      "int32",
						},
					},
					"servicePrefix": {
						SchemaProps: spec.SchemaProps{
							Description: "Prefix to set on the PagerDuty Service name.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"pagerdutyApiKeySecretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Reference to the secret containing PAGERDUTY_API_KEY.",
							Ref:         ref("k8s.io/api/core/v1.SecretReference"),
						},
					},
					"clusterDeploymentSelector": {
						SchemaProps: spec.SchemaProps{
							Description: "A label selector used to find which clusterdeployment CRs receive a PD integration based on this configuration.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"),
						},
					},
					"targetSecretRef": {
						SchemaProps: spec.SchemaProps{
							Description: "Name and namespace in the target cluster where the secret is synced.",
							Ref:         ref("k8s.io/api/core/v1.SecretReference"),
						},
					},
				},
				Required: []string{"escalationPolicy", "servicePrefix", "pagerdutyApiKeySecretRef", "clusterDeploymentSelector", "targetSecretRef"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.SecretReference", "k8s.io/apimachinery/pkg/apis/meta/v1.LabelSelector"},
	}
}

func schema_pkg_apis_pagerduty_v1alpha1_PagerDutyIntegrationStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "PagerDutyIntegrationStatus defines the observed state of PagerDutyIntegration",
				Type:        []string{"object"},
			},
		},
	}
}
