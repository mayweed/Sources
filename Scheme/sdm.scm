;; sdm: nbres -> un nbre
;; (sdm n m) prend 2 nombres et renvoie ...
(define (sdm n m)
  (if ( = n 0)
      0
      (+ (* m n)
         (sdm (- n 1)m))))
