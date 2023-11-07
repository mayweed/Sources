library("sf")
library("mapsf")
library("readxl")

#get everything 
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : add data
bibPrets <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/PRETS_POP.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibPrets,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

villeSansBib <- st_read("C:/Users/Raimondeaug/Documents/SHP/villeSansBib.shp")

x11() #use windows() or quartz() for mac

png("C:/Users/Raimondeaug/Documents/CARTES_R/cartePrets.png",width=800,height=800,res=100)

mf_init(GEOFLA_COMMUNE_2019_l93,expandBB=c(0,0.15,0,0))

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_choro(x=GEOFLA_COMMUNE_2019_l93,
         var="RATIO",
         leg_title = "Taux d’emprunt par habitant",
         leg_no_data="Données non communiquées",
         add=T)

# pas de bib dans la commune
mf_typo(x = villeSansBib,
        var ="INSEE_COM",
        pal= "lightgrey",
        leg_pos=NA,
        add=TRUE)

mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=5)

mf_map(x=GEOFLA_COMMUNE_2019_l93,
       var="PRETS",
       type="prop",
       inches=.20,
       col="yellow",
       symbol="circle",
       leg_pos="bottomright",
       leg_title="Nombre de prêts par bib.",
       add=T)

mf_title(txt = "Prêts par bibliothèque et taux d’emprunt/hab. (hors collectivités)")

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

mf_legend_t(title = NA,val = c("Pas de bib."), pal = "lightgrey")

# pour pdf
#dev.off()

#wait please!!
locator(1)
