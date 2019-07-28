set number
set linebreak
set showbreak=+++
set textwidth=100
set showmatch
set nocompatible

set hlsearch
set smartcase
set ignorecase
set incsearch
 
set autoindent
set expandtab
set shiftwidth=2
set smartindent
set smarttab
set softtabstop=2
 
set ruler
 
set undolevels=1000
set backspace=indent,eol,start

filetype plugin indent on
set ofu=syntaxcomplete#Complete

" Use 24-bit (true-color) mode in Vim/Neovim when outside tmux or screen.
" If you're using tmux version 2.2 or later, you can remove the outermost $TMUX
" check and use tmux's 24-bit color support
" (http://sunaku.github.io/tmux-24bit-color.html#usage for more information.)
if empty($TMUX) && empty($STY)
  " See https://gist.github.com/XVilka/8346728.
  if $COLORTERM =~# 'truecolor' || $COLORTERM =~# '24bit'
    if has('termguicolors')
      " See :help xterm-true-color
      if $TERM =~# '^screen'
        let &t_8f = "\<Esc>[38;2;%lu;%lu;%lum"
        let &t_8b = "\<Esc>[48;2;%lu;%lu;%lum"
      endif
      set termguicolors
    endif
  endif
endif

set background=dark
colorscheme open-color
syntax on
highlight LineNr term=bold cterm=NONE ctermfg=DarkGrey ctermbg=NONE gui=NONE guifg=DarkGrey guibg=darkgreen
set cursorline
" set cursorcolumn

set guifont=Monospace\ 14

" nerdtree
let g:NERDTreeWinPos = "right"
autocmd bufenter * if (winnr("$") == 1 && exists("b:NERDTree") && b:NERDTree.isTabTree()) | q | endif
let gRNERDTreeNodeDelimiter = "\u00a0"
nnoremap <F4> :NERDTreeToggle<CR>  
let g:NERDTreeFileExtensionHighlightFullName = 1
let g:NERDTreeExactMatchHighlightFullName = 1
let g:NERDTreePatternMatchHighlightFullName = 1
let g:NERDTreeHighlightFolders = 1 " enables folder icon highlighting using exact match
let g:NERDTreeHighlightFoldersFullName = 1 " highlights the folder name
let NERDTreeShowHidden=1

" airline
let g:airline_powerline_fonts = 1
let g:airline_theme='distinguished'
