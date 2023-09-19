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

// ShipNavStatus The current status of the ship
type ShipNavStatus string

// List of ShipNavStatus
const (
	IN_TRANSIT ShipNavStatus = "IN_TRANSIT"
	IN_ORBIT   ShipNavStatus = "IN_ORBIT"
	DOCKED     ShipNavStatus = "DOCKED"
)

// All allowed values of ShipNavStatus enum
var AllowedShipNavStatusEnumValues = []ShipNavStatus{
	"IN_TRANSIT",
	"IN_ORBIT",
	"DOCKED",
}

func (v *ShipNavStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ShipNavStatus(value)
	for _, existing := range AllowedShipNavStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ShipNavStatus", value)
}

// NewShipNavStatusFromValue returns a pointer to a valid ShipNavStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewShipNavStatusFromValue(v string) (*ShipNavStatus, error) {
	ev := ShipNavStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ShipNavStatus: valid values are %v", v, AllowedShipNavStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ShipNavStatus) IsValid() bool {
	for _, existing := range AllowedShipNavStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ShipNavStatus value
func (v ShipNavStatus) Ptr() *ShipNavStatus {
	return &v
}

type NullableShipNavStatus struct {
	value *ShipNavStatus
	isSet bool
}

func (v NullableShipNavStatus) Get() *ShipNavStatus {
	return v.value
}

func (v *NullableShipNavStatus) Set(val *ShipNavStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableShipNavStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableShipNavStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipNavStatus(val *ShipNavStatus) *NullableShipNavStatus {
	return &NullableShipNavStatus{value: val, isSet: true}
}

func (v NullableShipNavStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipNavStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
