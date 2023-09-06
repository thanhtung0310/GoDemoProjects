package database

import "fmt"

func (db Database) GetListCourse() (error, any) {
	var err error
	err = db.SqlDb.Ping()
	if err != nil {
		return err, ""
	}

	queryStatement := `select c.id, ct.name, c.product_type, c.created, c.modified from course c left join course_translation ct on c.id = ct.course_id`

	data, queryErr := db.SqlDb.Prepare(queryStatement)
	if queryErr != nil {
		return queryErr, ""
	}

	defer data.Close()

	for data.Next() {
		var id, name, productType, created, modified string
		nErr := data.Scan(&id, &name, &productType, &created, &modified)
		if nErr != nil {
			return nErr
		}
		fmt.Printf(id, name, productType, created, modified)
		return nil, id
	}

	return nil, ""
}
