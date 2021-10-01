package request

import (
	"wenku/model/common/request"
	"wenku/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
