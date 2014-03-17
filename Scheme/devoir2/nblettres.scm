;;; nblettres: liste[alpha] -> nbre
;;; (nblettres L) prend une liste et renvoie son nbre d'élts.
;;; HYPOTHESE: ces éléments sont des lettres...

(define (nblettres L)
  (if (pair? L)
      (+ 1 (nblettres (cdr L)))
      0))