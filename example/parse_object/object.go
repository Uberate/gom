package main

type Human struct {
	Name   string  `json:"name" expr:"((Tom)|(XiaoMing)|(GOM)) ((Jack)|(Test)|(Lily))"`
	Sex    bool    `json:"sex"`
	Height float64 `json:"height" min:"1.1" max:"2.3"`
}

type Student struct {
	Age   uint8  `json:"age" min:"6" max:"30"`
	Email string `json:"email" expr:"^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$" expr-rc:"0"`
	Human `json:"human"`
}
