'use strict';
module.exports = function(sequelize, DataTypes) {
  var UserPurchase = sequelize.define('UserPurchase', {
    UserPurchaseID:  {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true
    },
    ItemID: DataTypes.INTEGER,
    Date: DataTypes.DATE,
    Price: DataTypes.INTEGER,
    IsFirst: DataTypes.BOOLEAN
  }, {
    classMethods: {
      associate: function(models) {
        UserPurchase.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'UserPropertyID'
        });
      }
    }
  });
  return UserPurchase;
};