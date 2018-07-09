# portscanner
test port scanner

## Background

nmap is a utility that scans remote network hosts and attempts to open
TCP/UDP connections to determine open ports.

This program is a similar. It attempts to open connections to determine
open ports. It also leverages a semaphore to limit the number of concurrent
connections. 
