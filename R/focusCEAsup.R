library("sf")
library("mapsf")
library("readxl")
library("png")

# POur mémoire
#> GEOFLA_EPCI91_2019_l93[8-11]
#   CODE_DEPT NOM_DEPT                          EPCI        TAD
#1         91  ESSONNE               CA Paris Saclay Nord-Ouest
#2         91  ESSONNE            CA Grand Paris Sud        Est
#3         91  ESSONNE           CC du Val d'Essonne    Sud-Est
#4         91  ESSONNE      CA Versailles Grand Parc Nord-Ouest
#5         91  ESSONNE CC du Dourdannais en Hurepoix  Sud-Ouest
#6         91  ESSONNE       CC de l'Orée de la Brie   Nord-Est
#7         91  ESSONNE       Grand-Orly Seine Bièvre       Nord
#8         91  ESSONNE         CC du Pays de Limours Nord-Ouest
#9         91  ESSONNE Coeur d'Essonne Agglomération     Centre
#10        91  ESSONNE  CA de l'Etampois Sud Essonne  Sud-Ouest
#11        91  ESSONNE  CA Val d'Yerres Val de Seine   Nord-Est
#12        91  ESSONNE     CC Entre Juine et Renarde  Sud-Ouest
#13        91  ESSONNE           CC des Deux Vallées    Sud-Est

# ttes les communes de cea
# GEOFLA_COMMUNE_2019_l93[GEOFLA_COMMUNE_2019_l93$EPCI == "200057859",]

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : add data
bibSup <- read_excel(path="/home/guillaume/DONNEES_R/BIB_SUPERFICIE.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibSup,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

villeSansBib <- st_read("/home/guillaume/SHP/villeSansBib.shp")

x11() #use windows() or quartz() for mac
png("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/focusCEA.png",width=800,height=800,res=100)

target <- GEOFLA_EPCI91_2019_l93[9,]
mf_init(target,expandBB=c(0,0.15,0,0))

#mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
#       col=NA,
#       border="black",
#       add=TRUE,
#       lwd=3)

mf_choro(x= GEOFLA_COMMUNE_2019_l93[GEOFLA_COMMUNE_2019_l93$EPCI == "200057859",],
         var="RATIO",
         pal="Sunset",
         leg_title = "Superficie par habitant",
         leg_no_data="Données non communiquées",
         add=T)

# pas de bib dans la commune
#mf_typo(x = villeSansBib,
 #       var ="INSEE_COM",
 ##       pal= "lightgrey",
  #      leg_pos=NA,
   #     add=TRUE)

mf_map(x=GEOFLA_EPCI91_2019_l93[9,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=5)

mf_map(x= GEOFLA_COMMUNE_2019_l93[GEOFLA_COMMUNE_2019_l93$EPCI == "200057859",],
       var="SUP",
       type="prop",
       inches=.20,
       col="lightblue",
       symbol="circle",
       leg_pos="bottomright",
       leg_title="Superficie bib./commune",
       add=T)

mf_title(txt = "Superficie des bib. par commune et superficie/hab. CEA")

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

#mf_legend_t(title = NA,val = c("Pas de bib."), pal = "lightgrey")

# pour pdf
dev.off()

#wait please!!
#locator(1)
