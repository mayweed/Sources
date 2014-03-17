;;; last-elt : LISTE [NombreNat] -> NbreNat
;;; (last-elt L) prd une liste en entrée et renvoie son dernier elt.

(define (last-elt L)
  (if (pair? L)
      ;; le cas de base est pour L et non cdr L
      ;; (pair? (list 2) donne #t mais (pair? (list)) donne faux donc sur (cdr L)
      ;; on a toujours 0...cf trace...
      (if (equal? (cdr L) () )
          (list (car L))
      (last-elt (cdr L)))
      0))