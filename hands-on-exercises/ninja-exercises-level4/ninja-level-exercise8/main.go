package main

import "fmt"

func main() {
	m := map[string][]string{
		"nakazima_andre":   []string{`video games`, `hambugers`, `ninita`},
		"gonçalves_amanda": []string{`puppies`, `shopping`, `ninito`},
		"lino_dobby":       []string{`cheese`, `sofa`, `sleeping`},
		"quita_paço":       []string{`food`, `chewing`, `cozy`},
	}

	fmt.Println(m)

	for k, v := range m {
		fmt.Println("This is the record for:", k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}
}
