library("sf")
library("mapsf")
library("readxl")
library("cartography")

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

#pdf(file="/home/guillaume/carteSigb.pdf",
#    width=8,
#    height=12,
#    paper="a4") 

mf_map(x=GEOFLA_DEP_2019_l93[GEOFLA_DEP_2019_l93$CODE_DEPT == 91,],
     border="black",
     lwd=3)

mf_init(x=GEOFLA_COMMUNE_2019_l93,
        expandBB=c(0.2,0.3,0,0))

# http://www.sthda.com/french/wiki/couleurs-dans-r
mf_typo(x=GEOFLA_COMMUNE_2019_l93,
        var="SIGB",
        pal = c("#00FFCC","lightgreen","lightblue","pink","yellow","#FF0033","#FF9966","grey","#CC0099","#3399CC","#33FF99","#3300FF","#FF6600","#FF3399","#FF99FF","#669933"),
        val_order=c("Orphée NX","PMB","Syracuse","Orphée.net 3.3","AFI Nanook 4.3.5","SIGB DECALOG","Atalante","BiblixNet","V-Smart","PERGAME v03.12.4","Paprika CS2","BGP 4.20","X-Theques ASSISTERE","Andosace v98.2","Registar AMJ 3.98","Agate 2.02"),
        leg_pos= "topleft", #NA, #waiting for a solution
        leg_title="SIGB utilisé",
        leg_frame = TRUE,
        leg_no_data = "Non communiqué",
        add=TRUE
)

mf_symb(x =GEOFLA_COMMUNE_2019_l93,
        var ="BIB",
        val_order=c("0","1"),
        pch=c(17,26), #26 to 31 are unassigned and that does not work with NA
        leg_pos=NA,
        add=TRUE)

mf_title(txt = "Liste des SIGB utilisés")


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
