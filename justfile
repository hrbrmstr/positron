build:
  GOOS=darwin GOARCH=amd64 go build -o positron-amd64 .
  GOOS=darwin GOARCH=arm64 go build -o positron-arm64 .
  lipo -create -output positron positron-amd64 positron-arm64
  codesign -s "${APPLE_SIGN}" positron

run:
  find /Applications -type f -name "*Electron Framework*" -exec ./positron "{}" \;