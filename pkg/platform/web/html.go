package web

import (
	"inspr.dev/primal/pkg/api"
)

type Html struct {
	meta api.Metadata
}

func NewHtml() *Html {
	return &Html{
		meta: api.NewMetadata(),
	}
}

func (h *Html) Meta() api.Metadata {
	return h.meta
}

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
</html>
`

func (h *Html) Apply(props api.OperatorProps, opts api.OperatorOptions) {
	var writeHtml = func() {
		html := htmlTmpl
		props.Files.Write("/index.html", []byte(html))
		h.meta.Messages <- " 🎉 compiled html file with success"
		h.meta.Done <- true
	}

	writeHtml()
Main:
	for {
		select {
		case <-h.meta.Close:
			break Main

		case <-h.meta.Refresh:
			writeHtml()
		}
	}
}
