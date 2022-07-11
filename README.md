


```shell
rm -rf kitex_gen 

go mod init

kx example.thrift

go get github.com/cloudwego/kitex
```


==================== reqs=10000 arrLen=10
======== RT total:  5.490286579s
======== RT total:  5.040514972s
======== RT total:  4.540713133s

==================== reqs=10000 arrLen=10
======== RT total:  10.335694455s
======== RT total:  11.44689794s
======== RT total:  8.62609145s
