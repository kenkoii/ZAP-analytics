'use strict';
module.exports = function(sequelize, DataTypes) {
  var UserDailyProperty = sequelize.define('UserDailyProperty', {
    userDailyPropertyId: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true
    },
    loginDate: DataTypes.DATE,
    gacha: DataTypes.JSON,
    quest: DataTypes.JSON,
    event: DataTypes.JSON,
    cardAmount: DataTypes.INTEGER,
    playerLevel: DataTypes.INTEGER,
    reachStage: DataTypes.INTEGER,
    questionAmount: DataTypes.INTEGER,
    factoryAmount: DataTypes.INTEGER,
    strengthen: DataTypes.INTEGER,
    progress: DataTypes.INTEGER,
    sell: DataTypes.INTEGER,
    exceedLimit: DataTypes.INTEGER
  }, {
    classMethods: {
      associate: function(models) {
        UserDailyProperty.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'userId'
        });
      }
    }
  });
  return UserDailyProperty;
};