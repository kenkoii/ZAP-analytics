'use strict';
module.exports = function(sequelize, DataTypes) {
  var UserProperty = sequelize.define('UserProperty', {
    userId: {
      type: DataTypes.INTEGER,
      primaryKey: true
    },
    downloadDate: DataTypes.DATE,
    loginDate: DataTypes.DATE,
    os: DataTypes.INTEGER,
    ver: DataTypes.INTEGER,
    downloadVer: DataTypes.INTEGER,
    reachArea: DataTypes.INTEGER,
    wordLevel: DataTypes.INTEGER,
    grammarLevel: DataTypes.INTEGER,
    playerLevel: DataTypes.INTEGER,
    cardAmount: DataTypes.INTEGER,
    cardLimit: DataTypes.INTEGER,
    chocolate: DataTypes.INTEGER,
    gold: DataTypes.INTEGER,
    cookie: DataTypes.INTEGER,
    eventTicket: DataTypes.INTEGER,
    lottery: DataTypes.INTEGER,
    // deckInfo: DataTypes.JSON,
    lastHp: DataTypes.INTEGER,
    lastStr: DataTypes.INTEGER,
    lastTotal: DataTypes.INTEGER,
    bestHp: DataTypes.INTEGER,
    bestStr: DataTypes.INTEGER,
    bestTotal: DataTypes.INTEGER,
    leaderCard: DataTypes.INTEGER,
    helperCard: DataTypes.INTEGER,
    // settings: DataTypes.JSON,
    friend: DataTypes.INTEGER,
    friendLimit: DataTypes.INTEGER,
    questionTotal: DataTypes.INTEGER,
    factoryTotal: DataTypes.INTEGER,
    uraTotal: DataTypes.INTEGER,
    fast: DataTypes.INTEGER,
    light: DataTypes.INTEGER
  }, {
    classMethods: {
      associate: function(models) {
        
      }
    }
  });

  return UserProperty;
};

