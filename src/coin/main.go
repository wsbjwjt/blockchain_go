package main

import "core"

//const targetBits = 24
//
//func proof() {
//	data1 := []byte("I like donuts")
//	data2 := []byte("I like donutsca07ca")
//	target := big.NewInt(1)
//	target.Lsh(target, uint(256 - targetBits))
//	fmt.Printf("%x\n", sha256.Sum256(data1))
//	fmt.Printf("%064x\n", target)
//	fmt.Printf("%x\n", sha256.Sum256(data2))
//}

func main() {
	//proof()
	//bc := core.NewBlockchain()
	//defer bc.Db.Close()

	cli := core.CLI{}
	cli.Run()
}
