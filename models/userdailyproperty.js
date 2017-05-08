'use strict';
module.exports = function(sequelize, DataTypes) {
  var UserDailyProperty = sequelize.define('UserDailyProperty', {
    UserDailyPropertyID: {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true
    },
    LoginDate: DataTypes.DATE,
    Gacha: DataTypes.JSON,
    Quest: DataTypes.JSON,
    Event: DataTypes.JSON,
    CardAmount: DataTypes.INTEGER,
    PlayerLevel: DataTypes.INTEGER,
    ReachStage: DataTypes.INTEGER,
    QuestionAmount: DataTypes.INTEGER,
    FactoryAmount: DataTypes.INTEGER,
    Strengthen: DataTypes.INTEGER,
    Progress: DataTypes.INTEGER,
    Sell: DataTypes.INTEGER,
    ExceedLimit: DataTypes.INTEGER
  }, {
    classMethods: {
      associate: function(models) {
        UserDailyProperty.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'UserPropertyID'
        });
      }
    }
  });
  return UserDailyProperty;
};