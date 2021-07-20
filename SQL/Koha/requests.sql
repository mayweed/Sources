/*toutes les requetes faites pendant le stage théma */

//nbre d’ouvrages par fond et par année d’acq
SELECT v.lib as Fonds,count(*) as 'Total'
FROM items i left join authorised_values v on i.ccode=v.authorised_value
where v.category='CCODE'
and YEAR(i.dateaccessioned)=<<année>>
GROUP BY v.lib,YEAR(dateaccessioned)

//nbre de prêts par an et par structure
SELECT c.description as 'Type d’usager',
b.surname AS 'Bibliothèque',
i.itype,
/*case month(s.datetime)
when 1 then 'Janvier'
when 2 then 'Février'
when 3 then 'Mars'
when 4 then 'Avril'
when 5 then 'Mai'
when 6 then 'Juin'
when 7 then 'Juillet'
when 8 then 'Août'
when 9 then 'Septembre'
when 10 then 'Octobre'
when 11 then 'Novembre'
when 12 then 'Décembre'
end as Mois ,*/
count(distinct s.itemnumber)as Total
FROM  statistics s 
inner join borrowers b on s.borrowernumber=b.borrowernumber
left join items i on s.itemnumber=i.itemnumber
left join deleteditems d on s.itemnumber=d.itemnumber
left join categories c on b.categorycode=c.categorycode
where s.type='issue'
and year(s.datetime) = <<année>> 
and b.categorycode != 'PERS_BDY' 
AND b.categorycode !='RELIURE'
AND b.categorycode !='BDY'
GROUP BY b.categorycode, b.surname,i.itype/*,b.surname,month(s.datetime)s.borrowernumber,YEAR(s.datetime)*/
order by b.surname /*Total desc b.categorycode,b.surname,YEAR(s.datetime)*/

// Nombre d’ouvrages non sortis par année d’acq et par fonds
select year(dateaccessioned) as 'Année d’acquisition',ccode as Fonds,count(*) as Total
from items
where issues=0
group by year(dateaccessioned),ccode

//Leaderboard des prêts par année
select i.itemnumber,
i.location,
i.ccode,
b.title,
b.part_number,
count(*) as Total
from items i
left join statistics s on s.itemnumber=i.itemnumber
left join biblio b on i.biblionumber=b.biblionumber
where s.type='issue' 
and year(s.datetime)='2020'
group by i.ccode,b.title
order by Total desc
limit 10

//tx de renouvellement
select t.*,t.Total2020/t.Total as 'Taux de renouvellement'
from
(select v.lib as Fonds,
sum(if (i.dateaccessioned < '2020-12-31',1,0)) as 'Total', /*count(*) is inaccurate, should not count the ones accessioned after the 31/12/2021…*/
sum(case when i.dateaccessioned between '2020-01-01' and '2020-12-31' then 1 else 0 end) as 'Total2020'
from items i left join authorised_values v on i.ccode=v.authorised_value
where v.category='CCODE'
group by v.lib)t

//pilon par année et par fonds
select year(timestamp),v.lib,count(*)
from deleteditems i left join authorised_values v on i.ccode=v.authorised_value
where v.category='CCODE'
group by year(timestamp),v.lib

//tx de rotation
SELECT t.*,t.Pret/t.Total AS 'Taux de rotation'
FROM
(SELECT /*i.location as Section,*/
v.lib AS Fonds,
COUNT(DISTINCT if( year(i.dateaccessioned)<year(s.datetime), i.itemnumber,NULL)) AS Total,
sum(case when s.type = 'issue' and year(s.datetime)=<<année>> then 1 else 0 end) AS Pret
FROM items i 
LEFT JOIN statistics s ON i.itemnumber=s.itemnumber
LEFT JOIN authorised_values v ON i.ccode=v.authorised_value
where v.category='CCODE'
GROUP BY /*i.location,*/v.lib) t

//resa par année et par bib (incomplete)
select b.categorycode as "Code bib.",
b.surname as "Bibliothèque",
sum(case when ro.cancellationdate is null and it.onloan is not null then 1 else 0 end) as "Nombre de réservations",
sum(case when ro.waitingdate is not null and ro.cancellationdate is null then 1 else 0 end) as "Nbre de résas en attente",
sum(case when ro.cancellationdate is not null then 1 else 0 end) as "Nbre de résas annulées",
/*sum(case when s.type='issue' and s.datetime between ro.reservedate and ro.expirationdate then 1 else 0 end) as "Nbre de prêts"*/
sum(case when i.issuedate between ro.reservedate and ro.expirationdate then 1 else 0 end) as "Nbre de prêts"
from old_reserves ro
left join borrowers b on b.borrowernumber=ro.borrowernumber
left join reserves r on r.reserve_id=ro.reserve_id
left join issues i on i.itemnumber=ro.itemnumber and i.borrowernumber=ro.borrowernumber
left join items it on it.itemnumber=ro.itemnumber /*and it.borrowernumber=ro.borrowernumber*/
/*left join statistics s on s.itemnumber=ro.itemnumber and s.borrowernumber=ro.borrowernumber*/
where year(ro.reservedate) = <<Année>>
and b.categorycode != 'BDY'
group by year(ro.reservedate),b.surname
order by year(ro.reservedate),b.categorycode,b.surname

//cout moyen d’un doc par fonds
SELECT v.lib as Fonds, 
sum(i.price) as 'Prix Agrégé',
count(*) as 'Nb total d’ex.',
sum(i.price) div count(*) as 'Prix moyen (en €)'
from items i left join authorised_values v on i.ccode=v.authorised_value
group by v.lib

//nbre d’ouvrages et de prets par cote dewey
SELECT CONCAT(SUBSTRING(i.itemcallnumber,1,1), '00') AS "Cote Dewey",
COUNT(DISTINCT i.itemnumber) AS "Nbre d’ouvrages",
sum(case when s.type = 'issue' and year(s.datetime)=<<année>> then 1 else 0 end) AS "Pret"
FROM items i left join statistics s on s.itemnumber=i.itemnumber
WHERE SUBSTRING(i.itemcallnumber,1,1) REGEXP '^[0-9].*'
AND i.itemlost = '0' 
AND i.damaged ='0'
AND i.ccode='FDS_DOC'
GROUP BY SUBSTRING(i.itemcallnumber,1,1)
ORDER BY SUBSTRING(i.itemcallnumber,1,1) ASC

//nbre d’usagers par type
select c.description as "Catégorie Usager",
count(*) as "Total"
from borrowers b
left join categories c on b.categorycode=c.categorycode
group by c.description


