library("sf")
library("readxl")
library("mapsf")
library("tidyverse")

# ne marche pas à cause de crs diff et st_transform() ne passe pas :(
#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

bibT <- read_excel(path="/home/guillaume/DONNEES_R/STATUT_BIB.xlsx",
                   sheet=1,
                   col_names=TRUE)

bibC <- st_transform(st_read("/home/guillaume/SHP/adresses-des-bibliotheques-publiques.shp"),2154) 


bib91 <- bibC[bibC$dept == 91,] 

dep91 <- GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,]

x11()

plot(st_geometry(dep91))
#plot(st_geometry(bib91),add=T)

plot(st_geometry(GEOFLA_COMMUNE_2019_l93),add=T)
buff <- st_buffer(bib91,5000)
plot(st_geometry(buff),add=T)
 #ggplot() +
 # geom_sf(data = com91, colour = "grey20") +
 # geom_sf(data = biblio91,
  #        fill = "#FDE725", colour = "#FDE725")

#mf_map(com91,
 #      lwd=3,
  #     add=TRUE)
locator(1)
