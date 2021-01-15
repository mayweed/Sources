"https://github.com/sheoak/sheoak-neovim
let s:vimpath = "~/.config/nvim/"

if has('nvim')
    let s:plug_path=$HOME . '/.local/share/nvim/plugged'
else
    let s:plug_path=$HOME . '/.vim/plugged'
endif

call plug#begin(s:plug_path)
    execute "source " . s:vimpath . "plugin.vim"
call plug#end()

execute "source" . s:vimpath . "bepo.vim"
