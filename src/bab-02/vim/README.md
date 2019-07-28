# Struktur direktori pada Vim

Berikut ini adalah direktori yang diperlukan jika menggunakan fasilitas `native packaging` dari Vim
(mulai versi 8.0). $HOME adalah direktori home dari user yang bersangkutan (/home/namauser). Pada
direktori `$HOME/.vim/pack/default/start` letakkan berbagai repo yang diperlukan (bisa diperoleh
dari GitHub).


```bash
$HOME/.vim
├── pack
│   └── default
│       └── start
│           ├── nerdtree
│           ├── nerdtree-git-plugin
│           ├── vim-airline
│           ├── vim-airline-themes
│           ├── vim-go
│           ├── vim-open-color
│           └── vim-polyglot
└── vimrc
```
