# A tiny command line tool to make proxy management a little bit more user friendly

The tool simply adds or removes the `HTTP_PROXY` and `HTTPS_PROXY` environment variables. In addition all necessary informations
(proxy url, proxy port and proxy user; _but not proxy password_) are stored in a config file (`.proxy-manager.json`) in your 
home directory. If no home directory is present, the current directory is used or you can specify a directory of your choice.

You can also specify a timeout after that proxy-manager automatically removes the environment variables automatically. That becomes
especially handy when you use a proxy password, so the password is only visible (as environment variable) for the specified time period.

## Note
proxy-manager only works for programs that support the above mentioned environment variables.

## `proxy-manager -help`
