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

// checks if the Extraction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Extraction{}

// Extraction Extraction details.
type Extraction struct {
	// Symbol of the ship that executed the extraction.
	ShipSymbol string          `json:"shipSymbol"`
	Yield      ExtractionYield `json:"yield"`
}

// NewExtraction instantiates a new Extraction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtraction(shipSymbol string, yield ExtractionYield) *Extraction {
	this := Extraction{}
	this.ShipSymbol = shipSymbol
	this.Yield = yield
	return &this
}

// NewExtractionWithDefaults instantiates a new Extraction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtractionWithDefaults() *Extraction {
	this := Extraction{}
	return &this
}

// GetShipSymbol returns the ShipSymbol field value
func (o *Extraction) GetShipSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipSymbol
}

// GetShipSymbolOk returns a tuple with the ShipSymbol field value
// and a boolean to check if the value has been set.
func (o *Extraction) GetShipSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipSymbol, true
}

// SetShipSymbol sets field value
func (o *Extraction) SetShipSymbol(v string) {
	o.ShipSymbol = v
}

// GetYield returns the Yield field value
func (o *Extraction) GetYield() ExtractionYield {
	if o == nil {
		var ret ExtractionYield
		return ret
	}

	return o.Yield
}

// GetYieldOk returns a tuple with the Yield field value
// and a boolean to check if the value has been set.
func (o *Extraction) GetYieldOk() (*ExtractionYield, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Yield, true
}

// SetYield sets field value
func (o *Extraction) SetYield(v ExtractionYield) {
	o.Yield = v
}

func (o Extraction) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Extraction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["shipSymbol"] = o.ShipSymbol
	toSerialize["yield"] = o.Yield
	return toSerialize, nil
}

type NullableExtraction struct {
	value *Extraction
	isSet bool
}

func (v NullableExtraction) Get() *Extraction {
	return v.value
}

func (v *NullableExtraction) Set(val *Extraction) {
	v.value = val
	v.isSet = true
}

func (v NullableExtraction) IsSet() bool {
	return v.isSet
}

func (v *NullableExtraction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtraction(val *Extraction) *NullableExtraction {
	return &NullableExtraction{value: val, isSet: true}
}

func (v NullableExtraction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtraction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
