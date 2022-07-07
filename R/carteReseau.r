library("sf")
library("mapsf")
library("readxl")
library("cartography")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
bibT <- read_excel(path="/home/guillaume/DONNEES_R/STATUT_BIB.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibT,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

x11() #use windows() or quartz() for mac

pdf(file="/home/guillaume/carteReseau.pdf",
    width=8,
    height=12,
    paper="a4") 

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
     border="black",
     lwd=3)

mf_typo(x=GEOFLA_COMMUNE_2019_l93,
        var="TRANSFERT",
        pal = c("pink","lightgreen","lightblue"),
        val_order=c("Non-transférée","Transférée","Bib. associative"),
        leg_pos="bottomright",
        leg_title="Statut des bibs",
        leg_frame = TRUE,
        leg_no_data = "Pas de données",
        add=TRUE
)

# pch_na look at : http://www.sthda.com/french/wiki/les-differents-types-de-points-dans-r-comment-utiliser-pch
mf_symb(
  x =GEOFLA_COMMUNE_2019_l93, var ="TYPOLOGIE", pch = c(21:23), pal = c("red","yellow", "tan1", "#990066","white"),
  border = "grey20", cex = c(1.5, 1, .9,.7,.5), lwd = .5,
  val_order = c("B1", "B2", "B3","B4","B5"),
  #pch_na = 18,#
  leg_no_data="Pas de bibliothèque",col_na= "black",leg_frame = TRUE,leg_pos = "topleft",
  leg_title = "Typologie des bibs"
)

mf_title(txt = "Réseau de lecture publique du département de l’Essonne")


#plot(st_geometry(GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,]),
#     border="black",
#     add=TRUE,
#     lwd=3)
mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Données issues du rapport SCRIB 2020")

#plot(st_geometry(GEOFLA_EPCI91_2019_l93),
#     border="black",
#     add=TRUE,
#     lwd=5)
dev.off()
#labelLayer(x=GEOFLA_EPCI91_2019_l93,txt="EPCI",font=4)
#mf_export(GEOFLA_COMMUNE_2019_l93,filename=test.png)
#wait please!!
locator(1)
