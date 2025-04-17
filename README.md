# NETWORK ACL CHECKER
## Description
* This is a simple application that checks the network ACLs.


## Configuration
* `config.toml` has the following configuration options:
  * `server port` - echo server port.
  * `clinet address` - The address of the server to check.

## Usage
* The application will check the network ACLs and print the results to the console.
* To run the server application, use the following command:
```shell
./acl -type server -protocol tcp 
```


* To run the client check application, use the following command:
```shell
./acl -type clent -protocol tcp 
```
### results
```shell
+----+--------------------+---------+------------------------+----------------------------+
|  # | ADDRESS            | SUCCESS | SENT                   | RECEIVE                    |
+----+--------------------+---------+------------------------+----------------------------+
|  0 | localhost:10000    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  1 | localhost:10001    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  2 | localhost:10002    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  3 | localhost:10003    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  4 | localhost:10004    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  5 | localhost:10005    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  6 | localhost:10006    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  7 | localhost:10007    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  8 | localhost:10008    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
|  9 | localhost:10009    | true    | GET / HTTP/1.1         | GET / HTTP/1.1             |
|    |                    |         | Host: localhost        | Host: localhost            |
|    |                    |         | User-Agent: aclchecker | User-Agent: aclchec        |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
| 10 | localhost:10011    | false   | GET / HTTP/1.1         | EOF                        |
|    |                    |         | Host: localhost        |                            |
|    |                    |         | User-Agent: aclchecker |                            |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+
| 11 | 10.161.64.77:31971 | true    | GET / HTTP/1.1         | HTTP/1.1 404 Not Found     |
|    |                    |         | Host: localhost        | Date: Thu, 17 Apr 2025 13: |
|    |                    |         | User-Agent: aclchecker |                            |
|    |                    |         | Accept: */*            |                            |
|    |                    |         |                        |                            |
|    |                    |         |                        |                            |
+----+--------------------+---------+------------------------+----------------------------+

```
