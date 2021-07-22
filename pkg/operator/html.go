package operator

import (
	"inspr.dev/primal/pkg/filesystem"
)

type Html struct{}

func NewHtml() *Html {
	return &Html{}
}

func (*Html) Handler(fs filesystem.FileSystem) {
	var htmlTmpl = `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	    <meta charset="UTF-8">
	    <meta http-equiv="X-UA-Compatible" content="IE=edge">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta name="theme-color" content="white">
		<meta name="theme-color" media="(prefers-color-scheme: light)" content="white">
		<meta name="theme-color" media="(prefers-color-scheme: dark)" content="black">
		<link rel="preload" href="/entry-client.css" as="style">
		<link rel="modulepreload" href="/entry-client.js">
		<link rel="modulepreload" href="/react-dom.RT5KN4QJ.js">
	    <link rel="stylesheet" href="/entry-client.css">
		<title>Primal</title>
	</head>
	<body>
	    <div id="root"></div>
	</body>
	<script type="module" src="/entry-client.js" ></script>
	</html>`
	fs.Write("/index.html", []byte(htmlTmpl))
}
