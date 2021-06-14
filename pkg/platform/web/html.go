package web

import (
	"context"

	"inspr.dev/primal/pkg/api"
)

type Html struct {
	progress chan float32
	messages chan string
	done     chan bool
}

func NewHtml() *Html {
	return &Html{
		progress: make(chan float32),
		messages: make(chan string),
		done:     make(chan bool),
	}
}

func (h *Html) Progress() <-chan float32 {
	return h.progress
}

func (h *Html) Messages() <-chan string {
	return h.messages
}

func (h *Html) Done() <-chan bool {
	return h.done
}

var htmlTmpl = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">

	<link rel="preload" href="/client.css" as="style">
	<link rel="preload" href="/client.js" as="script">
	<link rel="preload" href="/assets/logo.VWJGXQZ7.png" as="image">
	<link rel="preload" href="/assets/bg.J2FRSW2E.png" as="image">

    <link rel="stylesheet" href="/client.css">
    <title>Primal</title>
</head>
<body>
    <div id="root"></div>
</body>
<script src="/client.js" ></script>
</html>
`

func (h *Html) Apply(ctx context.Context, spec api.Spec) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		html := htmlTmpl
		spec.Files.Write("/index.html", []byte(html))

		h.progress <- 1.0
		h.messages <- " 🎉 compiled html file with success"
		h.done <- true
		return nil
	}
}
