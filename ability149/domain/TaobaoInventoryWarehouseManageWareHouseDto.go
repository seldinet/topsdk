package domain


type TaobaoInventoryWarehouseManageWareHouseDto struct {
    /*
        详细地址描述 defalutValue:空    */
    Address  *string `json:"address,omitempty" `

    /*
        仓库地址信息,格式 :省~市~区县     */
    AddressAreaName  *string `json:"address_area_name,omitempty" `

    /*
        别名     */
    AliasName  *string `json:"alias_name,omitempty" `

    /*
        联系人 defalutValue:空    */
    Contact  *string `json:"contact,omitempty" `

    /*
        电话号码 defalutValue:空    */
    Phone  *string `json:"phone,omitempty" `

    /*
        邮编 defalutValue:空    */
    PostCode  *string `json:"post_code,omitempty" `

    /*
        仓库编码，仅允许使用字母、数字、下划线，并且在6-30个字符内     */
    StoreCode  *string `json:"store_code,omitempty" `

    /*
        仓库名称     */
    StoreName  *string `json:"store_name,omitempty" `

    /*
        操作类型，新增:ADD;修改:UPDATE     */
    OperateType  *string `json:"operate_type,omitempty" `

}

func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetAddress(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.Address = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetAddressAreaName(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.AddressAreaName = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetAliasName(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.AliasName = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetContact(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.Contact = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetPhone(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.Phone = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetPostCode(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.PostCode = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetStoreCode(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.StoreCode = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetStoreName(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.StoreName = &v
    return s
}
func (s *TaobaoInventoryWarehouseManageWareHouseDto) SetOperateType(v string) *TaobaoInventoryWarehouseManageWareHouseDto {
    s.OperateType = &v
    return s
}
