package request

import (
        "topsdk/defaultability/domain"
        "topsdk/util"
    )

type TaobaoFulfillmentOrderAssembleRequest struct {
    /*
        批量回传集合,一次接口最多40     */
    AssembleOrders  *[]domain.TaobaoFulfillmentOrderAssembleAssembleOrder `json:"assemble_orders" required:"true" `
    /*
        操作类型，支持参数为MERGE、CANCEL_MERGE。当进行CANCEL_MERGE操作时，只需要传入groupId即可，order_list可以为空     */
    Type  *string `json:"type" required:"true" `
}

func (s *TaobaoFulfillmentOrderAssembleRequest) SetAssembleOrders(v []domain.TaobaoFulfillmentOrderAssembleAssembleOrder) *TaobaoFulfillmentOrderAssembleRequest {
    s.AssembleOrders = &v
    return s
}
func (s *TaobaoFulfillmentOrderAssembleRequest) SetType(v string) *TaobaoFulfillmentOrderAssembleRequest {
    s.Type = &v
    return s
}

func (req *TaobaoFulfillmentOrderAssembleRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.AssembleOrders != nil) {
        paramMap["assemble_orders"] = util.ConvertStructList(*req.AssembleOrders)
    }
    if(req.Type != nil) {
        paramMap["type"] = *req.Type
    }
    return paramMap
}

func (req *TaobaoFulfillmentOrderAssembleRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}