library("sf")
library("readxl")
library("mapsf")
library("tidyverse")

# ne marche pas à cause de crs diff et st_transform() ne passe pas :(
#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

bibT <- read_excel(path="/home/guillaume/DONNEES_R/NO_BIB.xlsx",
                   sheet=1,
                   col_names=TRUE)
#zb <- GEOFLA_COMMUNE_2019_l93 %>%
 #   filter merge(GEOFLA_COMMUNE_2019_l93,
  #                               bibT,
   #                              by.x="INSEE_COM",
    #                             by.y="INSEE_COM",
     #                            all.x=TRUE)
#pas merge: filter!! pour ne garder que les lignes de geofla en zb

bibC <- st_transform(st_read("/home/guillaume/SHP/adresses-des-bibliotheques-publiques.shp"),2154) 


bib91 <- bibC[bibC$dept == 91,] 

dep91 <- GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,]

x11()

plot(st_geometry(dep91))

plot(st_geometry(GEOFLA_COMMUNE_2019_l93),add=T)
plot(st_geometry(GEOFLA_EPCI91_2019_l93),add=T)

buff <- st_buffer(bib91,5000)
plot(st_geometry(buff),add=T)
#plot(GEOFLA_COMMUNE_2019_l93[GEOFLA_COMMUNE_2019_l93$NO_BIB,])
locator(1)
