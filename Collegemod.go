package models

import (
	"src/config"
	"src/entites"
	//"honnef.co/go/tools/config"
)

type CollegesModel struct {
}

func (*CollegesModel) FindAll() ([]entites.Colleges, error) {

	db, err := config.GetDB()

	if err != nil {
		panic(err.Error())
	} else {
		rows, err2 := db.Query("select * from colleges")

		if err2 != nil {
			panic(err2.Error())
		} else {

			var college []entites.Colleges
			for rows.Next() {

				var colleges entites.Colleges
				rows.Scan(&colleges.COL_ID, colleges.COL_ADD, colleges.COL_TYPE)
				college = append(college, colleges)
			}
			return college, nil
		}

	}

}
func (*CollegesModel) Find(col_id int64) (entites.Colleges, error) {

	db, err := config.GetDB()

	if err != nil {
		panic(err.Error())
	} else {
		rows, err2 := db.Query("select * from colleges where id =?", col_id)

		if err2 != nil {
			panic(err2.Error())
		} else {

			var college entites.Colleges
			for rows.Next() {

				var colleges entites.Colleges
				rows.Scan(&colleges.COL_ID, colleges.COL_ADD, colleges.COL_TYPE)
				college = append(college, colleges)
			}
			return colleges, nil
		}

	}

}
func (*CollegesModel) Create(colleges *entites.Colleges) bool {

	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("insert into college(col_id,col_add,col_type)values(?,?,?)", colleges.COL_ID, colleges.COL_ADD, colleges.COL_TYPE)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
func (*collegesModel) Delete(id int64) bool {
	db, err := config.GetDB()
	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("delete from colleges where id=?", id)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}

	}
}
func (*CollegesModel) update(colleges entites.Colleges) bool {

	db, err := config.GetDB()

	if err != nil {
		return false
	} else {
		result, err2 := db.Exec("update  college set col_add=?,col_type =? where id =?", colleges.COL_ADD, colleges.COL_TYPE, colleges.COL_ID)
		if err2 != nil {
			return false
		} else {
			rowsAffected, _ := result.RowsAffected()
			return rowsAffected > 0
		}
	}
}
