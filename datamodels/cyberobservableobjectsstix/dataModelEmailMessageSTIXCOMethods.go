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

/* --- EmailMessageCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e EmailMessageCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e EmailMessageCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "Email Message", по терминалогии STIX, содержит экземпляр email сообщения
// Обязательные значения в полях IsMultipart
func (e *EmailMessageCyberObservableObjectSTIX) Get() (*EmailMessageCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- IsMultipart property ---------
func (a *EmailMessageCyberObservableObjectSTIX) GetIsMultipart() bool {
	return a.IsMultipart
}

// SetValueIsMultipart устанавливает BOOL значение для поля IsMultipart
func (a *EmailMessageCyberObservableObjectSTIX) SetValueIsMultipart(v bool) {
	a.IsMultipart = v
}

// SetAnyIsMultipart устанавливает ЛЮБОЕ значение для поля IsMultipart
func (a *EmailMessageCyberObservableObjectSTIX) SetAnyIsMultipart(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IsMultipart = v
	}
}

// -------- ContentType property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetContentType() string {
	return e.ContentType
}

// SetValueContentType устанавливает значение для поля ContentType
func (e *EmailMessageCyberObservableObjectSTIX) SetValueContentType(v string) {
	e.ContentType = v
}

// SetAnyContentType устанавливает ЛЮБОЕ значение для поля ContentType
func (e *EmailMessageCyberObservableObjectSTIX) SetAnyContentType(i interface{}) {
	e.ContentType = fmt.Sprint(i)
}

// -------- MessageID property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetMessageID() string {
	return e.MessageID
}

// SetValueMessageID устанавливает значение для поля MessageID
func (e *EmailMessageCyberObservableObjectSTIX) SetValueMessageID(v string) {
	e.MessageID = v
}

// SetAnyMessageID устанавливает ЛЮБОЕ значение для поля MessageID
func (e *EmailMessageCyberObservableObjectSTIX) SetAnyMessageID(i interface{}) {
	e.MessageID = fmt.Sprint(i)
}

// -------- Subject property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetSubject() string {
	return e.Subject
}

// SetValueSubject устанавливает значение для поля Subject
func (e *EmailMessageCyberObservableObjectSTIX) SetValueSubject(v string) {
	e.Subject = v
}

// SetAnySubject устанавливает ЛЮБОЕ значение для поля Subject
func (e *EmailMessageCyberObservableObjectSTIX) SetAnySubject(i interface{}) {
	e.Subject = fmt.Sprint(i)
}

// -------- Body property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetBody() string {
	return e.Body
}

// SetValueBody устанавливает значение для поля Body
func (e *EmailMessageCyberObservableObjectSTIX) SetValueBody(v string) {
	e.Body = v
}

// SetAnyBody устанавливает ЛЮБОЕ значение для поля Body
func (e *EmailMessageCyberObservableObjectSTIX) SetAnyBody(i interface{}) {
	e.Body = fmt.Sprint(i)
}

// -------- Date property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetDate() string {
	return e.Date
}

// SetValueDate устанавливает значение в формате RFC3339 для поля Date
func (e *EmailMessageCyberObservableObjectSTIX) SetValueDate(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.Date = v

	return nil
}

// SetAnyDate устанавливает ЛЮБОЕ значение для поля Date
func (e *EmailMessageCyberObservableObjectSTIX) SetAnyDate(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueDate(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.Date = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- ReceivedLines property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetReceivedLines() []string {
	return e.ReceivedLines
}

// SetValueReceivedLines устанавливает значение для поля ReceivedLines
func (e *EmailMessageCyberObservableObjectSTIX) SetValueReceivedLines(v string) {
	e.ReceivedLines = append(e.ReceivedLines, v)
}

// SetAnyReceivedLines устанавливает ЛЮБОЕ значение для поля ReceivedLines
func (e *EmailMessageCyberObservableObjectSTIX) SetAnyReceivedLines(i interface{}) {
	e.ReceivedLines = append(e.ReceivedLines, fmt.Sprint(i))
}

// -------- FromRef property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetFromRef() stixhelpers.IdentifierTypeSTIX {
	return e.FromRef
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueFromRef(v stixhelpers.IdentifierTypeSTIX) {
	e.FromRef = v
}

// -------- SenderRef property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetSenderRef() stixhelpers.IdentifierTypeSTIX {
	return e.SenderRef
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueSenderRef(v stixhelpers.IdentifierTypeSTIX) {
	e.SenderRef = v
}

// -------- RawEmailRef property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetRawEmailRef() stixhelpers.IdentifierTypeSTIX {
	return e.RawEmailRef
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueRawEmailRef(v stixhelpers.IdentifierTypeSTIX) {
	e.RawEmailRef = v
}

// -------- ToRefs property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetToRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.ToRefs
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueToRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.ToRefs = append(e.ToRefs, v)
}

func (e *EmailMessageCyberObservableObjectSTIX) SetFullValueToRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.ToRefs = v
}

// -------- CcRefs property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetCcRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.CcRefs
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueCcRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.CcRefs = append(e.CcRefs, v)
}

func (e *EmailMessageCyberObservableObjectSTIX) SetFullValueCcRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.CcRefs = v
}

// -------- BccRefs property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetBccRefs() []stixhelpers.IdentifierTypeSTIX {
	return e.BccRefs
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueBccRefs(v stixhelpers.IdentifierTypeSTIX) {
	e.BccRefs = append(e.BccRefs, v)
}

func (e *EmailMessageCyberObservableObjectSTIX) SetFullValueBccRefs(v []stixhelpers.IdentifierTypeSTIX) {
	e.BccRefs = v
}

// -------- BodyMultipart property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetBodyMultipart() []somecomplextypesstixco.EmailMIMEPartTypeSTIX {
	return e.BodyMultipart
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueBodyMultipart(v somecomplextypesstixco.EmailMIMEPartTypeSTIX) {
	e.BodyMultipart = append(e.BodyMultipart, v)
}

func (e *EmailMessageCyberObservableObjectSTIX) SetFullValueBodyMultipart(v []somecomplextypesstixco.EmailMIMEPartTypeSTIX) {
	e.BodyMultipart = v
}

// -------- AdditionalHeaderFields property ---------
func (e *EmailMessageCyberObservableObjectSTIX) GetAdditionalHeaderFields() map[string]stixhelpers.DictionaryTypeSTIX {
	return e.AdditionalHeaderFields
}

func (e *EmailMessageCyberObservableObjectSTIX) SetValueAdditionalHeaderFields(k string, v stixhelpers.DictionaryTypeSTIX) {
	e.AdditionalHeaderFields[k] = v
}

func (e *EmailMessageCyberObservableObjectSTIX) SetFullValueAdditionalHeaderFields(v map[string]stixhelpers.DictionaryTypeSTIX) {
	e.AdditionalHeaderFields = v
}

// ValidateStruct является валидатором параметров содержащихся в типе EmailMessageCyberObservableObjectSTIX
func (e EmailMessageCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(email-message--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.OptionalCommonPropertiesCyberObservableObjectSTIX.ValidateStructCommonFields() {
		return false
	}

	if !e.FromRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if !e.SenderRef.CheckIdentifierTypeSTIX() {
		return false
	}

	if len(e.ToRefs) > 0 {
		for _, v := range e.ToRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(e.CcRefs) > 0 {
		for _, v := range e.CcRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(e.BccRefs) > 0 {
		for _, v := range e.BccRefs {
			if !v.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if len(e.BodyMultipart) > 0 {
		for _, v := range e.BodyMultipart {
			if !v.BodyRawRef.CheckIdentifierTypeSTIX() {
				return false
			}
		}
	}

	if !e.RawEmailRef.CheckIdentifierTypeSTIX() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e EmailMessageCyberObservableObjectSTIX) SanitizeStruct() EmailMessageCyberObservableObjectSTIX {
	e.ContentType = commonlibs.StringSanitize(e.ContentType)
	e.MessageID = commonlibs.StringSanitize(e.MessageID)
	e.Subject = commonlibs.StringSanitize(e.Subject)

	if len(e.ReceivedLines) > 0 {
		tmp := make([]string, 0, len(e.ReceivedLines))
		for _, v := range e.ReceivedLines {
			tmp = append(tmp, commonlibs.StringSanitize(v))
		}

		e.ReceivedLines = tmp
	}

	if len(e.AdditionalHeaderFields) > 0 {
		tmp := make(map[string]stixhelpers.DictionaryTypeSTIX, len(e.AdditionalHeaderFields))
		for k, v := range e.AdditionalHeaderFields {
			switch v := v.Dictionary.(type) {
			case string:
				tmp[k] = stixhelpers.DictionaryTypeSTIX{Dictionary: commonlibs.StringSanitize(string(v))}
			default:
				tmp[k] = stixhelpers.DictionaryTypeSTIX{Dictionary: v}
			}
		}

		e.AdditionalHeaderFields = tmp
	}

	e.Body = commonlibs.StringSanitize(e.Body)

	if len(e.BodyMultipart) > 0 {
		tmp := make([]somecomplextypesstixco.EmailMIMEPartTypeSTIX, 0, len(e.BodyMultipart))
		for _, v := range e.BodyMultipart {
			tmp = append(tmp, somecomplextypesstixco.EmailMIMEPartTypeSTIX{
				Body:               commonlibs.StringSanitize(v.Body),
				BodyRawRef:         v.BodyRawRef,
				ContentType:        commonlibs.StringSanitize(v.ContentType),
				ContentDisposition: commonlibs.StringSanitize(v.ContentDisposition),
			})
		}

		e.BodyMultipart = tmp
	}

	return e
}

// GetID возвращает ID STIX объекта
func (e EmailMessageCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e EmailMessageCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e EmailMessageCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'is_multipart': '%v'\n", ws, e.IsMultipart))
	str.WriteString(fmt.Sprintf("%s'date': '%v'\n", ws, e.Date))
	str.WriteString(fmt.Sprintf("%s'content_type': '%s'\n", ws, e.ContentType))
	str.WriteString(fmt.Sprintf("%s'from_ref': '%v'\n", ws, e.FromRef))
	str.WriteString(fmt.Sprintf("%s'sender_ref': '%v'\n", ws, e.SenderRef))
	str.WriteString(fmt.Sprintf("%s'to_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'to_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.ToRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'cc_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'cc_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.CcRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'bcc_refs': \n%v", ws, func(l []stixhelpers.IdentifierTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'bcc_ref '%d'': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.BccRefs, num+1)))
	str.WriteString(fmt.Sprintf("%s'message_id': '%v'\n", ws, e.MessageID))
	str.WriteString(fmt.Sprintf("%s'subject': '%v'\n", ws, e.Subject))
	str.WriteString(fmt.Sprintf("%s'received_lines': \n%v", ws, func(l []string, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'received_line '%d'': '%s'\n", ws, k, v))
		}

		return str.String()
	}(e.ReceivedLines, num+1)))
	str.WriteString(fmt.Sprintf("%s'additional_header_fields': \n%v", ws, func(l map[string]stixhelpers.DictionaryTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'%s': '%v'\n", ws, k, v))
		}

		return str.String()
	}(e.AdditionalHeaderFields, num+1)))
	str.WriteString(fmt.Sprintf("%s'body': '%v'\n", ws, e.Body))
	str.WriteString(fmt.Sprintf("%s'body_multipart': \n%v", ws, func(l []somecomplextypesstixco.EmailMIMEPartTypeSTIX, num int) string {
		str := strings.Builder{}
		ws := commonlibs.GetWhitespace(num)
		dubleWs := commonlibs.GetWhitespace(num + 1)

		for k, v := range l {
			str.WriteString(fmt.Sprintf("%s'body_multipart '%d'':\n", ws, k))
			str.WriteString(fmt.Sprintf("%s'body': '%s'\n", dubleWs, v.Body))
			str.WriteString(fmt.Sprintf("%s'body_raw_ref': '%s'\n", dubleWs, v.BodyRawRef))
			str.WriteString(fmt.Sprintf("%s'content_type': '%s'\n", dubleWs, v.ContentType))
			str.WriteString(fmt.Sprintf("%s'content_disposition': '%s'\n", dubleWs, v.ContentDisposition))
		}

		return str.String()
	}(e.BodyMultipart, num+1)))
	str.WriteString(fmt.Sprintf("%s'raw_email_ref': '%v'\n", ws, e.RawEmailRef))

	return str.String()
}

// GeneratingDataForIndexing выполняет генерацию данных для их последующей индексации
func (e EmailMessageCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
