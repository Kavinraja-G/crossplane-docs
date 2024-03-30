# crossplane-docs
## What is XDocs?
- We have XR, XRD, XRC ;) Why not XDocs?
- Inspired from [terraform-docs](https://github.com/terraform-docs/terraform-docs) but for [Crossplane](https://www.crossplane.io/)
- Generate markdown based docs for your Compositions, it also included your linked XRDs and Claims

## Installation
**XDocs** Generator can be installed in several ways to suit your preferences and development environments
### Homebrew Tap (macOS & Linux)

If you're using Homebrew, you can install the tool via our custom Homebrew tap

```bash
brew tap Kavinraja-G/tap
brew install crossplane-docs
```
### Standalone Binary
For macOS, Linux, and Windows, standalone binaries are available. Download the appropriate binary for your operating system from the [releases](https://github.com/Kavinraja-G/crossplane-docs/releases/) page, then move it to a directory in your PATH.

#### macOS/Linux
```bash
# Feel free to change the release/arch names accordingly
curl -Lo crossplane-docs https://github.com/Kavinraja-G/crossplane-docs/releases/download/v0.1.0/crossplane-docs_v0.1.0_darwin_amd64.tar.gz
chmod +x crossplane-docs
sudo mv crossplane-docs /usr/local/bin/
```
#### Windows
Download the `.exe` file and add it to your `PATH`

## Usage
Currently xDocs supports only markdown output, but more in pipeline. To generate markdown docs for your compositions & XRDs
```bash
crossplane-docs md [INPUT_PATH|INPUT_FILE] -o [OUTPUT_FILE]
```
For example:
```bash
crossplane-docs md ./samples -o samples/README.md
```
Check [README.md](./samples/README.md) for the output.

## License
Distributed under the Apache-2.0 License. See [LICENSE](./LICENSE) for more information.
