## Starting
`go get`

## Building and Running the script

Edit whatever you need to and then build the binary
`go build -o scripts/run_roku ./roku`

Execute the binary
`./scripts/run_roku`

## Setting the roku IP

Still need to do this manually (before running the script) - I would be thrilled if a PR appeared to extend the script to do this 
automatically (arp scripts are incomplete/non-functional).

`arp -a`

(you should be able to see a roku hostname in the arp table, with an ip address next to it)

`export ROKU_DEV_TARGET=192.xxx.yyy.zzz`