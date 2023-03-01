library("sf")
library("mapsf")
library("readxl")

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : search for ods
bibT <- read_excel(path="/home/guillaume/DONNEES_R/BIB_SIGB.xlsx",
                   sheet=1,
                   col_names=TRUE)

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibT,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

x11() #use windows() or quartz() for mac

png("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/focusSigbCPS.png",width=800,height=800,res=100)

target <- GEOFLA_EPCI91_2019_l93[1,]
mf_init(target,expandBB=c(0,0.15,0,0))

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_typo(x= GEOFLA_COMMUNE_2019_l93[ GEOFLA_COMMUNE_2019_l93$EPCI == "200056232",] ,
        var="SIGB",
        pal = c("#00FFCC","yellow","blue","#CC6600","#FF0033","purple","lightgrey","#669933"),
        val_order=c("Orphée NX","PMB","Syracuse","Orphée.net 3.3","SIGB DECALOG","BiblixNet","Paprika CS2","Agate 2.02"),
        leg_title=NA,
        leg_pos= "topleft",
        leg_no_data = "Non communiqué",
        add=TRUE)

mf_label(x=GEOFLA_COMMUNE_2019_l93[ GEOFLA_COMMUNE_2019_l93$EPCI == "200056232",] ,
         var="LIBGEO",
         #TODO:changer cette police et la mettre en plus gros
         overlap=FALSE)

mf_title(txt = "Agglo’ Paris-Saclay - Système Informatisé de Gestion de Bibliothèques")


mf_map(x=GEOFLA_EPCI91_2019_l93[1,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_credits(txt="Réalisation: G. Raimondeau - Agglo’ Paris-Saclay",pos="rightbottom")

# pour pdf
dev.off()

#wait please!!
#locator(1)
