package request

import (
	"wenku/model/autocode"
	"wenku/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}