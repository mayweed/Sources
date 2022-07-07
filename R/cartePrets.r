library("sf")
library("mapsf")
library("readxl")
library("cartography")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : add data
bibT <- read_excel(path="/home/guillaume/DONNEES_R/PRETS_POP.xlsx",
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

#pdf(file="/home/guillaume/carteBudget.pdf",
#    width=8,
#    height=12,
#    paper="a4") 


mf_init(GEOFLA_COMMUNE_2019_l93,expandBB=c(0,0.15,0,0))

#mf_background("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/fondEssonne.png")

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

GEOFLA_COMMUNE_2019_l93$pretstypo <- cut(GEOFLA_COMMUNE_2019_l93$PRETS,
                                         breaks=c(0,1000,5000,10000,50000,100000,250000),
                                         labels=c("0-1000","1000-5000","5000-10000","10000-50000","50000-100000","> 100000"),
                                         include.lowest=TRUE)

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_map(x=GEOFLA_COMMUNE_2019_l93,
       var="pretstypo",
       type="typo",
       #breaks=c(0,1000,5000,10000,50000,100000,250000),
       val_order=c("0-1000","1000-5000","5000-10000","10000-50000","50000-100000","> 100000"),
       pal= c("red","orange","yellow","lightblue","#00CC33","green"),
       leg_no_data="Données non communiquées",
       leg_pos= "topleft", #NA, #waiting for a solution
       leg_title="Prêts",
       leg_frame = TRUE,
       add=TRUE,
       lwd=2
)

mf_title(txt = "Nombre de prêts par commune")

mf_symb(x =GEOFLA_COMMUNE_2019_l93,
        var ="BIB",
        val_order=c("0","1"),
        pch=c(17,26), #26 to 31 are unassigned and that does not work with NA
        leg_pos=NA,
        add=TRUE)

mf_map(x=GEOFLA_EPCI91_2019_l93,
       col=NA,
       border="black",
       add=TRUE,
       lwd=5)

mf_credits(txt="Données issues du rapport SCRIB 2020")

# pour pdf
#dev.off()

#wait please!!
locator(1)
