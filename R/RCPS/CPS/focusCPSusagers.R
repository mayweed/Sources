library("sf")
library("mapsf")
library("readxl")
library("png")

#xls : add data
bibSup <- read_excel(path="C:/Users/Raimondeaug/Documents/DONNEES_R/USAGERS_TX.xlsx",
                   sheet=1,
                   col_names=TRUE)
bibT <- st_read("C:/Users/Raimondeaug/Documents/SHP/TRANSFERT_BIBLIOTHEQUES.shp")

#all is very important for NA!!
bibM <- merge(bibT,
              bibSup,
              by.x="INSEE",
              by.y="INSEE",
              all.x=TRUE)

x11() #use windows() or quartz() for mac
png("C:/Users/Raimondeaug/Documents/CARTES_R/CPS/focusCPSusagers.png",width=800,height=800,res=100)

target <- bibM
mf_init(target,expandBB=c(0,0.15,0,0))

mf_map(x=bibM,
       col=NA,
       border="black",
       add=TRUE,
       lwd=5)

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

mf_map(x= bibM,
       var="INSCRITS",
       type="prop",
       inches=.20,
       col="lightblue",
       symbol="circle",
       leg_pos="bottomright",
       leg_title="Nombre d’inscrits par commune.",
       add=T)

mf_title(txt = "Nombre d’inscrits par commune - CPS")

mf_legend_t(pal = c("pink","lightgreen"),
            val=c("Bib. non-transférée","Bib. transférée"),
            pos="topleft",
            title=NA ) 

mf_credits(txt="Réalisation: G. Raimondeau - Agglo’ Paris-Saclay",pos="rightbottom")

# pour pdf
dev.off()

#wait please!!
#locator(1)
