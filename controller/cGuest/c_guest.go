package cGuest

import (
	"fmt"
	"net/http"
	"tugas/controller"
	"tugas/model/mAccounts"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Saldo int
}

type Result2 struct {
	Saldo int
}

func CreateAccount(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_create-account.html`, gin.H{})
		return
	}
	a := mAccounts.Account{}
	err := ctx.ShouldBindJSON(&a)
	res := map[string]interface{}{}
	if err == nil {
		err = a.IsValid()
	}
	if err == nil {
		tx := ctx.Db.Create(&a)
		err = tx.Error
		if err == nil {
			res[`Berhasil di-create`] = a
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func Login(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_login.html`, gin.H{})
		return
	}
	a := mAccounts.Account{}
	err := ctx.ShouldBindJSON(&a)
	res := map[string]interface{}{}
	if err == nil {
		err = a.IsValidLogin()
	}
	if err == nil {
		tx := ctx.Db.Where("id_account= ? AND password = ? ", &a.IdAccount, &a.Password).First(&a)
		err = tx.Error
		if err == nil {
			res[`Berhasil login`] = a
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func Deposit(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_deposit.html`, gin.H{})
		return
	}
	a := mAccounts.Account{}
	t := mAccounts.Transaction{}
	err := ctx.ShouldBindJSON(&t)
	var resl Result
	//var tambah int
	res := map[string]interface{}{}
	if err == nil {
		//err = a.IsValidLogin()
	}
	if err == nil {
		tx := ctx.Db.Create(&t)
		ctx.Db.Table("accounts").Select("saldo").Where("id_account= ?", &t.Sender).Scan(&resl)
		fmt.Println(resl)

		ctx.Db.Model(&a).Where("id_account = ?", &t.Sender).Update("saldo", t.Amount+resl.Saldo)
		err = tx.Error
		if err == nil {
			res[`Berhasil Deposit`] = t
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func Withdraw(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_withdraw.html`, gin.H{})
		return
	}
	a := mAccounts.Account{}
	t := mAccounts.Transaction{}
	err := ctx.ShouldBindJSON(&t)
	var resl Result
	//var tambah int
	res := map[string]interface{}{}
	if err == nil {
		//err = a.IsValidLogin()
	}
	if err == nil {
		tx := ctx.Db.Create(&t)
		ctx.Db.Table("accounts").Select("saldo").Where("id_account= ?", &t.Sender).Scan(&resl)
		fmt.Println(resl)

		ctx.Db.Model(&a).Where("id_account = ?", &t.Sender).Update("saldo", resl.Saldo-t.Amount)
		err = tx.Error
		if err == nil {
			res[`Berhasil Withdraw`] = t
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func Transfer(ctx *controller.Ctx) {
	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_transfer.html`, gin.H{})
		return
	}
	a := mAccounts.Account{}
	t := mAccounts.Transaction{}
	err := ctx.ShouldBindJSON(&t)
	var resl Result
	var resl2 Result2
	//var tambah int
	res := map[string]interface{}{}
	if err == nil {
		//err = a.IsValidLogin()
	}
	if err == nil {
		tx := ctx.Db.Create(&t)

		ctx.Db.Table("accounts").Select("saldo").Where("id_account= ?", &t.Sender).Scan(&resl)
		fmt.Println(resl)

		ctx.Db.Table("accounts").Select("saldo").Where("id_account= ?", &t.Recipient).Scan(&resl2)
		fmt.Println(resl2)

		ctx.Db.Model(&a).Where("id_account = ?", &t.Sender).Update("saldo", resl.Saldo-t.Amount)
		ctx.Db.Model(&a).Where("id_account = ?", &t.Recipient).Update("saldo", resl2.Saldo+t.Amount)

		err = tx.Error
		if err == nil {
			res[`Berhasil Transfer`] = t
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}

func GetAccountDet(ctx *controller.Ctx) {

	if ctx.Context.Request.Method == `GET` {
		ctx.HTML(http.StatusOK, `guest_detail.html`, gin.H{})
		return
	}
	a := mAccounts.Account{}
	t := mAccounts.Transaction{}
	err := ctx.ShouldBindJSON(&a)

	res := map[string]interface{}{}
	result1 := map[string]interface{}{}
	if err == nil {
		//err = a.IsValidLogin()
	}
	if err == nil {
		ctx.Db.Where("sender = ?", &a.IdAccount).Or("recipient= ?", &a.IdAccount).Find(&t).Scan(result1)
		if err == nil {
			res[`Detail Account`] = result1
		}
	}
	if err != nil {
		res[`error`] = err.Error()
	}
	ctx.JSON(http.StatusOK, res)
}
