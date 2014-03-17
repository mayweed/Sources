;;;add-greater: nat*nat*nat -> nat
;;;(add-greater a b c) prd 3 nbres et rend la somme des 2 plus grds.

(define (add-greater a b c)
 (if (and (and (> a b) (> a c))
          (and (< b a) (> b c)))
     (+ a b)
     (if (and (< c a) (> c b))
         (+ a c)
         (+ b c))))
     