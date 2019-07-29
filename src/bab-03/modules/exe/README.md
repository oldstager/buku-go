# Tentang

Direktori ini berisi kode sumber yang telah menggunakan fasilitas *modules* dari Go 1.12. Untuk
keperluan ini, inisialisasi terlebih dahulu menggunakan:

```bash
$ go mod init github.com/oldstager/go-to-hell-o
$ cat go.mod
module github.com/oldstager/go-to-hell-o

go 1.12
$
```

Setelah itu baru buat kode sumber hello.go. Untuk mengkompilasi:

```bash
$ go build
go: finding rsc.io/quote v1.5.2
go: downloading rsc.io/quote v1.5.2
go: extracting rsc.io/quote v1.5.2
go: finding rsc.io/sampler v1.3.0
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: downloading rsc.io/sampler v1.3.0
go: extracting rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: extracting golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
$
```

Hasil:

```bash
$ ls -la
total 2296
drwxr-xr-x 2 bpdp bpdp    4096 Jul 28 06:55 ./
drwxr-xr-x 4 bpdp bpdp    4096 Jul 28 06:45 ../
-rw------- 1 bpdp bpdp      79 Jul 28 06:55 go.mod
-rw------- 1 bpdp bpdp     499 Jul 28 06:55 go.sum
-rwxr-xr-x 1 bpdp bpdp 2327743 Jul 28 06:55 go-to-hell-o*
-rw-r--r-- 1 bpdp bpdp      93 Jul 28 06:54 hello.go
$
```

Module yang sudah di-*get* dan di-*build* berada di:

```bash
$ tree -L 3 ~/go/pkg/mod/
/home/bpdp/go/pkg/mod/
├── cache
│   ├── download
│   │   ├── golang.org
│   │   └── rsc.io
│   ├── lock
│   └── vcs
│       ├── 0c8659d2f971b567bc9bd6644073413a1534735b75ea8a6f1d4ee4121f78fa5b
│       ├── 0c8659d2f971b567bc9bd6644073413a1534735b75ea8a6f1d4ee4121f78fa5b.info
│       ├── 0c8659d2f971b567bc9bd6644073413a1534735b75ea8a6f1d4ee4121f78fa5b.lock
│       ├── 4db0c9594744360b0eaa452d2ccfbd45b05dffb9810882957d10d69e61e66382
│       ├── 4db0c9594744360b0eaa452d2ccfbd45b05dffb9810882957d10d69e61e66382.info
│       ├── 4db0c9594744360b0eaa452d2ccfbd45b05dffb9810882957d10d69e61e66382.lock
│       ├── 5b03666c2d7b526129bad48c5cea095aad8b83badc1daa202e7b0279e3a5d861
│       ├── 5b03666c2d7b526129bad48c5cea095aad8b83badc1daa202e7b0279e3a5d861.info
│       └── 5b03666c2d7b526129bad48c5cea095aad8b83badc1daa202e7b0279e3a5d861.lock
├── golang.org
│   └── x
│       └── text@v0.0.0-20170915032832-14c0d48ead0c
└── rsc.io
    ├── quote@v1.5.2
    │   ├── buggy
    │   ├── go.mod
    │   ├── LICENSE
    │   ├── quote.go
    │   ├── quote_test.go
    │   └── README.md
    └── sampler@v1.3.0
        ├── glass.go
        ├── glass_test.go
        ├── go.mod
        ├── hello.go
        ├── hello_test.go
        ├── LICENSE
        └── sampler.go

15 directories, 19 files
$
```

