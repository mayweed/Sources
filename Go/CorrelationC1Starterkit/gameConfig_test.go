package main

//idea: test the config before testing the turn

var testStringUi = []byte(`{"debug":{"printMapString":false,"printTStrings":false,"printActStrings":false,"printHitStrings":false,"printPlayerInputStrings":true,"printBotErrors":true,"printPlayerGetHitStrings":false},"unitInformation":[{"damage":0.0,"cost":1.0,"getHitRadius":0.51,"display":"Filter","range":0.0,"shorthand":"FF","stability":60.0},{"damage":0.0,"cost":4.0,"getHitRadius":0.51,"shieldAmount":10.0,"display":"Encryptor","range":3.0,"shorthand":"EF","stability":30.0},{"damage":4.0,"cost":3.0,"getHitRadius":0.51,"display":"Destructor","range":3.0,"shorthand":"DF","stability":75.0},{"damageI":1.0,"damageToPlayer":1.0,"cost":1.0,"getHitRadius":0.51,"damageF":1.0,"display":"Ping","range":3.0,"shorthand":"PI","stability":15.0,"speed":0.5},{"damageI":3.0,"damageToPlayer":1.0,"cost":3.0,"getHitRadius":0.51,"damageF":3.0,"display":"EMP","range":5.0,"shorthand":"EI","stability":5.0,"speed":0.25},{"damageI":10.0,"damageToPlayer":1.0,"cost":1.0,"getHitRadius":0.51,"damageF":0.0,"display":"Scrambler","range":3.0,"shorthand":"SI","stability":40.0,"speed":0.25},{"display":"Remove","shorthand":"RM"}],"timingAndReplay":{"waitTimeBotMax":35000.0,"playWaitTimeBotMax":40000.0,"waitTimeManual":1820000.0,"waitForever":false,"waitTimeBotSoft":5000.0,"playWaitTimeBotSoft":10000.0,"replaySave":1.0,"playReplaySave":0.0,"storeBotTimes":true,"waitTimeStartGame":3000.0,"waitTimeEndGame":3000.0},"resources":{"turnIntervalForBitCapSchedule":10.0,"turnIntervalForBitSchedule":10.0,"bitRampBitCapGrowthRate":5.0,"roundStartBitRamp":10.0,"bitGrowthRate":1.0,"startingHP":30.0,"maxBits":999999.0,"bitsPerRound":5.0,"coresPerRound":4.0,"coresForPlayerDamage":1.0,"startingBits":5.0,"bitDecayPerRound":0.33333,"startingCores":25.0},"mechanics":{"basePlayerHealthDamage":1.0,"damageGrowthBasedOnY":0.0,"bitsCanStackOnDeployment":true,"destroyOwnUnitRefund":0.5,"destroyOwnUnitsEnabled":true,"stepsRequiredSelfDestruct":5.0,"selfDestructRadius":1.5,"shieldDecayPerFrame":0.15,"meleeMultiplier":0.0,"destroyOwnUnitDelay":1.0,"rerouteMidRound":true,"firewallBuildTime":0.0}}`)
