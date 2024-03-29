library("sf")
library("mapsf")
library("readxl")

#get everything 
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")

# bib SHP
bibAdresse <- st_read("C:/Users/Raimondeaug/Documents/SHP/adresses-des-bibliotheques-publiques.shp")
bibT <- st_read("C:/Users/Raimondeaug/Documents/SHP/bibTransferee.shp")
bibNT <- st_read("C:/Users/Raimondeaug/Documents/SHP/bibNonTransferee.shp")
bibAssoc <- st_read("C:/Users/Raimondeaug/Documents/SHP/bibAssoc.shp")
bibTiersLieu <- st_read("C:/Users/Raimondeaug/Documents/SHP/bibTiersLieu.shp")
villeSansBib <- st_read("C:/Users/Raimondeaug/Documents/SHP/villeSansBib.shp")

# typologie ABD 
bibTypo <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/STATUT_BIB.xlsx",
                   sheet=1,
                   col_names=TRUE)

GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibTypo,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

x11() #use windows() or quartz() for mac

png("C:/Users/Raimondeaug/Documents/CARTES_R/carteRéseau.png",width=800,height=800,res=100)

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
     border="black",
     lwd=3)

mf_typo(x=bibT,
        var="INSEE_COM",
        pal = c("lightgreen"),
        leg_pos = NA,
        add=TRUE)

mf_typo(x=bibNT,
        var="INSEE_COM",
        pal = c("pink"),
        leg_pos = NA,
        add=TRUE)

mf_typo(x=bibAssoc,
        var="INSEE_COM",
        pal = c("lightblue"),
        leg_pos = NA,
        add=TRUE)

mf_typo(x=bibTiersLieu,
        var="INSEE_COM",
        pal = c("orange"),
        leg_pos = NA,
        add=TRUE)

mf_typo(x=villeSansBib,
        var="INSEE_COM",
        pal = c("lightgrey"),
        leg_pos = NA,
        add=TRUE)

mf_legend_t(pal = c("pink","lightgreen","lightblue","orange","lightgrey"),
            val=c("Bib. non-transférée","bib. transférée","Bib. associative","Bib. Tiers-lieu", "Pas de bibliothèque"),
            pos="bottomright",
            title=NA ) #"Statut des bibs")

# pch_na look at : http://www.sthda.com/french/wiki/les-differents-types-de-points-dans-r-comment-utiliser-pch
mf_symb(x =GEOFLA_COMMUNE_2019_l93, 
        var ="TYPOLOGIE", 
        val_order = c("B1", "B2", "B3","B4","B5"),
        pal = c("red","yellow", "tan1", "#990066","white"),
        pch = c(21), 
        border = "grey20", cex = c(1.5, 1, .9,.7,.5), lwd = .5,
        leg_pos = "topleft",
        leg_title = NA,
        leg_no_data="Pas de données")

mf_title(txt = "Typologie du réseau de lecture publique de l’Essonne")


mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

#wait please!!
locator(1)
