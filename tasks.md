[] Describe a mechanism for pipeline operations in the compiler
for example, take config for the project and apply transformations in steps,
step1 compile / bundle, step 2 minify, step3 gzip, step 4 lunch client or do any other optimizations.

The goal of primal now is only to compile for the web and later add different compilation targets such as react native and electron, windows

[] Create a server in go that can run / serve the assets for web

[] Create a development server in go, that can run in watch mode

[] Update the web client and server to support PWA

[] Update the types for primitives and state

[] Create a basic navigator for route based routing but using state

[] Add bindings to state in golang

[] Generate the HTML template as an step in the Platform web