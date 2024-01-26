package utils


/*
ch -> chinese
en -> english
ja ->  japanese
ko -> korea
ru -> russia
ge -> german

*/

//ch -> en -> ch
func Method1(query string) string {
	return Convert(Convert(query, "zh", "en"), "en", "zh")
}


//ch -> ja -> ch
func Method2(query string) string {
	return Convert(Convert(query, "zh", "jp"), "jp", "zh")
}

//ch -> ko -> ch
func Method3(query string) string {
	return Convert(Convert(query, "zh", "kor"), "kor", "zh")
}

//ch -> ru -> ch
func Method4(query string) string {
	return Convert(Convert(query, "zh", "ru"), "ru", "zh")
}

//ch -> ge -> ch
func Method5(query string) string {
	return Convert(Convert(query, "zh", "de"), "de", "zh")
}

//ch -> en -> ge -> ch
func Method6(query string) string {
	return Convert(Convert(Convert(query, "zh", "en"), "en", "de"), "de", "zh")
}

//ch -> en -> ge -> ja -> ch
func Method7(query string) string {
	return Convert(Convert(Convert(Convert(query, "zh", "en"), "en", "de"), "de", "jp"), "jp", "zh")
}