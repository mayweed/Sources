library("sf")
library("mapsf")
library("readxl")

#get everything 
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
bibT <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/BUDGET_POP.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibT,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

villeSansBib <- st_read("C:/Users/Raimondeaug/Documents/SHP/villeSansBib.shp")

x11() #use windows() or quartz() for mac

mf_init(GEOFLA_COMMUNE_2019_l93)

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="grey",
       add=TRUE,
       lwd=3)

mf_map(x=GEOFLA_COMMUNE_2019_l93,
       col=NA,
       add=TRUE,
       lwd=2)

# pas de bib dans la commune
mf_typo(x =villeSansBib,
        var ="INSEE_COM",
        pal=c("lightgrey"),
        leg_pos=NA,
        add=TRUE)

mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_prop_choro(
  x = GEOFLA_COMMUNE_2019_l93, var = c("BUDGET", "RATIO"), inches = .18, 
  val_max = 90000, symbol = "circle", pal = "Viridis",
  breaks = "equal", nbreaks = 4, 
  leg_pos = c("bottomright", "topleft"),
  leg_title = c("Budget/bibliothèque (en €)", "Budget/habitant (en €)"),
  leg_no_data = "Données non communiquées",
  add = TRUE
)

mf_title(txt = "Budget d’acquisition 2020 (en €)")

mf_legend_t(val = "Pas de bib.",pal="lightgrey",pos="topright",title=NA)

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

#décommenter pour obtenir une image png
#png("C:/Users/Raimondeaug/Documents/CARTES_R/carteBudget.png",width=800,height=800,res=100)
#wait please
locator(1)
