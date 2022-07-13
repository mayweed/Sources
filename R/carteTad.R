library("sf")
library("mapsf")
library("readxl")
library("cartography")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")
tad <-st_read("/home/guillaume/SHP/TAD_MDE.shp")

x11() #use windows() or quartz() for mac

mf_init(x=GEOFLA_COMMUNE_2019_l93)

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
     border="black",
     add=TRUE,
     lwd=3)

mf_typo(x=tad,
        var = "TAD",
        val_order = c("SUD-EST","NORD-OUEST","NORD-EST","SUD-OUEST"),
        pal=c("lightblue","pink","yellow","red"),
        add=T)

mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_title(txt = "Les territoires d’action départementale de la MDE")



#wait please!!
locator(1)
