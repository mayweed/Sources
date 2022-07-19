library("sf")
library("mapsf")
library("readxl")
#library("cartography")

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

villeSansBib <- st_read("/home/guillaume/SHP/villeSansBib.shp")

x11() #use windows() or quartz() for mac
png("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/focusSigbCPS.png",width=800,height=800,res=100)

target <- GEOFLA_EPCI91_2019_l93[1,]
mf_init(target,expandBB=c(0,0.15,0,0))

#mf_init(x=GEOFLA_COMMUNE_2019_l93,
 #       expandBB=c(0.2,0.3,0,0))

#mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
#     border="black",
#     add=TRUE,
#     lwd=3)


# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_typo(x= GEOFLA_COMMUNE_2019_l93[ GEOFLA_COMMUNE_2019_l93$EPCI == "200056232",] ,
        var="SIGB",
        pal = c("#00FFCC","lightgreen","blue","pink","#FF0033","purple","#33FF99","#FF99FF","#669933"),
        val_order=c("Orphée NX","PMB","Syracuse","Orphée.net 3.3","SIGB DECALOG","BiblixNet","Paprika CS2","Registar AMJ 3.98","Agate 2.02"),
        leg_pos= "topleft", #NA, #waiting for a solution
        leg_title="SIGB utilisé",
        leg_no_data = "Non communiqué",
        add=TRUE
)

# pas de bib dans la commune
#mf_typo(x = villeSansBib,
#        var ="INSEE_COM",
#        pal= "lightgrey",
#        leg_pos=NA,
#        add=TRUE)

mf_title(txt = "Liste des SIGB utilisés - Paris-Saclay")


mf_map(x=GEOFLA_EPCI91_2019_l93[1,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

#mf_legend_t(title = NA,val = c("Pas de bib."), pal = c("lightgrey"),pos="bottomleft")

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="rightbottom")

# pour pdf
dev.off()

#wait please!!
#locator(1)
