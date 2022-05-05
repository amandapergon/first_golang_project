package main

import "fmt"

func main() {
	m := map[string][]string{
		"nakazima_andre":   []string{`video games`, `hambugers`, `ninita`},
		"gonçalves_amanda": []string{`puppies`, `shopping`, `ninito`},
		"lino_dobby":       []string{`cheese`, `sofa`, `sleeping`},
		"quita_paço":       []string{`food`, `chewing`, `cozy`},
		"vampire_lazlo":    []string{`blood`, `sex`, `Nadja`},
	}

	fmt.Println(m)

	delete(m, "vampire_lazlo")
	fmt.Println(m)
}
