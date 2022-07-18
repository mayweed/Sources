library("sf")
library("readxl")
library("mapsf")

# ne marche pas à cause de crs diff et st_transform() ne passe pas :(
#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

bibT <- read_excel(path="/home/guillaume/DONNEES_R/NO_BIB.xlsx",
                   sheet=1,
                   col_names=TRUE)

bibC <- st_transform(st_read("/home/guillaume/SHP/adresses-des-bibliotheques-publiques.shp"),2154) 


bib91 <- bibC[bibC$dept == 91,] 

dep91 <- GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,]

zb <- st_read("/home/guillaume/SHP/zoneBlanche.shp")
x11()

plot(st_geometry(dep91))

mf_map(GEOFLA_COMMUNE_2019_l93)
#plot(st_geometry(GEOFLA_COMMUNE_2019_l93),add=T)
#plot(st_geometry(GEOFLA_EPCI91_2019_l93),add=T)

buff <- st_buffer(bib91,5000)

plot(st_geometry(buff),add=T)

mf_typo(x=zb,
        var="INSEE_COM",
        pal="red",
        leg_pos=NA,
        add=T)

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="rightbottom")

mf_title(txt="Villes n’ayant pas de bibliothèque dans un rayon de 5km")
locator(1)
