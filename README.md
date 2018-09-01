# ewp - a tiny command line tool that executes commands with proxy setup

ewp stands for _execute with proxy_. The tool simply starts command line programs in a new environment with set `HTTP_PROXY` and `HTTPS_PROXY` environment variables. In addition all necessary informations
(proxy url, proxy port and proxy user; _but not proxy password_) are stored in a config file (`.ewp-config.json`) in your home directory for easier reuse of the proxy configuration. If no home directory is present, the current directory is used or you can specify a directory of your choice.

You can also specify a timeout after that proxy-manager automatically removes the environment variables automatically. That becomes handy especially when you use a proxy password, so the password is only visible (as environment variable) for the specified time period.

## Note
proxy-manager only works for programs that support the above mentioned environment variables.
