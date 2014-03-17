;;; nbre ->/int/ nbre
;;; sum renvoie la somme des nombres entre eux jusqu'à n

(define (sum n)
  (if (= n 0)
      0
  (+ n (sum ( - n 1)))))