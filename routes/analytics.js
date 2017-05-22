var express = require('express');
var router = express.Router();
var models = require('../models/index');
var sequelize = models.sequelize;

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'FreCre Analytics' });
});

router.post('/', function(req, res, next) {
  const items = JSON.parse(req.body);
  for(var i = 0; i < items.length; i++){
    var item = items[i];
    switch(item.className){
      case 'Stage':
        insertStage(item.params);
        break;
      case 'UserProperty':
        insertUserProperty(item.params);
        break;
      case 'UserPurchase':
        insertUserPurchase(item.params);
        break;
      case 'UserDailyProperty':
        insertUserDailyProperty(item.params);
        break;
      case 'Tutorial':
        insertTutorial(item.params);
        break;
    }
  }
});

function insertStage(item) {
  models.StageData.create(
     item
  ).then(function(stageData) {
    // res.json(stageData);
  });
}

function insertUserProperty(item) {
  models.UserProperty.upsert(
    item,
      {
      where: {
        UserPropertyID: item.UserPropertyID
      }
  }).then(function(userProperty) {
    // res.json(userProperty);
  });
}

function insertUserPurchase(item) {
  models.UserPurchase.create(
    item
  ).then(function(userPurchase) {
    // res.json(userPurchase);
  });
}

function insertUserDailyProperty(item) {
  models.UserDailyProperty.create(
    item
  ).then(function(userDailyProperty) {
    // res.json(userDailyProperty);
  });
}

function insertTutorial(item) {
  models.TutorialData.create({
    TutorialDataID: parseInt(item.UserPropertyID.toString() + item.TutorialID.toString()),
    UserPropertyID: item.UserPropertyID,
    TutorialID: item.TutorialID,
    Date: item.Date
  }).then(function(tutorialData) {
    // res.json(tutorialData);
  });
}


module.exports = router;

