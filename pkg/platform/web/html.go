package web

import (
	"context"

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

func (h *Html) Apply(ctx context.Context, opts api.OperatorOptions) error {
	h.meta.State <- api.WORKING

	select {
	case <-ctx.Done():
		return nil
	default:
		html := htmlTmpl
		opts.Files.Write("/index.html", []byte(html))

		h.meta.Progress <- 1.0
		h.meta.Messages <- " 🎉 compiled html file with success"
		h.meta.State <- api.DONE
		return nil
	}
}
