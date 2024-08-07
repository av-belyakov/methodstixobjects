package domainobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- LocationDomainObjectsSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e LocationDomainObjectsSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e LocationDomainObjectsSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Location", по терминалогии STIX, содержит описание географического местоположения
func (e *LocationDomainObjectsSTIX) Get() (*LocationDomainObjectsSTIX, error) {
	return e, nil
}

// -------- Latitude property ---------
func (e *LocationDomainObjectsSTIX) GetLatitude() float32 {
	return e.Latitude
}

// SetValueLatitude устанавливает значение для поля Latitude
func (e *LocationDomainObjectsSTIX) SetValueLatitude(v float32) {
	e.Latitude = v
}

// -------- Longitude property ---------
func (e *LocationDomainObjectsSTIX) GetLongitude() float32 {
	return e.Longitude
}

// SetValueLongitude устанавливает значение для поля Longitude
func (e *LocationDomainObjectsSTIX) SetValueLongitude(v float32) {
	e.Longitude = v
}

// -------- Precision property ---------
func (e *LocationDomainObjectsSTIX) GetPrecision() float32 {
	return e.Precision
}

// SetValuePrecision устанавливает значение для поля Precision
func (e *LocationDomainObjectsSTIX) SetValuePrecision(v float32) {
	e.Precision = v
}

// -------- Name property ---------
func (e *LocationDomainObjectsSTIX) GetName() string {
	return e.Name
}

// SetValueName устанавливает значение для поля Name
func (e *LocationDomainObjectsSTIX) SetValueName(v string) {
	e.Name = v
}

// SetAnyName устанавливает ЛЮБОЕ значение для поля Name
func (e *LocationDomainObjectsSTIX) SetAnyName(i interface{}) {
	e.Name = fmt.Sprint(i)
}

// -------- Description property ---------
func (e *LocationDomainObjectsSTIX) GetDescription() string {
	return e.Description
}

// SetValueDescription устанавливает значение для поля Description
func (e *LocationDomainObjectsSTIX) SetValueDescription(v string) {
	e.Description = v
}

// SetAnyDescription устанавливает ЛЮБОЕ значение для поля Description
func (e *LocationDomainObjectsSTIX) SetAnyDescription(i interface{}) {
	e.Description = fmt.Sprint(i)
}

// -------- Country property ---------
func (e *LocationDomainObjectsSTIX) GetCountry() string {
	return e.Country
}

// SetValueCountry устанавливает значение для поля Country
func (e *LocationDomainObjectsSTIX) SetValueCountry(v string) {
	e.Country = v
}

// SetAnyCountry устанавливает ЛЮБОЕ значение для поля Country
func (e *LocationDomainObjectsSTIX) SetAnyCountry(i interface{}) {
	e.Country = fmt.Sprint(i)
}

// -------- AdministrativeArea property ---------
func (e *LocationDomainObjectsSTIX) GetAdministrativeArea() string {
	return e.AdministrativeArea
}

// SetValueAdministrativeArea устанавливает значение для поля AdministrativeArea
func (e *LocationDomainObjectsSTIX) SetValueAdministrativeArea(v string) {
	e.AdministrativeArea = v
}

// SetAnyAdministrativeArea устанавливает ЛЮБОЕ значение для поля AdministrativeArea
func (e *LocationDomainObjectsSTIX) SetAnyAdministrativeArea(i interface{}) {
	e.AdministrativeArea = fmt.Sprint(i)
}

// -------- City property ---------
func (e *LocationDomainObjectsSTIX) GetCity() string {
	return e.City
}

// SetValueCity устанавливает значение для поля City
func (e *LocationDomainObjectsSTIX) SetValueCity(v string) {
	e.City = v
}

// SetAnyCity устанавливает ЛЮБОЕ значение для поля City
func (e *LocationDomainObjectsSTIX) SetAnyCity(i interface{}) {
	e.City = fmt.Sprint(i)
}

// -------- StreetAddress property ---------
func (e *LocationDomainObjectsSTIX) GetStreetAddress() string {
	return e.StreetAddress
}

// SetValueStreetAddress устанавливает значение для поля StreetAddress
func (e *LocationDomainObjectsSTIX) SetValueStreetAddress(v string) {
	e.StreetAddress = v
}

// SetAnyStreetAddress устанавливает ЛЮБОЕ значение для поля StreetAddress
func (e *LocationDomainObjectsSTIX) SetAnyStreetAddress(i interface{}) {
	e.StreetAddress = fmt.Sprint(i)
}

// -------- PostalCode property ---------
func (e *LocationDomainObjectsSTIX) GetPostalCode() string {
	return e.PostalCode
}

// SetValuePostalCode устанавливает значение для поля PostalCode
func (e *LocationDomainObjectsSTIX) SetValuePostalCode(v string) {
	e.PostalCode = v
}

// SetAnyPostalCode устанавливает ЛЮБОЕ значение для поля PostalCode
func (e *LocationDomainObjectsSTIX) SetAnyPostalCode(i interface{}) {
	e.PostalCode = fmt.Sprint(i)
}

// -------- Region property ---------
func (e *LocationDomainObjectsSTIX) GetRegion() stixhelpers.OpenVocabTypeSTIX {
	return e.Region
}

func (e *LocationDomainObjectsSTIX) SetValueRegion(v stixhelpers.OpenVocabTypeSTIX) {
	e.Region = v
}

// ValidateStruct является валидатором параметров содержащихся в типе LocationDomainObjectsSTIX
func (e LocationDomainObjectsSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(location--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if (e.Latitude > 90.0) || (e.Latitude < -90.0) {
		return false
	}

	if (e.Longitude > 180.0) || (e.Longitude < -180.0) {
		return false
	}

	if e.Country != "" {
		if !(regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(e.Country)) {
			return false
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e LocationDomainObjectsSTIX) SanitizeStruct() LocationDomainObjectsSTIX {
	e.CommonPropertiesDomainObjectSTIX = e.CommonPropertiesDomainObjectSTIX.SanitizeStruct()

	e.Name = commonlibs.StringSanitize(e.Name)
	e.Description = commonlibs.StringSanitize(e.Description)
	e.Region = stixhelpers.OpenVocabTypeSTIX(commonlibs.StringSanitize(string(e.Region)))
	e.AdministrativeArea = commonlibs.StringSanitize(e.AdministrativeArea)
	e.City = commonlibs.StringSanitize(e.City)
	e.StreetAddress = commonlibs.StringSanitize(e.StreetAddress)
	e.PostalCode = commonlibs.StringSanitize(e.PostalCode)

	return e
}

// GetID возвращает ID STIX объекта
func (e LocationDomainObjectsSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e LocationDomainObjectsSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e LocationDomainObjectsSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.CommonPropertiesDomainObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'name': '%s'\n", ws, e.Name))
	str.WriteString(fmt.Sprintf("%s'description': '%s'\n", ws, e.Description))
	str.WriteString(fmt.Sprintf("%s'latitude': '%v'\n", ws, e.Latitude))
	str.WriteString(fmt.Sprintf("%s'longitude': '%v'\n", ws, e.Longitude))
	str.WriteString(fmt.Sprintf("%s'precision': '%v'\n", ws, e.Precision))
	str.WriteString(fmt.Sprintf("%s'region': '%s'\n", ws, e.Region))
	str.WriteString(fmt.Sprintf("%s'country': '%s'\n", ws, e.Country))
	str.WriteString(fmt.Sprintf("%s'administrative_area': '%s'\n", ws, e.AdministrativeArea))
	str.WriteString(fmt.Sprintf("%s'city': '%s'\n", ws, e.City))
	str.WriteString(fmt.Sprintf("%s'street_address': '%s'\n", ws, e.StreetAddress))
	str.WriteString(fmt.Sprintf("%s'postal_code': '%s'\n", ws, e.PostalCode))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e LocationDomainObjectsSTIX) GeneratingDataForIndexing() map[string]string {
	dataForIndex := map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}

	if e.Name != "" {
		dataForIndex["name"] = e.Name
	}

	if e.Description != "" {
		dataForIndex["description"] = e.Description
	}

	if e.StreetAddress != "" {
		dataForIndex["street_address"] = e.StreetAddress
	}

	return dataForIndex
}
