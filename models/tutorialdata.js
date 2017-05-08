'use strict';
module.exports = function(sequelize, DataTypes) {
  var TutorialData = sequelize.define('TutorialData', {
    TutorialDataID: {
      primaryKey: true,
      type: DataTypes.INTEGER
    },
    TutorialID: DataTypes.INTEGER,
    Date: DataTypes.DATE
  }, {
    classMethods: {
      associate: function(models) {
        TutorialData.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'UserPropertyID'
        });
      }
    }
  });
  return TutorialData;
};