library("sf")
library("mapsf")
library("readxl")
#library("maptiles")
library("png")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
bibT <- read_excel(path="/home/guillaume/DONNEES_R/USAGERS_TX.xlsx", sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibT,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

x11() #use windows() or quartz() for mac

mf_init(GEOFLA_COMMUNE_2019_l93)

#mf_background("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/fondEssonne.png")
# ,q='https://tile.openstreetmap.org/${z}/${x}/${y}.png')
# https://rdrr.io/cran/maptiles/man/get_tiles.html
# OKI : this is epsg 2154 we need 3857
# cf https://github.com/riatelab/maptiles/issues/5#issuecomment-810237958
#com3857 <- st_transform(GEOFLA_COMMUNE_2019_l93,3857)
#tiles3857 <- get_tiles(x=com3857,zoom=10)
#plot_tiles(tiles3857,add=TRUE)

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_map(x=GEOFLA_COMMUNE_2019_l93,
       col=NA,
       add=TRUE,
       lwd=2
)

mf_prop_choro(
  x = GEOFLA_COMMUNE_2019_l93, var = c("INSCRITS", "TX"), inches = .5, 
  val_max = 90000, symbol = "circle", col_na = "grey", pal = "Viridis",
  breaks = "equal", nbreaks = 4, 
  leg_pos = c("bottomright", "topleft"),
  leg_title = c("Inscrits/bibliothèque", "Taux/habitant"),
  leg_no_data = "Données non communiquées",
  leg_frame = c(TRUE, TRUE),
  add = TRUE
)

mf_title(txt = "Inscrits par commune")


mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="rightbottom")

# EXPORT
#png("txusagers.png")
#pdf(file="/home/guillaume/carteBudget.pdf",
#    width=8,
#    height=12,
#    paper="a4") 

#dev.off()

#wait please!!
locator(1)
