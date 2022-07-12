library("sf")
library("mapsf")
library("readxl")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
bibT <- read_excel(path="/home/guillaume/DONNEES_R/BUDGET_POP.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibT,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

bibS <- read_excel(path="/home/guillaume/DONNEES_R/BIB_SIGB.xlsx",
                   sheet=1,
                   col_names=TRUE)
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibS,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

x11() #use windows() or quartz() for mac


mf_init(GEOFLA_COMMUNE_2019_l93)

#mf_background("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/fondEssonne.png")

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="grey",
       add=TRUE,
       lwd=3)

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_map(x=GEOFLA_COMMUNE_2019_l93,
       col=NA,
       add=TRUE,
       lwd=2)

# pas de bib dans la commune
mf_typo(x =GEOFLA_COMMUNE_2019_l93,
        var ="BIB",
        val_order=c("0","1"),
        #pch=c(4,26), #26 to 31 are unassigned and that does not work with NA
        pal=c("lightgrey","white"),
        col="black",
        leg_pos=NA,
        add=TRUE)

mf_prop_choro(
  x = GEOFLA_COMMUNE_2019_l93, var = c("BUDGET", "RATIO"), inches = .18, 
  val_max = 90000, symbol = "circle", col_na = "grey", pal = "Viridis",
  breaks = "equal", nbreaks = 4, 
  leg_pos = c("bottomright", "topleft"),
  leg_title = c("Budget/bibliothèque (en €)", "Budget/habitant (en €)"),
  leg_no_data = "Données non communiquées",
  leg_frame = c(TRUE, TRUE),
  add = TRUE
)

mf_title(txt = "Budget d’acquisition 2020 (en €)")


mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Données issues du rapport SCRIB 2020")
#wait please!!
locator(1)
