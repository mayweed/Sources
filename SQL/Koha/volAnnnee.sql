SELECT YEAR(dateaccessioned) as 'Année d’acquisition',count(*) as 'Total'
FROM items 
GROUP BY YEAR(dateaccessioned)
