// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// UnmarshalJSON implements json.Unmarshaler.
func (j *DevfileComponentsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain DevfileComponentsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if v, ok := raw["mountSources"]; !ok || v == nil {
		plain.MountSources = "false"
	}
	*j = DevfileComponentsElem(plain)
	return nil
}

// This schema describes the structure of the devfile object
type Devfile struct {
	// ApiVersion corresponds to the JSON schema field "apiVersion".
	ApiVersion interface{} `json:"apiVersion"`

	// Attributes corresponds to the JSON schema field "attributes".
	Attributes DevfileAttributes `json:"attributes,omitempty"`

	// Description of the predefined commands to be available in workspace
	Commands []DevfileCommandsElem `json:"commands,omitempty"`

	// Description of the workspace components, such as editor and plugins
	Components []DevfileComponentsElem `json:"components,omitempty"`

	// Metadata corresponds to the JSON schema field "metadata".
	Metadata DevfileMetadata `json:"metadata"`

	// Description of the projects, containing names and sources locations
	Projects []DevfileProjectsElem `json:"projects,omitempty"`
}

type DevfileAttributes map[string]interface{}

type DevfileCommandsElemActionsElem struct {
	// The actual action command-line string
	Command *string `json:"command,omitempty"`

	// Describes component to which given action relates
	Component *string `json:"component,omitempty"`

	// the path relative to the location of the devfile to the configuration file
	// defining one or more actions in the editor-specific format
	Reference *string `json:"reference,omitempty"`

	// The content of the referenced configuration file that defines one or more
	// actions in the editor-specific format
	ReferenceContent *string `json:"referenceContent,omitempty"`

	// Describes action type
	Type *string `json:"type,omitempty"`

	// Working directory where the command should be executed
	Workdir *string `json:"workdir,omitempty"`
}

type DevfileCommandsElem struct {
	// List of the actions of given command. Now the only one command must be
	// specified in list but there are plans to implement supporting multiple actions
	// commands.
	Actions []DevfileCommandsElemActionsElem `json:"actions"`

	// Additional command attributes
	Attributes Attributes `json:"attributes,omitempty"`

	// Describes the name of the command. Should be unique per commands set.
	Name string `json:"name"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DevfileCommandsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["actions"]; !ok || v == nil {
		return fmt.Errorf("field actions: required")
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	type Plain DevfileCommandsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DevfileCommandsElem(plain)
	return nil
}

type DevfileComponentsElemType string

// UnmarshalJSON implements json.Unmarshaler.
func (j *DevfileProjectsElem) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["name"]; !ok || v == nil {
		return fmt.Errorf("field name: required")
	}
	if v, ok := raw["source"]; !ok || v == nil {
		return fmt.Errorf("field source: required")
	}
	type Plain DevfileProjectsElem
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DevfileProjectsElem(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DevfileComponentsElemType) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_DevfileComponentsElemType {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_DevfileComponentsElemType, v)
	}
	*j = DevfileComponentsElemType(v)
	return nil
}

const DevfileComponentsElemTypeCheEditor DevfileComponentsElemType = "cheEditor"
const DevfileComponentsElemTypeChePlugin DevfileComponentsElemType = "chePlugin"

type Attributes map[string]interface{}

// UnmarshalJSON implements json.Unmarshaler.
func (j *DevfileProjectsElemSource) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["location"]; !ok || v == nil {
		return fmt.Errorf("field location: required")
	}
	if v, ok := raw["type"]; !ok || v == nil {
		return fmt.Errorf("field type: required")
	}
	type Plain DevfileProjectsElemSource
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = DevfileProjectsElemSource(plain)
	return nil
}

const DevfileComponentsElemTypeDockerimage DevfileComponentsElemType = "dockerimage"

type DevfileComponentsElem struct {
	// The name using which other places of this devfile (like commands) can refer to
	// this component. This attribute is optional but must be unique in the devfile if
	// specified.
	Alias *string `json:"alias,omitempty"`

	// Describes whether projects sources should be mount to the component.
	// `CHE_PROJECTS_ROOT` environment variable should contains a path where projects
	// sources are mount
	MountSources bool `json:"mountSources,omitempty"`

	// Describes type of the component, e.g. whether it is an plugin or editor or
	// other type
	Type DevfileComponentsElemType `json:"type"`
}

const DevfileComponentsElemTypeKubernetes DevfileComponentsElemType = "kubernetes"
const DevfileComponentsElemTypeOpenshift DevfileComponentsElemType = "openshift"

type DevfileMetadata struct {
	// Workspaces created from devfile, will use it as base and append random suffix.
	// It's used when name is not defined.
	GenerateName *string `json:"generateName,omitempty"`

	// The name of the devfile. Workspaces created from devfile, will inherit this
	// name
	Name *string `json:"name,omitempty"`
}

type DevfileProjectsElem struct {
	// The path relative to the root of the projects to which this project should be
	// cloned into. This is a unix-style relative path (i.e. uses forward slashes).
	// The path is invalid if it is absolute or tries to escape the project root
	// through the usage of '..'. If not specified, defaults to the project name.
	ClonePath *string `json:"clonePath,omitempty"`

	// Name corresponds to the JSON schema field "name".
	Name string `json:"name"`

	// Describes the project's source - type and location
	Source DevfileProjectsElemSource `json:"source"`
}

// Describes the project's source - type and location
type DevfileProjectsElemSource struct {
	// The name of the of the branch to check out after obtaining the source from the
	// location. The branch has to already exist in the source otherwise the default
	// branch is used. In case of git, this is also the name of the remote branch to
	// push to.
	Branch *string `json:"branch,omitempty"`

	// The id of the commit to reset the checked out branch to. Note that this is
	// equivalent to 'startPoint' and provided for convenience.
	CommitId *string `json:"commitId,omitempty"`

	// Project's source location address. Should be URL for git and github located
	// projects, or file:// for zip.
	Location string `json:"location"`

	// Part of project to populate in the working directory.
	SparseCheckoutDir *string `json:"sparseCheckoutDir,omitempty"`

	// The tag or commit id to reset the checked out branch to.
	StartPoint *string `json:"startPoint,omitempty"`

	// The name of the tag to reset the checked out branch to. Note that this is
	// equivalent to 'startPoint' and provided for convenience.
	Tag *string `json:"tag,omitempty"`

	// Project's source type.
	Type string `json:"type"`
}

type Selector map[string]interface{}

var enumValues_DevfileComponentsElemType = []interface{}{
	"cheEditor",
	"chePlugin",
	"kubernetes",
	"openshift",
	"dockerimage",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Devfile) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if v, ok := raw["apiVersion"]; !ok || v == nil {
		return fmt.Errorf("field apiVersion: required")
	}
	if v, ok := raw["metadata"]; !ok || v == nil {
		return fmt.Errorf("field metadata: required")
	}
	type Plain Devfile
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Devfile(plain)
	return nil
}
