dlv attach 59515 --headless --listen=:2345 --api-version=2

launch.json

{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Attach to Remote",
            "type": "go",
            "request": "attach",
            "mode": "remote",
            "remotePath": "/home/tmolyakov/projects/projectx",
            "port": 2345,
            "host": "127.0.0.1",
            "showLog": true,
            "trace": "verbose"
        }
    ]
}