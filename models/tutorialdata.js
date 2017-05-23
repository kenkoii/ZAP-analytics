'use strict';
module.exports = function(sequelize, DataTypes) {
  var TutorialData = sequelize.define('TutorialData', {
    tutorialDataId: {
      primaryKey: true,
      type: DataTypes.INTEGER
    },
    tutorialId: DataTypes.INTEGER,
    date: DataTypes.DATE
  }, {
    classMethods: {
      associate: function(models) {
        TutorialData.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'userId'
        });
      }
    }
  });
  return TutorialData;
};