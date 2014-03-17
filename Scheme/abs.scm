;;; abs: nbre -> nbre
;;; (abs x) rend la valeur absolue de x

(define (abs x)
  (cond (( = x 0) 0)
        ((< x 0) (- 0 x))
        (else x)))