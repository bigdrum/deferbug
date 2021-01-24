Sample output:

```
$ go version
go version go1.15.7 darwin/amd64
```

```
$ go test
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> github.com/bigdrum/deferbug.TestDeferBug.func1 (3 handlers)
[GIN] 2021/01/24 - 11:39:22 | 200 |      19.706µs |       192.0.2.1 | GET      "/"
[GIN] 2021/01/24 - 11:39:22 | 200 |      10.544µs |       192.0.2.1 | GET      "/"
[GIN] 2021/01/24 - 11:39:22 | 200 |      45.414µs |       192.0.2.1 | GET      "/"
runtime: unexpected return pc for github.com/gin-gonic/gin.RecoveryWithWriter.func1 called from 0x7160f70
stack: frame={sp:0xc0000f7b70, fp:0xc0000f7bb0} stack=[0xc0000f6000,0xc0000f8000)
000000c0000f7a70:  0000000001c3ffff  0000000007160f70
000000c0000f7a80:  0000000007160f70  0000000000000030
000000c0000f7a90:  0000000000000001  000000c0000f7b10
000000c0000f7aa0:  000000000101a625 <runtime.(*mcentral).cacheSpan+517>  0000000007160f70
000000c0000f7ab0:  0000000000000000  000000c0000f7a78
000000c0000f7ac0:  000000000000006c  000000c0000a4060
000000c0000f7ad0:  0000000001464be0  0000000000000000
000000c0000f7ae0:  0000000000000064  0000000007160f70
000000c0000f7af0:  0000000000000003  0000000000000008
000000c0000f7b00:  0000000000002000  000000c0000f7b30
000000c0000f7b10:  0000000000203000  00000000000000aa
000000c0000f7b20:  0000000001716fb8  0000000000203000
000000c0000f7b30:  0000000000203000  0000000000203000
000000c0000f7b40:  0000000007160f70  0000000000000000
000000c0000f7b50:  000000c000386000  000000c0000f7bc0
000000c0000f7b60:  0000000007160f70  000000c0000f7c08
000000c0000f7b70: <000000000100f650 <runtime.mallocgc+880>  000000c00038e000
000000c0000f7b80:  0000000000000030  0000000000000030
000000c0000f7b90:  0000000001424640  0000000000000001
000000c0000f7ba0:  0001000000000000 !0000000007160f70
000000c0000f7bb0: >0000000000000030  0000000000000030
000000c0000f7bc0:  0000000001901560  000000c00038e000
000000c0000f7bd0:  00000000014cbfc0  000000c000030900
000000c0000f7be0:  0000000000000000  0000000000000001
000000c0000f7bf0:  0000000000000001  000000000000006c
000000c0000f7c00:  0000000000000000  000000c0000f7c38
000000c0000f7c10:  000000000100ffb8 <runtime.newobject+56>  0000000000000030
000000c0000f7c20:  0000000001424640  0000000000000001
000000c0000f7c30:  000000c00038e000  000000c0000f7c80
000000c0000f7c40:  0000000001010e1c <runtime.makemap+412>  0000000001424640
000000c0000f7c50:  000000c0000f7c98  00000000010178ef <runtime.heapBits.forwardOrBoundary+111>
000000c0000f7c60:  0000000001a5c800  0020300000000000
000000c0000f7c70:  0000000001c3ffff  0000000000000400
000000c0000f7c80:  0000000001a5c900  0020300000000000
000000c0000f7c90:  0000000001c3ffff  000000c0000f7d18
000000c0000f7ca0:  00000000010185bb <runtime.heapBits.initSpan+315>  0000000001a5c800
fatal error: unknown caller pc
runtime: unexpected return pc for github.com/gin-gonic/gin.RecoveryWithWriter.func1 called from 0x1d89f30
stack: frame={sp:0xc00003ab70, fp:0xc00003abb0} stack=[0xc00003a000,0xc00003b000)
000000c00003aa70:  000000000000006c  000000000000006c
000000c00003aa80:  000000000000006c  000000c0000a4078
000000c00003aa90:  00000000010994e0 <internal/poll.(*FD).Write.func1+0>  000000c0000a4060
000000c00003aaa0:  000000c0002a8000  000000000000006c
000000c00003aab0:  0000000000000080  000000c00003aa78
000000c00003aac0:  000000000000006c  000000c0000a4060
000000c00003aad0:  0000000001464be0  000000c00003ab58
000000c00003aae0:  000000000109afb7 <os.(*File).Write+119>  000000000107a965 <sync.(*Pool).pin+37>
000000c00003aaf0:  0000000000000002  0000000000000008
000000c00003ab00:  0000000000000080  000000c00003ab30
000000c00003ab10:  0000000000203000  00000000016de7a0
000000c00003ab20:  000000c00009ab00  0000000000203000
000000c00003ab30:  0000000000203000  0000000000203000
000000c00003ab40:  00000000016de7a0  0000000001436900
000000c00003ab50:  000000c0002a4000  000000c00003abc0
000000c00003ab60:  00000000010d6d77 <fmt.Fprint+183>  000000c00003ac08
000000c00003ab70: <000000000100f650 <runtime.mallocgc+880>  000000c000282300
000000c00003ab80:  0000000000000030  0000000000000030
000000c00003ab90:  0000000001424640  0000000000000000
000000c00003aba0:  0000000000000000 !0000000001d89f30
000000c00003abb0: >0000000000000030  0000000000000030
000000c00003abc0:  0000000001900e98  000000c000282300
000000c00003abd0:  00000000014cbfc0  000000c000030d80
000000c00003abe0:  0000000000000000  0000000000000001
000000c00003abf0:  0000000000000001  000000000000006c
000000c00003ac00:  0000000000000000  000000c00003ac38
000000c00003ac10:  000000000100ffb8 <runtime.newobject+56>  0000000000000030
000000c00003ac20:  0000000001424640  0000000000000001
000000c00003ac30:  000000c000282300  000000c00003ac80
000000c00003ac40:  0000000001010e1c <runtime.makemap+412>  0000000001424640
000000c00003ac50:  000000c000282300  000000000104ff4c <runtime.makeslice+108>
000000c00003ac60:  0000000000000000  00000000013c0ea0
000000c00003ac70:  0000000000000000  000000000171f260
000000c00003ac80:  000000c00003ad60  0000000001269ffa <net/http.Header.Clone+378>
000000c00003ac90:  0000000001424280  000000c0001cf4d0
000000c00003aca0:  000000c00003ad00  000000c000282300
fatal error: unknown caller pc
runtime: unexpected return pc for github.com/gin-gonic/gin.RecoveryWithWriter.func1 called from 0x190b308
stack: frame={sp:0xc00018bb70, fp:0xc00018bbb0} stack=[0xc000188000,0xc00018c000)
000000c00018ba70:  000000000000006c  000000000000006c
000000c00018ba80:  000000000000006c  000000c0000a4078
000000c00018ba90:  00000000010994e0 <internal/poll.(*FD).Write.func1+0>  000000c0000a4060
000000c00018baa0:  000000c00031a000  000000000000006c
000000c00018bab0:  00000000000000e0  000000c00018ba78
000000c00018bac0:  000000000000006c  000000c0000a4060
000000c00018bad0:  0000000001464be0  000000c00018bb58
000000c00018bae0:  000000000109afb7 <os.(*File).Write+119>  000000000107a965 <sync.(*Pool).pin+37>
000000c00018baf0:  0000000000000000  0000000000000008
000000c00018bb00:  00000000000000e0  000000c00018bb30
000000c00018bb10:  0000000000203000  00000000016de7a0
000000c00018bb20:  000000c00009aa00  0000000000203000
000000c00018bb30:  0000000000203000  0000000000203000
000000c00018bb40:  00000000016de7a0  0000000001436900
000000c00018bb50:  000000c0000685b0  000000c00018bbc0
000000c00018bb60:  00000000010d6d77 <fmt.Fprint+183>  000000c00018bc08
000000c00018bb70: <000000000100f650 <runtime.mallocgc+880>  000000c000064b40
000000c00018bb80:  0000000000000030  0000000000000030
000000c00018bb90:  0000000001424640  0000000000000000
000000c00018bba0:  0000000000000000 !000000000190b308
000000c00018bbb0: >0000000000000030  0000000000000030
000000c00018bbc0:  0000000001900108  000000c000064b40
000000c00018bbd0:  00000000014cbfc0  00000000016ee8a0
000000c00018bbe0:  0000000000000000  0000000000000001
000000c00018bbf0:  0000000000000001  000000000000006c
000000c00018bc00:  0000000000000000  000000c00018bc38
000000c00018bc10:  000000000100ffb8 <runtime.newobject+56>  0000000000000030
000000c00018bc20:  0000000001424640  0000000000000001
000000c00018bc30:  000000c000064b40  000000c00018bc80
000000c00018bc40:  0000000001010e1c <runtime.makemap+412>  0000000001424640
000000c00018bc50:  000000c000064b40  000000000104ff4c <runtime.makeslice+108>
000000c00018bc60:  0000000000000000  00000000013c0ea0
000000c00018bc70:  0000000000000000  000000000171f260
000000c00018bc80:  000000c00018bd60  0000000001269ffa <net/http.Header.Clone+378>
000000c00018bc90:  0000000001424280  000000c000064a50
000000c00018bca0:  000000c00018bd00  000000c000064b40
fatal error: unknown caller pc

runtime stack:
runtime.throw(0x144b659, 0x11)
	/usr/local/go/src/runtime/panic.go:1116 +0x72
runtime.gentraceback(0x138bea5, 0xc0000f7b70, 0x0, 0xc000001800, 0x0, 0x0, 0x7fffffff, 0x70000a669f18, 0x0, 0x0, ...)
	/usr/local/go/src/runtime/traceback.go:273 +0x1aec
runtime.addOneOpenDeferFrame.func1()
	/usr/local/go/src/runtime/panic.go:721 +0x91
runtime.systemstack(0x1800000)
	/usr/local/go/src/runtime/asm_amd64.s:370 +0x66
runtime.mstart()
	/usr/local/go/src/runtime/proc.go:1116

goroutine 6 [running]:
runtime.systemstack_switch()
	/usr/local/go/src/runtime/asm_amd64.s:330 fp=0xc0000f7e18 sp=0xc0000f7e10 pc=0x106e9c0
runtime.addOneOpenDeferFrame(0xc000001800, 0x0, 0x0)
	/usr/local/go/src/runtime/panic.go:720 +0x7b fp=0xc0000f7e58 sp=0xc0000f7e18 pc=0x1037bbb
panic(0x13c0ea0, 0x14c3dc0)
	/usr/local/go/src/runtime/panic.go:971 +0x417 fp=0xc0000f7f20 sp=0xc0000f7e58 pc=0x10385b7
github.com/bigdrum/deferbug.TestDeferBug.func3.1()
	/Users/bigdrum/build/deferbug/deferbug_test.go:57 +0x5b fp=0xc0000f7f50 sp=0xc0000f7f20 pc=0x138d6db
github.com/bigdrum/deferbug.TestDeferBug.func3(0xc000016640, 0xc00007e280)
	/Users/bigdrum/build/deferbug/deferbug_test.go:58 +0xf4 fp=0xc0000f7fd0 sp=0xc0000f7f50 pc=0x138d7f4
runtime.goexit()
	/usr/local/go/src/runtime/asm_amd64.s:1374 +0x1 fp=0xc0000f7fd8 sp=0xc0000f7fd0 pc=0x1070781
created by github.com/bigdrum/deferbug.TestDeferBug
	/Users/bigdrum/build/deferbug/deferbug_test.go:50 +0x12a

goroutine 1 [chan receive]:
testing.(*T).Run(0xc000001680, 0x144953e, 0xc, 0x1463ff0, 0x108e766)
	/usr/local/go/src/testing/testing.go:1169 +0x2da
testing.runTests.func1(0xc000001500)
	/usr/local/go/src/testing/testing.go:1439 +0x78
testing.tRunner(0xc000001500, 0xc00021fde0)
	/usr/local/go/src/testing/testing.go:1123 +0xef
testing.runTests(0xc00000e320, 0x16dd720, 0x1, 0x1, 0xbffb87249d391ff8, 0x8bb3027341, 0x16ee1c0, 0x100f650)
	/usr/local/go/src/testing/testing.go:1437 +0x2fe
testing.(*M).Run(0xc000066100, 0x0)
	/usr/local/go/src/testing/testing.go:1345 +0x1eb
main.main()
	_testmain.go:43 +0x138

goroutine 5 [semacquire]:
sync.runtime_Semacquire(0xc000016648)
	/usr/local/go/src/runtime/sema.go:56 +0x45
sync.(*WaitGroup).Wait(0xc000016640)
	/usr/local/go/src/sync/waitgroup.go:130 +0x65
github.com/bigdrum/deferbug.TestDeferBug(0xc000001680)
	/Users/bigdrum/build/deferbug/deferbug_test.go:67 +0x146
testing.tRunner(0xc000001680, 0x1463ff0)
	/usr/local/go/src/testing/testing.go:1123 +0xef
created by testing.(*T).Run
	/usr/local/go/src/testing/testing.go:1168 +0x2b3

goroutine 7 [running]:
	goroutine running on other thread; stack unavailable
created by github.com/bigdrum/deferbug.TestDeferBug
	/Users/bigdrum/build/deferbug/deferbug_test.go:50 +0x12a

goroutine 8 [running]:
	goroutine running on other thread; stack unavailable
created by github.com/bigdrum/deferbug.TestDeferBug
	/Users/bigdrum/build/deferbug/deferbug_test.go:50 +0x12a

runtime stack:
runtime.throw(0x144b659, 0x11)
	/usr/local/go/src/runtime/panic.go:1116 +0x72
runtime.gentraceback(0x138bea5, 0xc00018bb70, 0x0, 0xc000001b00, 0x0, 0x0, 0x7fffffff, 0x70000a6ecf18, 0x0, 0x0, ...)
	/usr/local/go/src/runtime/traceback.go:273 +0x1aec
runtime.addOneOpenDeferFrame.func1()
	/usr/local/go/src/runtime/panic.go:721 +0x91
runtime.systemstack(0x1800000)
	/usr/local/go/src/runtime/asm_amd64.s:370 +0x66
runtime.mstart()
	/usr/local/go/src/runtime/proc.go:1116

runtime stack:
runtime.throw(0x144b659, 0x11)
	/usr/local/go/src/runtime/panic.go:1116 +0x72
runtime.gentraceback(0x138bea5, 0xc00003ab70, 0x0, 0xc000001980, 0x0, 0x0, 0x7fffffff, 0x70000a76ff18, 0x0, 0x0, ...)
	/usr/local/go/src/runtime/traceback.go:273 +0x1aec
runtime.addOneOpenDeferFrame.func1()
	/usr/local/go/src/runtime/panic.go:721 +0x91
runtime.systemstack(0x1e00000)
	/usr/local/go/src/runtime/asm_amd64.s:370 +0x66
runtime.mstart()
	/usr/local/go/src/runtime/proc.go:1116
exit status 2
FAIL	github.com/bigdrum/deferbug	0.198s
```