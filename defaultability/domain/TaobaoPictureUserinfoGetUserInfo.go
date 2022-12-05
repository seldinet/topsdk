package domain


type TaobaoPictureUserinfoGetUserInfo struct {
    /*
        用户订购的图片空间容量     */
    OrderSpace  *string `json:"order_space,omitempty" `

    /*
        已使用的图片空间容量     */
    UsedSpace  *string `json:"used_space,omitempty" `

    /*
        剩余的图片空间容量     */
    RemainingSpace  *string `json:"remaining_space,omitempty" `

    /*
        用户的可用容量，即订购量与免费量之和     */
    AvailableSpace  *string `json:"available_space,omitempty" `

    /*
        图片空间的免费容量     */
    FreeSpace  *string `json:"free_space,omitempty" `

    /*
        图片空间的订购有效期     */
    OrderExpiryDate  *string `json:"order_expiry_date,omitempty" `

    /*
        用户自定义的水印参数，通过"|"分割开，如果用户没有定义则为""
具体水印参数组合方法，用"|"分开，顺序按"是否全局设置|水印文字|是否文字水印优先|透明度|字体|字体大小|字体是否加粗|字体是否斜体|字体是否加下划线|字体颜色|旋转角度|是否带阴影|水印位置|图片水印URL|reference水印相对位置" reference取值有左上（1）/中间（3）/右下（2）,其中的null代表为空     */
    WaterMark  *string `json:"water_mark,omitempty" `

}

func (s *TaobaoPictureUserinfoGetUserInfo) SetOrderSpace(v string) *TaobaoPictureUserinfoGetUserInfo {
    s.OrderSpace = &v
    return s
}
func (s *TaobaoPictureUserinfoGetUserInfo) SetUsedSpace(v string) *TaobaoPictureUserinfoGetUserInfo {
    s.UsedSpace = &v
    return s
}
func (s *TaobaoPictureUserinfoGetUserInfo) SetRemainingSpace(v string) *TaobaoPictureUserinfoGetUserInfo {
    s.RemainingSpace = &v
    return s
}
func (s *TaobaoPictureUserinfoGetUserInfo) SetAvailableSpace(v string) *TaobaoPictureUserinfoGetUserInfo {
    s.AvailableSpace = &v
    return s
}
func (s *TaobaoPictureUserinfoGetUserInfo) SetFreeSpace(v string) *TaobaoPictureUserinfoGetUserInfo {
    s.FreeSpace = &v
    return s
}
func (s *TaobaoPictureUserinfoGetUserInfo) SetOrderExpiryDate(v string) *TaobaoPictureUserinfoGetUserInfo {
    s.OrderExpiryDate = &v
    return s
}
func (s *TaobaoPictureUserinfoGetUserInfo) SetWaterMark(v string) *TaobaoPictureUserinfoGetUserInfo {
    s.WaterMark = &v
    return s
}
