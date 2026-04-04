library("sf")
library("mapsf")
library("readxl")

bibT <- st_read("C:/Users/Raimondeaug/Documents/SHP/TRANSFERT_BIBLIOTHEQUES.shp")

x11() #use windows() or quartz() for mac

png("C:/Users/Raimondeaug/Documents/CARTES_R/CPS/carteCPS.png",width=600,height=600,res=100)

target <- bibT
mf_init(target,expandBB=c(0,0.15,0,0))
mf_theme("candy",fg = "black", pos = "center", font = 1, tab = FALSE,bg="white")
mf_map(x=bibT,
       col=NA,
       border="black",
       add=TRUE,
       lwd=3)

mf_typo(x=bibT[bibT$transfert == "OUI",],
        var="INSEE",
        pal = c("lightgreen"),
        leg_pos = NA,
        add=TRUE)

mf_typo(x=bibT[bibT$transfert == "NON",],
        var="INSEE",
        pal = c("pink"),
        leg_pos = NA,
        add=TRUE)

mf_label(bibT,
         var="nom",
         cex=.6,
         overlap=FALSE)

mf_legend_t(pal = c("pink","lightgreen"),
            val=c("Bib. non-transférée","Bib. transférée"),
            pos="bottomleft",
            title=NA ) 

mf_title(txt = "Agglo’ Paris-Saclay")
mf_credits(txt="Réalisation: G. Raimondeau",pos="rightbottom")

# pour pdf
dev.off()
