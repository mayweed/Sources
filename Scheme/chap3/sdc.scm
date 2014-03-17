;;; sdc : nat -> nat
;;; (sdc x) renvoie sdc x

(define (sdc x)
  (if (= x 0)
      0
      (+ (remainder x 10)
         (sdc (quotient x 10)))))
         

;;; sds : nat -> nat
;;; (sds x) renvoie la somme de la somme de x

(define (sds y)
  (if (= y 0)
      0
      (sdc (sdc y))))