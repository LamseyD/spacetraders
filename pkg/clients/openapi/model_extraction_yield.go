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

// checks if the ExtractionYield type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtractionYield{}

// ExtractionYield A yield from the extraction operation.
type ExtractionYield struct {
	Symbol TradeSymbol `json:"symbol"`
	// The number of units extracted that were placed into the ship's cargo hold.
	Units int32 `json:"units"`
}

// NewExtractionYield instantiates a new ExtractionYield object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtractionYield(symbol TradeSymbol, units int32) *ExtractionYield {
	this := ExtractionYield{}
	this.Symbol = symbol
	this.Units = units
	return &this
}

// NewExtractionYieldWithDefaults instantiates a new ExtractionYield object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractionYieldWithDefaults() *ExtractionYield {
	this := ExtractionYield{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ExtractionYield) GetSymbol() TradeSymbol {
	if o == nil {
		var ret TradeSymbol
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ExtractionYield) GetSymbolOk() (*TradeSymbol, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ExtractionYield) SetSymbol(v TradeSymbol) {
	o.Symbol = v
}

// GetUnits returns the Units field value
func (o *ExtractionYield) GetUnits() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *ExtractionYield) GetUnitsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Units, true
}

// SetUnits sets field value
func (o *ExtractionYield) SetUnits(v int32) {
	o.Units = v
}

func (o ExtractionYield) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtractionYield) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["units"] = o.Units
	return toSerialize, nil
}

type NullableExtractionYield struct {
	value *ExtractionYield
	isSet bool
}

func (v NullableExtractionYield) Get() *ExtractionYield {
	return v.value
}

func (v *NullableExtractionYield) Set(val *ExtractionYield) {
	v.value = val
	v.isSet = true
}

func (v NullableExtractionYield) IsSet() bool {
	return v.isSet
}

func (v *NullableExtractionYield) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtractionYield(val *ExtractionYield) *NullableExtractionYield {
	return &NullableExtractionYield{value: val, isSet: true}
}

func (v NullableExtractionYield) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtractionYield) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
