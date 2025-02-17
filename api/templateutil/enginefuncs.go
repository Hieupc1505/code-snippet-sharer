package templateutil

import (
	"github.com/gofiber/template/html/v2"
)

func RegisterEngineFuncs(engine *html.Engine) {
	//engine.AddFunc("NoEscape", NoEscape)
	engine.AddFunc("GetFromMap", GetFromMap)
	engine.AddFunc("WithComData", WithComData)
	engine.AddFunc("Slice", TmplSlice)
	engine.AddFunc("Map", TmplMap)
	engine.AddFunc("substr", TmplSubstr)
}
