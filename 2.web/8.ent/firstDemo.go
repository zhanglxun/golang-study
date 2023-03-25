package main

import (
	"context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"golangStudy/2.web/8.ent/ent/accounts"
	"golangStudy/2.web/8.ent/ent/cwebsite"

	//"entgo.io/ent"
	//_ "github.com/go-sql-driver/mysql"
	"golangStudy/2.web/8.ent/ent"
	"log"
)

func main() {

	client, err := ent.Open("mysql", "root:85894113@(106.55.235.27:3306)/susense-platform?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	//ctx := context.Background()

	//CreateAccount(context.Background(), client)
	//
	//SelectAccount(context.Background(), client)

	//CreateCWebsite(context.Background(), client)

	SelectCWebsite(context.Background(), client)

}

func CreateAccount(ctx context.Context, client *ent.Client) (*ent.Accounts, error) {
	account := client.Accounts.
		Create().
		SetID(5068).
		SetAccount("hanhan").
		SetPwd("123456").
		SetEmail("hanhan@gmail.com").
		SetNickname("hanhanLike").
		SetMobile("13677778888").SaveX(ctx)

	//if err != nil {
	//	return nil, fmt.Errorf("failed creating user: %v", err)
	//}
	log.Println("user was created: ", account)
	return account, nil
}

func SelectAccount(ctx context.Context, client *ent.Client) (*ent.Accounts, error) {

	accountObj, _ := client.Accounts.Query().Where(accounts.Account("hanhan")).First(ctx)

	println(accountObj.Account + "-" + accountObj.Email + "-" + accountObj.Mobile)
	return accountObj, nil
}

func CreateCWebsite(ctx context.Context, client *ent.Client) (*ent.CWebsite, error) {

	cwebsite := client.CWebsite.
		Create().
		SetID(3688).
		SetWebsiteURL("http://baidu2.com").
		SetWebsiteName("交易系统2").
		SetWebsiteIcon("http://abc2.com").
		SetDescription("测试系统2").SaveX(ctx)

	return cwebsite, nil
}

func SelectCWebsite(ctx context.Context, client *ent.Client) (*ent.CWebsite, error) {

	cwebsiteObj, _ := client.CWebsite.Query().Where(cwebsite.ID(3689)).First(ctx)

	println(cwebsiteObj.WebsiteName + "-" + cwebsiteObj.WebsiteURL)
	return cwebsiteObj, nil
}
