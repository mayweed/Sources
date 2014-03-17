(define (sdm3 x)
  (if (= x 1)
      3
      (+ (* 3 x) 
         (si ( - x 1)))))