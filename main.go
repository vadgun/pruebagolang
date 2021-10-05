package main

import (
	"github.com/kataras/iris/v12"
	"fmt"
	"strconv"
	"errors"
)


type Respuesta struct{
	Cr3 int32 `json:"credit_300"`
	Cr5 int32 `json:"credit_500"`
	Cr7 int32 `json:"credit_700"`

}

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

func main(){
	app:= iris.New()
	app.Get("/credit-assignment/{inv}", AsignadorDeCreditos)
	app.Listen(":8080")
}

func AsignadorDeCreditos(ctx iris.Context){
	var res Respuesta
	var err1 error
	inv :=  ctx.Params().GetString("inv")
	inv2, err := strconv.Atoi(inv)
	if err != nil{
		fmt.Println("No se puede convertir este digito")
		ctx.StatusCode(400)
	}

	res.Cr3,res.Cr5,res.Cr7, err1 = res.Assign(int32(inv2))
	fmt.Println("Respuesta", res)

	if err1 != nil{
		ctx.StatusCode(400)
	}else{
		ctx.StatusCode(200)
		ctx.JSON(res)
	}
}

func (res Respuesta) Assign(inv int32)(int32, int32, int32, error){
	var err1 error

	fmt.Println("inversion", inv)
	if inv > 100 {
		return 0,0,0,errors.New("400")
	}

	return inv+1,inv+2,inv+3, err1
}