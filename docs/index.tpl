<!DOCTYPE html>
<html lang="en">
    <head>
        <meta charset="utf-8" />
        <title>{{ .Title }}</title>
        <link rel="stylesheet" type="text/css" href="//unpkg.com/swagger-ui-dist@3.40.0/swagger-ui.css" />
        <style>
      html
      {
        box-sizing: border-box;
        overflow: -moz-scrollbars-vertical;
        overflow-y: scroll;
      }

      *,
      *:before,
      *:after
      {
        box-sizing: inherit;
      }

      body
      {
        margin:0;
        background: #fafafa;
      }
    </style>
    </head>
    <body>
        <div id="swagger-ui"></div>
        <script src="//unpkg.com/swagger-ui-dist@3.40.0/swagger-ui-bundle.js"></script>
        <script>
            // init Swagger for faucet's openapi.yml.
            /*
            window.onload = function() {
              window.ui = SwaggerUIBundle({
                url: {{ .URL }},
                dom_id: "#swagger-ui",
                deepLinking: true,
                layout: "BaseLayout",
              });
            }*/
            window.onload = function() {
              // Build a system
              const ui = SwaggerUIBundle({
                url: {{ .URL }},
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                  SwaggerUIBundle.presets.apis,
                  SwaggerUIStandalonePreset
                ],
                plugins: [
                  SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
              })
              window.ui = ui
            }
          </script>
    </body>
</html>