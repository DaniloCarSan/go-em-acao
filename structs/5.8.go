//erro de compilação ao atribuir valor de tipos diferentes
package main

type Duration int64

func main() {

	var dur Duration

	dur = int(1000)

}
