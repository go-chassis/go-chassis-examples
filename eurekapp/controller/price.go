package controller

import (
	"net/http"

	goRestful "github.com/emicklei/go-restful"
	rf "github.com/go-chassis/go-chassis/v2/server/restful"
)

type Data struct {
	ID         string    `json:"priceID"`
	Category   string    `json:"type"`
	Value      string    `json:"value"`
	CreateTime string    `json:"-"`
	err        ErrorCode `json:"-"`
}

type ErrorCode struct {
	code uint32
}

type Result struct {
	Code    ErrorCode
	Message string
}

func (d *Data) GetPrice(c *rf.Context) {
	var xx ErrorCode
	xx.code = 777
	x := &Data{
		c.ReadPathParameter("id"),
		"internal",
		"104.5",
		"2019-09-09",
		xx,
	}
	c.WriteHeaderAndJSON(http.StatusOK, x, goRestful.MIME_JSON)
}

func (d *Data) URLPatterns() []rf.Route {

	return []rf.Route{
		{
			Method:           http.MethodGet,
			Path:             "/price/{id}",
			ResourceFuncName: "GetPrice",
			Consumes:         []string{goRestful.MIME_JSON, goRestful.MIME_XML},
			Produces:         []string{goRestful.MIME_JSON},
			Returns:          []*rf.Returns{{Code: http.StatusOK, Message: "true", Model: Data{}}},
			Parameters: []*rf.Parameters{
				{Name: "x-auth-token", DataType: "string", ParamType: goRestful.HeaderParameterKind, Desc: "this is a token", Required: true},
				{Name: "x-auth-token2", DataType: "string", ParamType: goRestful.HeaderParameterKind, Desc: "this is a token", Required: true},
			},
		},
	}
}
