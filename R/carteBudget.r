library("sf")
library("mapsf")
library("readxl")
library("maptiles")

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

x11() #use windows() or quartz() for mac

#pdf(file="/home/guillaume/carteBudget.pdf",
#    width=8,
#    height=12,
#    paper="a4") 


#mf_init(GEOFLA_COMMUNE_2019_l93,expandBB=c(0,0,0.1,0))

#mf_background("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/fondEssonne.png")
# ,q='https://tile.openstreetmap.org/${z}/${x}/${y}.png')
# https://rdrr.io/cran/maptiles/man/get_tiles.html
fondOSM <- get_tiles(GEOFLA_COMMUNE_2019_l93,provider='OpenStreetMap.France') #,crop=TRUE,zoom=4)
plot_tiles(fondOSM)

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_map(x=GEOFLA_COMMUNE_2019_l93,
       col=NA,
       var="RATIO",
       type="prop",
       #breaks="quantile",
       leg_pos= "topleft", #NA, #waiting for a solution
       leg_title="Budget par population",
       leg_frame = TRUE,
       add=TRUE,
       lwd=2
)

mf_title(txt = "Budget d’acquisition par habitant")


mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Données issues du rapport SCRIB 2020")

# pour pdf
#dev.off()

#wait please!!
locator(1)
