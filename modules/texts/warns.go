package texts

import (
	"fmt"
)

const ICanNotGiveWarnToAdministrator = "من نمی توانم به مدیر اخطار بدهم."

const WritePleaseInteger = "لطفاً یک عدد وارد کنید."

const AdministratorAlwaysIsClean = "مدیران همیشه فاقد وارن اند."

const RemoveWarn = "حذف اخطار"

const WarnWasRemoved = "هشدار حذف شد."

func WarnsQuantityOfUser(FirstName string, quantity int, maxQuantity int) string {
	return fmt.Sprintf("تعداد هشدارها  %s: %d/%d", FirstName, quantity, maxQuantity)
}

func NewWarnsQuantity(quantity string) string {
	return fmt.Sprintf("تعداد جدید حداکثر هشدارها: %s.", quantity)
}

func UserDoesNotHaveWarns(FirstName string) string {
	return fmt.Sprintf(" کاربر %s وارن ندارد.", FirstName)
}
