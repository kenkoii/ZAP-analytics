'use strict';
module.exports = function(sequelize, DataTypes) {
  var UserPurchase = sequelize.define('UserPurchase', {
    userPurchaseId:  {
      type: DataTypes.INTEGER,
      primaryKey: true,
      autoIncrement: true
    },
    itemId: DataTypes.INTEGER,
    date: DataTypes.DATE,
    price: DataTypes.DECIMAL,
    isFirst: DataTypes.BOOLEAN
  }, {
    classMethods: {
      associate: function(models) {
        UserPurchase.belongsTo(models.UserProperty, {
          onDelete: "CASCADE",
          foreignKey: 'userId'
        });
      }
    }
  });
  return UserPurchase;
};