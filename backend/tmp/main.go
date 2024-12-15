package main

import (
	"fmt"
	"time"

	"github.com/telegram-mini-apps/init-data-golang"
)

func main() {
	// Init data in raw format.
	initData := "user=%7B%22id%22%3A2200607311%2C%22first_name%22%3A%22RootTon%22%2C%22last_name%22%3A%22%22%2C%22username%22%3A%22rootton_vf%22%2C%22language_code%22%3A%22ru%22%2C%22allows_write_to_pm%22%3Atrue%2C%22photo_url%22%3A%22https%3A%5C%2F%5C%2Fa-ttgme.stel.com%5C%2Fi%5C%2Fuserpic%5C%2F320%5C%2Fizvwowz_50RovbvAzVg3dItUvIwVkfsp6EkA5g47VZRE9KLTgtf5EpY-RIxT_jMk.svg%22%7D&chat_instance=2450319013474712658&chat_type=private&auth_date=1734215432&signature=I9DwoZqueBMKw2wyxvY4BKm6Kk2PMhUmFuN8bQONG42689zuVwu74fK5OfjeaeHlcgw_aXsEPX3pTSV8dXeaAg&hash=732ed8ffcbcdf2d8f7ed3bfc642b54a373b3ead451ac2e8f815a9b131206e03c"

	// Telegram Bot secret key.
	token := "2200176272:AAHPlJmLm5EdUFqnPrS8AdzDoD540Hm37nQ"

	// Define how long since init data generation date init data is valid.
	expIn := 240 * time.Hour

	// Will return error in case, init data is invalid.
	if err := initdata.Validate(initData, token, expIn); err != nil {
		fmt.Println(err)
		return
	}

	user, _ := initdata.Parse(initData)

	fmt.Println("Init data is valid")
}
