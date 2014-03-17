 ;--------------- template for recursion  (bored with factorial :)--------
;
(define [remove-negatives numbers]
  (cond
    ;; Base case: No numbers provided, so return empty list
    [(empty? numbers)
     empty]
    ;; Recursive case: First number is negative, so return a
    ;; list constructed the rest of the
    ;; numbers
    [(negative? (first numbers))
     (remove-negatives (rest numbers))]
    ;; Recursive case: First number is non-negative, so
    ;; return a list constructed from this number followed by
    ;; the rest of the numbers with negatives removed
    [else
     (cons (first numbers) (remove-negatives (rest numbers)))]))

