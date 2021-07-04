# snakego

## Usage

Run your own server with command:

```bash
docker run -d -p 8080:8080 --name snake ghcr.io/arthur-snake/snakego:latest
```

Then open http://localhost:8080/ in the browser.

<details>
<summary>Development</summary>

Make sure you have:
- Go 1.16, [install](https://golang.org/doc/install)
- GoLand / VSCode / other IDE, [install goland](https://www.jetbrains.com/go/)
- golangci-lint 1.40, [install](https://golangci-lint.run/usage/install/)


### EnvFile plugin

EnvFile plugin for GoLand is useful for applying conf from .env files. Install [here](https://plugins.jetbrains.com/plugin/7861-envfile).

To use it:
- Open [Run configuration]
- Select EnvFile tab
- Add file .env from repo root
  * On macOS press shirt+cmd+. to display hidden files
</details>