A quick little script to initiate a developer environment

Enables pass-through, appending to path, and setting environment variables

By default it passes-through these environment variables: HOME, USERPROFILE, APPDATA, SSH_AUTH_SOCK, OS

Also passes windows's Proccessor ID info

Environment variables (append+set) have replacement functions

current replacement funcs:

$(PWD) -> replaces with the current working directory

$(OS) -> runtime.GOOS

$(ARCH) -> runtime.GOARCH

$(USER) -> user.Current().Username

$(HOME) -> user.Current().HomeDir

$(EDIR) -> Exe's dir

example configuration:
```json
{
    "Terminal": "cmd.exe",
    "path": [
        "$(PWD)\\BinX\\go\\bin",
        "$(PWD)\\BinX\\mingw64\\bin",
        "$(PWD)\\go\\bin"
    ],
    "vars": [{
            "Name": "GOROOT",
            "Value": "$(PWD)\\BinX\\go"
        },
        {
            "Name": "GOHOME",
            "Value": "$(PWD)\\go"
        }
    ]
}
```
for a portable cgo-enabled go installation


cmd/exec is a proxy command runner for apps like vs code

example exetab
```json
{
    "code": "VSCode/Code"
}
```