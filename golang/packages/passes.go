package passes

func Certs(rodzaj string) string {
	var db_username = "postgres://miecieks:ZAQ!2wsx@biblioteka-baza-1:5432/biblioteka"
	var returnable = "null"
	switch rodzaj {
	case "nazwa":
		returnable = db_username
	}
	return returnable
}
