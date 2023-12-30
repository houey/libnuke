package config

import (
	"github.com/ekristen/cloud-nuke-sdk/pkg/filter"
	"github.com/ekristen/cloud-nuke-sdk/pkg/types"
)

type Account struct {
	Filters       filter.Filters `yaml:"filters"`
	ResourceTypes ResourceTypes  `yaml:"resource-types"`
	Presets       []string       `yaml:"presets"`
}

type ResourceTypes struct {
	Targets  types.Collection `yaml:"targets"`
	Excludes types.Collection `yaml:"excludes"`
}

type Config interface {
	ResolveBlocklist() []string
	HasBlocklist() bool
	InBlocklist(searchID string) bool
	Validate(id string) error
	Filters(id string) (filter.Filters, error)
	Presets() map[string]PresetDefinitions
	ResourceTypes() ResourceTypes
	FeatureFlags() FeatureFlags
	ResolveDeprecations() error
}

type FeatureFlags struct {
	DisableDeletionProtection  DisableDeletionProtection `yaml:"disable-deletion-protection"`
	ForceDeleteLightsailAddOns bool                      `yaml:"force-delete-lightsail-addons"`
}

type DisableDeletionProtection struct {
	RDSInstance         bool `yaml:"RDSInstance"`
	EC2Instance         bool `yaml:"EC2Instance"`
	CloudformationStack bool `yaml:"CloudformationStack"`
}

type PresetDefinitions struct {
	Filters filter.Filters `yaml:"filters"`
}