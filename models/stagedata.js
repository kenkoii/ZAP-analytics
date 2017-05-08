'use strict';
module.exports = function(sequelize, DataTypes) {
  var StageData = sequelize.define('StageData', {
    StageDataID: {
      primaryKey: true,
      type: DataTypes.INTEGER,
      autoIncrement: true
    },
    Date: DataTypes.DATE,
    State: DataTypes.INTEGER,
    IsFirst: DataTypes.BOOLEAN,
    Continue: DataTypes.INTEGER,
    PlayerLevel: DataTypes.INTEGER,
    StageID: DataTypes.INTEGER,
    ProvinceID: DataTypes.INTEGER,
    AreaID: DataTypes.INTEGER,
    Cookie: DataTypes.INTEGER,
    Chocolate: DataTypes.INTEGER,
    FriendID: DataTypes.INTEGER,
    Turn: DataTypes.INTEGER,
    Correct: DataTypes.INTEGER,
    Achievement1: DataTypes.BOOLEAN,
    Achievement2: DataTypes.BOOLEAN,
    Achievement3: DataTypes.BOOLEAN
  }, {
    classMethods: {
      associate: function(models) {
        StageData.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'UserPropertyID'
        });
      }
    }
  });
  return StageData;
};