


```shell
rm -rf kitex_gen 

go mod init

kx example.thrift

go get github.com/daidai21/kitex@issue_433__benchmark
```


==================== reqs=10000 arrLen=10
======== RT total:  7.790511724s
======== RT total:  4.47568792s
======== RT total:  3.720081286s

==================== reqs=10000 arrLen=10
======== RT total:  2.755217676s
======== RT total:  10.982978524s
======== RT total:  11.531568287s
======== RT total:  6.905616514s
