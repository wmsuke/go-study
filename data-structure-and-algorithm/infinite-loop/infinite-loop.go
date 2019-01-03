package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/urfave/cli"
)

//リスト1-2

func main() {
	app := cli.NewApp()

	app.Action = func(context *cli.Context) error {
		var n, err = strconv.Atoi(context.Args().Get(0))
		if err != nil {
			fmt.Println("正の整数を入力してください")
			return nil
		}
		fmt.Println("1+2+...+n-1の計算")
		fmt.Println(" nの値：" + strconv.Itoa(n))
		var sum int
		for i := 1; i < n; i++ {
			sum += i
		}
		fmt.Println("1+2+...+n-1は" + strconv.Itoa(sum) + "です")
		return nil
	}

	app.Run(os.Args)
}
