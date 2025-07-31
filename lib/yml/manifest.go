package yml

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type ManifestGroup int

const (
	ManifestGroupCRD ManifestGroup = iota + 1
	ManifestGroupNamespace
	ManifestGroupRBAC
	ManifestGroupOther
)

type ManifestHeader struct {
	APIVersion string `yaml:"apiVersion"`
	Kind       string `yaml:"kind"`
	Metadata   struct {
		Name      string `yaml:"name"`
		Namespace string `yaml:"namespace,omitempty"`
	} `yaml:"metadata"`
}

func ParseHeader(manifest string) ManifestHeader {
	var header ManifestHeader
	var dec = yaml.NewDecoder(strings.NewReader(manifest))
	dec.KnownFields(false)
	if err := dec.Decode(&header); err != nil {
		panic(err)
	}
	if header.APIVersion == "" {
		panic("apiVersion is required")
	}
	if header.Kind == "" {
		panic("kind is required")
	}
	if header.Metadata.Name == "" {
		panic("name is required")
	}
	if header.Metadata.Namespace == "" {
		header.Metadata.Namespace = "default"
	}
	return header
}

func (m ManifestHeader) Hash() string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(
		fmt.Sprintf("%s|%s|%s|%s", m.APIVersion, m.Kind, m.Metadata.Name, m.Metadata.Namespace),
	)))
}

func (m ManifestHeader) Group() ManifestGroup {
	if m.APIVersion == "apiextensions.k8s.io/v1" && m.Kind == "CustomResourceDefinition" {
		return ManifestGroupCRD
	} else if m.APIVersion == "v1" && m.Kind == "Namespace" {
		return ManifestGroupNamespace
	} else if strings.HasPrefix(m.APIVersion, "rbac.authorization.k8s.io/") &&
		(m.Kind == "ClusterRole" ||
			m.Kind == "Role" ||
			m.Kind == "ClusterRoleBinding" ||
			m.Kind == "RoleBinding") {
		return ManifestGroupRBAC
	} else {
		return ManifestGroupOther
	}
}
