library("sf")
library("mapsf")
library("readxl")

#get the maps
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")

bibP <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/PRETS_USAGERS_2020.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibP,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

#  mf_get_breaks(x=bibP$USAGERS,nbreaks=5,breaks="equal")
#[1]      0  35238  70476 105714 140952 176190

#  mf_get_breaks(x=bibP$PRETS,nbreaks=5,breaks="equal")
#[1]    0 1168 2336 3504 4672 5840

#GEOFLA_COMMUNE_2019_l93$pretstypo <- cut(GEOFLA_COMMUNE_2019_l93$PRETS,
#                                         breaks=c(0,5000,10000,20000,300000),
#                                        labels = c("0-5000","5000-10000","10000-20000","> 20000"),
#                                         include.lowest=TRUE)
# mf_typo(x=GEOFLA_COMMUNE_2019_l93,
#        var="pretstypo",
#        pal=c("red","orange","yellow","green"),
#        val_order=c("0-5000","5000-10000","10000-20000","> 20000"),
#        leg_pos="topleft",
#        leg_title="Prêt en milliers de documents",
#)

x11() #use windows() or quartz() for maic

mf_map(GEOFLA_COMMUNE_2019_l93)

mf_prop_choro(
  x = GEOFLA_COMMUNE_2019_l93, var = c("USAGERS", "PRETS"), inches = .35, border = "black",
  val_max = 250000, symbol = "circle", col_na = "grey", pal = "Green-Orange",
  breaks = "equal", nbreaks = 8, lwd = 2,
  leg_pos = c("bottomright", "topleft"),
  leg_title = c("Nombre d’usagers", "Nombre de prêts"),
  leg_title_cex = c(0.8, 1),
  leg_val_cex = c(.7, .9),
  leg_val_rnd = c(0, 0),
  leg_no_data = "Pas de données",
  leg_frame = c(TRUE, TRUE),
  add = TRUE
)


mf_title(txt="Ratio prêts de documents/usagers par ville")

#instead of plot, won’t work
#mf_map(x=GEOFLA_EPCI91_2019_l93,add=TRUE,lwd=3)
#plot(st_geometry(GEOFLA_EPCI91_2019_l93),add=TRUE,lwd=3)
#instead of labelLayer() won’t work…
#mf_label(x=GEOFLA_EPCI91_2019_l93,
#         var="LIBEPCI",
#         cex=0.8,
#         overlap=FALSE,
#         lines=FALSE,
#)

#wait please!!
locator(1)
