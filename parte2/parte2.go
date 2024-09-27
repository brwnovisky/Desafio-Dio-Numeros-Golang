package parte2

func Multiplo3e5() {
	for counter := 1; counter <= 100; counter++ {
		if counter%3 == 0 {
			println("Pin")
			continue
		}

		if counter%5 == 0 {
			println("Pan")
			continue
		}

		println(counter)
	}
}
