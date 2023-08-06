package main

import (
	"database/sql"
	"reflect"
	"strings"
)

type Category struct {
	Id   int
	Name string
}

type Product struct {
	Id   int
	Name string
	Category
	Price float64
}

var insertNewCategory *sql.Stmt
var changeProductCategory *sql.Stmt

func queryDatabase(db *sql.DB, categoryName string) []Product {
	products := []Product{}
	rows, err := db.Query(`
      SELECT Products.Id, Products.Name, Products.Price, Categories.Id as Cat_Id,
      Categories.Name as CatName FROM Products, Categories
      WHERE Products.Category = Categories.Id
      AND Categories.Name = ?`, categoryName)
	if err == nil {
		for rows.Next() {
			// var id, category int
			// var name string
			// var price float64
			// scanErr := rows.Scan(&id, &name, &category, &price)
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
			if scanErr == nil {
				// Printfln("Row: %v %v %v %v", id, name, category, price)
				products = append(products, p)
			} else {
				Printfln("Scan Error: %v", scanErr.Error())
				break
			}
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
	return products
}

func querySingleRow(db *sql.DB, id int) (p Product) {
	row := db.QueryRow(`
      SELECT Products.Id, Products.Name, Products.Price, Categories.Id as Cat_Id,
      Categories.Name as CatName FROM Products, Categories
      WHERE Products.Category = Categories.Id AND Products.Id = ?`, id)
	if row.Err() == nil {
		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
		if scanErr != nil {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row Error: %v", row.Err().Error())
	}
	return
}

func insertRow(db *sql.DB, p *Product) (id int64) {
	res, err := db.Exec(`
    INSERT INTO Products (Name, Category, Price)
    VALUES (?, ?, ?)`, p.Name, p.Category.Id, p.Price)
	if err == nil {
		id, err = res.LastInsertId()
		if err != nil {
			Printfln("Exec error: %v", err.Error())
		}
	} else {
		Printfln("Exec error: %v", err.Error())
	}
	return
}

// func insertAndUseCategory(name string, productIDs ...int) {
// 	result, err := insertNewCategory.Exec(name)
// 	if err == nil {
// 		newID, _ := result.LastInsertId()
// 		for _, id := range productIDs {
// 			changeProductCategory.Exec(int(newID), id)
// 		}
// 	} else {
// 		Printfln("Prepared statement error: %v", err.Error())
// 	}
// }

func queryDatabaseWithReflection(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query(`
    SELECT Products.Id, Products.Name, Products.Price, Categories.Id as "Category.Id", Categories.Name as "Category.Name"
    FROM Products, Categories
    WHERE Products.Category = Categories.Id
    `)
	if err == nil {
		results, err := scanIntoStruct(rows, &Product{})
		if err == nil {
			products = (results).([]Product)
		} else {
			Printfln("Scanning error: %v", err.Error())
		}
	} else {
		return
	}
	return
}

func insertAndUseCategory(db *sql.DB, name string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	if err == nil {
		catResult, err := tx.Stmt(insertNewCategory).Exec(name)
		if err == nil {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(changeProductCategory)
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if err == nil {
					changes, _ := changeResult.RowsAffected()
					if changes == 0 {
						updatedFailed = true
						break
					}
				}
			}
		}
	}
	if err != nil || updatedFailed {
		Printfln("Aborting transaction %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func scanIntoStruct(rows *sql.Rows, target interface{}) (results interface{}, err error) {
	targetVal := reflect.ValueOf(target)
	if targetVal.Kind() == reflect.Ptr {
		targetVal = targetVal.Elem()
	}
	if targetVal.Kind() != reflect.Struct {
		return
	}
	colNames, _ := rows.Columns()
	colTypes, _ := rows.ColumnTypes()
	references := []interface{}{}
	fieldVal := reflect.Value{}
	// var placeholder interface{}

	for i, colName := range colNames {
		colNameParts := strings.Split(colName, ".")
		fieldVal = targetVal.FieldByName(colNameParts[0])
		if fieldVal.IsValid() && fieldVal.Kind() == reflect.Struct && len(colNameParts) > 1 {
			var namePart string
			for _, namePart = range colNameParts[1:] {
				compFunction := matchColName(namePart)
				fieldVal = fieldVal.FieldByNameFunc(compFunction)
			}
		}

		if !fieldVal.IsValid() || !colTypes[i].ScanType().ConvertibleTo(fieldVal.Type()) {
			references = append(references, fieldVal.Interface())
		} else if fieldVal.Kind() != reflect.Ptr && fieldVal.CanAddr() {
			fieldVal = fieldVal.Addr()
			references = append(references, fieldVal.Interface())
		}
	}

	resultSlice := reflect.MakeSlice(reflect.SliceOf(targetVal.Type()), 0, 10)
	for rows.Next() {
		err := rows.Scan(references...)
		if err == nil {
			resultSlice = reflect.Append(resultSlice, targetVal)
		} else {
			break
		}
	}
	results = resultSlice.Interface()
	return
}

func matchColName(colName string) func(string) bool {
	return func(fieldName string) bool {
		return strings.EqualFold(colName, fieldName)
	}
}

func main() {
	// listDrivers()
	db, err := openDatabase()
	if err == nil {
		// products := queryDatabase(db)
		// for i, p := range products {
		// 	Printfln("#%v: %v", i, p)
		// }
		// for _, cat := range []string{"Soccer", "Watersports"} {
		// 	Printfln("--- %v Results ---", cat)
		// 	products := queryDatabase(db, cat)
		// 	for i, p := range products {
		// 		Printfln("#%v: %v %v %v", i, p.Name, p.Category.Name, p.Price)
		// 	}
		// }

		// for _, id := range []int{1, 3, 10} {
		// 	p := querySingleRow(db, id)
		// 	Printfln("Product: %v", p)
		// }

		// newProduct := Product{Name: "Stadium", Category: Category{Id: 2}, Price: 79500}
		// newID := insertRow(db, &newProduct)
		// p := querySingleRow(db, int(newID))
		// Printfln("New Product: %v", p)

		// insertNewCategory, _ = db.Prepare("INSERT INTO Categories (Name) Values (?)")
		// changeProductCategory, _ = db.Prepare("UPDATE Products SET Category = ? WHERE Id = ?")
		// insertAndUseCategory("Misc Products", 2)
		// insertAndUseCategory(db, "Category_1", 2)
		// p := querySingleRow(db, 2)
		// Printfln("Product: %v", p)
		// insertAndUseCategory(db, "Category_2", 100)

		products, _ := queryDatabaseWithReflection(db)
		for _, p := range products {
			Printfln("Product: %v", p)
		}
		db.Close()
	} else {
		panic(err)
	}
}
