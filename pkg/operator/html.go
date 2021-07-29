package operator

import "inspr.dev/primal/pkg/workflow"

// Html is the html operator
type Html struct {
	*Operator
}

// NewHtml returns a new html operator
func (op *Operator) NewHtml() *Html {
	return &Html{
		op,
	}
}

// Task returns a html operator's workflow task
func (html *Html) Task() workflow.Task {
	return workflow.Task{
		ID:    "html",
		State: workflow.IDLE,
		Run: func(self *workflow.Task) {
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
				<link rel="modulepreload" href="/react-dom.WPNHM4JB.js">
				<link rel="stylesheet" href="/entry-client.css">
				<title>Primal</title>
			</head>
			<body>
				<div id="root"></div>
			</body>
			<script type="module" src="/entry-client.js" ></script>
			</html>`

			html.Fs.Write("/index.html", []byte(htmlTmpl))

			self.State = workflow.DONE
		},
	}
}
