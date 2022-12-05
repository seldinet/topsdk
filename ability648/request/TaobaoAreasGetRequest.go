package request

import (
        "topsdk/util"
    )

type TaobaoAreasGetRequest struct {
    /*
        需返回的字段列表.可选值:Area 结构中的所有字段;多个字段之间用","分隔.如:id,type,name,parent_id,zip.     */
    Fields  *[]string `json:"fields" required:"true" `
}

func (s *TaobaoAreasGetRequest) SetFields(v []string) *TaobaoAreasGetRequest {
    s.Fields = &v
    return s
}

func (req *TaobaoAreasGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    return paramMap
}

func (req *TaobaoAreasGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}