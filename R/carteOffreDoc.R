library("sf")
library("mapsf")
library("readxl")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
bibT <- read_excel(path="/home/guillaume/DONNEES_R/BIB_DOCS.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibT,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

villeSansBib <- st_read("/home/guillaume/SHP/villeSansBib.shp")
bibAssoc <- st_read("/home/guillaume/SHP/bibAssoc.shp")

x11() #use windows() or quartz() for mac

mf_init(GEOFLA_COMMUNE_2019_l93,expandBB=c(0.1,0.1,0,0.1))

#mf_background("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/fondEssonne.png")

# pas de bib dans la commune
mf_typo(x = villeSansBib,
        var ="INSEE_COM",
        pal= "lightgrey",
        leg_pos=NA,
        add=TRUE)

# bib Assoc
mf_typo(x = bibAssoc,
        var ="INSEE_COM",
        pal= "lightblue",
        leg_pos=NA,
        add=TRUE)

mf_legend_t(title = NA,val = c("Pas de bib.", "Bib. associative"), pal = c("lightgrey", "lightblue"))

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="grey",
       add=TRUE,
       lwd=3)

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_map(x=GEOFLA_COMMUNE_2019_l93,
       col=NA,
       add=TRUE,
       lwd=2
)

mf_prop_choro(
  x = GEOFLA_COMMUNE_2019_l93, var = c("DOCS", "RATIO"), inches = .18, 
  val_max = 90000, symbol = "circle", col_na = "grey", pal = "TealGrn",
  breaks = "equal", nbreaks = 4, 
  leg_pos = c("bottomright", "topleft"),
  leg_title = c("Nombre de documents/bibliothèque", "Nombre de documents/habitant"),
  leg_no_data = "Données non communiquées",
  add = TRUE
)

mf_title(txt = "Offre documentaire des bibliothèques de l’Essonne")

mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

# pour pdf
#dev.off()

#wait please!!
locator(1)
