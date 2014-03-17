;;; puiss: nat -> nat
;;; (puiss x i) rend la puissance de x par i
;;; HYPOTHESE: utilisation de la récursion linéaire plutot que de la rec dicho

(define (puiss x i)
  (if (= i 0)
      1
      (* x (puiss x (- i 1)))))