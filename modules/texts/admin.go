package texts

import "fmt"

const PleaseReplyToTheMessageYouWantToPin = "لطفا به پیام ریپلای کنید."

const ThisChatIsPrivateICanNotToPinMessage = "در گروه های خصوصی نمیتوانم پیامی را پین کنم."

const ThisChatIsPrivateICanNotToUnpinMessage = "در گروه های خصوصی نمیتوانم پیامی را از پین در بیارم."

const PleaseReplyToTheMessageOfThePersonYouWantToGrantAdministratorRightsToOrEnterTheirID = "لطفا ریپلای کنید " +
	"یا شناسه آن کاربر را وارد کنید."

const PleaseReplyToTheMessageOfThePersonYouWantToRemoveAdministratorRightsToOrEnterTheirID = "لطفا ریپلای کنید  " +
	"یا شناسه آن کاربر را وارد کنید."

const PurgeCompleted = "پاکسازی کامل شد پس از ۵ ثانیه پیام حذف می شود."

const AllUserCommandsAreDisabled = "همه دستورات سفارشی غیرفعال هستند."

const AllUserCommandsAreEnabled = "همه دستورات سفارشی فعال هستند."

const YouDidNotWriteAnyUserCommands = "هیچ دستور سفارشی ننوشته اید."

func DisabledUserCommandsList(disabledCommands string) string {
	return fmt.Sprintf("دستورات سفارشی زیر غیرفعال هستند: %s", disabledCommands)
}

func ReportMessage(chatName string, from string, per string) string {
	return fmt.Sprintf("گزارش چت جدید %s\nفرستنده: %s\nچه کسی: %s\nپیام:", chatName, from, per)
}
