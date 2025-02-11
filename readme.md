# [gio ui](https://gioui.org/) caching mechanism

### gio-cache serves 3 main purposes

1. Caches your WASM binary in a users browser based on a semantic version in the ETag
![gio cache running example](./.rsrc/gio-cache.png)

2. Shows a simple css spinner while your binary is being loaded & initialized in the browser
![simple css spinner](./.rsrc/simple_css_spinner.gif)

3. Compresses the main.wasm using [brotli](https://github.com/google/brotli) & serves the compressed version

# Usage
- place your wasm binary under this project directory in the bin dir
    - ![bin dir structure](./.rsrc/bin_dir_structure.png)
- `gio-cache` will fail fast if you don't set the environment variable `WASM_BINARY_VERSION`
    - this prevents you from accidentally deploying your app without a version (which results in some heavy network traffic)
- build the binary
```
make
```
- override any static files for your loading screen
```
make setup
```
- serve the app
```
WASM_BINARY_VERSION=v0.0.3 ./gio-cache
```
