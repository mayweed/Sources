library("sf")
library("mapsf")
library("readxl")
library("ggpubr")

#get everything 
load("C:/Users/Raimondeaug/Documents/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
#bibT <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/STATS_NAVETTE.xlsx",
 #                  sheet=1,
  #                 col_names=TRUE)

#all is very important for NA!!
#GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
#                                 bibT,
#                                 by.x="INSEE_COM",
#                                 by.y="INSEE_COM",
#                                 all.x=TRUE)

x11() #use windows() or quartz() for mac

#png("C:/Users/Raimondeaug/Documents/carteCPS.png",width=600,height=600,res=100)

target <- GEOFLA_EPCI91_2019_l93[1,]
mf_init(target,expandBB=c(0,0.15,0,0))

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_map(x=GEOFLA_EPCI91_2019_l93[1,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_map(x=GEOFLA_COMMUNE_2019_l93[ GEOFLA_COMMUNE_2019_l93$EPCI == "200056232",] ,
       col=NA,
       border="black",
       overlap=FALSE)

mf_label(x=GEOFLA_COMMUNE_2019_l93[ GEOFLA_COMMUNE_2019_l93$EPCI == "200056232",] ,
         var="LIBGEO",
#         font("var", face="bold"),
         cex=.6,
         overlap=FALSE)


mf_title(txt = "Agglo’ Paris-Saclay")
mf_credits(txt="Réalisation: G. Raimondeau",pos="rightbottom")

# pour pdf
dev.off()

#wait please!!
locator(1)
