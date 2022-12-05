package request


type TaobaoFenxiaoProductImportFromAuctionRequest struct {
    /*
        导入产品需要支持的交易类型：[1 代销][ 2 经销 ][3 代销和经销]     */
    TradeType  *int64 `json:"trade_type" required:"true" `
    /*
        店铺宝贝id     */
    AuctionId  *int64 `json:"auction_id" required:"true" `
    /*
        产品线id     */
    ProductLineId  *int64 `json:"product_line_id" required:"true" `
}

func (s *TaobaoFenxiaoProductImportFromAuctionRequest) SetTradeType(v int64) *TaobaoFenxiaoProductImportFromAuctionRequest {
    s.TradeType = &v
    return s
}
func (s *TaobaoFenxiaoProductImportFromAuctionRequest) SetAuctionId(v int64) *TaobaoFenxiaoProductImportFromAuctionRequest {
    s.AuctionId = &v
    return s
}
func (s *TaobaoFenxiaoProductImportFromAuctionRequest) SetProductLineId(v int64) *TaobaoFenxiaoProductImportFromAuctionRequest {
    s.ProductLineId = &v
    return s
}

func (req *TaobaoFenxiaoProductImportFromAuctionRequest) ToMap() map[string]interface{} {
    paramMap := make(map[string]interface{})
    if(req.TradeType != nil) {
        paramMap["trade_type"] = *req.TradeType
    }
    if(req.AuctionId != nil) {
        paramMap["auction_id"] = *req.AuctionId
    }
    if(req.ProductLineId != nil) {
        paramMap["product_line_id"] = *req.ProductLineId
    }
    return paramMap
}

func (req *TaobaoFenxiaoProductImportFromAuctionRequest) ToFileMap() map[string]interface{} {
    fileMap := make(map[string]interface{})
    return fileMap
}