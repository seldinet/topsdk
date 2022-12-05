package request

import (
        "topsdk/util"
    )

type TaobaoProductsGetRequest struct {
    /*
        需返回的字段列表.可选值:Product数据结构中的所有字段;多个字段之间用","分隔     */
    Fields  *[]string `json:"fields" required:"true" `
    /*
        用户昵称     */
    Nick  *string `json:"nick" required:"true" `
    /*
        页码.传入值为1代表第一页,传入值为2代表第二页,依此类推.默认返回的数据是从第一页开始.     */
    PageNo  *int64 `json:"page_no,omitempty" required:"false" `
    /*
        每页条数.每页返回最多返回100条,默认值为40     */
    PageSize  *int64 `json:"page_size,omitempty" required:"false" `
}

func (s *TaobaoProductsGetRequest) SetFields(v []string) *TaobaoProductsGetRequest {
    s.Fields = &v
    return s
}
func (s *TaobaoProductsGetRequest) SetNick(v string) *TaobaoProductsGetRequest {
    s.Nick = &v
    return s
}
func (s *TaobaoProductsGetRequest) SetPageNo(v int64) *TaobaoProductsGetRequest {
    s.PageNo = &v
    return s
}
func (s *TaobaoProductsGetRequest) SetPageSize(v int64) *TaobaoProductsGetRequest {
    s.PageSize = &v
    return s
}

func (req *TaobaoProductsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.Fields != nil) {
        paramMap["fields"] = util.ConvertBasicList(*req.Fields)
    }
    if(req.Nick != nil) {
        paramMap["nick"] = *req.Nick
    }
    if(req.PageNo != nil) {
        paramMap["page_no"] = *req.PageNo
    }
    if(req.PageSize != nil) {
        paramMap["page_size"] = *req.PageSize
    }
    return paramMap
}

func (req *TaobaoProductsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}