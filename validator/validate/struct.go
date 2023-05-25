package validate

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Name        string            `v:"required,alphaunicode"`
	Age         uint8             `v:"gte=10,lte=130"`
	Phone       string            `v:"required,e164"`
	Email       string            `v:"required,email"`
	Address     *Address          `v:"required"`
	ContactUser []*ContactUser    `v:"required,gte=1,dive"`
	Hobby       []string          `v:"required,gte=2,dive,required,gte=2,alphaunicode"`
	Data        map[string]string `v:"required,gte=2,dive,keys,len=2,alpha,endkeys,required,gte=2,alphaunicode"`
}

type ContactUser struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=10,lte=130"`
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"`
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"`
	Address *Address `v:"required_without_all=Email Phone"`
}

type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}

func StructValidation() {
	v := validate
	address := &Address{
		Province: "湖南",
		City:     "长沙",
	}
	contactUser1 := &ContactUser{
		Name:  "张三",
		Age:   18,
		Phone: "+8612345678910",
	}
	contactUser2 := &ContactUser{
		Name:  "李四",
		Age:   21,
		Email: "666@666.com",
	}
	user := &User{
		Name:        "nick",
		Age:         18,
		Phone:       "+8612345678910",
		Email:       "email@email.com",
		Address:     address,
		ContactUser: []*ContactUser{contactUser1, contactUser2},
		Hobby:       []string{"篮球", "羽毛球"},
		Data: map[string]string{
			"AA": "乒乓球",
			"BB": "羽毛球",
		},
	}
	err := v.Struct(user)
	if err != nil {
		if errors, ok := err.(validator.ValidationErrors); ok {
			for _, err := range errors {
				fmt.Println(err)
			}
		}
	}
}
