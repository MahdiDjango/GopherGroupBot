package texts

import "fmt"

const DescriptionOfRulesCommand = "این دستور به شما امکان می دهد قوانین را تنظیم کنید."

const NewRulesWereAdded = "قوانین وضع شده است! می توانید آنها را با دستور مشاهده کنید /rules."

func ErrorOfGettingRules(errText string) string {
	return fmt.Sprintf("خطا در دریافت قوانین.\n%s", errText)
}
