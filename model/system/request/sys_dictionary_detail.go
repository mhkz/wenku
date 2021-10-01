package request

import (
	"wenku/model/common/request"
	"wenku/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
