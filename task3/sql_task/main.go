package main

import (
	"errors"
	"fmt"
	"metonode-golang/task3/sql_task/database"
	"metonode-golang/task3/sql_task/models"
	"metonode-golang/task3/sql_task/repository"
	"sort"

	"github.com/jinzhu/gorm"
)

func crudStudent() {
	database.InitDB()
	defer database.CloseDB()
	re := repository.NewStudentRepository(database.DB)
	student := models.Students{Id: 1, Name: "张三", Age: 20, Grade: "三年级"}
	if err := re.Insert(&student); err != nil {
		fmt.Println(err)
	}
	var stu []models.Students
	if err := database.DB.Raw("select * from students where age > 18").Scan(&stu); err != nil {
		fmt.Println(err)
	}
	updateStudent := models.Students{Id: 1, Name: "张三", Age: 20, Grade: "四年级"}
	if err := re.Update(&updateStudent); err != nil {
		fmt.Println(err)
	}

	if err := database.DB.Unscoped().Where("age < ?", 15).Delete(&models.Students{}); err != nil {
		fmt.Println(err)
	}
}

type accounts struct {
	id      uint
	balance float64
}

type transactions struct {
	fromAccountId uint
	toAccountId   uint
	amount        float64
}

func transfer(aId, bId uint, amount float64) {
	database.InitDB()
	defer database.CloseDB()

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		var accA accounts

		//  1.检查A账户是否存在
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id = ?", aId).First(&accA).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				return errors.New("该A账户不存在，无法转账")
			}
			return err
		}
		// 2.检查A账户余额是否足够
		if accA.balance < amount {
			return errors.New("余额不足")
		}

		// 更新A账户余额
		accA.balance -= amount
		if er := tx.Save(&accA).Error; er != nil {
			return er
		}

		result := tx.Model(&accounts{}).Where("id = ?", bId).
			Update("balance", gorm.Expr("balance + ?", amount))
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected == 0 {
			return errors.New("该B账户不存在")
		}

		// 转账记录
		transaction := transactions{fromAccountId: aId, toAccountId: bId, amount: amount}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

func transfer2(fromId, toId uint, amount float64) {
	database.InitDB()
	defer database.CloseDB()

	if err := database.DB.Transaction(func(tx *gorm.DB) error {
		var account []accounts
		ids := []uint{fromId, toId}
		sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
		if err := tx.Set("gorm:query_option", "FOR UPDATE").
			Where("id IN (?)", ids).Find(&account).Error; err != nil {
			return err
		}
		if len(account) == 0 {
			return errors.New("两个账户不存在")
		}
		if len(account) == 1 {
			if account[0].id == fromId {
				return errors.New("目标账户不存在")
			} else if account[0].id == toId {
				return errors.New("转账账户不存在")
			}
		}

		// 找到对应的账户
		var fromAcc, toAcc *accounts
		for _, acc := range account {
			if acc.id == fromId {
				fromAcc = &acc
			} else if acc.id == toId {
				toAcc = &acc
			}
		}
		// 检查转账余额
		if fromAcc.balance < amount {
			return errors.New("转账余额不足")
		}

		// 修改金额
		fromAcc.balance -= amount
		toAcc.balance += amount
		if err := tx.Save(fromAcc).Error; err != nil {
			return err
		}
		if err := tx.Save(toAcc).Error; err != nil {
			return err
		}
		// 记录交易记录
		transaction := transactions{fromAccountId: fromAcc.id, toAccountId: toAcc.id}
		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		fmt.Println(err)
	}
}

func main() {

}
