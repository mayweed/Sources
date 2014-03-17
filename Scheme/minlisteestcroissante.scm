(define (minliste L)
  ;;ne pas oublier ceci sinon marche po évidemment
  ;;on procède par réduction de la liste jusqu'a ce qu'il n'en reste plus qu'un
  (if (pair? L)
  (if (= (length L) 1)
      (car L)
      (if(< (car L) (minliste (cdr L)))
         (car L)
         (minliste (cdr L))))
  100000))

(define (estcroissante? L)
  (if (<= (length L) 1)
      #t
      (if (< (car L) (cadr L))
          (estcroissante? (cdr L))
          #f)))
      