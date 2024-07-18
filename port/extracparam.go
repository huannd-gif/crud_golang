package port

import (
	"api_crud/app/query"
	"api_crud/domain"
	"errors"
	"strings"
	"time"
)

type TypeResult string

var allResult = [6]string{"INIT", "QUEUEING", "SUCCESS", "FAIL", "NOT_ANSWERED", "CANT_CONNECT"}

type ParamCallRequest struct {
	PhoneNumber          string `form:"phone_number"`
	MetadataDisplayField string `form:"metadata_display_field"`
	PageNum              *int   `form:"page_num"`
	PageSize             *int   `form:"page_size"`
}

type AddCallRequest struct {
	PhoneNumber string                 `json:"phone_number"`
	Result      TypeResult             `json:"result"`
	CallAt      *time.Time             `json:"call_at"`
	EndAt       *time.Time             `json:"end_at"`
	CallPress   *time.Time             `json:"call_press"`
	ReceiverAt  *time.Time             `json:"receiver_at"`
	Metadata    map[string]interface{} `json:"metadata"`
}

type UpdateCallRequest struct {
	PhoneNumber string                 `json:"phone_number"`
	Result      TypeResult             `json:"result"`
	UpdateAt    *time.Time             `json:"update_at"`
	CallAt      *time.Time             `json:"call_at"`
	EndAt       *time.Time             `json:"end_at"`
	CallPress   *time.Time             `json:"call_press"`
	ReceiverAt  *time.Time             `json:"receiver_at"`
	Metadata    map[string]interface{} `json:"metadata"`
}

func parseStrToTypeResult(s string) (TypeResult, error) {
	for _, v := range allResult {
		if v == s {
			return TypeResult(s), nil
		}
	}
	return "", errors.New("Invalid result")
}

func (typeResult *TypeResult) UnmarshalJSON(data []byte) error {
	str := strings.ReplaceAll(string(data), "\"", "")

	result, err := parseStrToTypeResult(str)
	if err != nil {
		return err
	}
	*typeResult = result

	return nil
}

func (typeResult TypeResult) String() string {
	return string(typeResult)
}

func (p *ParamCallRequest) extractToQuery() query.CallRequest {
	if p.PageNum == nil {
		defaultPageNum := 1
		p.PageNum = &defaultPageNum
	}
	if p.PageSize == nil {
		defaultPageSize := 1
		p.PageSize = &defaultPageSize
		//*p.PageSize = 15
	}
	return query.CallRequest{
		PhoneNumber:          p.PhoneNumber,
		MetadataDisplayField: p.MetadataDisplayField,
		PageNum:              *p.PageNum,
		PageSize:             *p.PageSize,
	}
}

func (request *AddCallRequest) extractToModelCreate() domain.Call {
	var call domain.Call = *domain.NewCallNoArgument()

	call.SetPhoneNumber(request.PhoneNumber)
	call.SetResult(request.Result.String())
	call.SetCallAt(request.CallAt)
	call.SetEndAt(request.EndAt)
	call.SetCallPress(request.CallPress)
	call.SetReceiverAt(request.ReceiverAt)
	call.SetMetadata(request.Metadata)

	return call

}

func (request *UpdateCallRequest) extractToModelUpdate(id int) domain.Call {
	var call domain.Call = *domain.NewCallNoArgument()

	call.SetId(id)
	call.SetPhoneNumber(request.PhoneNumber)
	call.SetResult(request.Result.String())
	call.SetCallAt(request.CallAt)
	call.SetEndAt(request.EndAt)
	call.SetCallPress(request.CallPress)
	call.SetReceiverAt(request.ReceiverAt)
	call.SetMetadata(request.Metadata)

	return call

}
