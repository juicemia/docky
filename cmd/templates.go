package cmd

const root = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">

    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/css/bootstrap.min.css"
        integrity="sha384-Smlep5jCw/wG7hdkwQ/Z5nLIefveQRIY9nfy6xoR1uRYBtpZgI6339F5dgvm/e9B"
        crossorigin="anonymous"
        rel="stylesheet">
    
    <link rel="stylesheet"
        href="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/styles/default.min.css">
    
    <!-- Putting this in here to avoid having to load any local files. -->
    <style>
    .container.base {
        padding: 15px;
    }
    </style>

    <title>{{.AppName}} API Documentation</title>
</head>
<body>
    <nav class="navbar navbar-expand-lg navbar-dark bg-dark">
        <a href="/" class="navbar-brand">{{.AppName}} API</a>

        <div class="nav-item dropdown">
            <a class="nav-link dropdown-toggle text-light" href="#" role="button" data-toggle="dropdown">Resources</a>

            <div class="dropdown-menu">
                {{range $index, $element := .Resources}}
                <a class="dropdown-item" href="./{{$element.GetLinkName}}.html">{{$element.Name}}</a>
                {{end}}                
            </div>
        </div>
    </nav>
    
    {{ template "body" . }}

    <script src="https://code.jquery.com/jquery-3.3.1.slim.min.js"
        integrity="sha384-q8i/X+965DzO0rT7abK41JStQIAqVgRVzpbzo5smXKp4YfRvH+8abtTE1Pi6jizo"
        crossorigin="anonymous">
    </script>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.14.3/umd/popper.min.js"
        integrity="sha384-ZMP7rVo3mIykV+2+9J3UJ46jBk0WLaUAdn689aCwoqbBJiSnjAK/l8WvCWPIPm49"
        crossorigin="anonymous">
    </script>
    <script src="https://stackpath.bootstrapcdn.com/bootstrap/4.1.2/js/bootstrap.min.js"
        integrity="sha384-o+RDsa0aLu++PJvFqy8fFScvbHFLtbvScb8AjopnFD+iEQ7wo/CG0xlczd+2O/em"
        crossorigin="anonymous">
    </script>
    <script src="//cdnjs.cloudflare.com/ajax/libs/highlight.js/9.12.0/highlight.min.js">
        </script>
    
    <!-- Putting this in here to avoid loading any local files. -->
    <script type="">
        hljs.initHighlightingOnLoad();
    </script>
</body>
</html>`

const route = `{{ define "body" }}
<div class="base container">
    {{range $index, $element := .Resource.Routes}}
    <div class="row">
        <div class="col">
            <a name="{{$element.GetLinkName}}"><h2 id="{{$element.GetLinkName}}">{{$element.Method}} {{$element.Path}}</h2></a>

            <p>{{$element.Description}}</p>
        </div>
    </div>

    <div class="row">
        <div class="col"><h3>Request</h3></div>
    </div>

    <div class="row">
        <div class="col">
            <h5>Headers</h5>

            {{range $key, $val := $element.Headers}}
            <pre><code>{{$key}}: {{$val}}</code></pre>
            {{else}}
            <p>None</p>
            {{end}}
        </div>

        <div class="col">
            <h5>Parameters</h5>

            {{range $key, $val := $element.Parameters}}
            <pre><code>{{$key}}: {{$val}}</code></pre>
            {{else}}
            <p>None</p>
            {{end}}
        </div>
    </div>

    {{range $index, $element := $element.Responses}}
    <div class="row">
        <div class="col"><h3>{{$element.Status}} - {{$element.Description}}</h3></div>
    </div>

    <div class="row">
        <div class="col-6">
            <h5>Headers</h5>

            {{range $key, $val := $element.Headers}}
            <pre><code>{{$key}}: {{$val}}</code></pre>
            {{else}}
            <p>None</p>
            {{end}}
        </div>

        <div class="col-6">
            <h5>Body</h5>

            <pre><code class="json">{{$element.BodyJSON}}</code></pre>
        </div>
    </div>
    {{end}}
    {{end}}
</div>
{{ end }}`

const index = `{{ define "body" }}
<div class="base container">
    {{range $index, $element := .Resources}}
    {{$baseLinkName := $element.GetLinkName}}
    <div class="row">
        <div class="col">
            <h2><a href="./{{$baseLinkName}}.html" class="text-body"><u>{{$element.Name}}</u></a></h2>

            <p>{{$element.Description}}</p>
        </div>
    </div>

    <div class="row">
        <div class="col"><h3>Routes</h3></div>
    </div>

    {{range $index, $element := $element.Routes}}
    <div class="row">
        <div class="col">
            <p><a href="./{{$baseLinkName}}.html#{{$element.GetLinkName}}" class="text-body"><u>{{$element.Method}} {{$element.Path}}</u></a></p>
        </div>
    </div>
    {{end}}
    {{end}}
</div>
{{ end }}`
