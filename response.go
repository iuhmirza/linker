package main

import (
	"html/template"
	"net/http"
)

const SHORT_LINK_PAGE_TEMPLATE string = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <link rel="stylesheet" href="styles.css">
    <title>Document</title>
</head>
<body>
    <code>localhost:3000/{{.}}</code>
</body>
</html>
`

var SHORT_LINK_PAGE *template.Template

func init() {
	SHORT_LINK_PAGE, _ = template.New("short_page").Parse(SHORT_LINK_PAGE_TEMPLATE)
}

func GenerateResponse(res http.ResponseWriter, short string) {
	SHORT_LINK_PAGE.Execute(res, short)
}
