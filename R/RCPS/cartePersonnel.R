library("sf")
library("mapsf")
library("readxl")

#get everything 
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
bibT <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/PERSONNEL_BÉNÉVOLES.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibT,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

bibAssoc <- st_read("C:/Users/Raimondeaug/Documents/SHP/bibAssoc.shp")
villeSansBib <- st_read("C:/Users/Raimondeaug/Documents/SHP/villeSansBib.shp")

x11() #use windows() or quartz() for mac

png("C:/Users/Raimondeaug/Documents/CARTES_R/cartePersonnel.png",width=800,height=800,res=100)

mf_init(GEOFLA_COMMUNE_2019_l93,expandBB=c(0.1,0,0,0.1))

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="grey",
       add=TRUE,
       lwd=3)

mf_map(x=GEOFLA_COMMUNE_2019_l93,
       col=NA,
       add=TRUE,
       lwd=1)

mf_typo(x = GEOFLA_COMMUNE_2019_l93, var = c("NB"),
        val_order = c("Oui","Non"),pal=c("green","orange"),
        leg_pos=NA,
        add=TRUE)


mf_typo(x =villeSansBib,
        var ="INSEE_COM",
        pal= c("lightgrey"),
        leg_pos=NA,
        add=TRUE)

mf_typo(x =bibAssoc,
        var ="INSEE_COM",
        pal= c("lightblue"),
        leg_pos=NA,
        add=TRUE)
mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=2)


mf_map(x = GEOFLA_COMMUNE_2019_l93, var = c("SALARIES"),type="prop",inches = .20, 
       leg_pos = c("topleft"),
       col = c("yellow"),
       leg_title = c("Nombre de salariés par commune"),
       add = TRUE)

mf_legend_t(title="Présence de bénévoles",pos="bottomright",
            val=c("Oui","Non","Bib. Associative","Pas de bib.","Pas de données communiqués"),
            pal=c("green","orange","lightblue","lightgrey","white"))

mf_title(txt = "Personnel dans les bibliothèques")

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

# pour pdf
#dev.off()

#wait please!!
locator(1)
