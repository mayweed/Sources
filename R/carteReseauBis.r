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

com91 <- GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,]

biblio91 <- com91 %>%
 st_intersection(bib91)

biblio91 <- merge(biblio91,
                  bibT,
                  by.x="insee",
                  by.y="INSEE_COM",
                  all.x=TRUE)

x11()

mf_init(com91)
mf_typo(x=biblio91,
        var="TRANSFERT",
        pal = c("pink","lightgreen","lightblue"),
        val_order=c("Non-transférée","Transférée","Bib. associative"),
        leg_pos="bottomright",
        leg_title="Statut des bibs",
        leg_frame = TRUE,
        leg_no_data = "Pas de données",
        add=TRUE
)

 #ggplot() +
 # geom_sf(data = com91, colour = "grey20") +
 # geom_sf(data = biblio91,
  #        fill = "#FDE725", colour = "#FDE725")

#mf_map(com91,
 #      lwd=3,
  #     add=TRUE)
