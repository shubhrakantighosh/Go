package main

import "log"

func main() {

	// animals := []string{"cow", "dog", "fish", "bird"}

	// for i, j := range animals {
	// 	log.Println(i, j)
	// }

	// for _, animals := range animals {
	// 	log.Println(animals)
	// }

	animals1 := make(map[string]string)

	animals1["dog"] = "faido"
	animals1["fish"] = "tush"
	animals1["horse"] = "miky"
	animals1["lion"] = "rambo"

	// for typeAnimal, animal := range animals1 {
	// 	log.Println(typeAnimal, animal)
	// }

	for j, i := range animals1 {
		log.Println(j, i)
	}

}
