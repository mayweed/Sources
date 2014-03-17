;;; est?: alpha * LISTE[alpha] -> bool
;;; (est? x L) permet de déterminer si x est dans L

(define (est? x L)
  (if (pair? L)
      (or (equal? x (car L))
          (est? x (cdr L)))
      #f))
                 
;;;est-voyelles?: quoted -> bool
;;;(est-voyelles? x) prend une lettre et vous dit si c'est une voyelle (ne pas oublier le ' avant!)

(define (est-voyelles? x)
  (let ((voyelles '(a e i o u y)))
        (est? x voyelles)))

;;; nblettres: liste[alpha] -> nbre
;;; (nblettres L) prend une liste et renvoie son nbre d'élts.
;;; HYPOTHESE: ces éléments sont des lettres...

(define (nblettres L)
  (if (pair? L)
      (+ 1 (nblettres (cdr L)))
      0))

;;; nbvoyelles : LISTE[alpha] -> nbre
;;; (nbvoyelles L) renvoie le nbre de voyelles dans les elts de la liste.

(define (nbvoyelles L)
  (if (pair? L)
      (let ((voyelles '(a e i o u y)))
        (if (est? (car L) voyelles)
            (-(nblettres L) 1)
            (nbvoyelles (cdr L))))
      0))

;;;nbconsonnes: LISTE[alpha] -> nbre
;;; (nbconsonnes L) renvoie le nbre de consonnes parmi les elts de L

(define (nbconsonnes L)
  (if (pair? L)
      (- (nblettres L) (nbvoyelles L))
      0))

;;;listevoyelles : LISTE[alpha] -> LISTE[alpha]
;;; (listevoyelles L) renvoie sous forme de liste les voyelles de L
;;; HYPOTHESE: non-utilisation de filtre...
(define (listevoyelles L)
  (if (pair? L)
      (if (est-voyelles? (car L))
          (cons (car L)
                (listevoyelles (cdr L)))
          (listevoyelles (cdr L)))
          '()))

;;; avec filtre
; (filtre est-voyelle? L)

