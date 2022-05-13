package models

type Staff struct {
    Model

    Name      string `json:"name"`
    CreatedBy string `json:"created_by"`
    Introduce string `json:"introduce"`
    Avatar    string `json:"avatar"`
}

func GetStaffs(pageNum int, pageSize int, maps interface{}) (Staffs []Staff) {
    db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&Staffs)
    return
}

func GetStaffTotal(maps interface{}) (count int64) {
    db.Model(&Staff{}).Where(maps).Count(&count)
    return
}

func ExistStaffByName(name string) bool {
    var Staff Staff
    db.Select("id").Where("name = ?", name).First(&Staff)
    return Staff.ID > 0
}

func AddStaff(s Staff) bool {
    db.Create(&s)
    return true
}
