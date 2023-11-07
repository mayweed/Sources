library("sf")
library("mapsf")
library("readxl")

bibS <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/BIB_SIGB.xlsx",
                   sheet=1,
                   col_names=TRUE)

bibT <- st_read("C:/Users/Raimondeaug/Documents/SHP/TRANSFERT_BIBLIOTHEQUES.shp")

#all is very important for NA!!
bibM <- merge(bibT,
              bibS,
              by.x="INSEE",
              by.y="INSEE_COM",
              all.x=TRUE)


x11() #use windows() or quartz() for mac

#décommenter pour obtenir une image png
#
png("C:/Users/Raimondeaug/Documents/CARTES_R/CPS/focusSigbCPS.png",width=800,height=800,res=100)

target <- bibM
mf_init(target,expandBB=c(0,0.15,0,0))

mf_theme("candy",fg = "black", pos = "center", font = 1, tab = FALSE,bg="white")

# http://www.sthda.com/french/wiki/couleurs-dans-r et https://htmlcolorcodes.com/
mf_typo(x= bibM ,
        var="SIGB",
        pal = c("#E0FFFF","#FFD700","lightblue","#E1C16E","#f1807e","#CBC3E3","lightgrey"),
        val_order=c("Orphée NX","PMB","Syracuse","Orphée.net 3.3","SIGB DECALOG","BiblixNet","Paprika CS2"),
        leg_title=NA,
        leg_pos= "bottomleft",
        leg_no_data = "Non communiqué",
        add=TRUE)

mf_label(x=bibM ,
         var="nom",
         #TODO:changer cette police et la mettre en plus gros
         overlap=FALSE)

mf_title(txt = "Logiciels de gestion de bibliothèque")


mf_map(x=bibM,
       col=NA,
       border="black",
       add=TRUE,
       lwd=1)

# question : vaut il mieux mettre les limites des villes en gras ou leur nom?
mf_map(bibM,
       col=NA,
       border="black",
       add=TRUE,
       lwd=2)

mf_credits(txt="Réalisation: G. Raimondeau - Agglo’ Paris-Saclay",pos="rightbottom")

#pour pdf
#dev.off()

#wait please!!
locator(1)
