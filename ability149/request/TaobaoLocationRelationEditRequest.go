package request

import (
        "topsdk/ability149/domain"
        "topsdk/util"
    )

type TaobaoLocationRelationEditRequest struct {
    /*
        关系对象列表     */
    LocationRelationList  *[]domain.TaobaoLocationRelationEditLocationRelationDto `json:"location_relation_list" required:"true" `
}

func (s *TaobaoLocationRelationEditRequest) SetLocationRelationList(v []domain.TaobaoLocationRelationEditLocationRelationDto) *TaobaoLocationRelationEditRequest {
    s.LocationRelationList = &v
    return s
}

func (req *TaobaoLocationRelationEditRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.LocationRelationList != nil) {
        paramMap["location_relation_list"] = util.ConvertStructList(*req.LocationRelationList)
    }
    return paramMap
}

func (req *TaobaoLocationRelationEditRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}