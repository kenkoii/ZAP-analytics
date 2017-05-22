'use strict';
module.exports = function(sequelize, DataTypes) {
  var UserProperty = sequelize.define('UserProperty', {
    UserPropertyID: {
      type: DataTypes.INTEGER,
      primaryKey: true
    },
    DownloadDate: DataTypes.DATE,
    LoginDate: DataTypes.DATE,
    OS: DataTypes.INTEGER,
    Version: DataTypes.INTEGER,
    DownloadVersion: DataTypes.INTEGER,
    ReachArea: DataTypes.INTEGER,
    WordLevel: DataTypes.INTEGER,
    GrammarLevel: DataTypes.INTEGER,
    PlayerLevel: DataTypes.INTEGER,
    CardAmount: DataTypes.INTEGER,
    CardLimit: DataTypes.INTEGER,
    Chocolate: DataTypes.INTEGER,
    Gold: DataTypes.INTEGER,
    Cookie: DataTypes.INTEGER,
    EventTicket: DataTypes.INTEGER,
    Lottery: DataTypes.INTEGER,
    DeckInfo: DataTypes.JSON,
    LeaderCard: DataTypes.INTEGER,
    HelperCard: DataTypes.INTEGER,
    Settings: DataTypes.JSON,
    Friend: DataTypes.INTEGER,
    FriendLimit: DataTypes.INTEGER,
    QuestionTotal: DataTypes.INTEGER,
    FactoryTotal: DataTypes.INTEGER,
    UraTotal: DataTypes.INTEGER,
    Fast: DataTypes.INTEGER,
    Light: DataTypes.INTEGER
  }, {
    classMethods: {
      associate: function(models) {
        
      }
    }
  });

  return UserProperty;
};

