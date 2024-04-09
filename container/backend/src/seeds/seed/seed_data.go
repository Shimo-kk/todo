package seed

import (
	"fmt"
	"reflect"
	"todo/app/domain/priority"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ModelInterface interface{}

func SeedData(url string) error {
	db, err := gorm.Open(postgres.Open(fmt.Sprintf("postgres://%s", url)), &gorm.Config{})
	if err != nil {
		return err
	}

	// prioritys
	if err := seedPriority(db); err != nil {
		return err
	}

	return nil
}

func seed(db *gorm.DB, records []ModelInterface) error {
	for _, record := range records {
		result := db.First(record)
		if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
			// その他のエラーが発生した場合
			return result.Error
		}

		// エラーがない場合、もしくはレコードが見つからない場合は挿入を試みる
		if result.Error == gorm.ErrRecordNotFound {
			err := db.Create(record).Error
			if err != nil {
				return err
			}
			fmt.Printf("Inserted: %+v\n", record)
			continue
		}

		existingRecord := reflect.ValueOf(record).Elem()

		// 同じIDのものが存在し、かつ変更されている場合は更新
		newRecord := reflect.ValueOf(record).Elem()
		if existingRecord.FieldByName("Id").Interface() == newRecord.FieldByName("Id").Interface() && !reflect.DeepEqual(existingRecord.Interface(), newRecord.Interface()) {
			err := db.Model(record).Updates(record).Error
			if err != nil {
				return err
			}
			fmt.Printf("Updated: %+v\n", record)
		} else {
			// 同じIDのものが存在していて、かつ変更されていない場合は何もしない
			fmt.Printf("Skipped: %+v\n", record)
		}
	}

	return nil
}

func seedPriority(db *gorm.DB) error {
	testdata := []ModelInterface{
		&priority.Priority{
			Id:   1,
			Name: "低",
		},
		&priority.Priority{
			Id:   2,
			Name: "中",
		},
		&priority.Priority{
			Id:   3,
			Name: "高",
		},
	}
	if err := seed(db, testdata); err != nil {
		return err
	}
	if err := db.Exec("SELECT setval('priorities_id_seq', (SELECT MAX(id) FROM priorities));").Error; err != nil {
		return err
	}
	return nil
}
