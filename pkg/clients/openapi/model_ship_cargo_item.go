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

// checks if the ShipCargoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipCargoItem{}

// ShipCargoItem The type of cargo item and the number of units.
type ShipCargoItem struct {
	// The unique identifier of the cargo item type.
	Symbol string `json:"symbol"`
	// The name of the cargo item type.
	Name string `json:"name"`
	// The description of the cargo item type.
	Description string `json:"description"`
	// The number of units of the cargo item.
	Units int32 `json:"units"`
}

// NewShipCargoItem instantiates a new ShipCargoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipCargoItem(symbol string, name string, description string, units int32) *ShipCargoItem {
	this := ShipCargoItem{}
	this.Symbol = symbol
	this.Name = name
	this.Description = description
	this.Units = units
	return &this
}

// NewShipCargoItemWithDefaults instantiates a new ShipCargoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipCargoItemWithDefaults() *ShipCargoItem {
	this := ShipCargoItem{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ShipCargoItem) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ShipCargoItem) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ShipCargoItem) SetSymbol(v string) {
	o.Symbol = v
}

// GetName returns the Name field value
func (o *ShipCargoItem) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipCargoItem) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipCargoItem) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value
func (o *ShipCargoItem) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ShipCargoItem) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ShipCargoItem) SetDescription(v string) {
	o.Description = v
}

// GetUnits returns the Units field value
func (o *ShipCargoItem) GetUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *ShipCargoItem) GetUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *ShipCargoItem) SetUnits(v int32) {
	o.Units = v
}

func (o ShipCargoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipCargoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["name"] = o.Name
	toSerialize["description"] = o.Description
	toSerialize["units"] = o.Units
	return toSerialize, nil
}

type NullableShipCargoItem struct {
	value *ShipCargoItem
	isSet bool
}

func (v NullableShipCargoItem) Get() *ShipCargoItem {
	return v.value
}

func (v *NullableShipCargoItem) Set(val *ShipCargoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableShipCargoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableShipCargoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipCargoItem(val *ShipCargoItem) *NullableShipCargoItem {
	return &NullableShipCargoItem{value: val, isSet: true}
}

func (v NullableShipCargoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipCargoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
