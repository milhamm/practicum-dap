package main

import "fmt"

/* (A)
   Name: Muhammad Ilham Mubarak
   Sid: 1301194276

   (B) The output for the second input example:

	hopefully it is a farewell and not a good bye!
	good memories and laughters surely we will share

*/
const MAXCHAR = 256

type dataType struct {
	thurs, day int
}

func readDict(dict *[MAXCHAR]byte, ndict *int, magic *int, base *int) {
	var ch byte
	var i int
	/* (C) read characters one at a time, end marker is a '.'
	      store the chars to dict, and
	the number of characters to ndict
	      then read two integers for magic and base */
	i = 0
	fmt.Scanf("%c", &ch)
	for ch != '.' {
		dict[i] = ch
		i++
		*ndict++
		fmt.Scanf("%c", &ch)
	}

	fmt.Scan(magic, base)

}

func readData(dat *[MAXCHAR]dataType, ndat *int) {
	var i, val1, val2 int
	/* (D) read pairs of integers, ended with a pair of 9999
	   store them into thurs and day fields of array dat
	   store the number of data in ndat */

	i = 0
	fmt.Scan(&val1, &val2)

	for val1 != 9999 && val2 != 9999 {
		dat[i].thurs = val1
		dat[i].day = val2
		i++
		fmt.Scan(&val1, &val2)
	}

	*ndat = i
}

func sortOnThurs(dat *[MAXCHAR]dataType, ndat int) {
	/* (F) sort on thurs fields in descending order
	   using selection sort */
	var i int = 0

	for i < ndat {
		var j = i
		var max = i
		for j < ndat {
			if dat[j].thurs > dat[max].thurs {
				max = j
			}
			j++
		}
		temp := dat[i]
		dat[i] = dat[max]
		dat[max] = temp
		i++
	}
}

func sortOnDay(dat *[MAXCHAR]dataType, ndat int) {
	/* (G) sort on day fields in descending order
	   using insertion sort */
	var i int = 1

	for i < ndat {
		var j int = i - 1
		temp := dat[i]

		for j >= 0 && dat[j].day < temp.day {
			dat[j+1] = dat[j]
			j--
		}

		dat[j+1] = temp
		i++
	}
}

func keyOne(dat [MAXCHAR]dataType, ndat int, magic int, base int) int {
	var pos int
	/* (H) set pos to index where magic is found in thurs fields,
	   using binary search method.
	   The array dat is in descending order on thurs */

	var left int = 0
	var right int = ndat
	var found bool = false

	for left < right && !found {
		var med int = (left + right) / 2
		if dat[med].thurs < magic {
			right = med
		} else if dat[med].thurs > magic {
			left = med + 1
		} else {
			found = true
			pos = med
		}
	}
	return dat[pos].day % base
}

func keyTwo(dat [MAXCHAR]dataType, ndat int, base int) int {
	var pos int
	/* (J) set pos to median index,
	   the array dat is in descending order on day */
	pos = ndat / 2
	return dat[pos].day % base
}

func printSecretOne(dat [MAXCHAR]dataType, ndat int, dict [MAXCHAR]byte, ndict int, magic, key int) {
	for i := 0; i < ndat; i++ {
		if dat[i].thurs != magic {
			/* (K) print a character from array dict.
			   The index is modulo key of the day field in dat */

			fmt.Printf("%c", dict[dat[i].day%key])
		}
	}
	fmt.Println()
}

func printSecretTwo(dat [MAXCHAR]dataType, ndat int, dict [MAXCHAR]byte, ndict int, magic, key int) {
	/* L do the same as printSecretOne, except now on thurs fields*/
	for i := 0; i < ndat; i++ {
		if dat[i].thurs != magic {
			fmt.Printf("%c", dict[dat[i].thurs%key])
		}
	}
	fmt.Println()
}

func main() {
	var dict [MAXCHAR]byte
	var ndict int
	var magic, base int
	var dat [MAXCHAR]dataType
	var ndat int
	var key int

	readDict(&dict, &ndict, &magic, &base)
	readData(&dat, &ndat)
	sortOnThurs(&dat, ndat)
	key = keyOne(dat, ndat, magic, base)
	printSecretOne(dat, ndat, dict, ndict, magic, key)
	sortOnDay(&dat, ndat)
	key = keyTwo(dat, ndat, base)
	printSecretTwo(dat, ndat, dict, ndict, magic, key)
}
