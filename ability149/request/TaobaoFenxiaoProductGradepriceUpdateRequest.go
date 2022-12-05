package request

import (
        "topsdk/util"
    )

type TaobaoFenxiaoProductGradepriceUpdateRequest struct {
    /*
        产品Id     */
    ProductId  *int64 `json:"product_id" required:"true" `
    /*
        skuId，如果产品有skuId,必须要输入skuId;没有skuId的时候不必选     */
    SkuId  *int64 `json:"sku_id,omitempty" required:"false" `
    /*
        交易类型： AGENT(代销)、DEALER(经销)，ALL(代销和经销)     */
    TradeType  *string `json:"trade_type,omitempty" required:"false" `
    /*
        选择折扣方式：GRADE（按等级进行设置）;DISCITUTOR(按分销商进行设置）。例如"GRADE,DISTRIBUTOR"     */
    TargetType  *string `json:"target_type" required:"true" `
    /*
        会员等级的id或者分销商id，例如：”1001,2001,1002”     */
    Ids  *[]int64 `json:"ids" required:"true" `
    /*
        优惠价格,大小为0到100000000之间的整数或两位小数，例：优惠价格为：100元2角5分，传入的参数应写成：100.25     */
    Prices  *[]string `json:"prices" required:"true" `
}

func (s *TaobaoFenxiaoProductGradepriceUpdateRequest) SetProductId(v int64) *TaobaoFenxiaoProductGradepriceUpdateRequest {
    s.ProductId = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceUpdateRequest) SetSkuId(v int64) *TaobaoFenxiaoProductGradepriceUpdateRequest {
    s.SkuId = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceUpdateRequest) SetTradeType(v string) *TaobaoFenxiaoProductGradepriceUpdateRequest {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceUpdateRequest) SetTargetType(v string) *TaobaoFenxiaoProductGradepriceUpdateRequest {
    s.TargetType = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceUpdateRequest) SetIds(v []int64) *TaobaoFenxiaoProductGradepriceUpdateRequest {
    s.Ids = &v
    return s
}
func (s *TaobaoFenxiaoProductGradepriceUpdateRequest) SetPrices(v []string) *TaobaoFenxiaoProductGradepriceUpdateRequest {
    s.Prices = &v
    return s
}

func (req *TaobaoFenxiaoProductGradepriceUpdateRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.ProductId != nil) {
        paramMap["product_id"] = *req.ProductId
    }
    if(req.SkuId != nil) {
        paramMap["sku_id"] = *req.SkuId
    }
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    if(req.TargetType != nil) {
        paramMap["target_type"] = *req.TargetType
    }
    if(req.Ids != nil) {
        paramMap["ids"] = util.ConvertBasicList(*req.Ids)
    }
    if(req.Prices != nil) {
        paramMap["prices"] = util.ConvertBasicList(*req.Prices)
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductGradepriceUpdateRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}