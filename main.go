package main

import (
	"tugas/controller"
	"tugas/controller/cGuest"
)

func main() {
	server := controller.InitServer()

	server.AssignHandler(`/guest/create-account`, cGuest.CreateAccount)
	server.AssignHandler(`/guest/login`, cGuest.Login)
	server.AssignHandler(`/guest/deposit`, cGuest.Deposit)
	server.AssignHandler(`/guest/withdraw`, cGuest.Withdraw)
	server.AssignHandler(`/guest/transfer`, cGuest.Transfer)
	server.AssignHandler(`/guest/detailac`, cGuest.GetAccountDet)

	server.Listen(":8084")
}
