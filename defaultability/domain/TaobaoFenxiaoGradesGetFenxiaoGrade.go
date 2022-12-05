package domain

import (
        "topsdk/util"
    )

type TaobaoFenxiaoGradesGetFenxiaoGrade struct {
    /*
        主键     */
    GradeId  *int64 `json:"grade_id,omitempty" `

    /*
        分销商等级名称     */
    Name  *string `json:"name,omitempty" `

    /*
        记录创建时间     */
    Created  *util.LocalTime `json:"created,omitempty" `

    /*
        记录最后修改时间     */
    Modified  *util.LocalTime `json:"modified,omitempty" `

}

func (s *TaobaoFenxiaoGradesGetFenxiaoGrade) SetGradeId(v int64) *TaobaoFenxiaoGradesGetFenxiaoGrade {
    s.GradeId = &v
    return s
}
func (s *TaobaoFenxiaoGradesGetFenxiaoGrade) SetName(v string) *TaobaoFenxiaoGradesGetFenxiaoGrade {
    s.Name = &v
    return s
}
func (s *TaobaoFenxiaoGradesGetFenxiaoGrade) SetCreated(v util.LocalTime) *TaobaoFenxiaoGradesGetFenxiaoGrade {
    s.Created = &v
    return s
}
func (s *TaobaoFenxiaoGradesGetFenxiaoGrade) SetModified(v util.LocalTime) *TaobaoFenxiaoGradesGetFenxiaoGrade {
    s.Modified = &v
    return s
}
