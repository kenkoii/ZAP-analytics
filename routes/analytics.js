var express = require('express');
var router = express.Router();
var models = require('../models/index');
var sequelize = models.sequelize;

/* GET home page. */
router.get('/', function(req, res, next) {
  res.render('index', { title: 'FreCre Analytics' });
});

router.post('/', function(req, res, next) {
  const items = req.body;
  let count = 0;
  let promises = [];
  for(var i = 0; i < items.length; i++){
    var item = items[i];
    switch(item.className){
      case 'Stage':
        promises.push(insertStage(item.params));
        break;
      case 'UserProperty':
        promises.push(insertUserProperty(item.params));
        break;
      case 'UserPurchase':
        promises.push(insertUserPurchase(item.params));
        break;
      case 'UserDailyProperty':
        promises.push(insertUserDailyProperty(item.params));
        break;
      case 'Tutorial':
        promises.push(insertTutorial(item.params));
        break;
    }
  }
  Promise.all(promises)
    .then((results) => {
      res.json(results.length);
    }).catch((err) => {
      console.log(err);
      res.json(-1);
    });
});

function insertStage(item) {
  return models.StageData.create(
     item
  );
}

function insertUserProperty(item) {
  return models.UserProperty.upsert(
    item,
      {
      where: {
        userId: item.userId
      }
  });
}

function insertUserPurchase(item) {
  return models.UserPurchase.create(
    item
  );
}

function insertUserDailyProperty(item) {
  return models.UserDailyProperty.create(
    item
  );
}

function insertTutorial(item) {
  var pKey = parseInt(item.userId.toString() + item.tutorialId.toString());
  return models.TutorialData.upsert({
    tutorialDataId: pKey,
    userId: item.userId,
    tutorialId: item.tutorialId,
    date: item.date
  },{
    where: {
      tutorialDataId: pKey
    }
  });
}


module.exports = router;

