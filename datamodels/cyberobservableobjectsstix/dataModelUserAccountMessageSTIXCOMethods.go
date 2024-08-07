package cyberobservableobjectsstix

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/av-belyakov/methodstixobjects/commonlibs"
	"github.com/av-belyakov/methodstixobjects/datamodels"
	"github.com/av-belyakov/methodstixobjects/datamodels/someextensionsstixco"
	"github.com/av-belyakov/methodstixobjects/datamodels/stixhelpers"
)

/* --- UserAccountCyberObservableObjectSTIX --- */

// DecoderJSON выполняет декодирование JSON объекта
func (e UserAccountCyberObservableObjectSTIX) DecodeJSON(raw *json.RawMessage) (interface{}, error) {
	if err := json.Unmarshal(*raw, &e); err != nil {
		return nil, err
	}

	return e, nil
}

// EncoderJSON выполняет кодирование в JSON объект
func (e UserAccountCyberObservableObjectSTIX) EncodeJSON(interface{}) (*[]byte, error) {
	result, err := json.Marshal(e)

	return &result, err
}

// Get возвращает объект "User Account Object", по терминалогии STIX, содержит экземпляр любого типа учетной записи пользователя
// Обязательные значения в полях Name
func (e *UserAccountCyberObservableObjectSTIX) Get() (*UserAccountCyberObservableObjectSTIX, error) {
	return e, nil
}

// -------- IsServiceAccount property ---------
func (a *UserAccountCyberObservableObjectSTIX) GetIsServiceAccount() bool {
	return a.IsServiceAccount
}

// SetValueIsServiceAccount устанавливает BOOL значение для поля IsServiceAccount
func (a *UserAccountCyberObservableObjectSTIX) SetValueIsServiceAccount(v bool) {
	a.IsServiceAccount = v
}

// SetAnyIsServiceAccount устанавливает ЛЮБОЕ значение для поля IsServiceAccount
func (a *UserAccountCyberObservableObjectSTIX) SetAnyIsServiceAccount(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IsServiceAccount = v
	}
}

// -------- IsPrivileged property ---------
func (a *UserAccountCyberObservableObjectSTIX) GetIsPrivileged() bool {
	return a.IsPrivileged
}

// SetValueIsPrivileged устанавливает BOOL значение для поля IsPrivileged
func (a *UserAccountCyberObservableObjectSTIX) SetValueIsPrivileged(v bool) {
	a.IsPrivileged = v
}

// SetAnyIsPrivileged устанавливает ЛЮБОЕ значение для поля IsPrivileged
func (a *UserAccountCyberObservableObjectSTIX) SetAnyIsPrivileged(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IsPrivileged = v
	}
}

// -------- CanEscalatePrivs property ---------
func (a *UserAccountCyberObservableObjectSTIX) GetCanEscalatePrivs() bool {
	return a.CanEscalatePrivs
}

// SetValueCanEscalatePrivs устанавливает BOOL значение для поля CanEscalatePrivs
func (a *UserAccountCyberObservableObjectSTIX) SetValueCanEscalatePrivs(v bool) {
	a.CanEscalatePrivs = v
}

// SetAnyCanEscalatePrivs устанавливает ЛЮБОЕ значение для поля CanEscalatePrivs
func (a *UserAccountCyberObservableObjectSTIX) SetAnyCanEscalatePrivs(i interface{}) {
	if v, ok := i.(bool); ok {
		a.CanEscalatePrivs = v
	}
}

// -------- IsDisabled property ---------
func (a *UserAccountCyberObservableObjectSTIX) GetIsDisabled() bool {
	return a.IsDisabled
}

// SetValueIsDisabled устанавливает BOOL значение для поля IsDisabled
func (a *UserAccountCyberObservableObjectSTIX) SetValueIsDisabled(v bool) {
	a.IsDisabled = v
}

// SetAnyIsDisabled устанавливает ЛЮБОЕ значение для поля IsDisabled
func (a *UserAccountCyberObservableObjectSTIX) SetAnyIsDisabled(i interface{}) {
	if v, ok := i.(bool); ok {
		a.IsDisabled = v
	}
}

// -------- UserID property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetUserID() string {
	return e.UserID
}

// SetValueUserID устанавливает значение для поля UserID
func (e *UserAccountCyberObservableObjectSTIX) SetValueUserID(v string) {
	e.UserID = v
}

// SetAnyUserID устанавливает ЛЮБОЕ значение для поля UserID
func (e *UserAccountCyberObservableObjectSTIX) SetAnyUserID(i interface{}) {
	e.UserID = fmt.Sprint(i)
}

// -------- Credential property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetCredential() string {
	return e.Credential
}

// SetValueCredential устанавливает значение для поля Credential
func (e *UserAccountCyberObservableObjectSTIX) SetValueCredential(v string) {
	e.Credential = v
}

// SetAnyCredential устанавливает ЛЮБОЕ значение для поля Credential
func (e *UserAccountCyberObservableObjectSTIX) SetAnyCredential(i interface{}) {
	e.Credential = fmt.Sprint(i)
}

// -------- AccountLogin property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetAccountLogin() string {
	return e.AccountLogin
}

// SetValueAccountLogin устанавливает значение для поля AccountLogin
func (e *UserAccountCyberObservableObjectSTIX) SetValueAccountLogin(v string) {
	e.AccountLogin = v
}

// SetAnyAccountLogin устанавливает ЛЮБОЕ значение для поля AccountLogin
func (e *UserAccountCyberObservableObjectSTIX) SetAnyAccountLogin(i interface{}) {
	e.AccountLogin = fmt.Sprint(i)
}

// -------- DisplayName property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetDisplayName() string {
	return e.DisplayName
}

// SetValueDisplayName устанавливает значение для поля DisplayName
func (e *UserAccountCyberObservableObjectSTIX) SetValueDisplayName(v string) {
	e.DisplayName = v
}

// SetAnyDisplayName устанавливает ЛЮБОЕ значение для поля DisplayName
func (e *UserAccountCyberObservableObjectSTIX) SetAnyDisplayName(i interface{}) {
	e.DisplayName = fmt.Sprint(i)
}

// -------- AccountCreated property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetAccountCreated() string {
	return e.AccountCreated
}

// SetValueAccountCreated устанавливает значение в формате RFC3339 для поля AccountCreated
func (e *UserAccountCyberObservableObjectSTIX) SetValueAccountCreated(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.AccountCreated = v

	return nil
}

// SetAnyAccountCreated устанавливает ЛЮБОЕ значение для поля AccountCreated
func (e *UserAccountCyberObservableObjectSTIX) SetAnyAccountCreated(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueAccountCreated(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.AccountCreated = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- AccountExpires property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetAccountExpires() string {
	return e.AccountExpires
}

// SetValueAccountExpires устанавливает значение в формате RFC3339 для поля AccountExpires
func (e *UserAccountCyberObservableObjectSTIX) SetValueAccountExpires(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.AccountExpires = v

	return nil
}

// SetAnyAccountExpires устанавливает ЛЮБОЕ значение для поля AccountExpires
func (e *UserAccountCyberObservableObjectSTIX) SetAnyAccountExpires(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueAccountExpires(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.AccountExpires = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- CredentialLastChanged property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetCredentialLastChanged() string {
	return e.CredentialLastChanged
}

// SetValueCredentialLastChanged устанавливает значение в формате RFC3339 для поля CredentialLastChanged
func (e *UserAccountCyberObservableObjectSTIX) SetValueCredentialLastChanged(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.CredentialLastChanged = v

	return nil
}

// SetAnyCredentialLastChanged устанавливает ЛЮБОЕ значение для поля CredentialLastChanged
func (e *UserAccountCyberObservableObjectSTIX) SetAnyCredentialLastChanged(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueCredentialLastChanged(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.CredentialLastChanged = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- AccountFirstLogin property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetAccountFirstLogin() string {
	return e.AccountFirstLogin
}

// SetValueAccountFirstLogin устанавливает значение в формате RFC3339 для поля AccountFirstLogin
func (e *UserAccountCyberObservableObjectSTIX) SetValueAccountFirstLogin(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.AccountFirstLogin = v

	return nil
}

// SetAnyAccountFirstLogin устанавливает ЛЮБОЕ значение для поля AccountFirstLogin
func (e *UserAccountCyberObservableObjectSTIX) SetAnyAccountFirstLogin(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueAccountFirstLogin(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.AccountFirstLogin = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- AccountLastLogin property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetAccountLastLogin() string {
	return e.AccountLastLogin
}

// SetValueAccountLastLogin устанавливает значение в формате RFC3339 для поля AccountLastLogin
func (e *UserAccountCyberObservableObjectSTIX) SetValueAccountLastLogin(v string) error {
	if _, err := time.Parse(time.RFC3339, v); err != nil {
		return err
	}

	e.AccountLastLogin = v

	return nil
}

// SetAnyAccountLastLogin устанавливает ЛЮБОЕ значение для поля AccountLastLogin
func (e *UserAccountCyberObservableObjectSTIX) SetAnyAccountLastLogin(i interface{}) error {
	if str, ok := i.(string); ok {
		return e.SetValueAccountLastLogin(str)
	}

	tmp := commonlibs.ConversionAnyToInt(i)
	e.AccountLastLogin = commonlibs.GetDateTimeFormatRFC3339(int64(tmp))

	return nil
}

// -------- AccountType property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetAccountType() stixhelpers.OpenVocabTypeSTIX {
	return e.AccountType
}

func (e *UserAccountCyberObservableObjectSTIX) SetValueAccountType(v stixhelpers.OpenVocabTypeSTIX) {
	e.AccountType = v
}

// -------- Extensions property ---------
func (e *UserAccountCyberObservableObjectSTIX) GetExtensions() map[string]someextensionsstixco.UNIXAccountExtensionSTIX {
	return e.Extensions
}

// SetValueExtensions добаляет значение в Extensions
func (e *UserAccountCyberObservableObjectSTIX) SetValueExtensions(k string, v someextensionsstixco.UNIXAccountExtensionSTIX) {
	e.Extensions[k] = v
}

// SetFullValueExtensions добаляет полное значение в Extensions
func (e *UserAccountCyberObservableObjectSTIX) SetFullValueExtensions(v map[string]someextensionsstixco.UNIXAccountExtensionSTIX) {
	e.Extensions = v
}

// ValidateStruct является валидатором параметров содержащихся в типе UserAccountCyberObservableObjectSTIX
func (e UserAccountCyberObservableObjectSTIX) ValidateStruct() bool {
	if !(regexp.MustCompile(`^(user-account--)[0-9a-f|-]+$`).MatchString(e.ID)) {
		return false
	}

	if !e.ValidateStructCommonFields() {
		return false
	}

	return true
}

// SanitizeStruct для ряда полей, выполняет замену некоторых специальных символов на их HTML код
func (e UserAccountCyberObservableObjectSTIX) SanitizeStruct() UserAccountCyberObservableObjectSTIX {
	e.UserID = commonlibs.StringSanitize(e.UserID)
	e.Credential = commonlibs.StringSanitize(e.Credential)
	e.AccountLogin = commonlibs.StringSanitize(e.AccountLogin)
	e.AccountType = e.AccountType.SanitizeStructOpenVocabTypeSTIX()
	e.DisplayName = commonlibs.StringSanitize(e.DisplayName)

	esize := len(e.Extensions)
	tmp := make(map[string]someextensionsstixco.UNIXAccountExtensionSTIX, esize)
	for k, v := range e.Extensions {
		result := datamodels.SanitizeExtensionsSTIX(v)
		if ct, ok := result.(someextensionsstixco.UNIXAccountExtensionSTIX); ok {
			tmp[k] = ct
		}
	}
	e.Extensions = tmp

	return e
}

// GetID возвращает ID STIX объекта
func (e UserAccountCyberObservableObjectSTIX) GetID() string {
	return e.ID
}

// GetType возвращает Type STIX объекта
func (e UserAccountCyberObservableObjectSTIX) GetType() string {
	return e.Type
}

// ToStringBeautiful выполняет красивое представление информации содержащейся в типе
func (e UserAccountCyberObservableObjectSTIX) ToStringBeautiful(num int) string {
	str := strings.Builder{}
	ws := commonlibs.GetWhitespace(num)

	str.WriteString(e.CommonPropertiesObjectSTIX.ToStringBeautiful(num))
	str.WriteString(e.OptionalCommonPropertiesCyberObservableObjectSTIX.ToStringBeautiful(num))
	str.WriteString(fmt.Sprintf("%s'user_id': '%s'\n", ws, e.UserID))
	str.WriteString(fmt.Sprintf("%s'credential': '%s'\n", ws, e.Credential))
	str.WriteString(fmt.Sprintf("%s'account_login': '%s'\n", ws, e.AccountLogin))
	str.WriteString(fmt.Sprintf("%s'account_type': '%v'\n", ws, e.AccountType))
	str.WriteString(fmt.Sprintf("%s'display_name': '%s'\n", ws, e.DisplayName))
	str.WriteString(fmt.Sprintf("%s'is_service_account': '%v'\n", ws, e.IsServiceAccount))
	str.WriteString(fmt.Sprintf("%s'is_privileged': '%v'\n", ws, e.IsPrivileged))
	str.WriteString(fmt.Sprintf("%s'can_escalate_privs': '%v'\n", ws, e.CanEscalatePrivs))
	str.WriteString(fmt.Sprintf("%s'is_disabled': '%v'\n", ws, e.IsDisabled))
	str.WriteString(fmt.Sprintf("%s'account_created': '%v'\n", ws, e.AccountCreated))
	str.WriteString(fmt.Sprintf("%s'account_expires': '%v'\n", ws, e.AccountExpires))
	str.WriteString(fmt.Sprintf("%s'credential_last_changed': '%v'\n", ws, e.CredentialLastChanged))
	str.WriteString(fmt.Sprintf("%s'account_first_login': '%v'\n", ws, e.AccountFirstLogin))
	str.WriteString(fmt.Sprintf("%s'account_last_login': '%v'\n", ws, e.AccountLastLogin))
	str.WriteString(fmt.Sprintf("%s'extensions': \n%v", ws, func(l map[string]someextensionsstixco.UNIXAccountExtensionSTIX, num int) string {
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
func (e UserAccountCyberObservableObjectSTIX) GeneratingDataForIndexing() map[string]string {
	return map[string]string{
		"id":   e.ID,
		"type": e.Type,
	}
}
