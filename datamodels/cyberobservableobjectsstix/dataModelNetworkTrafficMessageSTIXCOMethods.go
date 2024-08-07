package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- NetworkTrafficCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e NetworkTrafficCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	var commonObject CommonNetworkTrafficCyberObservableObjectSTIX
	if err := json.Unmarshal(*raw, &commonObject); err != nil {
		return nil, err
	}

	e = NetworkTrafficCyberObservableObjectSTIX{
		CommonPropertiesObjectSTIX:                        commonObject.CommonPropertiesObjectSTIX,
		OptionalCommonPropertiesCyberObservableObjectSTIX: commonObject.OptionalCommonPropertiesCyberObservableObjectSTIX,
		Start:             commonObject.Start,
		End:               commonObject.End,
		IsActive:          commonObject.IsActive,
		SrcRef:            commonObject.SrcRef,
		DstRef:            commonObject.DstRef,
		SrcPort:           commonObject.SrcPort,
		DstPort:           commonObject.DstPort,
		Protocols:         commonObject.Protocols,
		SrcByteCount:      commonObject.SrcByteCount,
		DstByteCount:      commonObject.DstByteCount,
		SrcPackets:        commonObject.SrcPackets,
		DstPackets:        commonObject.DstPackets,
		SrcPayloadRef:     commonObject.SrcPayloadRef,
		DstPayloadRef:     commonObject.DstPayloadRef,
		EncapsulatesRefs:  commonObject.EncapsulatesRefs,
		EncapsulatedByRef: commonObject.EncapsulatedByRef,
	}

	if len(commonObject.Extensions) > 0 {
		ext := map[string]interface{}{}
		for key, value := range commonObject.Extensions {
			e, err := datamodels.DecodingExtensionsSTIX(key, value)
			if err != nil {
				continue
			}

			ext[key] = e
		}

		e.Extensions = ext
	}

	if len(commonObject.IPFix) > 0 {
		ipfix := map[string]string{}
		for key, value := range commonObject.IPFix {
			switch e := value.(type) {
			case string:
				ipfix[key] = e
			//case int, int16, int32, int64, float32, float64:
			case int:
				ipfix[key] = strconv.Itoa(e)
			case float64:
				ipfix[key] = strconv.Itoa(int(e))
			}
		}

		e.IPFix = ipfix
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e NetworkTrafficCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Network Traffic Object", по терминалогии STIX, содержит объект
// сетевого трафика представляющий собой произвольный сетевой трафик, который исходит из источника и адресуется адресату.
func (e *NetworkTrafficCyberObservableObjectSTIX) Get() (*NetworkTrafficCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- IsActive property ---------
func (a *NetworkTrafficCyberObservableObjectSTIX) GetIsActive() bool {
	return a.IsActive
}

// SetValueIsActive устанавливает BOOL значение для поля IsActive
func (a *NetworkTrafficCyberObservableObjectSTIX) SetValueIsActive(v bool) {
	a.IsActive = v
}

// SetAnyIsActive устанавливает ЛЮБОЕ значение для поля IsActive
func (a *NetworkTrafficCyberObservableObjectSTIX) SetAnyIsActive(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IsActive = v
	}
}

// -------- SrcPort property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetSrcPort() int {
	return e.SrcPort
}

// SetValueSrcPort устанавливает значение для поля SrcPort
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueSrcPort(v int) {
	e.SrcPort = v
}

// SetAnySrcPort устанавливает ЛЮБОЕ значение для поля SrcPort
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnySrcPort(i interface{}) {
	e.SrcPort = commonlibs.ConversionAnyToInt(i)
}

// -------- DstPort property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetDstPort() int {
	return e.DstPort
}

// SetValueDstPort устанавливает значение для поля DstPort
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueDstPort(v int) {
	e.DstPort = v
}

// SetAnyDstPort устанавливает ЛЮБОЕ значение для поля DstPort
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnyDstPort(i interface{}) {
	e.DstPort = commonlibs.ConversionAnyToInt(i)
}

// -------- SrcPackets property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetSrcPackets() int {
	return e.SrcPackets
}

// SetValueSrcPackets устанавливает значение для поля SrcPackets
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueSrcPackets(v int) {
	e.SrcPackets = v
}

// SetAnySrcPackets устанавливает ЛЮБОЕ значение для поля SrcPackets
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnySrcPackets(i interface{}) {
	e.SrcPackets = commonlibs.ConversionAnyToInt(i)
}

// -------- DstPackets property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetDstPackets() int {
	return e.DstPackets
}

// SetValueDstPackets устанавливает значение для поля DstPackets
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueDstPackets(v int) {
	e.DstPackets = v
}

// SetAnyDstPackets устанавливает ЛЮБОЕ значение для поля DstPackets
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnyDstPackets(i interface{}) {
	e.DstPackets = commonlibs.ConversionAnyToInt(i)
}

// -------- SrcByteCount property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetSrcByteCount() uint64 {
	return e.SrcByteCount
}

// SetValueSrcByteCount устанавливает значение для поля SrcByteCount
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueSrcByteCount(v uint64) {
	e.SrcByteCount = v
}

// SetAnySrcByteCount устанавливает ЛЮБОЕ значение для поля SrcByteCount
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnySrcByteCount(i interface{}) {
	e.SrcByteCount = uint64(commonlibs.ConversionAnyToInt(i))
}

// -------- DstByteCount property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetDstByteCount() uint64 {
	return e.DstByteCount
}

// SetValueDstByteCount устанавливает значение для поля DstByteCount
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueDstByteCount(v uint64) {
	e.DstByteCount = v
}

// SetAnyDstByteCount устанавливает ЛЮБОЕ значение для поля DstByteCount
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnyDstByteCount(i interface{}) {
	e.DstByteCount = uint64(commonlibs.ConversionAnyToInt(i))
}

// -------- Start property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetStart() string {
	return e.Start
}

// SetValueStart устанавливает значение в формате RFC3339 для поля Start
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueStart(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Start = v

	return nil
}

// SetAnyStart устанавливает ЛЮБОЕ значение для поля Start
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnyStart(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueStart(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.Start = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- End property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetEnd() string {
	return e.End
}

// SetValueEnd устанавливает значение в формате RFC3339 для поля End
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueEnd(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.End = v

	return nil
}

// SetAnyEnd устанавливает ЛЮБОЕ значение для поля End
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnyEnd(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueEnd(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.End = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- Protocols property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetProtocols() []string {
	return e.Protocols
}

// SetValueProtocols устанавливает значение для поля Protocols
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueProtocols(v string) {
	e.Protocols = append(e.Protocols, v)
}

// SetAnyProtocols устанавливает ЛЮБОЕ значение для поля Protocols
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnyProtocols(i interface{}) {
	e.Protocols = append(e.Protocols, fmt.Sprint(i))
}

// -------- SrcRef property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetSrcRef() stixhelpers.IdentifierTypeSTIX {
	return e.SrcRef
}

func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueSrcRef(v stixhelpers.IdentifierTypeSTIX) {
	e.SrcRef = v
}

// -------- DstRef property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetDstRef() stixhelpers.IdentifierTypeSTIX {
	return e.DstRef
}

func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueDstRef(v stixhelpers.IdentifierTypeSTIX) {
	e.DstRef = v
}

// -------- SrcPayloadRef property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetSrcPayloadRef() stixhelpers.IdentifierTypeSTIX {
	return e.SrcPayloadRef
}

func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueSrcPayloadRef(v stixhelpers.IdentifierTypeSTIX) {
	e.SrcPayloadRef = v
}

// -------- DstPayloadRef property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetDstPayloadRef() stixhelpers.IdentifierTypeSTIX {
	return e.DstPayloadRef
}

func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueDstPayloadRef(v stixhelpers.IdentifierTypeSTIX) {
	e.DstPayloadRef = v
}

// -------- EncapsulatedByRef property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetEncapsulatedByRef() stixhelpers.IdentifierTypeSTIX {
	return e.EncapsulatedByRef
}

func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueEncapsulatedByRef(v stixhelpers.IdentifierTypeSTIX) {
	e.EncapsulatedByRef = v
}

// -------- EncapsulatesRefs property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetEncapsulatesRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.EncapsulatesRefs
}

func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueEncapsulatesRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.EncapsulatesRefs = append(e.EncapsulatesRefs, v)
}

func (e *NetworkTrafficCyberObservableObjectSTIX) SetFullValueEncapsulatesRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.EncapsulatesRefs = v
}

// -------- IPFix property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetIPFix() map[string]string {
	return e.IPFix
}

// SetValueIPFix добаляет значение в IPFix
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueIPFix(k, v string) {
	e.IPFix[k] = v
}

// SetValueIPFix добаляет значение в IPFix
func (e *NetworkTrafficCyberObservableObjectSTIX) SetFullValueIPFix(v map[string]string) {
	e.IPFix = v
}

// SetAnyIPFix устанавливает ЛЮБОЕ значение для поля IPFix
func (e *NetworkTrafficCyberObservableObjectSTIX) SetAnyIPFix(k string, i interface{}) {
	e.SetValueIPFix(k, fmt.Sprint(i))
}

// -------- Extensions property ---------
func (e *NetworkTrafficCyberObservableObjectSTIX) GetExtensions() map[string]interface{} {
	return e.Extensions
}

// SetAnyExtensions устанавливает ЛЮБОЕ значение для поля Extensions
func (e *NetworkTrafficCyberObservableObjectSTIX) SetValueExtensions(k string, i interface{}) {
	e.Extensions[k] = i
}

// SetFullValueExtensions устанавливает полное значение для поля Extensions
func (e *NetworkTrafficCyberObservableObjectSTIX) SetFullValueExtensions(v map[string]interface{}) {
	e.Extensions = v
}

// ValidateStruct является валидатором параметров содержащихся в типе NetworkTrafficCyberObservableObjectSTIX
func (e NetworkTrafficCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(network-traffic--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	if len(e.Protocols) == 0 {
		return false
	}

	if !e.SrcRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !e.DstRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !e.SrcPayloadRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !e.DstPayloadRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(e.EncapsulatesRefs) > 0 {
		for _, v := range e.EncapsulatesRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !e.EncapsulatedByRef.CheckIdentifierTypeSTIX() {
		return false
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
func (e NetworkTrafficCyberObservableObjectSTIX) SanitizeStruct() NetworkTrafficCyberObservableObjectSTIX {
	esize := len(e.Extensions)
	if esize > 0 {
		tmp := make(map[string]interface{}, esize)
		for k, v := range e.Extensions {
			result := datamodels.SanitizeExtensionsSTIX(v)
			tmp[k] = result
		}

		e.Extensions = tmp
	}

	if len(e.IPFix) > 0 {
		ipfix := map[string]string{}

		for key, value := range e.IPFix {
			switch i := datamodels.SanitizeExtensionsSTIX(value).(type) {
			case string:
				ipfix[key] = i
			case int:
				ipfix[key] = strconv.Itoa(i)
			case float64:
				ipfix[key] = strconv.Itoa(int(i))
			}
		}

		e.IPFix = ipfix
	}

	sizeProtocols := len(e.Protocols)
	if sizeProtocols == 0 {
		tmp := make([]string, 0, sizeProtocols)
		for _, v := range e.Protocols {
			tmp = append(tmp, commonlibs.StringSanitize(v))
		}

		e.Protocols = tmp
	}

	return e
}

// GetID возвращает ID STIX объекта
func (e NetworkTrafficCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e NetworkTrafficCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e NetworkTrafficCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'start': '%v'\n", ws, e.Start))
	str.WriteString(fmt.Sprintf("%s'end': '%v'\n", ws, e.End))
	str.WriteString(fmt.Sprintf("%s'is_active': '%v'\n", ws, e.IsActive))
	str.WriteString(fmt.Sprintf("%s'src_ref': '%v'\n", ws, e.SrcRef))
	str.WriteString(fmt.Sprintf("%s'dst_ref': '%v'\n", ws, e.DstRef))
	str.WriteString(fmt.Sprintf("%s'src_port': '%d'\n", ws, e.SrcPort))
	str.WriteString(fmt.Sprintf("%s'dst_port': '%d'\n", ws, e.DstPort))
	str.WriteString(fmt.Sprintf("%s'protocols': \n%v", ws, func(l []string, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'protocol '%d'': '%s'\n", ws, k, v))
		}

		return str.String()
	}(e.Protocols, num+1)))
	str.WriteString(fmt.Sprintf("%s'src_byte_count': '%d'\n", ws, e.SrcByteCount))
	str.WriteString(fmt.Sprintf("%s'dst_byte_count': '%d'\n", ws, e.DstByteCount))
	str.WriteString(fmt.Sprintf("%s'src_packets': '%d'\n", ws, e.SrcPackets))
	str.WriteString(fmt.Sprintf("%s'dst_packets': '%d'\n", ws, e.DstPackets))
	str.WriteString(fmt.Sprintf("%s'ipfix': \n%v", ws, func(l map[string]string, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'%s': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.IPFix, num+1)))
	str.WriteString(fmt.Sprintf("%s'src_payload_ref': '%v'\n", ws, e.SrcPayloadRef))
	str.WriteString(fmt.Sprintf("%s'dst_payload_ref': '%v'\n", ws, e.DstPayloadRef))
	str.WriteString(fmt.Sprintf("%s'encapsulates_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'encapsulates_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.EncapsulatesRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'encapsulated_by_ref': '%v'\n", ws, e.EncapsulatedByRef))
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
func (e NetworkTrafficCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
