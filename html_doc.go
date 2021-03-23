package lib

// HTMLDoc fn
func HTMLDoc(title, head, body string) string {
	return `
<!DOCTYPE html>
<html lang="es">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>` + title + `</title>
	` + head + `
</head>
<body>
` + body + `
</body>
</html>	
	`
}
