package models

type Staff struct {
    Model

    Name       string `json:"name"`
    CreatedBy  string `json:"created_by"`
    ModifiedBy string `json:"modified_by"`
    State      int    `json:"state"`
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
    if Staff.ID > 0 {
        return true
    }
    return false
}

func AddStaff(name string, state int, createdBy string) bool {
    db.Create(&Staff{
        Name:      name,
        State:     state,
        CreatedBy: createdBy,
    })
    return true
}
