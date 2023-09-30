# positron

Electron checkup for macOS users

Ref: <https://infosec.exchange/@TomSellers/111126352647377681>

## Build

```bash
$ git clone https://gitlab.com/hrbrmstr/positron.git
$ cd positron
$ just build
```

## Run

```bash
$ find /Applications -type f -name "*Electron Framework*" -exec ./positron "{}" \;
/Applications/Signal.app: Chrome/114.0.5735.289 Electron/25.8.4 🟢
/Applications/Keybase.app: Chrome/87.0.4280.141 Electron/11.5.0 🔴
/Applications/Raindrop.io.app: Chrome/102.0.5005.167 Electron/19.0.17 🔴
/Applications/1Password.app: Chrome/114.0.5735.289 Electron/25.8.1 🟢
/Applications/Replit.app: Chrome/116.0.5845.188 Electron/26.2.1 🟢
/Applications/lghub.app: Chrome/104.0.5112.65 Electron/20.0.0 🔴
```