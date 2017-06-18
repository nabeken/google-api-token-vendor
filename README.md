# google-api-token-vendor

This tiny command-line app allows you to retrieve OAuth2 token for offline (aka CLI).
The token will be printed to STDOUT.

This code is extracted from https://developers.google.com/drive/v3/web/quickstart/go for ease of reuse between future projects.

## Usage

```sh
Usage of ./google-api-token-vendor:
  -c string
        JSON file for client secret (default "client_secret.json")
  -s string
        scope you want to have (default "https://www.googleapis.com/auth/drive.file")
```
