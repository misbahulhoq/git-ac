const binwrap = require("binwrap");
const path = require("path");
const packageInfo = require("./package.json");
const version = packageInfo.version;
const root = `https://github.com/misbahulhoq/gcli/releases/download/v${version}`;

module.exports = binwrap({
  dirname: __dirname,
  binaries: ["gcli"],
  urls: {
    // These must match the filenames you upload to GitHub Releases EXACTLY.
    // Format: "OS-Arch": "Download URL"
    // Windows 64-bit
    "win32-x64": root + "/gcli-windows.exe",

    // Linux 64-bit
    "linux-x64": root + "/gcli-linux",

    // Mac (Intel)
    "darwin-x64": root + "/gcli-mac",

    // Mac (M1/M2/Apple Silicon)
    "darwin-arm64": root + "/gcli-mac",
  },
});
