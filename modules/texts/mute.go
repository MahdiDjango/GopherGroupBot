package texts

import "fmt"

const TalkToMeHere = "اینجا با من صحبت کن..."

const ThisCommandForTemporaryMute = "این دستور برای سکوت موقت می باشد.\n" +
	"m - دقیقه, h - ساعت, d - روز.\nمثال: /tmute @user 1h 30m"

func UserIsMuted(name string) string {
	return fmt.Sprintf("کاربر %s خفه شد.", name)
}
