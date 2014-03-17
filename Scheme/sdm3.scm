(define (sdm3 x)
  (if ( = x 1)
      3
      (+ ( * 3 x)
         (sdm3 ( - x 3)))))
      