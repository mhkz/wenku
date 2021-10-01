package request

import (
	"wenku/model/common/request"
	"wenku/model/system"
)

type SysDictionarySearch struct {
	system.SysDictionary
	request.PageInfo
}
