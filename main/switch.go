package main

func Whatday(n int) string {
	if n != 0 && n <= 7 {
		switch n {
		case 1:
			return `週一`
		case 2:
			return `週二`
		case 3:
			return `週三`
		case 4:
			return `週四`
		case 5:
			return `週五`
		case 6:
			return `週六`
		case 7:
			return `週日`
		default:
			return `錯誤`
		}

	} else {
		return `輸入錯誤`
	}

}
