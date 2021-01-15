"https://github.com/sheoak/sheoak-neovim
let s:vimpath = "~/.config/nvim/"

if has('nvim')
    let s:plug_path=$HOME . '/.local/share/nvim/plugged'
else
    let s:plug_path=$HOME . '/.vim/plugged'
endif

"my general settings
execute "source" . s:vimpath . "settings.vim"

"my plugins
call plug#begin(s:plug_path)
    execute "source " . s:vimpath . "plugin.vim"
call plug#end()

"my life-saving bepo shortcuts
execute "source" . s:vimpath . "bepo.vim"
 
"some time you need mappings
execute "source" . s:vimpath . "mappings.vim"

" must put that in a special file!!
" this is vim-go only…
au FileType go set noexpandtab
au FileType go set shiftwidth=4
au FileType go set softtabstop=4
au FileType go set tabstop=4

"" Run goimports when running gofmt
let g:go_fmt_command = "goimports"
"
"" Enable syntax highlighting per default
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1
let g:go_highlight_structs = 1
let g:go_highlight_operators = 1
let g:go_highlight_build_constraints = 1
let g:go_highlight_extra_types = 1

" gopls
let g:go_def_mode='gopls'
let g:go_info_mode='gopls'

""mappings
map <C-n> :cnext<CR>
map <C-m> :cprevious<CR>
