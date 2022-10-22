package main

import "fmt"




func hesabla() {
	for j:=1;j<=10;j++ {
		fmt.Println("---------------------------")
		for i:=1;i<=10;i++{
			fmt.Println(j, "*", i, "=" , j,i,j*i)
		}
	}
}



func main()  {
	hesabla()
}