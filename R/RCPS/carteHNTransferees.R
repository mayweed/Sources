library("sf")
library("mapsf")
library("readxl")
library("png")

#get everything 
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : add data
bibH <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/BIB_HORAIRES.xlsx",
                   sheet=1,
                   col_names=TRUE)

bibT <- st_read("C:/Users/Raimondeaug/Documents/SHP/bibTransferee.shp")
bibNT <- st_read("C:/Users/Raimondeaug/Documents/SHP/bibNonTransferee.shp")
villeSansBib <- st_read("C:/Users/Raimondeaug/Documents/SHP/villeSansBib.shp")

bibNT <- merge(bibNT,
              bibH,
              by.x="INSEE_COM",
              by.y="INSEE_COM",
              all.x=TRUE)


x11() #use windows() or quartz() for mac

png("C:/Users/Raimondeaug/Documents/CARTES_R/carteHNTransferees.png",width=800,height=800,res=100)

mf_init(GEOFLA_COMMUNE_2019_l93,expandBB=c(0,0.15,0,0))

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_map(x=GEOFLA_COMMUNE_2019_l93,
       col=NA,
       add=T)

mf_choro(x=bibNT,
         var="SEM",
         pal= "RdPu",
         leg_pos="topleft",
         leg_title=NA,
         leg_no_data="Données non communiquées",
         add=T)

# les transférées 
mf_typo(x = bibT,
        var ="INSEE_COM",
        pal= "lightblue",
        leg_pos=NA,
        add=TRUE)

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


mf_title(txt = "Nombre d’heures d’ouverture par semaine des bib. non transférées")

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

mf_legend_t(title = NA,val = c("Bib. transférée","Pas de bib."), pal = c("lightblue","lightgrey"))

# pour pdf
dev.off()

#wait please!!
#locator(1)
