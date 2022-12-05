package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TaobaoLocationRelationQueryRequest struct {
    /*
        关系查询     */
    LocationRelation  *domain.TaobaoLocationRelationQueryLocationRelationDto `json:"location_relation" required:"true" `
}

func (s *TaobaoLocationRelationQueryRequest) SetLocationRelation(v domain.TaobaoLocationRelationQueryLocationRelationDto) *TaobaoLocationRelationQueryRequest {
    s.LocationRelation = &v
    return s
}

func (req *TaobaoLocationRelationQueryRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LocationRelation != nil) {
        paramMap["location_relation"] = util.ConvertStruct(*req.LocationRelation)
    }
    return paramMap
}

func (req *TaobaoLocationRelationQueryRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}