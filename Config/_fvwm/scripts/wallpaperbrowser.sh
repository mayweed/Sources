#!/bin/bash
# $1 = répertoire des wallpapers
# Ce script vient de la config de Taviso. Je l'ai légèrement modifié pour que
# ça marche avec les wallpapers dont les noms contiennent des espaces ou
# d'autres caractères spéciaux.

test ! -d "$1/.thumbs" && mkdir "$1/.thumbs";
  for i in "$1/"*; do
    test -f "$1/.thumbs/${i##*/}" -a "${i}" -ot "$1/.thumbs/${i##*/}" ||
      { convert -quality 0 -scale 64 "${i}" "png:$1/.thumbs/${i##*/}" 2>/dev/null || 
        continue;};
  done;
fvwm-menu-directory --title "$(basename "$1")" \
  --icon-title gnome-folder.png --icon-file __PIXMAP__ \
  --icon-dir  gnome-folder.png  --dir "$1" --exec-file "^fbsetbg -f" \
  --exec-t="^gthumb ." | \
  sed -e "s#FuncFvwmMenuDirectory#WallpaperBrowser#g" \
    -e "s#__PIXMAP__\(.*\)\"\(.*/\)\(.*\)\"#\2.thumbs/\3\1\"\2\3\"#g" \
    -e "s#item +100 c##g"
