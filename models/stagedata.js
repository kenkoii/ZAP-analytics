'use strict';
module.exports = function(sequelize, DataTypes) {
  var StageData = sequelize.define('StageData', {
    StageDataId: {
      primaryKey: true,
      type: DataTypes.INTEGER,
      autoIncrement: true
    },
    date: DataTypes.DATE,
    state: DataTypes.INTEGER,
    isFirst: DataTypes.BOOLEAN,
    continue: DataTypes.INTEGER,
    playerLevel: DataTypes.INTEGER,
    stageId: DataTypes.INTEGER,
    provinceId: DataTypes.INTEGER,
    areaId: DataTypes.INTEGER,
    cookie: DataTypes.INTEGER,
    chocolate: DataTypes.INTEGER,
    friendId: DataTypes.INTEGER,
    turn: DataTypes.INTEGER,
    correct: DataTypes.INTEGER,
    achievement1: DataTypes.BOOLEAN,
    achievement2: DataTypes.BOOLEAN,
    achievement3: DataTypes.BOOLEAN
  }, {
    classMethods: {
      associate: function(models) {
        StageData.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'userId'
        });
      }
    }
  });
  return StageData;
};