package texts

import "fmt"

const UserIsInTheChat = "این کاربر در حال چت کردن است."

func UserIsBanned(name string) string {
	return fmt.Sprintf("کاربر %s بن شده!", name)
}

func UserIsUnbanned(name string) string {
	return fmt.Sprintf("کاربر %s آنبن شده!", name)
}

func UserIsKicked(name string) string {
	return fmt.Sprintf("کاربر %s کیک شده!", name)
}

func UserIsPromoted(name string) string {
	return fmt.Sprintf("کاربر %s مدیر شد.", name)
}

func UserIsDemoted(name string) string {
	return fmt.Sprintf("کاربر %s صلب مدیر شد.", name)
}
