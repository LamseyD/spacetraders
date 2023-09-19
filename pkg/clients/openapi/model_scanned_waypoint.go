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
)

// checks if the ScannedWaypoint type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannedWaypoint{}

// ScannedWaypoint A waypoint that was scanned by a ship.
type ScannedWaypoint struct {
	// Symbol of the waypoint.
	Symbol string       `json:"symbol"`
	Type   WaypointType `json:"type"`
	// Symbol of the system.
	SystemSymbol string `json:"systemSymbol"`
	// Position in the universe in the x axis.
	X int32 `json:"x"`
	// Position in the universe in the y axis.
	Y int32 `json:"y"`
	// List of waypoints that orbit this waypoint.
	Orbitals []WaypointOrbital `json:"orbitals"`
	Faction  *WaypointFaction  `json:"faction,omitempty"`
	// The traits of the waypoint.
	Traits []WaypointTrait `json:"traits"`
	Chart  *Chart          `json:"chart,omitempty"`
}

// NewScannedWaypoint instantiates a new ScannedWaypoint object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannedWaypoint(symbol string, type_ WaypointType, systemSymbol string, x int32, y int32, orbitals []WaypointOrbital, traits []WaypointTrait) *ScannedWaypoint {
	this := ScannedWaypoint{}
	this.Symbol = symbol
	this.Type = type_
	this.SystemSymbol = systemSymbol
	this.X = x
	this.Y = y
	this.Orbitals = orbitals
	this.Traits = traits
	return &this
}

// NewScannedWaypointWithDefaults instantiates a new ScannedWaypoint object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannedWaypointWithDefaults() *ScannedWaypoint {
	this := ScannedWaypoint{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ScannedWaypoint) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ScannedWaypoint) SetSymbol(v string) {
	o.Symbol = v
}

// GetType returns the Type field value
func (o *ScannedWaypoint) GetType() WaypointType {
	if o == nil {
		var ret WaypointType
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetTypeOk() (*WaypointType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScannedWaypoint) SetType(v WaypointType) {
	o.Type = v
}

// GetSystemSymbol returns the SystemSymbol field value
func (o *ScannedWaypoint) GetSystemSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SystemSymbol
}

// GetSystemSymbolOk returns a tuple with the SystemSymbol field value
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetSystemSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SystemSymbol, true
}

// SetSystemSymbol sets field value
func (o *ScannedWaypoint) SetSystemSymbol(v string) {
	o.SystemSymbol = v
}

// GetX returns the X field value
func (o *ScannedWaypoint) GetX() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.X
}

// GetXOk returns a tuple with the X field value
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetXOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.X, true
}

// SetX sets field value
func (o *ScannedWaypoint) SetX(v int32) {
	o.X = v
}

// GetY returns the Y field value
func (o *ScannedWaypoint) GetY() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Y
}

// GetYOk returns a tuple with the Y field value
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetYOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Y, true
}

// SetY sets field value
func (o *ScannedWaypoint) SetY(v int32) {
	o.Y = v
}

// GetOrbitals returns the Orbitals field value
func (o *ScannedWaypoint) GetOrbitals() []WaypointOrbital {
	if o == nil {
		var ret []WaypointOrbital
		return ret
	}

	return o.Orbitals
}

// GetOrbitalsOk returns a tuple with the Orbitals field value
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetOrbitalsOk() ([]WaypointOrbital, bool) {
	if o == nil {
		return nil, false
	}
	return o.Orbitals, true
}

// SetOrbitals sets field value
func (o *ScannedWaypoint) SetOrbitals(v []WaypointOrbital) {
	o.Orbitals = v
}

// GetFaction returns the Faction field value if set, zero value otherwise.
func (o *ScannedWaypoint) GetFaction() WaypointFaction {
	if o == nil || IsNil(o.Faction) {
		var ret WaypointFaction
		return ret
	}
	return *o.Faction
}

// GetFactionOk returns a tuple with the Faction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetFactionOk() (*WaypointFaction, bool) {
	if o == nil || IsNil(o.Faction) {
		return nil, false
	}
	return o.Faction, true
}

// HasFaction returns a boolean if a field has been set.
func (o *ScannedWaypoint) HasFaction() bool {
	if o != nil && !IsNil(o.Faction) {
		return true
	}

	return false
}

// SetFaction gets a reference to the given WaypointFaction and assigns it to the Faction field.
func (o *ScannedWaypoint) SetFaction(v WaypointFaction) {
	o.Faction = &v
}

// GetTraits returns the Traits field value
func (o *ScannedWaypoint) GetTraits() []WaypointTrait {
	if o == nil {
		var ret []WaypointTrait
		return ret
	}

	return o.Traits
}

// GetTraitsOk returns a tuple with the Traits field value
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetTraitsOk() ([]WaypointTrait, bool) {
	if o == nil {
		return nil, false
	}
	return o.Traits, true
}

// SetTraits sets field value
func (o *ScannedWaypoint) SetTraits(v []WaypointTrait) {
	o.Traits = v
}

// GetChart returns the Chart field value if set, zero value otherwise.
func (o *ScannedWaypoint) GetChart() Chart {
	if o == nil || IsNil(o.Chart) {
		var ret Chart
		return ret
	}
	return *o.Chart
}

// GetChartOk returns a tuple with the Chart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScannedWaypoint) GetChartOk() (*Chart, bool) {
	if o == nil || IsNil(o.Chart) {
		return nil, false
	}
	return o.Chart, true
}

// HasChart returns a boolean if a field has been set.
func (o *ScannedWaypoint) HasChart() bool {
	if o != nil && !IsNil(o.Chart) {
		return true
	}

	return false
}

// SetChart gets a reference to the given Chart and assigns it to the Chart field.
func (o *ScannedWaypoint) SetChart(v Chart) {
	o.Chart = &v
}

func (o ScannedWaypoint) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannedWaypoint) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["type"] = o.Type
	toSerialize["systemSymbol"] = o.SystemSymbol
	toSerialize["x"] = o.X
	toSerialize["y"] = o.Y
	toSerialize["orbitals"] = o.Orbitals
	if !IsNil(o.Faction) {
		toSerialize["faction"] = o.Faction
	}
	toSerialize["traits"] = o.Traits
	if !IsNil(o.Chart) {
		toSerialize["chart"] = o.Chart
	}
	return toSerialize, nil
}

type NullableScannedWaypoint struct {
	value *ScannedWaypoint
	isSet bool
}

func (v NullableScannedWaypoint) Get() *ScannedWaypoint {
	return v.value
}

func (v *NullableScannedWaypoint) Set(val *ScannedWaypoint) {
	v.value = val
	v.isSet = true
}

func (v NullableScannedWaypoint) IsSet() bool {
	return v.isSet
}

func (v *NullableScannedWaypoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannedWaypoint(val *ScannedWaypoint) *NullableScannedWaypoint {
	return &NullableScannedWaypoint{value: val, isSet: true}
}

func (v NullableScannedWaypoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannedWaypoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
