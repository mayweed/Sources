;;; est?: alpha * LISTE[alpha] -> bool
;;; (est? x L) permet de déterminer si x est dans L

(define (est? x L)
  (if (pair? L)
      (or (equal? x (car L))
          (est? x (cdr L)))
      #f))
                 
;;;est-voyelles?: quoted -> bool
;;;(est-voyelles? x) prend une lettre et vous dit si c'est une voyelle (ne pas oublier le ' avant!)

(define (est-voyelles? x)
  (let ((voyelles '(a e i o u y)))
        (est? x voyelles)))