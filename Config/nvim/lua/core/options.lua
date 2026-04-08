local opt = vim.opt

-- numéros de ligne
opt.relativenumber = true --afficher nl relative à la pos actuelle du curseur
opt.nu = true

-- syntax related
opt.encoding="utf-8"
opt.background = "dark" 
opt.termguicolors = true
syntax = true
opt.cursorline = true
opt.fileformat=unix --avoid the ^M or ^V bullsh*t

-- backup, undo, swap
opt.backup = false
opt.writebackup = false
opt.swapfile = false
opt.undofile = true

-- tabs & indentation
opt.expandtab=true -- spaces instead of tabs
opt.tabstop=4 -- a tab = four spaces
opt.shiftwidth=4 -- number of spaces for auto-indent
opt.softtabstop=4 -- let backspace delete indent
opt.textwidth=85
opt.autoindent = true
opt.smartindent = true

-- recherche
opt.ignorecase = true -- ignore la casse qd on cherche
opt.smartcase = true -- sauf qd recherche en maj on rebascule en sensible à la casse
opt.hlsearch = true -- surligne les occurrences des recherches

-- gestion des fenêtres
opt.splitbelow = true -- split hz en bas
opt.splitright = true -- split vert à droite

opt.showmode = true
