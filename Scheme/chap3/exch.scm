;;; existe-chiffres : nat -> bool
;;; (ex-ch? x y) vérifie que x est dans y

(define (ex-ch? x y)
  (if (< y 10)
      ;; x étant forcément inf à 10 on peut direct comparer à x...
      (= x y)
      (or (= (remainder y 10) x)
      (ex-ch? x (quotient y 10)))))