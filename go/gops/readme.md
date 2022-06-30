
## 安装

## 使用

### gops

```
gops
703  588  gopls                        go1.18.1 /Users/yangtao/go/bin/gopls
3999 3998 ___go_build_GoNotes_go_gops* go1.18.1 /private/var/folders/02/9nw63rhx5n9bmj2ybxyhs8rw0000gn/T/GoLand/___go_build_GoNotes_go_gops
4329 756  gops                         go1.18.1 /Users/yangtao/go/bin/gops
```

#### 表头

- PID
- PPID
- Name of the program
- Go version used to build the program
- Location of the associated program

### gops <pid> [duration]

```
gops 703
parent PID:     588
threads:        17
memory usage:   2.386%
cpu usage:      0.149%
cpu usage (0s): NaN%
username:       yangtao
cmd+args:       /Users/yangtao/go/bin/gopls -mode=stdio
elapsed time:   01-02:05:21
```

```
gops 3998
parent PID:     3998
threads:        6
memory usage:   0.017%
cpu usage:      0.000%
cpu usage (0s): NaN%
username:       yangtao
cmd+args:       /private/var/folders/02/9nw63rhx5n9bmj2ybxyhs8rw0000gn/T/GoLand/___go_build_GoNotes_go_gops
elapsed time:   03:32
local/remote:   127.0.0.1:52979 <-> :0 (LISTEN)
```

### gops tree

```
gops tree
...
├── 3998
│   └── [*]  3999 (___go_build_GoNotes_go_gops) {go1.18.1}
├── 756
│   └── 4404 (gops) {go1.18.1}
└── 588
    └── 703 (gopls) {go1.18.1}
```

### gops stack (<pid>|<addr>)

    仅支持启动了 agent 的

```
gops stack 3999
goroutine 35 [running]:
runtime/pprof.writeGoroutineStacks({0x11a5470, 0xc000010010})
        /usr/local/go/src/runtime/pprof/pprof.go:694 +0x93
runtime/pprof.writeGoroutine({0x11a5470, 0xc000010010}, 0x2)
        /usr/local/go/src/runtime/pprof/pprof.go:683 +0x46
runtime/pprof.(*Profile).WriteTo(0x1254240, {0x11a5470, 0xc000010010}, 0x2)
        /usr/local/go/src/runtime/pprof/pprof.go:332 +0xa7
github.com/google/gops/agent.handle({0x17850c8, 0xc000010010}, {0xc00012e238, 0x1, 0x1})
        /Users/yangtao/go/pkg/mod/github.com/google/gops@v0.3.23/agent/agent.go:201 +0x11c
github.com/google/gops/agent.listen()
        /Users/yangtao/go/pkg/mod/github.com/google/gops@v0.3.23/agent/agent.go:145 +0x47e
created by github.com/google/gops/agent.Listen
        /Users/yangtao/go/pkg/mod/github.com/google/gops@v0.3.23/agent/agent.go:123 +0x765

goroutine 1 [sleep]:
time.Sleep(0x34630b8a000)
        /usr/local/go/src/runtime/time.go:194 +0x11b
main.main()
        /Users/yangtao/IdeaProjects/an/demo/job-interview-notes/go/gops/main.go:13 +0xfc
```

### gops memstats (<pid>|<addr>)

    仅支持启动了 agent 的

```
gops memstats 3999
alloc: 1.17MB (1226248 bytes)
total-alloc: 1.17MB (1226248 bytes)
sys: 9.70MB (10175752 bytes)
lookups: 0
mallocs: 382
frees: 8
heap-alloc: 1.17MB (1226248 bytes)
heap-sys: 3.66MB (3833856 bytes)
heap-idle: 2.04MB (2138112 bytes)
heap-in-use: 1.62MB (1695744 bytes)
heap-released: 2.01MB (2105344 bytes)
heap-objects: 374
stack-in-use: 352.00KB (360448 bytes)
stack-sys: 352.00KB (360448 bytes)
stack-mspan-inuse: 27.76KB (28424 bytes)
stack-mspan-sys: 31.88KB (32640 bytes)
stack-mcache-inuse: 14.06KB (14400 bytes)
stack-mcache-sys: 15.23KB (15600 bytes)
other-sys: 797.62KB (816764 bytes)
gc-sys: 3.50MB (3672248 bytes)
next-gc: when heap-alloc >= 4.00MB (4194304 bytes)
last-gc: -
gc-pause-total: 0s
gc-pause: 0
gc-pause-end: 0
num-gc: 0
num-forced-gc: 0
gc-cpu-fraction: 0
enable-gc: true
debug-gc: false
```

### gops gc (<pid>|<addr>)

    仅支持启动了 agent 的

### gops setgc (<pid>|<addr>)

    仅支持启动了 agent 的

```
gops setgc 3999 10
New GC percent set to 10. Previous value was 100.
```

### gops version (<pid>|<addr>)

    仅支持启动了 agent 的

```
gops version 3999 
go1.18.1
```

### gops stats (<pid>|<addr>)

    仅支持启动了 agent 的

```
gops stats 3999
goroutines: 2
OS threads: 10
GOMAXPROCS: 12
num CPU: 12
```

### Pprof

    仅支持启动了 agent 的

```
gops pprof-cpu (<pid>|<addr>)
gops pprof-heap (<pid>|<addr>) 
```

```
gops pprof-cpu 3999
Profiling CPU now, will take 30 secs...
Profile dump saved to: /var/folders/02/9nw63rhx5n9bmj2ybxyhs8rw0000gn/T/cpu_profile3997823624
Binary file saved to: /var/folders/02/9nw63rhx5n9bmj2ybxyhs8rw0000gn/T/binary883836728
File: binary883836728
Type: cpu
Time: Jun 23, 2022 at 9:53am (CST)
Duration: 30s, Total samples = 0 
No samples were found with the default sample value type.
Try "sample_index" command to analyze different sample values.
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) 
(pprof) web
(pprof) q
```

### gops trace (<pid>|<addr>)

    仅支持启动了 agent 的

```
gops trace 3999          
Tracing now, will take 5 secs...
Trace dump saved to: /var/folders/02/9nw63rhx5n9bmj2ybxyhs8rw0000gn/T/trace3332793368
2022/06/23 09:55:19 Parsing trace...
2022/06/23 09:55:19 Splitting trace...
2022/06/23 09:55:19 Opening browser. Trace viewer is listening on http://127.0.0.1:53093
```
