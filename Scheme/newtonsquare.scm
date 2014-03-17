;;; puiss: nat -> nat
;;; (puiss x i) rend la puissance de x par i
;;; HYPOTHESE: utilisation de la récursion linéaire plutot que de la rec dicho

(define (puiss x i)
  (if (= i 0)
      1
      (* x (puiss x (- i 1)))))

;;; abs: nbre -> nbre
;;; (abs x) rend la valeur absolue de x

(define (abs x)
  (cond (( = x 0) 0)
        ((< x 0) (- 0 x))
        (else x)))

;;;definition fausse car mal comprise, à revoir...
(define (newton-square x y)
  
         