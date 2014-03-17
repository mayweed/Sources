;;;convert: nbre -> nbre
;;;(convert t) prend une température en °C et renvoie l'équivalent Fahrenheit

(define (convert t)
      (+ ( / (* 9 t) 5)
         32))
      
;;;celsius2f : LISTE[alpha] -> LISTE[alpha]
;;; (celsius2f L) prend une liste en °C et en renvoie une en Fahrenheit
;;;HYPOTHESE: on n'utilise pas map

(define (celsius2f L)
  (if (pair? L)
       (cons (convert (car L))
             (celsius2f (cdr L)))
       '()))

;;;la même avec map
;(map convert '(0 10 20)

;;; minliste: LISTE[nbre] -> nbre
;;; (minliste L) prend une liste renvoie son mini, si vide alors 100000
(define (minliste L)
  ;;ne pas oublier ceci sinon marche po évidemment
  ;;on procède par réduction de la liste jusqu'a ce qu'il n'en reste plus qu'un
  (if (pair? L)
  (if (= (length L) 1)
      (car L)
      (if(< (car L) (minliste (cdr L)))
         (car L)
         (minliste (cdr L))))
  100000))

;;;min celsius avec reduce
(define (mincelsius L)
  (reduce minliste 100000 (map convert '(L))))