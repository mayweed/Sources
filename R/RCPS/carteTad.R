library("sf")
library("mapsf")
library("readxl")

#get everything 
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")
tad <-st_read("C:/Users/Raimondeaug/Documents/SHP/TAD_MDE.shp")

x11() #use windows() or quartz() for mac

png("C:/Users/Raimondeaug/Documents/CARTES_R/carteTad.png",width=800,height=800,res=100)

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

#mf_label(x=GEOFLA_EPCI91_2019_l93,
#         var="EPCI",
#         col="black")

mf_title(txt = "Les territoires d’action départementale de la MDE")

#wait please!!
locator(1)
