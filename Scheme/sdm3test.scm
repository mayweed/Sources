(define (sdm3 x)
  (let (( n ( * x 3)))
    (if ( = n 0)
        0
        (+ n
           (sdm3 ( - n 3))))))