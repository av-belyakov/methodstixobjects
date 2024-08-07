package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels/somecomplextypesstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- WindowsRegistryKeyCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &wrkstix); err != nil {
		return nil, err
	}

	return wrkstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (wrkstix WindowsRegistryKeyCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(wrkstix)

	return &result, err
}

// Get возвращает объект "Windows Registry Key Object", по терминалогии STIX, содержит описание значений полей
// раздела реестра Windows.
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) Get() (*WindowsRegistryKeyCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- NumberOfSubkeys property ---------
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) GetNumberOfSubkeys() int {
	return e.NumberOfSubkeys
}

// SetValueNumberOfSubkeys устанавливает значение для поля NumberOfSubkeys
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetValueNumberOfSubkeys(v int) {
	e.NumberOfSubkeys = v
}

// SetAnyNumberOfSubkeys устанавливает ЛЮБОЕ значение для поля NumberOfSubkeys
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetAnyNumberOfSubkeys(i interface{}) {
	e.NumberOfSubkeys = commonlibs.ConversionAnyToInt(i)
}

// -------- Key property ---------
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) GetKey() string {
	return e.Key
}

// SetValueKey устанавливает значение для поля Key
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetValueKey(v string) {
	e.Key = v
}

// SetAnyKey устанавливает ЛЮБОЕ значение для поля Key
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetAnyKey(i interface{}) {
	e.Key = fmt.Sprint(i)
}

// -------- ModifiedTime property ---------
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) GetModifiedTime() string {
	return e.ModifiedTime
}

// SetValueModifiedTime устанавливает значение в формате RFC3339 для поля ModifiedTime
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetValueModifiedTime(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.ModifiedTime = v

	return nil
}

// SetAnyModifiedTime устанавливает ЛЮБОЕ значение для поля ModifiedTime
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetAnyModifiedTime(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueModifiedTime(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.ModifiedTime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- CreatorUserRef property ---------
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) GetCreatorUserRef() stixhelpers.IdentifierTypeSTIX {
	return e.CreatorUserRef
}

func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetValueCreatorUserRef(v stixhelpers.IdentifierTypeSTIX) {
	e.CreatorUserRef = v
}

// -------- Values property ---------
func (e *WindowsRegistryKeyCyberObservableObjectSTIX) GetValues() []somecomplextypesstixco.WindowsRegistryValueTypeSTIX {
	return e.Values
}

func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetValueValues(v somecomplextypesstixco.WindowsRegistryValueTypeSTIX) {
	e.Values = append(e.Values, v)
}

func (e *WindowsRegistryKeyCyberObservableObjectSTIX) SetFullValueValues(v []somecomplextypesstixco.WindowsRegistryValueTypeSTIX) {
	e.Values = v
}

// ValidateStruct является валидатором параметров содержащихся в типе WindowsRegistryKeyCyberObservableObjectSTIX
func (e WindowsRegistryKeyCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(windows-registry-key--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if !e.CreatorUserRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e WindowsRegistryKeyCyberObservableObjectSTIX) SanitizeStruct() WindowsRegistryKeyCyberObservableObjectSTIX {
	e.Key = commonlibs.StringSanitize(e.Key)

	sizev := len(e.Values)
	if sizev > 0 {
		tmp := make([]somecomplextypesstixco.WindowsRegistryValueTypeSTIX, 0, sizev)

		for _, v := range e.Values {
			tmp = append(tmp, v.SanitizeStructWindowsRegistryValueTypeSTIX())
		}

		e.Values = tmp
	}

	return e
}

// GetID возвращает ID STIX объекта
func (e WindowsRegistryKeyCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e WindowsRegistryKeyCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e WindowsRegistryKeyCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'key': '%s'\n", ws, e.Key))
	str.WriteString(fmt.Sprintf("%s'values': \n%v", ws, func(l []somecomplextypesstixco.WindowsRegistryValueTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)
		dubleWs := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'value '%d'':\n", ws, k))
			str.WriteString(fmt.Sprintf("%s'name': '%s'\n", dubleWs, v.Name))
			str.WriteString(fmt.Sprintf("%s'data': '%s'\n", dubleWs, v.Data))
			str.WriteString(fmt.Sprintf("%s'data_type': '%s'\n", dubleWs, v.DataType))
		}

		return str.String()
	}(e.Values, num+1)))
	str.WriteString(fmt.Sprintf("%s'modified_time': '%v'\n", ws, e.ModifiedTime))
	str.WriteString(fmt.Sprintf("%s'creator_user_ref': '%v'\n", ws, e.CreatorUserRef))
	str.WriteString(fmt.Sprintf("%s'number_of_subkeys': '%d'\n", ws, e.NumberOfSubkeys))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e WindowsRegistryKeyCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
