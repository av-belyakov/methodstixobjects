package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- ProcessCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (pstix ProcessCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	var commonObject CommonProcessCyberObservableObjectSTIX
	if err := json.Unmarshal(*raw, &commonObject); err != nil {
		return nil, err
	}

	pstix = ProcessCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        commonObject.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: commonObject.OptionalCommonPropertiesCyberObservableObjectSTIX,
		IsHidden:             commonObject.IsHidden,
		PID:                  commonObject.PID,
		CreatedTime:          commonObject.CreatedTime,
		Cwd:                  commonObject.Cwd,
		CommandLine:          commonObject.CommandLine,
		EnvironmentVariables: commonObject.EnvironmentVariables,
		OpenedConnectionRefs: commonObject.OpenedConnectionRefs,
		CreatorUserRef:       commonObject.CreatorUserRef,
		ImageRef:             commonObject.ImageRef,
	}

	if len(commonObject.Extensions) == 0 {
		return pstix, nil
	}

	ext := map[string]interface{}{}
	for key, value := range commonObject.Extensions {
		e, err := datamodels.DecodingExtensionsSTIX(key, value)
		if err != nil {
			continue
		}

		ext[key] = e
	}
	pstix.Extensions = ext

	return pstix, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (pstix ProcessCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(pstix)

	return &result, err
}

// Get возвращает объект "Process Object", по терминалогии STIX, содержит общие свойства экземпляра компьютерной
// программы, выполняемой в операционной системе.
func (e *ProcessCyberObservableObjectSTIX) Get() (*ProcessCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- IsHidden property ---------
func (a *ProcessCyberObservableObjectSTIX) GetIsHidden() bool {
	return a.IsHidden
}

// SetValueIsHidden устанавливает BOOL значение для поля IsHidden
func (a *ProcessCyberObservableObjectSTIX) SetValueIsHidden(v bool) {
	a.IsHidden = v
}

// SetAnyIsHidden устанавливает ЛЮБОЕ значение для поля IsHidden
func (a *ProcessCyberObservableObjectSTIX) SetAnyIsHidden(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IsHidden = v
	}
}

// -------- PID property ---------
func (e *ProcessCyberObservableObjectSTIX) GetPID() int {
	return e.PID
}

// SetValuePID устанавливает значение для поля PID
func (e *ProcessCyberObservableObjectSTIX) SetValuePID(v int) {
	e.PID = v
}

// SetAnyPID устанавливает ЛЮБОЕ значение для поля PID
func (e *ProcessCyberObservableObjectSTIX) SetAnyPID(i interface{}) {
	e.PID = commonlibs.ConversionAnyToInt(i)
}

// -------- Cwd property ---------
func (e *ProcessCyberObservableObjectSTIX) GetCwd() string {
	return e.Cwd
}

// SetValueCwd устанавливает значение для поля Cwd
func (e *ProcessCyberObservableObjectSTIX) SetValueCwd(v string) {
	e.Cwd = v
}

// SetAnyCwd устанавливает ЛЮБОЕ значение для поля Cwd
func (e *ProcessCyberObservableObjectSTIX) SetAnyCwd(i interface{}) {
	e.Cwd = fmt.Sprint(i)
}

// -------- CommandLine property ---------
func (e *ProcessCyberObservableObjectSTIX) GetCommandLine() string {
	return e.CommandLine
}

// SetValueCommandLine устанавливает значение для поля CommandLine
func (e *ProcessCyberObservableObjectSTIX) SetValueCommandLine(v string) {
	e.CommandLine = v
}

// SetAnyCommandLine устанавливает ЛЮБОЕ значение для поля CommandLine
func (e *ProcessCyberObservableObjectSTIX) SetAnyCommandLine(i interface{}) {
	e.CommandLine = fmt.Sprint(i)
}

// -------- CreatedTime property ---------
func (e *ProcessCyberObservableObjectSTIX) GetCreatedTime() string {
	return e.CreatedTime
}

// SetValueCreatedTime устанавливает значение в формате RFC3339 для поля CreatedTime
func (e *ProcessCyberObservableObjectSTIX) SetValueCreatedTime(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.CreatedTime = v

	return nil
}

// SetAnyCreatedTime устанавливает ЛЮБОЕ значение для поля CreatedTime
func (e *ProcessCyberObservableObjectSTIX) SetAnyCreatedTime(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueCreatedTime(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.CreatedTime = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- CreatorUserRef property ---------
func (e *ProcessCyberObservableObjectSTIX) GetCreatorUserRef() stixhelpers.IdentifierTypeSTIX {
	return e.CreatorUserRef
}

func (e *ProcessCyberObservableObjectSTIX) SetValueCreatorUserRef(v stixhelpers.IdentifierTypeSTIX) {
	e.CreatorUserRef = v
}

// -------- ImageRef property ---------
func (e *ProcessCyberObservableObjectSTIX) GetImageRef() stixhelpers.IdentifierTypeSTIX {
	return e.ImageRef
}

func (e *ProcessCyberObservableObjectSTIX) SetValueImageRef(v stixhelpers.IdentifierTypeSTIX) {
	e.ImageRef = v
}

// -------- ParentRef property ---------
func (e *ProcessCyberObservableObjectSTIX) GetParentRef() stixhelpers.IdentifierTypeSTIX {
	return e.ParentRef
}

func (e *ProcessCyberObservableObjectSTIX) SetValueParentRef(v stixhelpers.IdentifierTypeSTIX) {
	e.ParentRef = v
}

// -------- ChildRefs property ---------
func (e *ProcessCyberObservableObjectSTIX) GetChildRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ChildRefs
}

func (e *ProcessCyberObservableObjectSTIX) SetValueChildRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.ChildRefs = append(e.ChildRefs, v)
}

func (e *ProcessCyberObservableObjectSTIX) SetFullValueChildRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ChildRefs = v
}

// -------- OpenedConnectionRefs property ---------
func (e *ProcessCyberObservableObjectSTIX) GetOpenedConnectionRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.OpenedConnectionRefs
}

func (e *ProcessCyberObservableObjectSTIX) SetValueOpenedConnectionRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.OpenedConnectionRefs = append(e.OpenedConnectionRefs, v)
}

func (e *ProcessCyberObservableObjectSTIX) SetFullValueOpenedConnectionRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.OpenedConnectionRefs = v
}

// -------- EnvironmentVariables property ---------
func (e *ProcessCyberObservableObjectSTIX) GetEnvironmentVariables() map[string]string {
	return e.EnvironmentVariables
}

// SetValueEnvironmentVariables добаляет значение в EnvironmentVariables
func (e *ProcessCyberObservableObjectSTIX) SetValueEnvironmentVariables(k, v string) {
	e.EnvironmentVariables[k] = v
}

// SetFullValueEnvironmentVariables добаляет 'полное' значение в EnvironmentVariables
func (e *ProcessCyberObservableObjectSTIX) SetFullValueEnvironmentVariables(v map[string]string) {
	e.EnvironmentVariables = v
}

// SetAnyEnvironmentVariables устанавливает ЛЮБОЕ значение для поля EnvironmentVariables
func (e *ProcessCyberObservableObjectSTIX) SetAnyEnvironmentVariables(k string, i interface{}) {
	e.SetValueEnvironmentVariables(k, fmt.Sprint(i))
}

// -------- Extensions property ---------
func (e *ProcessCyberObservableObjectSTIX) GetExtensions() map[string]interface{} {
	return e.Extensions
}

func (e *ProcessCyberObservableObjectSTIX) SetValueExtensions(k string, i interface{}) {
	e.Extensions[k] = i
}

func (e *ProcessCyberObservableObjectSTIX) SetFullValueExtensions(v map[string]interface{}) {
	e.Extensions = v
}

// ValidateStruct является валидатором параметров содержащихся в типе ProcessCyberObservableObjectSTIX
func (e ProcessCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(process--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if len(e.OpenedConnectionRefs) > 0 {
		for _, v := range e.OpenedConnectionRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !e.CreatorUserRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !e.ImageRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !e.ParentRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(e.ChildRefs) > 0 {
		for _, v := range e.ChildRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(e.Extensions) > 0 {
		for _, v := range e.Extensions {
			if !datamodels.CheckingExtensionsSTIX(v) {
				return false
			}
		}
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e ProcessCyberObservableObjectSTIX) SanitizeStruct() ProcessCyberObservableObjectSTIX {
	e.Cwd = commonlibs.StringSanitize(e.Cwd)
	e.CommandLine = commonlibs.StringSanitize(e.CommandLine)

	sizeEnv := len(e.EnvironmentVariables)
	if sizeEnv > 0 {
		//tmp := make(map[string]DictionaryTypeSTIX, sizeEnv)
		tmp := make(map[string]string, sizeEnv)
		for k, v := range e.EnvironmentVariables {
			tmp[k] = commonlibs.StringSanitize(string(v))
		}

		e.EnvironmentVariables = tmp
	}

	esize := len(e.Extensions)
	tmp := make(map[string]interface{}, esize)
	for k, v := range e.Extensions {
		result := datamodels.SanitizeExtensionsSTIX(v)
		tmp[k] = result
	}

	e.Extensions = tmp

	return e
}

// GetID возвращает ID STIX объекта
func (pstix ProcessCyberObservableObjectSTIX) GetID() string {
	return pstix.ID
}

// GetType возвращает Type STIX объекта
func (e ProcessCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e ProcessCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'is_hidden': '%v'\n", ws, e.IsHidden))
	str.WriteString(fmt.Sprintf("%s'pid': '%d'\n", ws, e.PID))
	str.WriteString(fmt.Sprintf("%s'created_time': '%v'\n", ws, e.CreatedTime))
	str.WriteString(fmt.Sprintf("%s'cwd': '%s'\n", ws, e.Cwd))
	str.WriteString(fmt.Sprintf("%s'command_line': '%s'\n", ws, e.CommandLine))
	str.WriteString(fmt.Sprintf("%s'environment_variables': \n%v", ws, func(l map[string]interface{}, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'%s':\n%v\n", ws, k, datamodels.ToStringBeautiful(v)))
		}

		return str.String()
	}(e.Extensions, num+1)))
	str.WriteString(fmt.Sprintf("%s'opened_connection_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'opened_connection_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.OpenedConnectionRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'creator_user_ref': '%v'\n", ws, e.CreatorUserRef))
	str.WriteString(fmt.Sprintf("%s'image_ref': '%v'\n", ws, e.ImageRef))
	str.WriteString(fmt.Sprintf("%s'parent_ref': '%v'\n", ws, e.ParentRef))
	str.WriteString(fmt.Sprintf("%s'child_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'child_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.ChildRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'extensions': \n%v", ws, func(l map[string]interface{}, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'%s':\n%v\n", ws, k, datamodels.ToStringBeautiful(v)))
		}

		return str.String()
	}(e.Extensions, num+1)))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e ProcessCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
