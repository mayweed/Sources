(define (carre x)
  (* x x))

(define (aucarre L)
  (if (pair? L)
      (list (carre (car L))
            (aucarre (cdr L)))
      0))