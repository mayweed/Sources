library("sf")
library("mapsf")
library("readxl")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

# bib SHP
bibAdresse <- st_read("/home/guillaume/SHP/adresses-des-bibliotheques-publiques.shp")
bibT <- st_read("/home/guillaume/SHP/bibTransferee.shp")
bibNT <- st_read("/home/guillaume/SHP/bibNonTransferee.shp")
bibAssoc <- st_read("/home/guillaume/SHP/bibAssoc.shp")
bibTiersLieu <- st_read("/home/guillaume/SHP/bibTiersLieu.shp")
villeSansBib <- st_read("/home/guillaume/SHP/villeSansBib.shp")

# typologie ABD 
bibTypo <- read_excel(path="/home/guillaume/DONNEES_R/STATUT_BIB.xlsx",
                   sheet=1,
                   col_names=TRUE)

GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibTypo,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

x11() #use windows() or quartz() for mac

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
            val=c("Non-transférée","Transférée","Bib. associative","Bib. Tiers-lieu", "Pas de bibliothèque"),
            pos="bottomright",
            title="Statut des bibs")

# pch_na look at : http://www.sthda.com/french/wiki/les-differents-types-de-points-dans-r-comment-utiliser-pch
mf_symb(
  x =GEOFLA_COMMUNE_2019_l93, var ="TYPOLOGIE", pch = c(21), pal = c("red","yellow", "tan1", "#990066","white"),
  border = "grey20", cex = c(1.5, 1, .9,.7,.5), lwd = .5,
  val_order = c("B1", "B2", "B3","B4","B5"),
  leg_pos = "topleft",
  leg_title = "Typologie des bibs"
)

mf_title(txt = "Réseau de lecture publique du département de l’Essonne")


mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Données issues du rapport SCRIB 2020")

#wait please!!
locator(1)
