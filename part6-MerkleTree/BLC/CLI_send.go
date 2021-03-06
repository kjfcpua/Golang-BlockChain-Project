package BLC

import (
	"fmt"
	"os"
)

func (cli *CLI) send(from []string, to []string, amount []string) {
	/*
	address:  1Rs9zcPDqosXucdJjGP4wjGrtA1SmYpwGnQBMCprE2TdvhUyhk	c
	address:  1YfMAGkzTU3P19DobiAiggGzzcymvJyePughP37efhVgCV4W8e	b
	address:  1Z4DNkwSLgQR8yhtTnZSyobdenW3FfjMtkAnHJdM9ZAVenYDsU  	a
	 */
	//go run main.go send -from '["1Z4DNkwSLgQR8yhtTnZSyobdenW3FfjMtkAnHJdM9ZAVenYDsU"]' -to '["1YfMAGkzTU3P19DobiAiggGzzcymvJyePughP37efhVgCV4W8e"]' -amount '["4"]'
	//go run main.go send -from '["1Z4DNkwSLgQR8yhtTnZSyobdenW3FfjMtkAnHJdM9ZAVenYDsU","1Z4DNkwSLgQR8yhtTnZSyobdenW3FfjMtkAnHJdM9ZAVenYDsU"]' -to '["1YfMAGkzTU3P19DobiAiggGzzcymvJyePughP37efhVgCV4W8e","1Rs9zcPDqosXucdJjGP4wjGrtA1SmYpwGnQBMCprE2TdvhUyhk"]' -amount '["2","1"]'
	//go run main.go send -from '["1YfMAGkzTU3P19DobiAiggGzzcymvJyePughP37efhVgCV4W8e","1Rs9zcPDqosXucdJjGP4wjGrtA1SmYpwGnQBMCprE2TdvhUyhk"]' -to '["1Rs9zcPDqosXucdJjGP4wjGrtA1SmYpwGnQBMCprE2TdvhUyhk","1Z4DNkwSLgQR8yhtTnZSyobdenW3FfjMtkAnHJdM9ZAVenYDsU"]' -amount '["3","1"]'
	//go run main.go send -from '["1Z4DNkwSLgQR8yhtTnZSyobdenW3FfjMtkAnHJdM9ZAVenYDsU"]' -to '["1Rs9zcPDqosXucdJjGP4wjGrtA1SmYpwGnQBMCprE2TdvhUyhk"]' -amount '["8"]'
	/*
	1/	a->b 4					a: 16 / b: 4 / c: 0
	2/	a->b 2  a->c 1			a: 23 / b: 6 / c: 1
	3/	b->c 3  c->a 1			a: 24 / b: 13 / c: 3
	4/  a->c 8					a: 26 / b: 13 / c: 11
	 */
	if DBExists() == false {
		fmt.Println("数据不存在.......")
		os.Exit(1)
	}

	bc := BlockchainObject()
	defer bc.DB.Close()

	bc.MineNewBlock(from, to, amount)

	utxoSet := &UTXOSet{bc}
	utxoSet.Update()
}
