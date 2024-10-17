package main

import "fmt"

func main() {
	sl := NewSkipList(16)

	// insert:
	// aaa -> 12312
	// bbb -> 1231231
	// c -> 32423sdf
	// d -> ksjdfk

	// delete:
	// bbb

	// insert:
	// abb -> ksjdfksd
	// aac -> slkdjfsd

	// delete:
	// abb

	// truncate list

	ins := [][]string{
		[]string{"aaa", "12312"},
		[]string{"bbb", "1231231"},
		[]string{"c", "32423sdf"},
		[]string{"d", "ksjdfk"},

		[]string{"abb", "ksjdfksd"},
		[]string{"aac", "slkdjfsd"},
	}

	for i := 0; i < 5; i++ {
		sl.Insert([]byte(ins[i][0]), []byte(ins[i][1]))
	}

	fmt.Println(sl.Delete([]byte(ins[1][0])))

	for i := 5; i < len(ins); i++ {
		sl.Insert([]byte(ins[i][0]), []byte(ins[i][1]))
	}

	fmt.Println(sl.Delete([]byte(ins[5][0])))

	sl.Print()

}
