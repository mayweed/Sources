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
#> select(GEOFLA_COMMUNE_2019_l93,LIBEPCI,EPCI)
#   LIBEPCI                                 EPCI                         geometry
#   <chr>                                   <chr>                   <POLYGON [m]>
# 1 CA Grand Paris Sud Seine Essonne Sénart 200059228 ((661187.8 6830962, 661097…
# 2 CA Grand Paris Sud Seine Essonne Sénart 200059228 ((663540.2 6836966, 663206…
# 3 CA Cœur d'Essonne Agglomération         200057859 ((647804.3 6839946, 647829…
# 4 CC Entre Juine et Renarde (CCEJR)       249100553 ((641175.4 6824331, 641115…
# 5 CA Grand Paris Sud Seine Essonne Sénart 200059228 ((655659.7 6834272, 655614…
# 6 CC Entre Juine et Renarde (CCEJR)       249100553 ((640876.4 6830906, 641023…
# 7 CA Étampois Sud-Essonne                 200017846 ((626105.7 6811057, 626149…
# 8 CA Grand Paris Sud Seine Essonne Sénart 200059228 ((662427.3 6834250, 662562…
# 9 CA Communauté Paris-Saclay              200056232 ((648033.6 6845229, 648095…
#10 CA Étampois Sud-Essonne                 200017846 ((632146.5 6818063, 632351…
# … with 184 more rows

# ttes les communes de cea
# GEOFLA_COMMUNE_2019_l93[GEOFLA_COMMUNE_2019_l93$EPCI == "200057859",]

#get everything 
load("/home/guillaume/DONNEES_R/GEOFLA_2019_l93.RData")

#xls : add data
bibSup <- read_excel(path="/home/guillaume/DONNEES_R/USAGERS_TX.xlsx",
                   sheet=1,
                   col_names=TRUE)
bibT <- st_read("/home/guillaume/SHP/bibTransferee.shp")
bibNT <- st_read("/home/guillaume/SHP/bibNonTransferee.shp")

#all is very important for NA!!
GEOFLA_COMMUNE_2019_l93 <- merge(GEOFLA_COMMUNE_2019_l93,
                                 bibSup,
                                 by.x="INSEE_COM",
                                 by.y="INSEE_COM",
                                 all.x=TRUE)

villeSansBib <- st_read("/home/guillaume/SHP/villeSansBib.shp")

x11() #use windows() or quartz() for mac
png("/home/guillaume/Desktop/INET/STAGES/Stage_pro/ESSONNE/CARTES_DIAG/focusCPSusagers.png",width=800,height=800,res=100)

target <- GEOFLA_EPCI91_2019_l93[1,]
mf_init(target,expandBB=c(0,0.15,0,0))

mf_typo(x=bibT[bibT$EPCI == "200056232",],
        var="INSEE_COM",
        pal = c("lightgreen"),
        leg_pos = NA,
        add=TRUE)

mf_typo(x=bibNT[bibNT$EPCI == "200056232",],
        var="INSEE_COM",
        pal = c("pink"),
        leg_pos = NA,
        add=TRUE)

mf_map(x=GEOFLA_EPCI91_2019_l93[1,],
       col=NA,
       border="black",
       add=TRUE,
       lwd=5)

mf_map(x= GEOFLA_COMMUNE_2019_l93[GEOFLA_COMMUNE_2019_l93$EPCI == "200056232",],
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

mf_credits(txt="Réalisation: MDE - Données issues du rapport SCRIB 2020",pos="bottomleft")

#mf_legend_t(title = NA,val = c("Pas de bib."), pal = "lightgrey")

# pour pdf
dev.off()

#wait please!!
#locator(1)
