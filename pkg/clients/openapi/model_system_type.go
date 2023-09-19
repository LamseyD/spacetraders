/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SystemType The type of waypoint.
type SystemType string

// List of SystemType
const (
	NEUTRON_STAR  SystemType = "NEUTRON_STAR"
	RED_STAR      SystemType = "RED_STAR"
	ORANGE_STAR   SystemType = "ORANGE_STAR"
	BLUE_STAR     SystemType = "BLUE_STAR"
	YOUNG_STAR    SystemType = "YOUNG_STAR"
	WHITE_DWARF   SystemType = "WHITE_DWARF"
	BLACK_HOLE    SystemType = "BLACK_HOLE"
	HYPERGIANT    SystemType = "HYPERGIANT"
	SYSTEM_NEBULA SystemType = "NEBULA"
	UNSTABLE      SystemType = "UNSTABLE"
)

// All allowed values of SystemType enum
var AllowedSystemTypeEnumValues = []SystemType{
	"NEUTRON_STAR",
	"RED_STAR",
	"ORANGE_STAR",
	"BLUE_STAR",
	"YOUNG_STAR",
	"WHITE_DWARF",
	"BLACK_HOLE",
	"HYPERGIANT",
	"NEBULA",
	"UNSTABLE",
}

func (v *SystemType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SystemType(value)
	for _, existing := range AllowedSystemTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SystemType", value)
}

// NewSystemTypeFromValue returns a pointer to a valid SystemType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSystemTypeFromValue(v string) (*SystemType, error) {
	ev := SystemType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SystemType: valid values are %v", v, AllowedSystemTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SystemType) IsValid() bool {
	for _, existing := range AllowedSystemTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SystemType value
func (v SystemType) Ptr() *SystemType {
	return &v
}

type NullableSystemType struct {
	value *SystemType
	isSet bool
}

func (v NullableSystemType) Get() *SystemType {
	return v.value
}

func (v *NullableSystemType) Set(val *SystemType) {
	v.value = val
	v.isSet = true
}

func (v NullableSystemType) IsSet() bool {
	return v.isSet
}

func (v *NullableSystemType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSystemType(val *SystemType) *NullableSystemType {
	return &NullableSystemType{value: val, isSet: true}
}

func (v NullableSystemType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSystemType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
