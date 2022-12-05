package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TmallChannelProductsGetRequest struct {
    /*
        top_query_product_d_o     */
    TopQueryProductDO  *domain.TmallChannelProductsGetTopQueryProductDo `json:"top_query_product_d_o,omitempty" required:"false" `
}

func (s *TmallChannelProductsGetRequest) SetTopQueryProductDO(v domain.TmallChannelProductsGetTopQueryProductDo) *TmallChannelProductsGetRequest {
    s.TopQueryProductDO = &v
    return s
}

func (req *TmallChannelProductsGetRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TopQueryProductDO != nil) {
        paramMap["top_query_product_d_o"] = util.ConvertStruct(*req.TopQueryProductDO)
    }
    return paramMap
}

func (req *TmallChannelProductsGetRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}